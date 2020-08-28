// Copyright 2020 Google LLC. All Rights Reserved.
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

package datastore

import (
	"cloud.google.com/go/datastore"
	"github.com/apigee/registry/server/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Query represents a query in a storage provider
type Query struct {
	query *datastore.Query
}

// NewQuery creates a new query.
func (c *Client) NewQuery(kind string) storage.Query {
	return &Query{query: datastore.NewQuery(kind)}
}

func (q *Query) Filter(filter string, value interface{}) storage.Query {
	return &Query{query: q.query.Filter(filter, value)}
}

func (q *Query) Distinct() storage.Query {
	return &Query{query: q.query.Distinct()}

}

func (q *Query) Order(order string) storage.Query {
	return &Query{query: q.query.Order(order)}
}

// QueryApplyCursor applies a cursor to a query so that results will start at the cursor.
func (c *Client) QueryApplyCursor(q storage.Query, cursorStr string) (storage.Query, error) {
	if cursorStr != "" {
		cursor, err := datastore.DecodeCursor(cursorStr)
		if err != nil {
			return nil, internalError(err)
		}
		q = &Query{query: q.(*Query).query.Start(cursor)}
	}
	return q, nil
}

// internalError ...
func internalError(err error) error {
	if err == nil {
		return nil
	}
	// TODO: selectively mask error details depending on caller privileges
	return status.Error(codes.Internal, err.Error())
}
