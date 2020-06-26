package main

import (
	"flag"
	"fmt"
	"myproject/myconverter/assert"
	"myproject/myconverter/convert"
	"os"
	"path/filepath"
)

var (
	iExt = flag.String("i", "jpg", "string_flag")
	oExt = flag.String("o", "png", "string_flag")
)

func main() {
	flag.Parse()
	dir := flag.Args()

	fmt.Println(dir)
	err := filepath.Walk(dir[0],
		func(path string, info os.FileInfo, err error) error {

			// iオプションで指定した拡張子のファイルのみ変換処理を実行
			ext := "." + *iExt
			if filepath.Ext(path) == ext {
				fmt.Println(path)
				err := convert.ConvertFile(path, oExt)
				assert.Assert(err, "Failed to convert file.")
			}
			return nil
		})

	assert.Assert(err, "Failed to walk path.")
}
