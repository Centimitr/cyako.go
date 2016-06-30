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
	ns "github.com/Centimitr/namespace"
	cyako "github.com/Cyako/Cyako.go"
)

type KVStore struct {
	ns.Interface
	ns.Namespace
	Service ns.Prefix
}

func (k *KVStore) Init() {
	k.Interface.Init()
	k.Namespace.Bind(k.Interface)
	_, k.Service = k.Namespace.Prefix("SERVICE")
}

func init() {
	kvstore := &KVStore{
		Interface: &ns.Map{},
	}
	kvstore.Init()
	cyako.LoadService(kvstore)
}
