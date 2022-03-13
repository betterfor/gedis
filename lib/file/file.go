package file

import (
	"fmt"
	"os"
	"path"
)

// MustOpen 打开的文件
func MustOpen(dir, filename string) (*os.File, error) {
	perm := checkPermission(dir)
	if perm {
		return nil, fmt.Errorf("permission denied dir: %s", dir)
	}

	err := mkdir(dir)
	if err != nil {
		return nil, fmt.Errorf("making dir: %s, error: %v", dir, err)
	}

	return os.OpenFile(path.Join(dir, filename), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
}

// checkNotExist 检查文件/目录是否存在
func checkNotExist(s string) bool {
	_, err := os.Stat(s)
	return os.IsNotExist(err)
}

// checkPermission 检查文件/目录权限是否存在
func checkPermission(s string) bool {
	_, err := os.Stat(s)
	return os.IsPermission(err)
}

// mkdir 创建目录
func mkdir(s string) error {
	if checkNotExist(s) {
		return os.MkdirAll(s, os.ModePerm)
	}
	return nil
}
