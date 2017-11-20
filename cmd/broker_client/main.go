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

package main

import (
	"context"
	"encoding/json"
	"fmt"

	flag "github.com/spf13/pflag"

	"github.com/getamis/service-broker/broker"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8080, "The server port")
}

func main() {
	client, err := broker.NewClient(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Failed to create client, %v\n", err)
		return
	}

	catalog, err := client.GetCatalog(context.TODO())
	if err != nil {
		fmt.Printf("Failed to get catalog, %v\n", err)
		return
	}

	jsonBytes, _ := json.MarshalIndent(catalog, "", "\t")
	fmt.Println(string(jsonBytes))
}
