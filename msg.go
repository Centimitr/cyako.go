package cyako

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
)

type middlewareConfig struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type config struct {
	Middleware []middlewareConfig `json:"middleware"`
}

type middlewareSupport struct {
	Name          string
	AfterReceive  bool
	BeforeProcess bool
	AfterProcess  bool
	BeforeSend    bool
	AfterSend     bool
}

type middleware struct {
	Map               map[string]interface{}
	Support           []middlewareSupport
	AfterReceiveFunc  []func(*Req)
	BeforeProcessFunc []func(*Ctx)
	AfterProcessFunc  []func(*Ctx)
	BeforeSendFunc    []func(*Res)
	AfterSendFunc     []func(*Res)
}

type Cyako struct {
	Config       config
	Middleware   middleware
	ProcessorMap map[string]*Processor
}

// return a http.Handler
func (c *Cyako) Server(ws *websocket.Conn) {
	var err error
	for {
		var req Req
		req.Init()
		if err = websocket.JSON.Receive(ws, &req); err != nil {
			break
		}
		go c.do(ws, &req)
	}
}

func (c *Cyako) loadConfig() {
	var err error
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Read config file error:", err)
	}
	err = json.Unmarshal(data, &c.Config)
	if err != nil {
		fmt.Println("Unmarshal config file content error:", err)
	}
}

/*
	init
*/

var cyako *Cyako

func init() {
	cyako = &Cyako{
		Middleware: middleware{
			Map: make(map[string]interface{}),
		},
		ProcessorMap: make(map[string]*Processor),
	}
	cyako.loadConfig()
}

/*
	global
*/

// return cyako package's global object: cyako
func Ins() *Cyako {
	return cyako
}

// used in Processor Module package to load itself
func LoadModule(x interface{}) {
	cyako.loadModule(x)
}

// used in Middleware package to load itself
func LoadMiddleware(x interface{}) {
	cyako.loadMiddleware(x)
}
