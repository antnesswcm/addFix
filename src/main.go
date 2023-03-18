package main

import (
	"fmt"
	"github.com/antnesswcm/stario"
	"path/filepath"

	"os"
)

type flags struct {
	Args          []string
	path          string
	connectSymbol string
	help          bool
}
type params struct {
	path          []string
	connectSymbol string
	skipFolder    bool
	after         bool
	preview       bool
}

var f = new(flags)
var p = params{}
var selfPath string

func main() {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		fmt.Println("未获取到程序路径，可能导致程序也被重命名(不重要)")
	}
	selfPath = path

	parseFlags()
	//fmt.Printf("%#v\n", f)
	processParams()
	//fmt.Printf("%#v\n", p)
	//os.Exit(0)
	files, err := getAllFiles(p.path[0], p.skipFolder) // todo 处理更多文件夹
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("成功加载文件夹:%s\n", p.path[0]) // todo 清除终端
	if !stario.GetYesNoInput(fmt.Sprintf("成功加载了%d个文件(夹)，是否进行处理(y/N)", len(files)), false) {
		fmt.Println("未进行任何更改！")
		pauseExit(0)
	}
	for _, file := range files {
		folderName, err := getFolderName(p.path[0])
		if err != nil {
			fmt.Println(err)
			break
		}
		// todo 预览
		if err := processFile(file, folderName, p.connectSymbol, p.after); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("Done!\n")
	pauseExit(0)
}

// D:\测试文件夹
func pauseExit(code int) {
	stario.WaitUntilString("任意键退出...", "", false)
	os.Exit(code)
}
