package main

import (
	"flag"
	"fmt"
	"github.com/antnesswcm/stario"
	"os"
	"strings"
)

func processParams() {
	if f.help {
		flag.Usage()
		os.Exit(0)
	}
	if len(f.Args) != 0 {
		fmt.Printf("错误用法！请参见下列帮助信息\n============================\n")
		flag.Usage()

	}
	if f.path != "" {
		p.path = strings.Split(f.path, ",")
	} else if len(f.Args) != 0 {
		p.path = f.Args
	} else {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println("没有指定文件夹且获取当前文件夹失败：\n", err)
			stario.WaitUntilString("任意键退出...", "", false)
			os.Exit(1)
		}
		p.path = append(p.path, path)
	}

	if (f.path != "") && (len(f.Args) != 0) {
		fmt.Fprintf(os.Stderr, "由于指定了-f参数，%s将被忽略\n", f.Args)
	}
	if f.connectSymbol != "" {
		p.connectSymbol = f.connectSymbol
	} else {
		fmt.Println(`(输入"/"表示无分隔符，回车默认"_")`)
		c := stario.MessageBox(`请指定连接符号:`, "_")
		if c.MustString() == "/" {
			p.connectSymbol = ""
			fmt.Println("无分隔符")
		} else {
			p.connectSymbol = c.MustString()
			fmt.Printf("分隔符是:\"%s\"\n", p.connectSymbol)
		}
	}
}
