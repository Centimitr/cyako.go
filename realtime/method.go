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
