package archiver

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// структура
type ZipArchiver struct{}

// Реалізація інтерфейсу Compress
func (z ZipArchiver) Compress(files []string, directory string) error {
	// создання архиву
	archiveFile, err := os.Create(directory)
	if err != nil {
		return fmt.Errorf("error creating archive")
	}

	// создання Zip-писателя для записі архіва
	zipWriter := zip.NewWriter(archiveFile)

	// додавання файлів в архив
	for _, file := range files {
		f, err := os.Open(file) // відкріття файлу
		if err != nil {
			zipWriter.Close()                       // закриття писателя
			archiveFile.Close()                     // закриття архіву
			return fmt.Errorf("error opening file") // вивід помилки
		}

		// створення нового запису в архіві з назвою файлу
		w, err := zipWriter.Create(filepath.Base(file))
		if err != nil {
			f.Close()                                         // закриття файлу
			zipWriter.Close()                                 // закриття писателя
			archiveFile.Close()                               // закриття архіву
			return fmt.Errorf("error creating archive entry") // вивід помилки
		}

		// копиювання файлу у архів
		_, err = io.Copy(w, f)
		f.Close() // закриття файлу

		if err != nil {
			zipWriter.Close()                       // закриття писателя
			archiveFile.Close()                     // закриття архіву
			return fmt.Errorf("error copying file") // вивід помилки
		}
	}
	// обробка помилок при закрітті файлів
	err = zipWriter.Close()
	if err != nil {
		archiveFile.Close()
		return fmt.Errorf("error closing zipWriter")
	}

	err = archiveFile.Close()
	if err != nil {
		return fmt.Errorf("error closing archive file")
	}

	return nil
}
