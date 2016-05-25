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
