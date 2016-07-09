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
)

func (c *Cyako) PrintLoadInfo() {
	fmt.Println()
	fmt.Println(" Loading...")

	fmt.Printf("\n %-35s %-21s %-10s %-10s\n", "Config", "Name", "Key", "Value")
	for _, config := range c.Config.Service {
		fmt.Printf(" %-35s %-21s %-10s %-10s\n", "Service", config.Name, config.Key, config.Value)
	}

	fmt.Printf("\n %-35s %-10s %-10s %-10s %-10s %-10s\n", "Service", "AR", "BP", "AP", "BS", "AS")
	for _, c := range c.Service.Support {
		fmt.Printf(" %-35s %-10v %-10v %-10v %-10v %-10v\n", c.Name, c.AfterReceive, c.BeforeProcess, c.AfterProcess, c.BeforeSend, c.AfterSend)
	}

	fmt.Printf("\n %-35s %-50s\n", "API", "Package Path")
	for _, proc := range c.ProcessorMap {
		fmt.Printf(" %-35s %-50s\n", proc.Module+"."+proc.Name, proc.PkgPath)
	}
}

func (c *Cyako) CheckModule() {
	fmt.Println()
	fmt.Println(" Checking...")
	for _, proc := range c.ProcessorMap {
		req := &Req{}
		req.Init()
		res := &Res{Id: req.Id, Method: req.Method, Temp: req.Temp}
		ctx := &Ctx{res: res, req: req, Method: req.Method, Data: req.Data, Temp: req.Temp}
		res.Init()
		ctx.Init()
		proc.Func(ctx)
	}
}

func (c *Cyako) PrintAPIDoc() {
	fmt.Println()
	type method struct {
		ParamConfigs []*ParamConfig `json:"ParamConfigs"`
		Processor
	}
	type APIDoc struct {
		Methods map[string]method `json:"method"`
	}
	doc := &APIDoc{
		Methods: make(map[string]method),
	}
	for methodName, proc := range c.ProcessorMap {
		req := &Req{}
		req.Init()
		res := &Res{Id: req.Id, Method: req.Method, Temp: req.Temp}
		ctx := &Ctx{res: res, req: req, Method: req.Method, Data: req.Data, Temp: req.Temp}
		res.Init()
		ctx.Init()
		proc.Func(ctx)
		doc.Methods[methodName] = method{
			ParamConfigs: ctx.ParamConfigs,
			Processor:    *proc,
		}
	}
	// bytes, err := json.Marshal(doc)
	// if err != nil {
	// 	fmt.Println(" Error APIDoc:", err)
	// }
	// fmt.Println(string(bytes))
	fmt.Println()
	fmt.Printf("\n %-35s %-10s %-40s\n", "API Detail", "Module", "Package Path")
	for _, proc := range doc.Methods {
		fmt.Printf(" %-35s %-10s %-40s\n", proc.Module+"."+proc.Name, proc.Module, proc.PkgPath)
		for _, cfg := range proc.ParamConfigs {
			fmt.Printf(" -%-10s %+v\n", cfg.Key, *cfg)
		}
	}
	fmt.Println()
}
