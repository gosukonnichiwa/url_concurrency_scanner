package scanner

import (
	"fmt"
	"os"
	"sync"
	"time"

	"concurrency-url-scanner/internal/config"
)

// Потокобезопасная запись в файл с повторами при ошибках
type FileWriter struct {
	fileMux sync.Mutex
	file    string
	retries int
}

// Создает новый экземпляр FileWriter
func NewFileWriter(cfg *config.Config) *FileWriter {
	return &FileWriter{
		file:    cfg.ResultsFile,
		retries: cfg.MaxRetries,
	}
}

// Пытается записать данные с несколькими попытками
func (fw *FileWriter) Write(data string) error {
	for i := 0; i < fw.retries; i++ {
		if err := fw.tryWrite(data); err == nil {
			return nil
		}
		time.Sleep(time.Second * time.Duration(i+1))
	}
	return fmt.Errorf("max retries exceeded")
}

// Выполняет одну попытку записи (с защитой мьютексом)
func (fw *FileWriter) tryWrite(data string) error {
	fw.fileMux.Lock()
	defer fw.fileMux.Unlock()

	file, err := os.OpenFile(fw.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data + "\n")
	return err
}
