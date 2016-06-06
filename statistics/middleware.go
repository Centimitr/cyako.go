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
	service hooked methods
*/

const (
	TEMP_SCOPE           = "Statistics"
	TEMP_KEY_TIME_RECORD = "RequestReceivedTime"
)

func (s StatisticsMap) AfterReceive(req *cyako.Req) {
	req.Temp.Put(TEMP_SCOPE, TEMP_KEY_TIME_RECORD, time.Now())
	s.recordReq(req.Method)
}

func (s StatisticsMap) BeforeProcess(ctx *cyako.Ctx) {
}

func (s StatisticsMap) AfterProcess(ctx *cyako.Ctx) {
}

func (s StatisticsMap) BeforeSend(res *cyako.Res) {

}

func (s StatisticsMap) AfterSend(res *cyako.Res) {
	t := res.Temp.Get(TEMP_SCOPE, TEMP_KEY_TIME_RECORD).(time.Time)
	duration := time.Now().Sub(t)
	s.recordResAndStat(res.Method, duration)
}

/*
	init
*/

// use struct Statistics to combime the service, so Statistics is the service name.
type Statistics struct {
	StatisticsMap
}

func init() {
	cyako.LoadService(Statistics{
		StatisticsMap{
			methodMap: make(map[string]*StatisticsItem),
		},
	})
}
