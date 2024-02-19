package extractor

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ExtractZip は指定されたZIPファイルを解凍し、指定されたディレクトリに内容を展開します。
// この関数は再帰的にネストされたZIPファイルも解凍します。
func ExtractZip(srcZipPath, destDir string) error {
	r, err := zip.OpenReader(srcZipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(destDir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		if strings.HasSuffix(f.Name, ".zip") {
			// ネストされたZIPファイルを一時的に保存
			tempZipPath := filepath.Join(destDir, f.Name)
			if err := extractFile(f, tempZipPath); err != nil {
				return err
			}
			// 再帰的に解凍
			if err := ExtractZip(tempZipPath, destDir); err != nil {
				return err
			}
			// 一時ファイルを削除
			os.Remove(tempZipPath)
		} else {
			if err := extractFile(f, fPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// extractFile はZIPファイル内のファイルを解凍します。
func extractFile(f *zip.File, outputPath string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	outFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, rc)
	return err
}
