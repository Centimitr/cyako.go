// Copyright 2016 Cyako Author

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required` by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cyako

import (
	// "fmt"
	"errors"
	"reflect"
	"strings"
)

/*
	Processor & Processor Module
*/

type Processor struct {
	Module  string
	PkgPath string
	Name    string
	Func    func(*Ctx) []reflect.Value
}

// Client use a part of the table's key (usually is ProcessorName) to match a processor.
// Currently, the key is "PkgName.ProcessorName" so team members should use different package names to do a replacement though their package path is not the same.
func (c *Cyako) matchProcessor(suffix string) (func(*Ctx) []reflect.Value, error) {
	matchedList := []string{}
	// key is "PkgName.ProcessorName"
	for key, _ := range c.ProcessorMap {
		if strings.HasSuffix(key, suffix) {
			matchedList = append(matchedList, key)
		}
	}
	if len(matchedList) == 1 {
		return c.ProcessorMap[matchedList[0]].Func, nil
	}
	return nil, errors.New("Cannot select 1 processor.")
}

// module packages use LoadModule to load itselft
func (c *Cyako) loadModule(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	for i := 0; i < v.NumMethod(); i++ {
		index := i
		c.ProcessorMap[t.Name()+"."+t.Method(i).Name] = &Processor{
			PkgPath: t.PkgPath(),
			Module:  t.Name(),
			Name:    t.Method(i).Name,
			Func: func(ctx *Ctx) []reflect.Value {
				return t.Method(index).Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
			},
		}
	}
}
