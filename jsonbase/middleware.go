package jsonbase

import (
	cyako "github.com/Cyako/Cyako.go"
)

/*
	middleware hooked methods
*/

// func (j JSONBase) AfterReceive(req *cyako.Req) {
// }

// func (j JSONBase) BeforeProcess(ctx *cyako.Ctx) {
// }

// func (j JSONBase) AfterProcess(ctx *cyako.Ctx) {
// }

// func (j JSONBase) BeforeSend(res *cyako.Res) {

// }

// func (j JSONBase) AfterSend(res *cyako.Res) {
// }

/*
	init
*/

func init() {
	cyako.LoadMiddleware(JSONBase{})
}
