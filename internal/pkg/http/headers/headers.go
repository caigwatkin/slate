/*
Copyright 2018 Cai Gwatkin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package headers

import (
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
)

var (
	KeyXCorrelationID = "X-Correlation-Id"
	KeyXTest          = "X-Test"

	ValXTest = "5c1bca85-9e09-4af4-96ac-7f353265838c" // This can be stored and used unencrypted, as anything running in test mode should be safe enough that it doesn't matter who knows it
)

func SetKeyXCorrelationID(serviceNameInCamelCase string) {
	if serviceNameInCamelCase != "" {
		KeyXCorrelationID = fmt.Sprintf("X-%s-Correlation-Id", camelToCanonical(serviceNameInCamelCase))
		return
	}
	KeyXCorrelationID = "X-Correlation-Id"
}

func SetKeyXTest(serviceNameInCamelCase string) {
	if serviceNameInCamelCase != "" {
		KeyXTest = fmt.Sprintf("X-%s-Test", camelToCanonical(serviceNameInCamelCase))
		return
	}
	KeyXTest = "X-Test"
}

func SetValXTest(val string) {
	ValXTest = val
}

func camelToCanonical(s string) string {
	words := camelcase.Split(s)
	for i, v := range words {
		lv := strings.ToLower(v)
		words[i] = strings.ToUpper(lv[:1]) + lv[1:]
	}
	return strings.Join(words, "-")
}
