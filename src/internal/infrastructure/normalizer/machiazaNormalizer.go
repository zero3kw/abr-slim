package normalizer

import (
	"fmt"
	"os"
)

func NormalizeMachiaza(srcFilePath, destFilePath string) error {
	columnConfigs := []ColumnConfig{
		{"全国地方公共団体コード", true, nil},
		{"町字id", true, nil},
		{"町字区分コード", true, nil},
		{"都道府県名", false, nil},
		{"都道府県名_カナ", false, nil},
		{"都道府県名_英字", false, nil},
		{"郡名", false, nil},
		{"郡名_カナ", false, nil},
		{"郡名_英字", false, nil},
		{"市区町村名", false, nil},
		{"市区町村名_カナ", false, nil},
		{"市区町村名_英字", false, nil},
		{"政令市区名", false, nil},
		{"政令市区名_カナ", false, nil},
		{"政令市区名_英字", false, nil},
		{"大字・町名", true, nil},
		{"大字・町名_カナ", false, nil},
		{"大字・町名_英字", false, nil},
		{"丁目名", true, nil},
		{"丁目名_カナ", false, nil},
		{"丁目名_数字", false, nil},
		{"小字名", true, nil},
		{"小字名_カナ", false, nil},
		{"小字名_英字", false, nil},
		//{"同一町字識別情報", true, nil},
		{"住居表示フラグ", true, nil},
		{"住居表示方式コード", true, nil},
		{"大字・町名_通称フラグ", true, nil},
		{"小字名_通称フラグ", true, nil},
		{"大字・町名_電子国土基本図外字", false, nil},
		{"小字名_電子国土基本図外字", false, nil},
		{"状態フラグ", true, nil},
		{"起番フラグ", true, nil},
		{"効力発生日", false, nil},
		{"廃止日", false, nil},
		{"原典資料コード", false, nil},
		{"郵便番号", false, nil},
		{"備考", false, nil},
	}

	err := NormalizeCSV(srcFilePath, destFilePath, columnConfigs, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error normalizing data in file %s: %v\n", srcFilePath, err)
	}

	return nil
}
