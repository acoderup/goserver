package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	modelDir      = flag.String("model-dir", "", "模型所在目录; 必需")
	modelNames    = flag.String("model-names", "", "模型结构体名称; 必需")
	modelPkgPath  = flag.String("model-pkg-path", "", "模型包名; 默认自动生成")
	modelPkgAlias = flag.String("model-pkg-alias", "", "模型包别名")
	daoDir        = flag.String("dao-dir", "", "生成代码所在目录; 必需")
	daoPkgPath    = flag.String("dao-pkg-path", "", "生成代码所在目录的包名; 默认自动生成")
	subPkgEnable  = flag.Bool("sub-pkg-enable", false, "每个模型创建一个子包; 默认关闭")
	subPkgStyle   = flag.String("sub-pkg-style", "kebab", "模型子包目录名称风格; 选项: kebab | underscore | lower | camel | pascal; 默认是 kebab")
	counterName   = flag.String("counter-name", "", "自增模型名称; 默认是 counter")
	fileNameStyle = flag.String("file-style", "underscore", "代码文件名称风格; 选项: kebab | underscore | lower | camel | pascal; 默认是 underscore")
)

// Usage is a replacement usage function for the flags package.
func usage() {
	fmt.Fprintf(os.Stderr, "用法 mongoctl:\n")
	fmt.Fprintf(os.Stderr, "\tmongoctl [flags] -model-dir=. -model-names=T,T -dao-dir=./dao\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

//go:generate mongoctl -type=Mail,User
func main() {
	log.SetFlags(0)
	log.SetPrefix("mongoctl: ")
	flag.Usage = usage
	flag.Parse()

	if len(*modelDir) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if len(*modelNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if len(*daoDir) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	newGenerator(&options{
		daoDir:        *daoDir,
		daoPkgPath:    *daoPkgPath,
		modelDir:      *modelDir,
		modelNames:    strings.Split(*modelNames, ","),
		modelPkgPath:  *modelPkgPath,
		modelPkgAlias: *modelPkgAlias,
		subPkgEnable:  *subPkgEnable,
		subPkgStyle:   style(*subPkgStyle),
		counterName:   *counterName,
		fileNameStyle: style(*fileNameStyle),
	}).makeDao()
}
