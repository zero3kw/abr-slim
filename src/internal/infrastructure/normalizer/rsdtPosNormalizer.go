package normalizer

import (
	"abr-slim/src/util"
	"fmt"
	"os"
)

func NormalizePosRsdt(srcFilePath, destFilePath string) error {
	columnConfigs := []ColumnConfig{
		{"全国地方公共団体コード", true, nil},
		{"町字id", true, nil},
		{"街区id", true, nil},
		{"住居id", true, nil},
		{"住居2id", true, nil},
		{"住居表示フラグ", true, nil},
		{"住居表示方式コード", false, nil},
		{"代表点_経度", true, util.RoundToSixDigits},
		{"代表点_緯度", true, util.RoundToSixDigits},
		{"代表点_座標参照系", false, nil},
		{"代表点_地図情報レベル", false, nil},
		{"代表点_原典資料コード", false, nil},
		{"電子国土基本図（地名情報）「住居表示住所」_住所コード（可読）", false, nil},
		{"電子国土基本図（地名情報）「住居表示住所」_データ整備日", false, nil},
		{"基礎番号・住居番号区分", false, nil},
	}

	err := NormalizeCSV(srcFilePath, destFilePath, columnConfigs, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error normalizing data in file %s: %v\n", srcFilePath, err)
	}

	return nil
}
