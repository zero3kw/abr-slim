// abr-slim/src/internal/infrastructure/downloader/webDownloader.go

package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url, destPath string) error {
	// 保存先のディレクトリを確認（存在しない場合は作成）
	dirPath := filepath.Dir(destPath)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	// 保存先にファイルが既に存在するかチェック
	if _, err := os.Stat(destPath); err == nil {
		// ファイルが存在する場合、ダウンロードをスキップ
		return nil
	} else if !os.IsNotExist(err) {
		// Statが失敗したが「ファイルがない」以外の理由である場合、エラーを返す
		return err
	}

	// HTTPリクエストを送信してファイルをダウンロード
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// ステータスコードのチェック
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s returned status %d", url, resp.StatusCode)
	}

	// 保存先ファイルを開く
	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// ダウンロードした内容をファイルに書き込み
	_, err = io.Copy(outFile, resp.Body)
	return err
}
