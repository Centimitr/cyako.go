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
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func (c *Cyako) Run(addr, pattern string) error {
	c.Handle(pattern)
	return http.ListenAndServe(addr, nil)
}

// return a http.Handler
func (c *Cyako) Handle(pattern string) {
	http.Handle(pattern, websocket.Handler(c.Server))
}

func (c *Cyako) Server(ws *websocket.Conn) {
	var err error
	for {
		var req Req
		req.Init()
		if err = websocket.JSON.Receive(ws, &req); err != nil {
			fmt.Println("RECEIVE ERR:", err)
			break
		}
		go c.handle(ws, &req)
	}
}

func (c *Cyako) handle(ws *websocket.Conn, req *Req) {
	var err error

	// Phase I: AfterReceive
	// - global, req relative methods
	c.AfterReceive(req)
	// - initial context and response
	res := &Res{Id: req.Id, Method: req.Method, Temp: req.Temp}
	ctx := &Ctx{res: res, req: req, Method: req.Method, Data: req.Data, Temp: req.Temp, Conn: ws}
	res.Init()
	ctx.Init()

	// Phase II: BeforeProcess
	// - global, ctx relative methods
	// - provide chance to manipulate context object for services
	c.BeforeProcess(ctx)

	// Phase III: Process
	// - match and select processor, and then execute it on ctx
	process, err := c.matchProcessor(req.Method)
	if err != nil {
		fmt.Println(err)
		return
	}
	process(ctx)

	// Phase IV: AfterProcess
	// - global, ctx relative methods
	c.AfterProcess(ctx)
	// - mainly handle response relative tasks
	res.Data = ctx.Data
	// data, err := json.Marshal(ctx.Data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// res.Data = string(data)
	ctx.setResParams()

	// Phase V: BeforeSend
	c.BeforeSend(res)

	// Phase VI: Send
	// - send
	if err := websocket.JSON.Send(ws, res); err != nil {
		// fmt.Println("SEND ERR:", err)
		return
	}

	// Phase VI: AfterSend
	c.AfterSend(res)
}
