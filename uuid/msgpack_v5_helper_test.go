//go:build go_tarantool_msgpack_v5
// +build go_tarantool_msgpack_v5

package uuid_test

import (
	"github.com/vmihailenco/msgpack/v5"
)

type decoder = msgpack.Decoder

func marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
