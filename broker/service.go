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
	"net"

	"google.golang.org/grpc"

	"github.com/getamis/service-broker/broker/pb"
)

func New(opts ...Option) (*broker, error) {
	o := &option{}

	for _, opt := range opts {
		o = opt(o)
	}

	return newBroker(o)
}

func newBroker(o *option) (b *broker, err error) {
	b = &broker{
		server: grpc.NewServer(o.grpcServerOptions...),
	}

	allMw := func(server pb.BrokerServer) pb.BrokerServer {
		for _, mw := range o.middlewares {
			server = mw(server)
		}

		return server
	}

	pb.RegisterBrokerServer(b.server, allMw(b))

	return b, nil
}

// ----------------------------------------------------------------------------

type broker struct {
	server  *grpc.Server
	options *option
}

func (b *broker) Serve(l net.Listener) error {
	return b.server.Serve(l)
}

func (b *broker) GetCatalog(ctx context.Context, req *pb.Empty) (*pb.Catalog, error) {
	return &pb.Catalog{
		Services: []*pb.Service{
			&pb.Service{
				Name:        "example",
				Id:          "1234567890",
				Description: "This is an example service",
			},
		},
	}, nil
}

func (b *broker) GetServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (*pb.ServiceInstanceResponse, error) {
	return &pb.ServiceInstanceResponse{}, nil
}

func (b *broker) CreateServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (*pb.ServiceInstanceResponse, error) {
	return &pb.ServiceInstanceResponse{}, nil
}

func (b *broker) RemoveServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (b *broker) Bind(ctx context.Context, req *pb.ServiceBindingRequest) (*pb.ServiceBindingResponse, error) {
	return &pb.ServiceBindingResponse{}, nil
}

func (b *broker) Unbind(ctx context.Context, req *pb.ServiceBindingRequest) (*pb.ServiceBindingResponse, error) {
	return &pb.ServiceBindingResponse{}, nil
}
