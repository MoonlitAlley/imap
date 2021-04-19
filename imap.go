package imap

import (
	"bytes"
	"encoding/gob"
	"errors"
)

func DeepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	gob.Register([]interface{}{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[string]string{})
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func DeepCopyStrInt(input map[string]int) (map[string]int, error) {
	if input == nil {
		return nil, errors.New("can't copy nil map")
	}
	newMap := make(map[string]int, len(input))
	err := DeepCopy(&input, &newMap)
	return newMap, err
}
