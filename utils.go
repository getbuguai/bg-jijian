package main

import (
	"fmt"
	"io/ioutil"
)

// SaveFile 保存文件到本地
func SaveFile(filename string, data []byte) (err error) {
	if len(filename) == 0 {
		return fmt.Errorf("filename length is 0")
	}
	if len(data) == 0 {
		return fmt.Errorf("data length is 0")
	}

	return ioutil.WriteFile(filename, data, 0644)
}
