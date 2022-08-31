package baiduyun

import (
	"crypto/md5"
	"fmt"
)

// 4 MB
const blockMaxSize = 4 * 1024 * 1024

func splitBytes(content []byte, size int) [][]byte {
	res := make([][]byte, 0)
	for len(content) > 0 {
		if len(content) > size {
			res = append(res, content[:size])
			content = content[size:]
		} else {
			res = append(res, content)
			content = nil
		}
	}
	return res
}

func getMd5(bs []byte) string {
	res := md5.Sum(bs)
	return fmt.Sprintf("%x", res)
}

func ptrInt64(v int64) *int64 {
	return &v
}
