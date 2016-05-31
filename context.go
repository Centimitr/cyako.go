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
	"errors"
	"fmt"
	"strings"
)

type temp map[string]interface{} // use for middleware maintain state

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
	Id     string `json:"id"`
	Method string `json:"method"`
	Params string `json:"params"`
	Data   string `json:"data"`
	Temp   temp   // use for middleware maintain state
}

type Res struct {
	Id     string `json:"id"`
	Method string `json:"method"`
	Params string `json:"params"`
	Data   string `json:"data"`
	Error  string `json:"error"`
	Temp   temp   `json:"-"` // use for middleware maintain state
}

type Ctx struct {
	res          *Res
	req          *Req
	reqParams    map[string]interface{}
	ParamConfigs []*ParamConfig
	echoParams   []string
	Method       string
	Middleware   map[string]interface{}
	Params       map[string]string
	Data         string
	Error        CtxError
	Temp         temp // use for middleware maintain state
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
	r.Temp = make(map[string]interface{})
}

func (r *Res) Init() {
}

func (c *Ctx) Init() {
	c.Middleware = cyako.Middleware.Map
	c.Params = make(map[string]string)
	c.reqParams = make(map[string]interface{})
	c.parseParams()
}

func (c *Ctx) parseParams() {
	s := c.req.Params
	err := json.Unmarshal([]byte(s), &c.reqParams)
	if err != nil {
		c.Error.NewFatal("Params parse error.")
	}
}

/*
	context methods used in processors
*/

func (c *Ctx) getReqParamString(key string) string {
	switch c.reqParams[key].(type) {
	case string:
		return c.reqParams[key].(string)
	case float64:
		return fmt.Sprint(c.reqParams[key].(float64))
	default:
		c.Error.NewWarn(fmt.Sprint("Param type error, not a known type."))
		return fmt.Sprint(c.reqParams[key])
	}
}

func (c *Ctx) Set(data interface{}) {
	var setWitchConfig = func(p *ParamConfig) {
		// add paramConfig to context for docgen etc.
		c.ParamConfigs = append(c.ParamConfigs, p)
		switch {
		case p.Echo:
			c.echoParams = append(c.echoParams, p.Key)
			fallthrough
		case p.Default != "":
			c.Params[p.Key] = p.Default
		case p.Required:
			c.Error.NewFatal("Lack required param.")
		default:
			c.Params[p.Key] = c.getReqParamString(p.Key)
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

func (c *Ctx) Get(key string) string {
	return c.Params[key]
}

/*
	set res
*/
func (c *Ctx) setResParams() {
	var toEscaped = func(s string) string {
		return strings.Replace(s, `"`, `\"`, -1)
	}
	// var params []string
	// var stringMapMarshal = func(m map[string]string) string {
	// 	var kvs []string
	// 	for k, v := range m {
	// 		kvs = append(kvs, `"`+toEscaped(k)+`":"`+toEscaped(v)+`"`)
	// 	}
	// 	return "{" + strings.Join(kvs, ",") + "}"
	// }

	// will be replaced with mature convert solution, here is a temporary process
	var stringMapPartlyMarshal = func(m map[string]string, keys []string) (string, error) {
		var kvs []string
		var err error
		for _, k := range keys {
			if v, ok := m[k]; ok {
				kvs = append(kvs, `"`+toEscaped(k)+`":"`+toEscaped(v)+`"`)
			} else {
				err = errors.New("Cannot find one given key in the map.")
			}
		}
		return "{" + strings.Join(kvs, ",") + "}", err
	}
	json, _ := stringMapPartlyMarshal(c.Params, c.echoParams)
	c.res.Params = json
}
