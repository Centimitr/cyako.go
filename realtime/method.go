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
	"golang.org/x/net/websocket"
)

func (r *RealtimeManager) RegisterListener() {
	// r.addListener(listName, ls)
}

func (r *RealtimeManager) Send(listName string, res *cyako.Res) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, listener := range r.List[listName] {
		if err := websocket.JSON.Send(listener.ws, res); err != nil {
			// fmt.Println("SEND ERR:", err)
			return
		}
	}
}
