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

package specvalue

// import (
// 	ns "github.com/Centimitr/namespace"
// 	cyako "github.com/Cyako/Cyako.go"
// 	"github.com/Cyako/Cyako.go/kvstore"
// )

// // Specific values include: string/struct sets or lists.

// type dep struct {
// 	KVStore *kvstore.KVStore
// }

// type SpecValue struct {
// 	Dependences dep
// 	Scope       ns.Scope
// }

// type Value struct {
// 	Value interface{}
// }

// // storage manipulations
// func (s *SpecValue) Set(key string, value interface{}) {
// 	// kvstore := s.Dependences.KVStore
// 	// kvstore.SetWithScoped(KVSTORE_SCOPE_SPECVALUE, key, value)
// 	// kvstore.Service.Extend(...)
// 	s.Scope.Handler(key).Set(value)
// }

// func (s *SpecValue) Get(key string) interface{} {
// 	// kvstore := s.Dependences.KVStore
// 	// return &Value{Value: kvstore.GetWithScoped(KVSTORE_SCOPE_SPECVALUE, key)}
// 	return s.Scope.Handler(key).Get()
// }

// type MatchFunc (func(interface{}, interface{}) bool)

// func (v *Value) MatchFunc(value interface{}, fn MatchFunc) bool {
// 	return fn(v.Value, value)
// }

// // // match methods

// // func (v *Value) HasInt(value interface{}) bool {
// // 	return HasInt(v.Value, value)
// // }

// // func (v *Value) HasFloat(value interface{}) bool {
// // 	return HasFloat(v.Value, value)
// // }

// // func (v *Value) HasString(value interface{}) bool {
// // 	return HasString(v.Value, value)
// // }

// // init
// func init() {
// 	specValue := &SpecValue{
// 		Dependences: dep{
// 			KVStore: cyako.Svc["KVStore"].(*kvstore.KVStore),
// 		},
// 	}
// 	// _, specValue. = specValue.Dependences.KVStore.Service.Extend("SpecValue")
// 	_, specValue.Scope = specValue.Dependences.KVStore.Service.Apply("SpecValue")
// 	cyako.LoadService(specValue)
// }
