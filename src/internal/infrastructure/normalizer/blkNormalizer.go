package normalizer

import (
	"fmt"
	"os"
)

func NormalizeBlk(srcFilePath, destFilePath string) error {
	columnConfigs := []ColumnConfig{
		{"全国地方公共団体コード", true, nil},
		{"町字id", true, nil},
		{"街区id", false, nil},
		{"市区町村名", false, nil},
		{"政令市区名", false, nil},
		{"大字・町名", false, nil},
		{"丁目名", false, nil},
		{"小字名", false, nil},
		{"街区符号", true, nil},
		{"住居表示フラグ", false, nil},
		{"住居表示方式コード", true, nil},
		{"大字・町_外字フラグ", false, nil},
		{"小字_外字フラグ", false, nil},
		{"状態フラグ", false, nil},
		{"効力発生日", false, nil},
		{"廃止日", false, nil},
		{"原典資料コード", false, nil},
		{"備考", false, nil},
	}

	err := NormalizeCSV(srcFilePath, destFilePath, columnConfigs, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error normalizing data in file %s: %v\n", srcFilePath, err)
	}

	return nil
}
