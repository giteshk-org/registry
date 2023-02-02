// Copyright 2022 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package patch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/apigee/registry/cmd/registry/core"
	"github.com/apigee/registry/gapic"
	"github.com/apigee/registry/pkg/connection"
	"github.com/apigee/registry/pkg/models"
	"github.com/apigee/registry/rpc"
	"github.com/apigee/registry/server/registry/names"
	"gopkg.in/yaml.v3"
)

// NewApiSpec allows an API spec to be individually exported as a YAML file.
func NewApiSpec(ctx context.Context, client *gapic.RegistryClient, message *rpc.ApiSpec, nested bool) (*models.ApiSpec, error) {
	specName, err := names.ParseSpecRevision(message.Name)
	if err != nil {
		return nil, err
	}
	var artifacts []*models.Artifact
	if nested {
		artifacts, err = collectChildArtifacts(ctx, client, specName.Artifact("-"))
		if err != nil {
			return nil, err
		}
	}
	return &models.ApiSpec{
		Header: models.Header{
			ApiVersion: RegistryV1,
			Kind:       "Spec",
			Metadata: models.Metadata{
				Name:        specName.SpecID,
				Parent:      names.ExportableName(specName.Parent(), specName.ProjectID),
				Labels:      message.Labels,
				Annotations: message.Annotations,
			},
		},
		Data: models.ApiSpecData{
			FileName:    message.Filename,
			Description: message.Description,
			MimeType:    message.MimeType,
			SourceURI:   message.SourceUri,
			Artifacts:   artifacts,
		},
	}, nil
}

func applyApiSpecPatchBytes(
	ctx context.Context,
	client connection.RegistryClient,
	bytes []byte,
	project string,
	filename string) error {
	var spec models.ApiSpec
	err := yaml.Unmarshal(bytes, &spec)
	if err != nil {
		return err
	}
	return applyApiSpecPatch(ctx, client, &spec, project, filename)
}

func specName(parent string, metadata models.Metadata) (names.Spec, error) {
	if metadata.Parent != "" {
		parent = parent + "/" + metadata.Parent
	}
	version, err := names.ParseVersion(parent)
	if err != nil {
		return names.Spec{}, err
	}
	return version.Spec(metadata.Name), nil
}

func applyApiSpecPatch(
	ctx context.Context,
	client connection.RegistryClient,
	spec *models.ApiSpec,
	parent string,
	filename string) error {
	name, err := specName(parent, spec.Metadata)
	if err != nil {
		return err
	}
	req := &rpc.UpdateApiSpecRequest{
		ApiSpec: &rpc.ApiSpec{
			Name:        name.String(),
			Filename:    spec.Data.FileName,
			Description: spec.Data.Description,
			MimeType:    spec.Data.MimeType,
			SourceUri:   spec.Data.SourceURI,
			Labels:      spec.Metadata.Labels,
			Annotations: spec.Metadata.Annotations,
		},
		AllowMissing: true,
	}
	// if the spec's filename points to a local file, use that as the spec's contents
	if filename != "" {
		body, err := os.ReadFile(filepath.Join(filepath.Dir(filename), spec.Data.FileName))
		if err == nil {
			if strings.Contains(spec.Data.MimeType, "+gzip") {
				body, err = core.GZippedBytes(body)
				if err != nil {
					return err
				}
			}
			req.ApiSpec.Contents = body
		}
	}
	// if we didn't find the spec body in a file, try to read it from the SourceURI
	if req.ApiSpec.Contents == nil && spec.Data.SourceURI != "" {
		u, err := url.ParseRequestURI(spec.Data.SourceURI)
		if err != nil {
			return err
		}
		switch u.Scheme {
		case "http", "https":
			resp, err := http.Get(spec.Data.SourceURI)
			if err != nil {
				return err
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			if strings.Contains(spec.Data.MimeType, "+gzip") {
				body, err = core.GZippedBytes(body)
				if err != nil {
					return err
				}
			}
			req.ApiSpec.Contents = body
		case "file":
			// Remove leading slash from path.
			// We expect to load from paths relative to the working directory,
			// but users can add an additional slash to specify a global path.
			path := strings.TrimPrefix(u.Path, "/")
			info, err := os.Stat(path)
			if err != nil {
				return err
			}
			if info.IsDir() {
				recursive := true
				d := u.Query()["recursive"]
				if len(d) > 0 {
					recursive, err = strconv.ParseBool(d[0])
					if err != nil {
						return err
					}
				}
				contents, err := core.ZipArchiveOfPath(path, "", recursive)
				if err != nil {
					return err
				}
				req.ApiSpec.Contents = contents.Bytes()
			} else {
				body, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				if strings.Contains(spec.Data.MimeType, "+gzip") {
					body, err = core.GZippedBytes(body)
					if err != nil {
						return err
					}
				}
				req.ApiSpec.Contents = body
			}
		}
	}
	_, err = client.UpdateApiSpec(ctx, req)
	if err != nil {
		return fmt.Errorf("UpdateApiSpec: %s", err)
	}
	for _, artifactPatch := range spec.Data.Artifacts {
		err = applyArtifactPatch(ctx, client, artifactPatch, name.String())
		if err != nil {
			return err
		}
	}
	return nil
}
