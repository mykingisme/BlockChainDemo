package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func IntToByte(i int64) []byte {
	var buff bytes.Buffer
	Err(binary.Write(&buff, binary.BigEndian, i))
	return buff.Bytes()
}

func Err(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
