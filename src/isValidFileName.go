package main

import (
	"regexp"
	"strings"
)

func isValidFileName(filename string) bool {

	// 文件名不能以点开头或仅由点组成，且长度不能超过 255 个字符
	if strings.HasPrefix(filename, ".") || strings.Trim(filename, ".") == "" || len(filename) > 255 {
		return false
	}

	// 文件名不能以空格或句号结尾
	if strings.HasSuffix(filename, " ") || strings.HasSuffix(filename, ".") {
		return false
	}

	// 文件名不能包含特殊的文件名
	specialFileNames := []string{"CON", "PRN", "AUX", "NUL", "COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9", "LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9"}
	for _, name := range specialFileNames {
		if strings.ToUpper(filename) == name {
			return false
		}
	}

	// 文件名不能包含以下字符
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1f]`)
	if invalidCharsRegex.MatchString(filename) {
		return false
	}

	// 所有检查都通过，则认为是有效的文件名
	return true
}
