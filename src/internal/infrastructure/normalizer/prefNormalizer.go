package normalizer

import (
	"fmt"
	"os"
	"strings"
)

func NormalizePref(srcFilePath string, destFilePath string) error {
	// 都道府県データ用のColumnConfigスライスを定義
	columnConfigs := []ColumnConfig{
		{"全国地方公共団体コード", true, nil},
		{"都道府県名", true, nil},
		{"都道府県名_カナ", false, nil},
		{"都道府県名_英字", false, strings.ToUpper},
		{"効力発生日", false, nil},
		{"廃止日", false, nil},
		{"備考", false, nil},
	}

	newColumns := []NewColumnConfig{
		{
			HeaderName: "都道府県コード",
			Generator: func(headers []string, record []string) string {
				value := getRecordValueByHeaderName(headers, record, "全国地方公共団体コード")
				if len(value) >= 2 {
					return value[:2]
				}
				return value
			},
		},
		//{
		//	HeaderName: "都道府県コード3桁",
		//	Generator: func(headers []string, record []string) string {
		//		value := getRecordValueByHeaderName(headers, record, "全国地方公共団体コード")
		//		if len(value) >= 3 {
		//			return value[:3] // 先頭2文字に縮小
		//		}
		//		return value
		//	},
		//},
	}

	// NormalizeCSV 関数を呼び出して都道府県データを正規化
	err := NormalizeCSV(srcFilePath, destFilePath, columnConfigs, newColumns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error normalizing prefecture data: %v\n", err)
	}

	return nil
}
