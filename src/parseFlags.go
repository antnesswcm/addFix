package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func parseFlags() {
	flag.StringVar(&f.path, "f", "", "需要处理的文件所在的文件夹路径, 多个路径之间用英文逗号分隔！")
	flag.StringVar(&f.connectSymbol, "c", "", "用于连接的字符")

	flag.BoolVar(&p.after, "a", false, "将文字连接到后面")
	flag.BoolVar(&p.preview, "p", false, "预览结果")
	flag.BoolVar(&p.skipFolder, "s", false, "跳过文件夹")

	flag.BoolVar(&f.help, "h", false, "显示帮助信息")
	flag.BoolVar(&f.help, "help", false, "显示帮助信息")

	// 自定义打印帮助信息的格式
	flag.Usage = usage

	flag.Parse()
	// 游离参数
	f.Args = flag.Args()
}
func usage() {
	fmt.Fprintf(os.Stderr, "用法: %s [选项] -f path1,path2,...\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(os.Stderr, "选项:")
	flag.PrintDefaults()
	// todo 添加更多提示
}
