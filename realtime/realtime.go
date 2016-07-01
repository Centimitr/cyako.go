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

package realtime

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/kvstore"

	"fmt"
	ns "github.com/Centimitr/namespace"
	"golang.org/x/net/websocket"
)

/*
	define
*/
type dep struct {
	KVStore *kvstore.KVStore
}

type Listener struct {
	Conn   *websocket.Conn
	Id     string
	Method string
}

func (l *Listener) Receive(res *cyako.Res) {
	fmt.Println("Receive:", l.Conn, res)
	if l.Conn == nil {
		return
	}
	if err := websocket.JSON.Send(l.Conn, res); err != nil {
		// fmt.Println("SEND ERR:", err)
		return
	}
}

type Realtime struct {
	Dependences dep
	Scope       ns.Scope
}

// realtime use the prefix to store data in KVStore
// const KVSTORE_SCOPE_LISTENER_GROUPS = "service.realtime.listnerGroups"

// This method add specific *websocket.Conn to listeners list
func (r *Realtime) AddListener(groupName string, conn *websocket.Conn, id string, method string) {
	// kvstore := r.Dependences.KVStore
	listeners := []Listener{}
	// if kvstore.HasWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName) {
	// 	listeners = kvstore.GetWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName).([]Listener)
	// }
	listeners = append(listeners, Listener{Conn: conn, Id: id})
	// kvstore.SetWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName, listeners)
	if r.Scope.Handler(groupName).Has() {
		r.Scope.Handler(groupName).Set(listeners)
	}
}

func (r *Realtime) AddListenerDefault(groupName string, ctx *cyako.Ctx) {
	r.AddListener(groupName, ctx.Conn, ctx.Id, ctx.Method)
}

// Send response to listeners in some group
func (r *Realtime) Send(groupName string, res *cyako.Res) {
	fmt.Println("Start Sending.")
	// kvstore := r.Dependences.KVStore
	listeners := []Listener{}
	// if kvstore.HasWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName) {
	// 	listeners = kvstore.GetWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName).([]Listener)
	// }
	if r.Scope.Handler(groupName).Has() {
		listeners = r.Scope.Handler(groupName).Get().([]Listener)
	}
	fmt.Println("listners:", listeners)
	for _, listener := range listeners {
		res.Id = listener.Id
		res.Method = listener.Method
		listener.Receive(res)
	}
}

/*
	init
*/

func init() {
	r := &Realtime{
		Dependences: dep{
			KVStore: cyako.Svc["KVStore"].(*kvstore.KVStore),
		},
	}
	_, r.Scope = r.Dependences.KVStore.Service.Apply("REALTIME")
	cyako.LoadService(r)
	fmt.Println("LOAD REALTIME.", r)
}
