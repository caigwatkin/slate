package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	go_testing "github.com/caigwatkin/go/testing"
	"github.com/caigwatkin/slate/internal/lib/dto"
)

func TestParseCreateGreeting(t *testing.T) {
	expectedResultCreateGreeting := dto.CreateGreeting{
		UserInput: dto.CreateGreetingUserInput{
			Message: "message",
		},
	}

	body, err := json.Marshal(expectedResultCreateGreeting.UserInput)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("", "", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req = req.WithContext(context.Background())

	type expected struct {
		result dto.CreateGreeting
	}
	var data = []struct {
		desc string
		client
		input *http.Request
		expected
	}{
		{
			desc:   "success",
			client: clientSuccess,
			input:  req,
			expected: expected{
				result: expectedResultCreateGreeting,
			},
		},
		// {
		// 	desc:   "schema error",
		// 	client: clientErrSchema,
		// 	input:  req,
		// 	expected: expected{
		// 		result: nil,
		// 		err:    &iam_mock.ExpectedErrSchemaClient,
		// 	},
		// },
	}

	for i, d := range data {
		result := d.client.ParseCreateGreeting(d.input)

		// if d.expected.err != nil {
		// 	if !reflect.DeepEqual(err, error(*d.expected.err)) {
		// 		var r interface{} = err
		// 		if err != nil {
		// 			r = err.Error()
		// 		}
		// 		t.Error(lib_testing.Errorf(lib_testing.Error{
		// 			Unexpected: "err not equal",
		// 			Desc:       d.desc,
		// 			At:         i,
		// 			Expected:   d.expected.err.Error(),
		// 			Result:     r,
		// 		}))
		// 	}
		// } else if err != nil {
		// 	t.Error(lib_testing.Errorf(lib_testing.Error{
		// 		Unexpected: "err exists",
		// 		Desc:       d.desc,
		// 		At:         i,
		// 		Expected:   nil,
		// 		Result:     err.Error(),
		// 	}))

		// } else {
		if !reflect.DeepEqual(result, d.expected.result) {
			t.Error(go_testing.Errorf(go_testing.Error{
				Unexpected: "result",
				Desc:       d.desc,
				At:         i,
				Expected:   d.expected.result,
				Result:     result,
			}))
		}
		// }
	}
}

// func TestParseSearchGreetingsOfIdentity(t *testing.T) {
// 	req, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	expectedResultGreetingsOfResourceSearch := dto.GreetingsOfResourceSearch{
// 		ResourceID: "xxxx-xxxx-xxxx-xxxx",
// 	}
// 	chiCtx := chi.NewRouteContext()
// 	chiCtx.URLParams.Add("identity_id", expectedResultGreetingsOfResourceSearch.ResourceID)
// 	req = req.WithContext(context.WithValue(lib_context.NewContext(context.Background()), chi.RouteCtxKey, chiCtx))

// 	reqMissingIdentityID, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	reqMissingIdentityID = reqMissingIdentityID.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, chi.NewRouteContext()))
// 	errMissingIdentityID := missingURLParamErr("identity_id")

// 	type expected struct {
// 		result     *dto.GreetingsOfResourceSearch
// 		pagination *lib_pagination.Pagination
// 		err        *lib_errors.Custom
// 	}
// 	var data = []struct {
// 		desc string
// 		client
// 		input *http.Request
// 		expected
// 	}{
// 		{
// 			desc:   "success",
// 			client: clientSuccess,
// 			input:  req,
// 			expected: expected{
// 				result:     &expectedResultGreetingsOfResourceSearch,
// 				pagination: lib_pagination.Default(),
// 				err:        nil,
// 			},
// 		},
// 		{
// 			desc:   "missing identity id",
// 			client: clientSuccess,
// 			input:  reqMissingIdentityID,
// 			expected: expected{
// 				result:     nil,
// 				pagination: nil,
// 				err:        &errMissingIdentityID,
// 			},
// 		},
// 	}

// 	for i, d := range data {
// 		result, pagination, err := d.client.ParseSearchGreetingsOfIdentity(d.input)

// 		if d.expected.err != nil {
// 			if !reflect.DeepEqual(err, error(*d.expected.err)) {
// 				var r interface{} = err
// 				if err != nil {
// 					r = err.Error()
// 				}
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "err not equal",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.err.Error(),
// 					Result:     r,
// 				}))
// 			}
// 		} else if err != nil {
// 			t.Error(lib_testing.Errorf(lib_testing.Error{
// 				Unexpected: "err exists",
// 				Desc:       d.desc,
// 				At:         i,
// 				Expected:   nil,
// 				Result:     err.Error(),
// 			}))

// 		} else {
// 			if !reflect.DeepEqual(result, d.expected.result) {
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "result",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.result,
// 					Result:     result,
// 				}))
// 			}
// 			if !reflect.DeepEqual(pagination, d.expected.pagination) {
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "pagination",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.pagination,
// 					Result:     pagination,
// 				}))
// 			}
// 		}
// 	}
// }

// func TestParseReadGreeting(t *testing.T) {
// 	req, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	expectedResultGreetingRead := dto.GreetingRead{
// 		ID: "xxxx-xxxx-xxxx-xxxx",
// 	}
// 	chiCtx := chi.NewRouteContext()
// 	chiCtx.URLParams.Add("group_id", expectedResultGreetingRead.ID)
// 	req = req.WithContext(context.WithValue(lib_context.NewContext(context.Background()), chi.RouteCtxKey, chiCtx))

// 	reqMissingGreetingID, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	reqMissingGreetingID = reqMissingGreetingID.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, chi.NewRouteContext()))
// 	errMissingGreetingID := missingURLParamErr("group_id")

// 	type expected struct {
// 		result *dto.GreetingRead
// 		err    *lib_errors.Custom
// 	}
// 	var data = []struct {
// 		desc string
// 		client
// 		input *http.Request
// 		expected
// 	}{
// 		{
// 			desc:   "success",
// 			client: clientSuccess,
// 			input:  req,
// 			expected: expected{
// 				result: &expectedResultGreetingRead,
// 				err:    nil,
// 			},
// 		},
// 		{
// 			desc:   "missing group ID",
// 			client: clientSuccess,
// 			input:  reqMissingGreetingID,
// 			expected: expected{
// 				result: nil,
// 				err:    &errMissingGreetingID,
// 			},
// 		},
// 	}

// 	for i, d := range data {
// 		result, err := d.client.ParseReadGreeting(d.input)

// 		if d.expected.err != nil {
// 			if !reflect.DeepEqual(err, error(*d.expected.err)) {
// 				var r interface{} = err
// 				if err != nil {
// 					r = err.Error()
// 				}
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "err not equal",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.err.Error(),
// 					Result:     r,
// 				}))
// 			}
// 		} else if err != nil {
// 			t.Error(lib_testing.Errorf(lib_testing.Error{
// 				Unexpected: "err exists",
// 				Desc:       d.desc,
// 				At:         i,
// 				Expected:   nil,
// 				Result:     err.Error(),
// 			}))

// 		} else {
// 			if !reflect.DeepEqual(result, d.expected.result) {
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "result",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.result,
// 					Result:     result,
// 				}))
// 			}
// 		}
// 	}
// }

// //revive:disable:cyclomatic
// func TestParseUpdateGreeting(t *testing.T) {
// 	expectedResultGreetingUpdate := dto.GreetingUpdate{
// 		Body: dto.GreetingUpdateBody{
// 			RoleIDs: []string{"group_id"},
// 		},
// 		ID:       "xxxx-xxxx-xxxx-xxxx",
// 		UserName: "user_name",
// 	}
// 	chiCtx := chi.NewRouteContext()
// 	chiCtx.URLParams.Add("group_id", expectedResultGreetingUpdate.ID)
// 	ctx := lib_context.WithName(lib_context.NewContext(context.Background()), expectedResultGreetingUpdate.UserName)

// 	body, err := json.Marshal(expectedResultGreetingUpdate.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, err := http.NewRequest("", "", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, chiCtx))

// 	reqMissingGreetingID, err := http.NewRequest("", "", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	reqMissingGreetingID = reqMissingGreetingID.WithContext(context.WithValue(ctx, chi.RouteCtxKey, chi.NewRouteContext()))
// 	errMissingGreetingID := missingURLParamErr("group_id")

// 	type expected struct {
// 		result *dto.GreetingUpdate
// 		err    *lib_errors.Custom
// 	}
// 	var data = []struct {
// 		desc string
// 		client
// 		input *http.Request
// 		expected
// 	}{
// 		{
// 			desc:   "success",
// 			client: clientSuccess,
// 			input:  req,
// 			expected: expected{
// 				result: &expectedResultGreetingUpdate,
// 				err:    nil,
// 			},
// 		},
// 		{
// 			desc:   "missing group ID",
// 			client: clientSuccess,
// 			input:  reqMissingGreetingID,
// 			expected: expected{
// 				result: nil,
// 				err:    &errMissingGreetingID,
// 			},
// 		},
// 		{
// 			desc:   "schema error",
// 			client: clientErrSchema,
// 			input:  req,
// 			expected: expected{
// 				result: nil,
// 				err:    &iam_mock.ExpectedErrSchemaClient,
// 			},
// 		},
// 	}

// 	for i, d := range data {
// 		result, err := d.client.ParseUpdateGreeting(d.input)

// 		if d.expected.err != nil {
// 			if !reflect.DeepEqual(err, error(*d.expected.err)) {
// 				var r interface{} = err
// 				if err != nil {
// 					r = err.Error()
// 				}
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "err not equal",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.err.Error(),
// 					Result:     r,
// 				}))
// 			}
// 		} else if err != nil {
// 			t.Error(lib_testing.Errorf(lib_testing.Error{
// 				Unexpected: "err exists",
// 				Desc:       d.desc,
// 				At:         i,
// 				Expected:   nil,
// 				Result:     err.Error(),
// 			}))

// 		} else {
// 			if !reflect.DeepEqual(result, d.expected.result) {
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "result",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.result,
// 					Result:     result,
// 				}))
// 			}
// 		}
// 	}
// 	//revive:enable:cyclomatic
// }

// func TestParseDeleteGreeting(t *testing.T) {
// 	req, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	expectedResultGreetingDelete := dto.GreetingDelete{
// 		ID: "xxxx-xxxx-xxxx-xxxx",
// 	}
// 	chiCtx := chi.NewRouteContext()
// 	chiCtx.URLParams.Add("group_id", expectedResultGreetingDelete.ID)
// 	req = req.WithContext(context.WithValue(lib_context.NewContext(context.Background()), chi.RouteCtxKey, chiCtx))

// 	reqMissingGreetingID, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	reqMissingGreetingID = reqMissingGreetingID.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, chi.NewRouteContext()))
// 	errMissingGreetingID := missingURLParamErr("group_id")

// 	type expected struct {
// 		result *dto.GreetingDelete
// 		err    *lib_errors.Custom
// 	}
// 	var data = []struct {
// 		desc string
// 		client
// 		input *http.Request
// 		expected
// 	}{
// 		{
// 			desc:   "success",
// 			client: clientSuccess,
// 			input:  req,
// 			expected: expected{
// 				result: &expectedResultGreetingDelete,
// 				err:    nil,
// 			},
// 		},
// 		{
// 			desc:   "missing group ID",
// 			client: clientSuccess,
// 			input:  reqMissingGreetingID,
// 			expected: expected{
// 				result: nil,
// 				err:    &errMissingGreetingID,
// 			},
// 		},
// 	}

// 	for i, d := range data {
// 		result, err := d.client.ParseDeleteGreeting(d.input)

// 		if d.expected.err != nil {
// 			if !reflect.DeepEqual(err, error(*d.expected.err)) {
// 				var r interface{} = err
// 				if err != nil {
// 					r = err.Error()
// 				}
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "err not equal",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.err.Error(),
// 					Result:     r,
// 				}))
// 			}
// 		} else if err != nil {
// 			t.Error(lib_testing.Errorf(lib_testing.Error{
// 				Unexpected: "err exists",
// 				Desc:       d.desc,
// 				At:         i,
// 				Expected:   nil,
// 				Result:     err.Error(),
// 			}))

// 		} else {
// 			if !reflect.DeepEqual(result, d.expected.result) {
// 				t.Error(lib_testing.Errorf(lib_testing.Error{
// 					Unexpected: "result",
// 					Desc:       d.desc,
// 					At:         i,
// 					Expected:   d.expected.result,
// 					Result:     result,
// 				}))
// 			}
// 		}
// 	}
// }
