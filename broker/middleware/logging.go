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

package middleware

import (
	"context"
	"encoding/json"

	"github.com/getamis/service-broker/broker"
	"github.com/getamis/service-broker/broker/pb"
	"github.com/getamis/service-broker/log"
)

func Logging(logger log.Logger) broker.Middleware {
	return func(srv pb.BrokerServer) pb.BrokerServer {
		return loggingMiddleware{
			logger: logger,
			next:   srv,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   pb.BrokerServer
}

func (mw loggingMiddleware) GetCatalog(ctx context.Context, req *pb.Empty) (resp *pb.Catalog, err error) {
	defer func() {
		reqStr, _ := json.MarshalIndent(req, "", "\t")
		respStr, _ := json.MarshalIndent(resp, "", "\t")
		mw.logger.Debug("", "method", "GetCatalog", "req", string(reqStr), "resp", string(respStr))
	}()

	return mw.next.GetCatalog(ctx, req)
}

func (mw loggingMiddleware) GetServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (resp *pb.ServiceInstanceResponse, err error) {
	defer func() {
		reqStr, _ := json.MarshalIndent(req, "", "\t")
		respStr, _ := json.MarshalIndent(resp, "", "\t")
		mw.logger.Debug("", "method", "GetServiceInstance", "req", string(reqStr), "resp", string(respStr))
	}()

	return mw.next.GetServiceInstance(ctx, req)
}

func (mw loggingMiddleware) CreateServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (resp *pb.ServiceInstanceResponse, err error) {
	defer func() {
		reqStr, _ := json.MarshalIndent(req, "", "\t")
		respStr, _ := json.MarshalIndent(resp, "", "\t")
		mw.logger.Debug("", "method", "CreateServiceInstance", "req", string(reqStr), "resp", string(respStr))
	}()

	return mw.next.CreateServiceInstance(ctx, req)
}

func (mw loggingMiddleware) RemoveServiceInstance(ctx context.Context, req *pb.ServiceInstanceRequest) (resp *pb.Empty, err error) {
	defer func() {
		reqStr, _ := json.MarshalIndent(req, "", "\t")
		respStr, _ := json.MarshalIndent(resp, "", "\t")
		mw.logger.Debug("", "method", "RemoveServiceInstance", "req", string(reqStr), "resp", string(respStr))
	}()

	return mw.next.RemoveServiceInstance(ctx, req)
}

func (mw loggingMiddleware) Bind(ctx context.Context, req *pb.ServiceBindingRequest) (resp *pb.ServiceBindingResponse, err error) {
	defer func() {
		reqStr, _ := json.MarshalIndent(req, "", "\t")
		respStr, _ := json.MarshalIndent(resp, "", "\t")
		mw.logger.Debug("", "method", "Bind", "req", string(reqStr), "resp", string(respStr))
	}()

	return mw.next.Bind(ctx, req)
}

func (mw loggingMiddleware) Unbind(ctx context.Context, req *pb.ServiceBindingRequest) (resp *pb.ServiceBindingResponse, err error) {
	defer func() {
		reqStr, _ := json.MarshalIndent(req, "", "\t")
		respStr, _ := json.MarshalIndent(resp, "", "\t")
		mw.logger.Debug("", "method", "Unbind", "req", string(reqStr), "resp", string(respStr))
	}()

	return mw.next.Unbind(ctx, req)
}
