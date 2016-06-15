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
	cyako "github.com/Cyako/Cyako.go"

	"fmt"
	"sync"
	"time"
)

type StatisticsItem struct {
	RequestTimes  int
	ResponseTimes int
	TotalDuration time.Duration
	MinDuration   time.Duration
	MaxDuration   time.Duration
	LastDuration  time.Duration
}

type StatisticsMap struct {
	mutex     sync.RWMutex
	methodMap map[string]*StatisticsItem
}

/*
	private methods
*/

func (s *StatisticsMap) Get() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	fmt.Printf("\n %-35s %-10s %-10s %-10s %-10s %-10s\n", "API", "Res/Req", "Avg", "Min", "Max", "Last")
	for k, item := range s.methodMap {
		fmt.Printf(" %-35s %-10v %-10.4f %-10.4f %-10.4f %-10.4f\n", k, fmt.Sprintf("%v/%v", item.ResponseTimes, item.RequestTimes),
			item.TotalDuration.Seconds()/float64(item.ResponseTimes)*1000,
			item.MinDuration.Seconds()*1000,
			item.MaxDuration.Seconds()*1000,
			item.LastDuration.Seconds()*1000,
		)
	}
}

func (s *StatisticsMap) recordReq(method string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.methodMap[method]; ok {
		s.methodMap[method].RequestTimes++
	} else {
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  1,
			ResponseTimes: 0,
			TotalDuration: 0,
			MaxDuration:   0,
			MinDuration:   time.Second,
			LastDuration:  0,
		}
	}
}

func (s *StatisticsMap) recordResAndStat(method string, duration time.Duration) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// record response
	if m, ok := s.methodMap[method]; ok {
		m.ResponseTimes++
		switch {
		case true:
			m.TotalDuration += duration
			m.LastDuration = duration
			fallthrough
		case m.MaxDuration < duration:
			m.MaxDuration = duration
			fallthrough
		case m.MinDuration > duration:
			m.MinDuration = duration
			// fallthrough
		}
	} else {
		// may not have a statistics item when in realtime send situation
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  0,
			ResponseTimes: 1,
			TotalDuration: duration,
			MinDuration:   duration,
			MaxDuration:   duration,
			LastDuration:  duration,
		}
	}
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
