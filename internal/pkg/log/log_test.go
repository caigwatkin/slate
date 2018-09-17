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

package log

import (
	"reflect"
	pkg_testing "slate/internal/pkg/testing"
	"testing"
)

func TestNewClient(t *testing.T) {
	var data = []struct {
		desc     string
		input    bool
		expected client
	}{
		{
			desc:  "debug enabled",
			input: true,
			expected: client{
				debug: true,
			},
		},

		{
			desc:  "debug disabled",
			input: false,
			expected: client{
				debug: false,
			},
		},
	}

	for i, d := range data {
		result := NewClient(d.input)

		if reflect.TypeOf(result) != reflect.TypeOf(d.expected) {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result",
				Desc:       d.desc,
				At:         i,
				Expected:   reflect.TypeOf(d.expected),
				Result:     reflect.TypeOf(result),
			}))
		}
		if v, ok := result.(client); !ok {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected,
				Result:     result,
			}))

		} else if v.debug != d.expected.debug {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result.debug",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected.debug,
				Result:     v.debug,
			}))
		}
	}
}
