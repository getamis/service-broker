// Copyright 2017 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package broker

import (
	"context"

	"github.com/inconshreveable/log15"
	"google.golang.org/grpc"

	"github.com/getamis/service-broker/broker/pb"
	"github.com/getamis/service-broker/log"
)

func NewClient(addr string, opts ...ClientOption) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	o := &clientOption{}

	for _, opt := range opts {
		o = opt(o)
	}

	return newClient(conn, o)
}

// ----------------------------------------------------------------------------

func newClient(conn *grpc.ClientConn, o *clientOption) (*Client, error) {
	c := &Client{
		conn:   conn,
		client: pb.NewBrokerClient(conn),
		logger: o.logger,
	}

	if c.logger == nil {
		c.logger = log15.New()
	}

	return c, nil
}

type Client struct {
	conn   *grpc.ClientConn
	logger log.Logger
	client pb.BrokerClient
}

func (c *Client) GetCatalog(ctx context.Context) (*pb.Catalog, error) {
	return c.client.GetCatalog(ctx, &pb.Empty{})
}

func (c *Client) GetServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (*pb.ServiceInstanceResponse, error) {
	return c.client.GetServiceInstance(ctx, req)
}

func (c *Client) CreateServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (*pb.ServiceInstanceResponse, error) {
	return c.client.CreateServiceInstance(ctx, req)
}

func (c *Client) RemoveServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) error {
	_, err := c.client.RemoveServiceInstance(ctx, req)
	return err
}

func (c *Client) Bind(ctx context.Context, req *pb.ServiceBindingRequest) (*pb.ServiceBindingResponse, error) {
	return c.client.Bind(ctx, req)
}

func (c *Client) Unbind(ctx context.Context, req *pb.ServiceBindingRequest) (*pb.ServiceBindingResponse, error) {
	return c.client.Unbind(ctx, req)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
