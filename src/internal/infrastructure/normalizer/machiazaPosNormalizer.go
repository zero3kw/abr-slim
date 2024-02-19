package normalizer

import (
	"abr-slim/src/util"
	"fmt"
	"os"
)

func NormalizePosMachiaza(srcFilePath, destFilePath string) error {
	columnConfigs := []ColumnConfig{
		{"全国地方公共団体コード", true, nil},
		{"町字id", true, nil},
		{"住居表示フラグ", true, nil},
		{"代表点_経度", true, util.RoundToSixDigits},
		{"代表点_緯度", true, util.RoundToSixDigits},
		{"代表点_座標参照系", false, nil},
		{"代表点_地図情報レベル", false, nil},
		{"ポリゴン_ファイル名", false, nil},
		{"ポリゴン_キーコード", false, nil},
		{"ポリゴン_データフォーマット", false, nil},
		{"ポリゴン_座標参照系", false, nil},
		{"ポリゴン_地図情報レベル", false, nil},
		{"位置参照情報_大字町丁目コード", false, nil},
		{"位置参照情報_データ整備年度", false, nil},
		{"国勢調査_境界_小地域（町丁・字等別）_KEY_CODE", false, nil},
		{"国勢調査_境界_データ整備年度", false, nil},
	}

	err := NormalizeCSV(srcFilePath, destFilePath, columnConfigs, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error normalizing data in file %s: %v\n", srcFilePath, err)
	}

	return nil
}
