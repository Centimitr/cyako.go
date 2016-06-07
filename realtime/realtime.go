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

	"golang.org/x/net/websocket"
)

/*
	define
*/
type dep struct {
	KVStore *kvstore.KVStore
}

type Listener struct {
	Conn *websocket.Conn
}

func (l *Listener) Receive(res *cyako.Res) {
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
}

// realtime use the prefix to store data in KVStore
const KVSTORE_SCOPE_LISTENER_GROUPS = "realtime.listnerGroups"

// This method add specific *websocket.Conn to listeners list
func (r *Realtime) AddListener(groupName string, conn *websocket.Conn) {
	kvstore := r.Dependences.KVStore
	listeners := []Listener{}
	if kvstore.HasWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName) {
		listeners = kvstore.GetWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName).([]Listener)
	}
	listeners = append(listeners, Listener{Conn: conn})
	kvstore.SetWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName, listeners)
}

// Send response to listeners in some group
func (r *Realtime) Send(groupName string, res *cyako.Res) {
	kvstore := r.Dependences.KVStore
	listeners := []Listener{}
	if kvstore.HasWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName) {
		listeners = kvstore.GetWithScoped(KVSTORE_SCOPE_LISTENER_GROUPS, groupName).([]Listener)
	}
	for _, listener := range listeners {
		listener.Receive(res)
	}
}

/*
	service hooked methods
*/

/*
	init
*/

func init() {
	r := &Realtime{
		Dependences: dep{
			KVStore: cyako.Svc["KVStore"].(*kvstore.KVStore),
		},
	}
	cyako.LoadService(r)
}
