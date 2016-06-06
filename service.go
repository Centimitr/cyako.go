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
	"fmt"
	"reflect"
	"strings"
)

func getName(t reflect.Type) string {
	if strings.HasPrefix(t.String(), "*") {
		return strings.Split(t.String(), ".")[1]
	} else {
		return t.Name()
	}
}

// service packages use LoadService to load itself
func (c *Cyako) loadService(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	serviceName := getName(t)
	c.Service.Map[serviceName] = x
	var support = serviceSupport{
		Name: serviceName,
	}
	for _, name := range []string{"AfterReceive", "BeforeProcess", "AfterProcess", "BeforeSend", "AfterSend"} {
		if method, ok := t.MethodByName(name); ok {
			switch name {
			case "AfterReceive":
				support.AfterReceive = true
				c.Service.AfterReceiveFunc = append(c.Service.AfterReceiveFunc, func(req *Req) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(req)})
				})
			case "BeforeProcess":
				support.BeforeProcess = true
				c.Service.BeforeProcessFunc = append(c.Service.BeforeProcessFunc, func(ctx *Ctx) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
				})
			case "AfterProcess":
				support.AfterProcess = true
				c.Service.AfterProcessFunc = append(c.Service.AfterProcessFunc, func(ctx *Ctx) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
				})
			case "BeforeSend":
				support.BeforeSend = true
				c.Service.BeforeSendFunc = append(c.Service.BeforeSendFunc, func(res *Res) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(res)})
				})
			case "AfterSend":
				support.AfterSend = true
				c.Service.AfterSendFunc = append(c.Service.AfterSendFunc, func(res *Res) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(res)})
				})
			default:
				fmt.Println("Service load logic error.")
			}
		}
	}
	c.Service.Support = append(c.Service.Support, support)
}

/*
	exec methods
*/

func (c *Cyako) AfterReceive(req *Req) {
	for _, fn := range c.Service.AfterReceiveFunc {
		fn(req)
	}
}

func (c *Cyako) BeforeProcess(ctx *Ctx) {
	for _, fn := range c.Service.BeforeProcessFunc {
		fn(ctx)
	}
}

func (c *Cyako) AfterProcess(ctx *Ctx) {
	for _, fn := range c.Service.AfterProcessFunc {
		fn(ctx)
	}
}

func (c *Cyako) BeforeSend(res *Res) {
	for _, fn := range c.Service.BeforeSendFunc {
		fn(res)
	}
}

func (c *Cyako) AfterSend(res *Res) {
	for _, fn := range c.Service.AfterSendFunc {
		fn(res)
	}
}
