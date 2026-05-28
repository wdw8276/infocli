package utils

import (
	"bufio"
	"fmt"
	"os"
)

func SaveToFile(filePath string, content string) bool {
	file, e := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		return false
	}
	defer file.Close()

	_, err := file.WriteString(content)
	return err == nil
}

func DirExist(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && info.IsDir()
}

func FileExist(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetFileLines(filePath string) []string {
	var lines []string
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file [%s] failed\n", filePath)
		return lines
	}

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		lines = append(lines, fs.Text())
	}

	f.Close()
	return lines
}

func GetFileAll(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("open file [%s] failed\n", filePath)
		return ""
	}
	return string(bytes)
}
