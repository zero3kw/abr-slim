package normalizer

import (
	"fmt"
	"os"
	"strings"
)

func NormalizeCity(srcFilePath, destFilePath string) error {
	columnConfigs := []ColumnConfig{
		{"全国地方公共団体コード", true, nil},
		{"都道府県名", true, nil},
		{"都道府県名_カナ", false, nil},
		{"都道府県名_英字", false, strings.ToUpper},
		{"郡名", true, nil},
		{"郡名_カナ", false, nil},
		{"郡名_英字", false, strings.ToUpper},
		{"市区町村名", true, nil},
		{"市区町村名_カナ", false, nil},
		{"市区町村名_英字", false, strings.ToUpper},
		{"政令市区名", true, nil},
		{"政令市区名_カナ", false, nil},
		{"政令市区名_英字", false, strings.ToUpper},
		{"効力発生日", false, nil},
		{"廃止日", false, nil},
		{"備考", false, nil},
	}

	err := NormalizeCSV(srcFilePath, destFilePath, columnConfigs, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error normalizing data in file %s: %v\n", srcFilePath, err)
	}

	return nil
}
