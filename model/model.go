package model

// структура FileEntry яка містить
type FileEntry struct {
	Index int    //  індекс файлу
	Name  string // назву
	Path  string // шлях
}

// інтерфейс який має метод Compress який:
type Archiver interface {
	Compress(files []string, directory string) error // архівуе файли
}
