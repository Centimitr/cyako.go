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

import ()

func HasInt(slice, item interface{}) bool {
	if slice, ok := slice.([]int); ok {
		if item, ok := item.(int); ok {
			for _, v := range slice {
				if v == item {
					return true
				}
			}
		}
	}
	return false
}

func HasFloat(slice, item interface{}) bool {
	if slice, ok := slice.([]float64); ok {
		if item, ok := item.(float64); ok {
			for _, v := range slice {
				diff := v - item
				if diff > -1e-10 && diff < 1e-10 {
					return true
				}
			}
		}
	}
	return false
}

func HasString(slice, item interface{}) bool {
	if slice, ok := slice.([]string); ok {
		if item, ok := item.(string); ok {
			for _, v := range slice {
				if v == item {
					return true
				}
			}
		}
	}
	return false
}
