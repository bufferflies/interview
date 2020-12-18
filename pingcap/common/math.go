package common

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
func ByteToInt64(bs []byte) int64 {
	if len(bs) < 4 {
		return 0
	}
	v := (int64(bs[7]) << 58) | (int64(bs[6]) << 48) | (int64(bs[5]) << 40) | (int64(bs[4]) << 32) |
		(int64(bs[3]) << 24) | (int64(bs[2]) << 16) | (int64(bs[1]) << 8) | (int64(bs[0]))
	return v
}
