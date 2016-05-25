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
)

// middleware packages use LoadMiddleware to load itself
func (c *Cyako) loadMiddleware(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	middlewareName := t.Name()
	c.Middleware.Map[middlewareName] = x
	var support = middlewareSupport{
		Name: middlewareName,
	}
	for _, name := range []string{"AfterReceive", "BeforeProcess", "AfterProcess", "BeforeSend", "AfterSend"} {
		if method, ok := t.MethodByName(name); ok {
			switch name {
			case "AfterReceive":
				support.AfterReceive = true
				c.Middleware.AfterReceiveFunc = append(c.Middleware.AfterReceiveFunc, func(req *Req) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(req)})
				})
			case "BeforeProcess":
				support.BeforeProcess = true
				c.Middleware.BeforeProcessFunc = append(c.Middleware.BeforeProcessFunc, func(ctx *Ctx) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
				})
			case "AfterProcess":
				support.AfterProcess = true
				c.Middleware.AfterProcessFunc = append(c.Middleware.AfterProcessFunc, func(ctx *Ctx) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
				})
			case "BeforeSend":
				support.BeforeSend = true
				c.Middleware.BeforeSendFunc = append(c.Middleware.BeforeSendFunc, func(res *Res) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(res)})
				})
			case "AfterSend":
				support.AfterSend = true
				c.Middleware.AfterSendFunc = append(c.Middleware.AfterSendFunc, func(res *Res) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(res)})
				})
			default:
				fmt.Println("Middleware load logic error.")
			}
		}
	}
	c.Middleware.Support = append(c.Middleware.Support, support)
}

/*
	exec methods
*/

func (c *Cyako) AfterReceive(req *Req) {
	for _, fn := range c.Middleware.AfterReceiveFunc {
		fn(req)
	}
}

func (c *Cyako) BeforeProcess(ctx *Ctx) {
	for _, fn := range c.Middleware.BeforeProcessFunc {
		fn(ctx)
	}
}

func (c *Cyako) AfterProcess(ctx *Ctx) {
	for _, fn := range c.Middleware.AfterProcessFunc {
		fn(ctx)
	}
}

func (c *Cyako) BeforeSend(res *Res) {
	for _, fn := range c.Middleware.BeforeSendFunc {
		fn(res)
	}
}

func (c *Cyako) AfterSend(res *Res) {
	for _, fn := range c.Middleware.AfterSendFunc {
		fn(res)
	}
}
