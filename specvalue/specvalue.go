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

// Specific values include: string/struct sets or lists.

type SpecValue struct {
	data map[string]interface{}
}

func (s *SpecValue) Get(key string) interface{} {
	return s.data[key]
}

func (s *SpecValue) Set(key string, value interface{}) {
	s.data[key] = value
}

type SpecValueType interface {
	Provide(interface{}) interface{}
	Match(interface{}, interface{}) bool
}

func (s *SpecValue) Provide(key string, t SpecValueType) {
	return t.Provide(s.Get(key))
}
func (s *SpecValue) Match(key string, value interface{}, t SpecValueType) bool {
	return t.Match(s.Get(key), value)
}
