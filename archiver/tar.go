package archiver

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// структура tarArchiver
type TarArchiver struct{}

// Реалізація інтерфейсу Compress
func (t TarArchiver) Compress(files []string, directory string) error {
	// создання архиву
	archiveFile, err := os.Create(directory)
	if err != nil {
		return fmt.Errorf("error creating archive")
	}

	// создання tar-писателя для записі архіва
	tarWriter := tar.NewWriter(archiveFile)

	// запис файлів в архів
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			tarWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error opening file")
		}
		// получення інформації про файл
		info, err := f.Stat()
		if err != nil {
			f.Close()
			tarWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error getting file info")
		}

		// создання заголовку для запису у файл
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			f.Close()
			tarWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error creating header")
		}

		// встановлення ім'я для файлу
		header.Name = filepath.Base(file)

		// запис заголовку
		err = tarWriter.WriteHeader(header)
		if err != nil {
			f.Close()
			tarWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error writing header")
		}

		// копиювання файлу у архів
		_, err = io.Copy(tarWriter, f)
		f.Close() // закриття файлу

		if err != nil {
			tarWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error copying file")
		}
	}

	// обробка помилок при закрітті файлів
	err = tarWriter.Close()
	if err != nil {
		archiveFile.Close()
		return fmt.Errorf("error closing tar writer")
	}

	err = archiveFile.Close()
	if err != nil {
		return fmt.Errorf("error closing archive file")
	}

	return nil
}
