// Copyright 2016 Cyako Author

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cyako

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type serviceConfig struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type config struct {
	Service []serviceConfig `json:"service"`
}

type serviceSupport struct {
	Name          string
	AfterReceive  bool
	BeforeProcess bool
	AfterProcess  bool
	BeforeSend    bool
	AfterSend     bool
}

type service struct {
	Map               map[string]interface{}
	Support           []serviceSupport
	AfterReceiveFunc  []func(*Req)
	BeforeProcessFunc []func(*Ctx)
	AfterProcessFunc  []func(*Ctx)
	BeforeSendFunc    []func(*Res)
	AfterSendFunc     []func(*Res)
}

type Cyako struct {
	Config       config
	Service      service
	ProcessorMap map[string]*Processor
	// Server       http.Server
}

func (c *Cyako) loadConfig() {
	var err error
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(" Error config.json:", err)
	} else {
		err = json.Unmarshal(data, &c.Config)
		if err != nil {
			fmt.Println(" Error config.json:", err)
		}
	}
}

/*
	init
*/

var cyako *Cyako

func init() {
	cyako = &Cyako{
		Service: service{
			Map: make(map[string]interface{}),
		},
		ProcessorMap: make(map[string]*Processor),
	}
	cyako.loadConfig()
}

/*
	global
*/

// return cyako package's global object: cyako
func Ins() *Cyako {
	return cyako
}

// Svc is an alias to Ins().Service.Map
var Svc map[string]interface{}

// Service is an alias to Ins().Service.Map
var Service map[string]interface{}

func init() {
	Svc = Ins().Service.Map
	Service = Ins().Service.Map
}

// used in Processor Module package to load itself
func LoadModule(x interface{}) {
	Ins().loadModule(x)
}

// used in Service package to load itself
func LoadService(x interface{}) {
	Ins().loadService(x)
}
