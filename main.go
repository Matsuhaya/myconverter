// This is my own extension converter for image file.
//
// Features
//
// Supports the option to specify the extension.
// Default is to convert JPG to PNG.
//
//  $ go run main.go -i=png -o=jpeg ~/Desktop
//
package main

import (
	"flag"
	"fmt"
	"myproject/myconverter/assert"
	"myproject/myconverter/convert"
	"os"
	"path/filepath"
)

type ext struct {
	input  string
	output string
}

func main() {
	var ext ext
	ext.input = "." + *flag.String("i", "jpg", "string_flag")
	ext.output = "." + *flag.String("o", "png", "string_flag")

	flag.Parse()
	dir := flag.Args()

	fmt.Println(dir)
	err := filepath.Walk(dir[0],
		func(path string, info os.FileInfo, err error) error {

			// iオプションで指定した拡張子のファイルのみ変換処理を実行
			if filepath.Ext(path) == ext.input {
				fmt.Println(path)
				err := convert.ConvertFile(path, ext.output)
				assert.Assert(err, "Failed to convert file.")
			}
			return nil
		})

	assert.Assert(err, "Failed to walk path.")
}
