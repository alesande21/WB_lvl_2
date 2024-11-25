package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		os.Exit(1)
	}

	siteURL := os.Args[1]
	outputDir := "downloaded_site"

	err := downloadSite(siteURL, outputDir)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Сайт успешно скачан в папку: %s\n", outputDir)
	}

}

func downloadSite(siteURL, outputDir string) error {
	parsedURL, err := url.Parse(siteURL)
	if err != nil {
		return fmt.Errorf("некорректный URL: %v", err)
	}

	resp, err := http.Get(siteURL)
	if err != nil {
		return fmt.Errorf("не удалось скачать сайт: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("сервер вернул статус: %d", resp.StatusCode)
	}

	filename := filepath.Join(outputDir, filepath.Base(parsedURL.Path))
	if filename == outputDir {
		filename = filepath.Join(outputDir, "index.html")
	}

	err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return fmt.Errorf("не удалось создать папку: %v", err)
	}

	err = saveToFile(filename, resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка сохранения файла: %v", err)
	}

	return nil
}

func saveToFile(filename string, data io.Reader) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	fmt.Printf("Файл сохранён: %s\n", filename)
	return nil
}
