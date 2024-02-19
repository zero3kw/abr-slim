package normalizer

import (
	"abr-slim/src/util"
	"fmt"
	"os"
)

func NormalizePosBlk(srcFilePath, destFilePath string) error {
	columnConfigs := []ColumnConfig{
		{"全国地方公共団体コード", true, nil},
		{"町字id", true, nil},
		{"街区id", true, nil},
		{"住居表示フラグ", true, nil},
		{"住居表示方式コード", false, nil},
		{"代表点_経度", true, util.RoundToSixDigits},
		{"代表点_緯度", true, util.RoundToSixDigits},
		{"代表点_座標参照系", false, nil},
		{"代表点_地図情報レベル", false, nil},
		{"ポリゴン_ファイル名", false, nil},
		{"ポリゴン_キーコード", false, nil},
		{"ポリゴン_データフォーマット", false, nil},
		{"ポリゴン_座標参照系", false, nil},
		{"ポリゴン_地図情報レベル", false, nil},
		{"位置参照情報_都道府県名", false, nil},
		{"位置参照情報_市区町村名", false, nil},
		{"位置参照情報_大字・町丁目名", false, nil},
		{"位置参照情報_小字・通称名", false, nil},
		{"位置参照情報_街区符号・地番", false, nil},
		{"位置参照情報_データ整備年度", false, nil},
		{"電子国土基本図（地名情報）「住居表示住所」_住所コード（可読）", false, nil},
		{"電子国土基本図（地名情報）「住居表示住所」_データ整備日", false, nil},
	}

	err := NormalizeCSV(srcFilePath, destFilePath, columnConfigs, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error normalizing data in file %s: %v\n", srcFilePath, err)
	}

	return nil
}
