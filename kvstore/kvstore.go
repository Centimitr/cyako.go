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

package kvstore

import (
	// "fmt"
	cyako "github.com/Cyako/Cyako.go"
)

type Interfacce interface {
	Get(string) interface{}
	Set(string, interface{})
	Delete(string)
	Disactive()
	Active()
}

type KVStore struct {
	Interfacce
}

func GetScopedKeyString(scope, name string) string {
	return scope + "." + name
}

func (k *KVStore) GetWithScoped(scope, name string) interface{} {
	return k.Get(GetScopedKeyString(scope, name))
}

func (k *KVStore) SetWithScoped(scope, name string, value interface{}) {
	k.Set(GetScopedKeyString(scope, name), value)
}

func init() {
	kvstore := &KVStore{
		MemoryKVStore{
			m: make(map[string]interface{}),
		},
	}
	cyako.LoadMiddleware(kvstore)
}
