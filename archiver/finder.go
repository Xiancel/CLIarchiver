package archiver

import (
	mod "cliarchiver/model"
	"fmt"
	"os"
	"path/filepath"
)

// пошук файлів в деректорії
func FindFile(directory string) ([]mod.FileEntry, error) {
	// читання директорії
	dir, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("error reading directory")
	}

	// змінна для зберегання файлів знайдених у директорії
	var files []mod.FileEntry
	index := 1 // змінна для індексації файлів

	// додавання файлів у змінну
	for _, d := range dir {
		// перевірка на файл
		if !d.IsDir() {
			files = append(files, mod.FileEntry{
				Index: index,
				Name:  d.Name(),
				Path:  filepath.Join(directory, d.Name()),
			})
			index++
		}
	}

	// перевірка на порожність директорії
	if len(files) == 0 {
		return nil, fmt.Errorf("directory is empty")
	}

	return files, nil
}
