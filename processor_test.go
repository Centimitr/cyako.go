// Copyright 2016 Cyako Author

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required` by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cyako

import (
	"testing"
)

func TestMethodMatch(t *testing.T) {
	var tests = []struct {
		processorKey string
		reqMethodStr string
		want         bool
	}{
		{"ArticleV1.GetArticleList", "GetArticleList", true},
		{"ArticleV1.GetArticleList", "getArticleList", false},
		{"ArticleV1.GetArticleList", ".GetArticleList", true},
		{"ArticleV1.GetArticleList", "ArticleV1.GetArticleList", true},
		{"ArticleV1.GetArticleList", "Article.GetArticleList", false},
		{"ArticleV1.GetArticleList", "ArticleList", true},
		{"ArticleV1.GetArticleList", "1.GetArticleList", true},
	}
	for _, test := range tests {
		if got := isMethodMatch(test.processorKey, test.reqMethodStr); got != test.want {
			t.Errorf("isMethodMatch(%q, %q) = %v", test.processorKey, test.reqMethodStr, test.want)
		}
	}
}
