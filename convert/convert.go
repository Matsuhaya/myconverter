package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"myproject/myconverter/assert"
	"os"
	"path/filepath"
)

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func ConvertFile(path string, ext string) error {

	// ファイルオープン
	file, err := os.Open(path)
	assert.Assert(err, "Failed to open file.")
	defer file.Close()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, err := image.Decode(file)
	assert.Assert(err, "Failed to decode file")

	// oオプションで指定した拡張子の出力ファイルパス生成
	dstPath := filepath.Join(filepath.Dir(path), getFileNameWithoutExt(path)+ext)
	fmt.Println(dstPath)

	// 出力ファイルを生成
	out, err := os.Create(dstPath)
	assert.Assert(err, "Failed to create file.")
	defer out.Close()

	jpeg.Encode(out, img, nil)

	return nil
}
