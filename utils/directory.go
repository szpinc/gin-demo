package utils

import "os"

// PathExists 检查文件是否存在
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return true, nil
}
