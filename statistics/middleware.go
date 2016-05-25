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

package statistics

import (
	// "fmt"
	cyako "github.com/Cyako/Cyako.go"
	"time"
)

/*
	middleware hooked methods
*/

func (s StatisticsMap) AfterReceive(req *cyako.Req) {
	req.Temp.Put("Stat", "start", time.Now())
	s.recordReq(req.Method)
}

func (s StatisticsMap) BeforeProcess(ctx *cyako.Ctx) {
}

func (s StatisticsMap) AfterProcess(ctx *cyako.Ctx) {
}

func (s StatisticsMap) BeforeSend(res *cyako.Res) {

}

func (s StatisticsMap) AfterSend(res *cyako.Res) {
	t := res.Temp.Get("Stat", "start").(time.Time)
	duration := time.Now().Sub(t)
	s.recordResAndStat(res.Method, duration)
}

/*
	init
*/

// use struct Statistics to combime the middleware, so Statistics is the middleware name.
type Statistics struct {
	StatisticsMap
}

func init() {
	cyako.LoadMiddleware(Statistics{
		StatisticsMap{
			methodMap: make(map[string]*StatisticsItem),
		},
	})
}
