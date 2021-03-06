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

package jsonbase

import (
	cyako "github.com/Cyako/Cyako.go"

	"fmt"
	"io/ioutil"
)

type JSONBase struct{}

func (j *JSONBase) Load(c *cyako.Ctx) {
	filename := c.Method
	data, err := ioutil.ReadFile(filename + ".json")
	if err != nil {
		fmt.Println(err)
	}
	c.Data = string(data)
}

func (j *JSONBase) Save(c *cyako.Ctx) {
	filename := c.Method
	err := ioutil.WriteFile(filename+".json", []byte(c.Data.(string)), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	cyako.LoadService(&JSONBase{})
}
