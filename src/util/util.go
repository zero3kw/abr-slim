// util.go
package util

import (
	"fmt"
	"strconv"
	"strings"
)

// RoundToSixDigits は与えられた文字列形式の数値を浮動小数点数に変換し、小数点以下6桁に丸めた文字列を返します。
func RoundToSixDigits(value string) string {
	//parts := strings.Split(value, ".")
	//if len(parts) > 1 && len(parts[1]) <= 6 {
	//	return value
	//}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		// 値が数値でない場合はそのまま返します。
		return value
	}

	// 数値を6桁の精度で四捨五入します。
	rounded := fmt.Sprintf("%.6f", floatValue)

	// 小数点以下が0で終わる場合でも、最低1桁の小数を保持します。
	if strings.Contains(rounded, ".") {
		split := strings.Split(rounded, ".")
		integerPart := split[0]
		decimalPart := split[1]

		// 小数部分が0のみで構成されているか、または6桁以下で0で終わる場合、末尾の0を削除しますが、最低1桁は保持します。
		trimmedDecimalPart := strings.TrimRight(decimalPart, "0")
		if len(trimmedDecimalPart) == 0 {
			trimmedDecimalPart = "0" // 小数部分が完全に0の場合は、最低1桁の0を保持します。
		}
		return fmt.Sprintf("%s.%s", integerPart, trimmedDecimalPart)
	}

	// 整数部のみの場合は、そのまま返します。
	return rounded
}
