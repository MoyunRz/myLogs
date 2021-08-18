package log

import (
	"bytes"
	"fmt"
	"my-log/config"
	"os"
	"path"
)

func check(e error) bool {
	if e != nil {
		fmt.Printf(e.Error())
		return false
	} else {
		return true
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func WriteLog(msg string) bool {
	var bt bytes.Buffer
	var err1 error
	var f *os.File
	filename := config.LogName()

	if checkFileIsExist(filename) {
		f, err1 = os.OpenFile(filename, os.O_WRONLY, 0666)
		bt.WriteString("\n")
		bt.WriteString(msg)
	} else {
		bt.WriteString(msg)
		EnsureBaseDir(filename)
		f, err1 = os.Create(filename)
	}
	defer f.Close()
	if check(err1) {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, wErr := f.WriteAt([]byte(bt.String()), n)
		if wErr != nil {
			fmt.Errorf(wErr.Error())
			return false
		}
	} else {

		return false
	}

	return true
}

func EnsureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}
