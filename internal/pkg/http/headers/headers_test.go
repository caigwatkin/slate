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
	pkg_testing "slate/internal/pkg/testing"
	"testing"
)

func TestSetKeyXCorrelationID(t *testing.T) {

	var data = []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "not empty service name",
			input:    "ServiceName",
			expected: "X-Service-Name-Correlation-Id",
		},

		{
			desc:     "empty service name",
			input:    "",
			expected: "X-Correlation-Id",
		},
	}

	for i, d := range data {
		SetKeyXCorrelationID(d.input)
		result := KeyXCorrelationID

		if result != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result",
				Desc:       d.desc,
				At:         i,
				Input:      d.input,
				Expected:   d.expected,
				Result:     result,
			}))
		}
	}
}

func TestSetKeyXTest(t *testing.T) {

	var data = []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "not empty service name",
			input:    "ServiceName",
			expected: "X-Service-Name-Test",
		},

		{
			desc:     "empty service name",
			input:    "",
			expected: "X-Test",
		},
	}

	for i, d := range data {
		SetKeyXTest(d.input)
		result := KeyXTest

		if result != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result",
				Desc:       d.desc,
				At:         i,
				Input:      d.input,
				Expected:   d.expected,
				Result:     result,
			}))
		}
	}
}

func TestSetValXTest(t *testing.T) {

	var data = []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "not empty val",
			input:    "SomeValue-xXxXx",
			expected: "SomeValue-xXxXx",
		},

		{
			desc:     "empty val",
			input:    "",
			expected: "",
		},
	}

	for i, d := range data {
		SetValXTest(d.input)
		result := ValXTest

		if result != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result",
				Desc:       d.desc,
				At:         i,
				Input:      d.input,
				Expected:   d.expected,
				Result:     result,
			}))
		}
	}
}

func TestCamelToCanonical(t *testing.T) {

	var data = []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "one character",
			input:    "S",
			expected: "S",
		},

		{
			desc:     "camel case one word",
			input:    "Service",
			expected: "Service",
		},

		{
			desc:     "camel case two words",
			input:    "ServiceName",
			expected: "Service-Name",
		},

		{
			desc:     "camel case acronym one word",
			input:    "API",
			expected: "Api",
		},

		{
			desc:     "camel case acronym two words",
			input:    "APIName",
			expected: "Api-Name",
		},

		{
			desc:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for i, d := range data {
		result := camelToCanonical(d.input)

		if result != d.expected {
			t.Error(pkg_testing.Errorf(pkg_testing.Error{
				Unexpected: "result",
				Desc:       d.desc,
				At:         i,
				Input:      d.input,
				Expected:   d.expected,
				Result:     result,
			}))
		}
	}
}
