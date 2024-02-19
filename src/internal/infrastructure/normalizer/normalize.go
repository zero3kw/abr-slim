package normalizer

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// ColumnConfig はカラムの設定を定義する構造体です。
type ColumnConfig struct {
	HeaderName string
	Output     bool
	Processor  func(string) string
}

type NewColumnConfig struct {
	HeaderName string
	// Generator関数がheaders配列も受け取れるようにシグネチャを変更します。
	Generator func(headers []string, record []string) string
}

// NormalizeCSV はCSVファイルを正規化する関数です。
func NormalizeCSV(srcFilePath, destFilePath string, columnConfigs []ColumnConfig, newColumns []NewColumnConfig) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	reader := csv.NewReader(srcFile)
	headers, err := reader.Read()
	if err != nil {
		return err
	}

	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	writer := csv.NewWriter(destFile)

	// 出力対象のヘッダーとそのインデックスを特定し、Outputがtrueのものだけを追加
	outputHeaders := []string{}
	for _, config := range columnConfigs {
		if config.Output {
			outputHeaders = append(outputHeaders, config.HeaderName)
		}
	}
	for _, newCol := range newColumns {
		outputHeaders = append(outputHeaders, newCol.HeaderName)
	}

	// 出力対象のヘッダー行を書き込む
	if err := writer.Write(outputHeaders); err != nil {
		return err
	}

	// データ行の処理
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("%s\n", err)
			fmt.Printf("%s\n", record)
			continue
			//return err
		}

		// 処理済みレコードを構築
		outputRecord := []string{}
		for _, config := range columnConfigs {
			if !config.Output {
				continue
			}
			index := indexOf(headers, config.HeaderName)
			if index == -1 || index >= len(record) {
				continue // ヘッダーが見つからない、または範囲外の場合はスキップ
			}
			value := record[index]
			if config.Processor != nil {
				value = config.Processor(value) // 値の加工
			}
			outputRecord = append(outputRecord, value)
		}

		// 新しいカラムの値を生成して追加
		for _, newCol := range newColumns {
			newValue := newCol.Generator(headers, record)
			outputRecord = append(outputRecord, newValue)
		}

		if err := writer.Write(outputRecord); err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}

// indexOf は指定された文字列がスライス内で最初に現れるインデックスを返します。
// 存在しない場合は-1を返します。
func indexOf(slice []string, item string) int {
	for i, sliceItem := range slice {
		if sliceItem == item {
			return i
		}
	}
	return -1
}

func getRecordValueByHeaderName(headers []string, record []string, headerName string) string {
	for i, h := range headers {
		if h == headerName {
			return record[i]
		}
	}
	return "" // ヘッダー名が見つからない場合は空文字を返します。
}
