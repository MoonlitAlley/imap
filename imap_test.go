package imap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImap(t *testing.T) {

	type Item struct {
		src      interface{}
		expected interface{}
		msg      string
	}

	items := []Item{
		{
			src: map[string]interface{}{"keyOfString": "stringValue", "KeyOfStringArr": []string{"stingValue", "stringValue"}},
			msg: "string to string and string to []string",
		},
		{
			src: map[string]interface{}{"keyOfInt": 1, "keyOfString": "test", "keyOfBool": true},
			msg: "int string bool",
		},
		{
			src: map[string]interface{}{"keyOfInt": map[string]string{"stringKey": "stringValue"}},
			msg: "map[string]string",
		},
	}

	for _, item := range items {
		var dst map[string]interface{}
		err := DeepCopy(item.src, &dst)
		assert.NoError(t, err, item.msg)
		assert.Equal(t, item.src, dst, item.msg)
	}

	/*
		{
			src: map[string]interface{}{"keyOfString": map[string]int{"stringKey": 1}},
			msg: "map[string]int",
		},
	*/
}
