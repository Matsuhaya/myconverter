package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

// errorオブジェクトをチェックし、nilの場合例外を送出
func assert(err error, msg string) {
	if err != nil {
		panic(err.Error() + ":" + msg)
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func convertFile(path string) error {

	// ファイルオープン
	file, err := os.Open(path)
	assert(err, "Failed to open file.")
	defer file.Close()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, err := image.Decode(file)
	assert(err, "Failed to decode file")

	// 出力ファイルの.pngパスを生成
	dstPath := filepath.Join(filepath.Dir(path), getFileNameWithoutExt(path)+".png")
	fmt.Println(dstPath)

	// 出力ファイルを生成
	out, err := os.Create(dstPath)
	assert(err, "Failed to create file.")
	defer out.Close()

	jpeg.Encode(out, img, nil)

	return nil
}

func main() {
	flag.Parse()
	dir := flag.Args()

	fmt.Println(dir)
	err := filepath.Walk(dir[0],
		func(path string, info os.FileInfo, err error) error {

			// .jpgのみ変換処理を実行
			if filepath.Ext(path) == ".jpg" {

				fmt.Println(path)
				err := convertFile(path)

				assert(err, "Failed to convert file.")
			}
			return nil
		})

	assert(err, "Failed to walk path.")
}
