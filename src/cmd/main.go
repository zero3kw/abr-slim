// abr-slim/src/cmd/main.go

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"abr-slim/src/config"
	"abr-slim/src/internal/infrastructure/downloader"
	"abr-slim/src/internal/infrastructure/extractor"
	"abr-slim/src/internal/infrastructure/normalizer"
)

func main() {
	configPath := "./config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 作業ディレクトリ内のサブディレクトリを作成
	downloadDir := filepath.Join(cfg.WorkingDir, "downloads")
	extractDir := filepath.Join(cfg.WorkingDir, "extracts")
	normalizeDir := filepath.Join(cfg.WorkingDir, "normalized")

	// サブディレクトリの作成
	createDir(downloadDir)
	createDir(extractDir)
	createDir(normalizeDir)

	for _, dataset := range cfg.Datasets {
		downloadURL := cfg.BaseURL + dataset.File
		savePath := filepath.Join(cfg.WorkingDir, "downloads", dataset.File)

		// download
		if err := downloader.DownloadFile(downloadURL, savePath); err != nil {
			log.Printf("Failed to download %s: %v\n", dataset.ID, err)
			continue
		}

		// unzip
		extractPath := filepath.Join(cfg.WorkingDir, "extracts") // サブディレクトリを使用しない
		if err := extractor.ExtractZip(savePath, extractPath); err != nil {
			log.Printf("Failed to extract %s: %v\n", dataset.ID, err)
			continue
		}

		// normalize
		var srcCSVPath, destCSVPath string
		switch dataset.ID {
		case "pref":
			srcCSVPath = filepath.Join(extractPath, "mt_pref_all.csv")
			destCSVPath = filepath.Join(cfg.WorkingDir, "normalized", "mt_pref_slim.csv")
			normalizer.NormalizePref(srcCSVPath, destCSVPath)
		case "city":
			srcCSVPath = filepath.Join(extractPath, "mt_city_all.csv")
			destCSVPath = filepath.Join(cfg.WorkingDir, "normalized", "mt_city_slim.csv")
			normalizer.NormalizeCity(srcCSVPath, destCSVPath)
		case "machiaza":
			srcCSVPath = filepath.Join(extractPath, "mt_town_all.csv")
			destCSVPath = filepath.Join(cfg.WorkingDir, "normalized", "mt_town_slim.csv")
			normalizer.NormalizeMachiaza(srcCSVPath, destCSVPath)
		case "machiaza-pos":
			for i := 1; i <= 47; i++ {
				srcCSVPath := filepath.Join(extractDir, fmt.Sprintf("mt_town_pos_pref%02d.csv", i))
				destCSVPath := filepath.Join(normalizeDir, fmt.Sprintf("mt_town_pos_pref%02d_slim.csv", i))
				normalizer.NormalizePosMachiaza(srcCSVPath, destCSVPath)
			}
		case "blk":
			for i := 1; i <= 47; i++ {
				srcCSVPath := filepath.Join(extractDir, fmt.Sprintf("mt_rsdtdsp_blk_pref%02d.csv", i))
				destCSVPath := filepath.Join(normalizeDir, fmt.Sprintf("mt_rsdtdsp_blk_pref%02d_slim.csv", i))
				normalizer.NormalizeBlk(srcCSVPath, destCSVPath)
			}
		case "blk-pos":
			for i := 1; i <= 47; i++ {
				srcCSVPath := filepath.Join(extractDir, fmt.Sprintf("mt_rsdtdsp_blk_pos_pref%02d.csv", i))
				destCSVPath := filepath.Join(normalizeDir, fmt.Sprintf("mt_rsdtdsp_blk_pos_pref%02d_slim.csv", i))
				normalizer.NormalizePosBlk(srcCSVPath, destCSVPath)
			}
		case "rsdt":
			for i := 1; i <= 47; i++ {
				srcCSVPath := filepath.Join(extractDir, fmt.Sprintf("mt_rsdtdsp_rsdt_pref%02d.csv", i))
				destCSVPath := filepath.Join(normalizeDir, fmt.Sprintf("mt_rsdtdsp_rsdt_pref%02d_slim.csv", i))
				normalizer.NormalizeRsdt(srcCSVPath, destCSVPath)
			}
		case "rsdt-pos":
			for i := 1; i <= 47; i++ {
				srcCSVPath := filepath.Join(extractDir, fmt.Sprintf("mt_rsdtdsp_rsdt_pos_pref%02d.csv", i))
				destCSVPath := filepath.Join(normalizeDir, fmt.Sprintf("mt_rsdtdsp_rsdt_pos_pref%02d_slim.csv", i))
				normalizer.NormalizePosRsdt(srcCSVPath, destCSVPath)
			}
		default:
			log.Printf("Unknown dataset ID: %s\n", dataset.ID)
			continue
		}

		fmt.Printf("Successfully processed dataset: %s\n", dataset.ID)
	}
}

// createDir creates a directory at the specified path. If the directory already exists, it does nothing.
func createDir(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalf("Failed to create directory %s: %v", path, err)
	}
}
