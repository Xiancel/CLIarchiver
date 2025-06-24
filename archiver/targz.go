package archiver

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// структура tarGzArchiver
type TarGzArchiver struct{}

// Реалізація інтерфейсу Compress
func (t TarGzArchiver) Compress(files []string, directory string) error {
	// создання архиву
	archiveFile, err := os.Create(directory)
	if err != nil {
		return fmt.Errorf("error creating archive")
	}

	// создання писачів для файлів
	gzipWriter := gzip.NewWriter(archiveFile)

	// создання писачя для gzipWriter
	tarWriter := tar.NewWriter(gzipWriter)

	// запис файлів в архів
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			tarWriter.Close()
			gzipWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error opening file")
		}

		// получення інформації про файли
		info, err := f.Stat()
		if err != nil {
			f.Close()
			tarWriter.Close()
			gzipWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error getting file info")
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			f.Close()
			tarWriter.Close()
			gzipWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error creating header")
		}

		// встановлює ім'я для файлу
		header.Name = filepath.Base(file)

		// запис заголовку
		err = tarWriter.WriteHeader(header)
		if err != nil {
			f.Close()
			tarWriter.Close()
			gzipWriter.Close()
			archiveFile.Close()
			return fmt.Errorf("error writing header")
		}

		// копиювання файлу у архів
		_, err = io.Copy(tarWriter, f)
		f.Close()

		if err != nil {
			tarWriter.Close()
			gzipWriter.Close()
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

	err = gzipWriter.Close()
	if err != nil {
		archiveFile.Close()
		return fmt.Errorf("error closing gzip writer")
	}
	err = archiveFile.Close()
	if err != nil {
		return fmt.Errorf("error closing archive file")
	}

	return nil
}
