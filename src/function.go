package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func processFile(filePath string, text string, delimiter string, controlParam bool) error {
	fileName := filepath.Base(filePath)
	newName := fileName
	if controlParam {
		newName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) + delimiter + text + filepath.Ext(fileName)
	} else {
		newName = text + delimiter + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + filepath.Ext(fileName)
	}
	nameNoExt := fileName[0 : len(fileName)-len(filepath.Ext(fileName))]
	if isProcessed(nameNoExt, text, controlParam) {
		return errors.New(fmt.Sprintf("%s\t该文件已经处理过！", fileName))
	}
	newPath := filepath.Join(filepath.Dir(filePath), newName) // todo 校验文件名是否合法
	err := os.Rename(filePath, newPath)
	if err != nil {
		return err
	}
	return nil
}

// getFolderName 函数接受一个路径，返回路径最后的文件夹名
func getFolderName(folderPath string) (string, error) {
	absPath, err := filepath.Abs(folderPath) // 将相对路径转化为绝对路径
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(absPath) // 返回路径文件对象
	if err != nil {
		return "", err
	}
	if !fi.IsDir() { // 判断路径是否为文件夹
		absPath = filepath.Dir(absPath) // 返回文件所在目录的绝对路径
	}
	return filepath.Base(absPath), nil // 返回路径最后一个元素
}

// getAllFiles 函数接受一个路径，返回路径下所有的文件(不包含文件夹)，
// recursive 参数控制是否递归字文件夹
func getAllFiles(path string, skipFolder bool) ([]string, error) {
	var files []string
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(path, fileInfo.Name())
		if filePath == selfPath { // 跳过自己
			continue
		}
		if fileInfo.IsDir() && !skipFolder {
			files = append(files, filePath)
		} else {
			files = append(files, filePath)
		}
	}
	return files, nil
}

func isProcessed(text string, fix string, after bool) (match bool) {
	if after {
		re := regexp.MustCompile(regexp.QuoteMeta(fix) + "$")
		match = re.MatchString(text)
	} else {
		re := regexp.MustCompile("^" + regexp.QuoteMeta(fix))
		match = re.MatchString(text)
	}
	return match
}
