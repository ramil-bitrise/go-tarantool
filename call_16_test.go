//go:build !go_tarantool_call_17
// +build !go_tarantool_call_17

package tarantool_test

import (
	"testing"

	. "github.com/tarantool/go-tarantool/v2"
	"github.com/tarantool/go-tarantool/v2/test_helpers"
)

func TestConnection_Call(t *testing.T) {
	var resp *Response
	var err error

	conn := test_helpers.ConnectWithValidation(t, server, opts)
	defer conn.Close()

	// Call16
	resp, err = conn.Call("simple_concat", []interface{}{"1"})
	if err != nil {
		t.Errorf("Failed to use Call")
	}
	if val, ok := resp.Data[0].([]interface{})[0].(string); !ok || val != "11" {
		t.Errorf("result is not {{1}} : %v", resp.Data)
	}
}

func TestCallRequest(t *testing.T) {
	var resp *Response
	var err error

	conn := test_helpers.ConnectWithValidation(t, server, opts)
	defer conn.Close()

	req := NewCallRequest("simple_concat").Args([]interface{}{"1"})
	resp, err = conn.Do(req).Get()
	if err != nil {
		t.Errorf("Failed to use Call")
	}
	if val, ok := resp.Data[0].([]interface{})[0].(string); !ok || val != "11" {
		t.Errorf("result is not {{1}} : %v", resp.Data)
	}
}

func TestCallRequestCode(t *testing.T) {
	req := NewCallRequest("simple_concat")
	code := req.Code()
	expected := Call16RequestCode
	if code != int32(expected) {
		t.Errorf("CallRequest actual code %v != %v", code, expected)
	}
}
