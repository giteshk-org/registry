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

package server

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/apigee/registry/rpc"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const verbose = false

const topicName = "changes"

func (s *RegistryServer) notify(change rpc.Notification_Change, resource string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, s.projectID)
	if err != nil {
		return err
	}
	// Ensure that topic exists.
	{
		_, err := client.CreateTopic(context.Background(), topicName)
		if err != nil {
			code := status.Code(err)
			if code != codes.AlreadyExists {
				return err
			}
		}
	}
	// Get the topic.
	topic := client.Topic(topicName)
	defer topic.Stop()
	// Create the notification
	n := &rpc.Notification{}
	n.Change = change
	n.Resource = resource
	n.ChangeTime, err = ptypes.TimestampProto(time.Now())
	// Marshal the notification.
	m, err := (&jsonpb.Marshaler{}).MarshalToString(n)
	if err != nil {
		return err
	}
	// Send the notification.
	log.Printf("sending %+s", m)
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(m),
	})
	results = append(results, r)
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			return err
		}
		if verbose {
			log.Printf("Published a message with a message ID: %s", id)
		}
	}
	return nil
}
