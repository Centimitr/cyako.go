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
	// "encoding/json"
	// "errors"
	"fmt"
	"golang.org/x/net/websocket"
	// "strings"
)

type temp map[string]interface{} // use for service maintain state

func (t *temp) getRealKey(scope, key string) string {
	return scope + "." + key
}

func (t temp) Get(scope, key string) interface{} {
	return t[t.getRealKey(scope, key)]
}

func (t temp) Put(scope, key string, v interface{}) {
	t[t.getRealKey(scope, key)] = v
}

type Req struct {
	Id     string                 `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
	Data   interface{}            `json:"data"`
	Temp   temp                   // use for service maintain state
}

type Res struct {
	Id     string                 `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
	Data   interface{}            `json:"data"`
	Error  interface{}            `json:"error"`
	Temp   temp                   `json:"-"` // use for service maintain state
}

type Ctx struct {
	Conn *websocket.Conn
	res  *Res
	req  *Req
	// reqParams    map[string]interface{}
	ParamConfigs []*ParamConfig
	echoParams   []string
	Id           string
	Method       string
	Service      map[string]interface{}
	Params       map[string]interface{}
	Data         interface{}
	Error        CtxError
	Temp         temp // use for service maintain state
}

type ParamConfig struct {
	Key      string `json:"Key"`
	Required bool   `json:"Required"`
	Default  string `json:"Default"`
	Echo     bool   `json:"Echo"`
}

/*
	error
*/

type CtxError struct {
	Warn  []string
	Fatal []string
}

func (c *CtxError) NewFatal(info string) {
	c.Fatal = append(c.Fatal, info)
}

func (c *CtxError) NewWarn(info string) {
	c.Warn = append(c.Warn, info)
}

/*
	init
*/
func (r *Req) Init() {
	r.Params = make(map[string]interface{})
	r.Temp = make(map[string]interface{})
}

func (r *Res) Init() {
	r.Params = make(map[string]interface{})
}

func (c *Ctx) Init() {
	c.Service = cyako.Service.Map
	c.Params = make(map[string]interface{})
	// c.reqParams = make(map[string]interface{})
	// c.parseParams()
}

// func (c *Ctx) parseParams() {
// 	s := c.req.Params
// 	err := json.Unmarshal([]byte(s), &c.reqParams)
// 	if err != nil {
// 		c.Error.NewFatal("Params parse error.")
// 	}
// }

/*
	context methods used in processors
*/

func (c *Ctx) getReqParamString(key string) string {
	// param := c.reqParams[key]
	param := c.req.Params[key]
	switch param.(type) {
	case string:
		return param.(string)
	case float64:
		return fmt.Sprint(param.(float64))
	default:
		c.Error.NewWarn(fmt.Sprint("Param type error, not a known type."))
		return fmt.Sprint(param)
	}
}

func (c *Ctx) Set(data interface{}) {
	fmt.Println(c.req.Params)
	var setWitchConfig = func(p *ParamConfig) {
		// add paramConfig to context for docgen etc.
		c.ParamConfigs = append(c.ParamConfigs, p)
		param, isParamExist := c.req.Params[p.Key]
		switch {
		case p.Echo:
			c.echoParams = append(c.echoParams, p.Key)
			fallthrough
		case isParamExist:
			c.Params[p.Key] = param
		case p.Default != "":
			c.Params[p.Key] = p.Default
		case p.Required:
			c.Error.NewFatal("Lack required param.")
			// default:
			// fmt.Print("Key:", c.getReqParamString(p.Key))
			// c.Params[p.Key] = c.getReqParamString(p.Key)
		}
	}
	switch d := data.(type) {
	case *ParamConfig:
		setWitchConfig(d)
	case []*ParamConfig:
		for _, c := range d {
			setWitchConfig(c)
		}
	default:
		c.Error.NewWarn("Error params to *Ctx.Set().")
	}
}

func (c *Ctx) Get(key string) interface{} {
	return c.Params[key]
}

/*
	set res
*/
func (c *Ctx) setResParams() {
	// var toEscaped = func(s string) string {
	// 	return strings.Replace(s, `"`, `\"`, -1)
	// }
	// // var params []string
	// // var stringMapMarshal = func(m map[string]string) string {
	// // 	var kvs []string
	// // 	for k, v := range m {
	// // 		kvs = append(kvs, `"`+toEscaped(k)+`":"`+toEscaped(v)+`"`)
	// // 	}
	// // 	return "{" + strings.Join(kvs, ",") + "}"
	// // }

	// // will be replaced with mature convert solution, here is a temporary process
	// var stringMapPartlyMarshal = func(m map[string]string, keys []string) (string, error) {
	// 	var kvs []string
	// 	var err error
	// 	for _, k := range keys {
	// 		if v, ok := m[k]; ok {
	// 			kvs = append(kvs, `"`+toEscaped(k)+`":"`+toEscaped(v)+`"`)
	// 		} else {
	// 			err = errors.New("Cannot find one given key in the map.")
	// 		}
	// 	}
	// 	return "{" + strings.Join(kvs, ",") + "}", err
	// }
	// json, _ := stringMapPartlyMarshal(c.Params, c.echoParams)
	// c.res.Params = json
	// var kvs []string
	// var err error
	for _, k := range c.echoParams {
		if v, ok := c.Params[k]; ok {
			c.res.Params[k] = v
		} else {
			// err = errors.New("EchoParams: Cannot find one given key in the map.")
		}
	}
}
