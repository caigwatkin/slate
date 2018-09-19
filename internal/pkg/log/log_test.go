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

func TestFmtAny(t *testing.T) {
	type input struct {
		value interface{}
		name  string
	}
	type valueStruct struct {
		X string `json:"x,omitempty"`
		Y int    `json:"y,omitempty"`
		Z bool   `json:"z,omitempty"`
	}
	var data = []struct {
		desc string
		input
		expected string
	}{
		{
			desc: "struct as json indented",
			input: input{
				value: valueStruct{
					X: "some string",
					Y: 2,
					Z: true,
				},
				name: "name",
			},
			expected: "\"name\": {\n\t\t\"x\": \"some string\",\n\t\t\"y\": 2,\n\t\t\"z\": true\n\t}",
		},

		{
			desc: "struct omitempty",
			input: input{
				value: valueStruct{
					X: "",
					Y: 0,
					Z: false,
				},
				name: "name",
			},
			expected: "\"name\": {}",
		},
	}

	for i, d := range data {
		result := FmtAny(d.input.value, d.input.name)

		if result.s != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result.s",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected,
				Result:     result.s,
			}))
		}
	}
}

func TestFmtBool(t *testing.T) {
	type input struct {
		value bool
		name  string
	}
	var data = []struct {
		desc string
		input
		expected string
	}{
		{
			desc: "false",
			input: input{
				value: false,
				name:  "name",
			},
			expected: "\"name\": false",
		},

		{
			desc: "true",
			input: input{
				value: true,
				name:  "name",
			},
			expected: "\"name\": true",
		},
	}

	for i, d := range data {
		result := FmtBool(d.input.value, d.input.name)

		if result.s != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result.s",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected,
				Result:     result.s,
			}))
		}
	}
}

func TestFmtBools(t *testing.T) {
	type input struct {
		value []bool
		name  string
	}
	var data = []struct {
		desc string
		input
		expected string
	}{
		{
			desc: "bool",
			input: input{
				value: []bool{false},
				name:  "name",
			},
			expected: "\"name\": [\n\t\tfalse\n\t]",
		},

		{
			desc: "bools",
			input: input{
				value: []bool{false, true},
				name:  "name",
			},
			expected: "\"name\": [\n\t\tfalse,\n\t\ttrue\n\t]",
		},

		{
			desc: "empty",
			input: input{
				value: []bool{},
				name:  "name",
			},
			expected: "\"name\": []",
		},
	}

	for i, d := range data {
		result := FmtBools(d.input.value, d.input.name)

		if result.s != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result.s",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected,
				Result:     result.s,
			}))
		}
	}
}

func TestFmtByte(t *testing.T) {
	type input struct {
		value byte
		name  string
	}
	var data = []struct {
		desc string
		input
		expected string
	}{
		{
			desc: "Character",
			input: input{
				value: 'A',
				name:  "name",
			},
			expected: "\"name\": 'A'",
		},

		{
			desc: "integer",
			input: input{
				value: 0,
				name:  "name",
			},
			expected: "\"name\": '\\x00'",
		},
	}

	for i, d := range data {
		result := FmtByte(d.input.value, d.input.name)

		if result.s != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result.s",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected,
				Result:     result.s,
			}))
		}
	}
}

func TestFmtBytes(t *testing.T) {
	type input struct {
		value []byte
		name  string
	}
	var data = []struct {
		desc string
		input
		expected string
	}{
		{
			desc: "empty",
			input: input{
				value: []byte{},
				name:  "name",
			},
			expected: "\"name\": \"\"",
		},

		{
			desc: "string",
			input: input{
				value: []byte("some text"),
				name:  "name",
			},
			expected: "\"name\": \"some text\"",
		},

		{
			desc: "characters",
			input: input{
				value: []byte{'A', 'a', 'B'},
				name:  "name",
			},
			expected: "\"name\": \"AaB\"",
		},

		{
			desc: "characters",
			input: input{
				value: []byte{0, 25, 18},
				name:  "name",
			},
			expected: "\"name\": \"\\x00\\x19\\x12\"",
		},
	}

	for i, d := range data {
		result := FmtBytes(d.input.value, d.input.name)

		if result.s != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result.s",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected,
				Result:     result.s,
			}))
		}
	}
}
