package config

import (
	"bufio"
	"cliarchiver/archiver"
	mod "cliarchiver/model"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// читачь строки
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// приймає текстове введеня
func getInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Отримує ціло числене введення
func getIntInput(prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		//перевірка на пробіли
		if strings.ContainsAny(input, " \t") {
			fmt.Println("\nНекоректне введення. Ведіть ціле число без пробілів. ❌")
			return -1
		}

		//перевірка на введеня числа
		var value int
		_, err := fmt.Sscanf(input, "%d", &value)
		if err != nil {
			fmt.Println("\nНекоректне введення. Ведіть ціле число. ❌")
			continue
		}
		return value
	}
}

func Input() {
	// получення діректорії від пользователя
	fmt.Println("Введіть путь к діректорії:")
	userDir := getInput("> ")

	// пошук файлів в діректорії
	files, err := archiver.FindFile(userDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// вивід всіх файлів у директорії
	fmt.Println("\nФайли у дерикторіі")
	displayFile(userDir)

	// получення номера файлів для добавлення
	fmt.Println("\nВведіть номера файлів які бажаете добавити(через пробіл):")
	number := getInput("> ")

	// формування обраних файлів
	selectFile := getSelectFiles(number, files)
	if len(selectFile) == 0 {
		fmt.Println("Жодного файлу не вибрано")
	}

	// получення формату архівації від користувача
	fmt.Println("\nОберіть формат:")
	format()
	format := getIntInput("> ")

	// получення назви для архуву від користувача
	fmt.Println("\nВведіть назву архіву:")
	archiveName := getInput("> ")

	// згідно формату формується архів
	switch format {
	case 1:
		// перевірка назви на наявність расширення
		if !strings.HasSuffix(archiveName, ".zip") {
			archiveName += ".zip"
		}

		// формування шляху для архіва
		path := filepath.Join(userDir, archiveName)

		// формування шляху для архіва
		zip := archiver.ZipArchiver{}
		zip.Compress(selectFile, path)
		fmt.Println("\n✅ Архів успішно створено:", path)
	case 2:
		// перевірка назви на наявність расширення
		if !strings.HasSuffix(archiveName, ".tar") {
			archiveName += ".tar"
		}

		// формування шляху для архіва
		path := filepath.Join(userDir, archiveName)

		// формування шляху для архіва
		tar := archiver.TarArchiver{}
		tar.Compress(selectFile, path)
		fmt.Println("\n✅ Архів успішно створено:", path)
	case 3:
		// перевірка назви на наявність расширення
		if !strings.HasSuffix(archiveName, ".tar.gz") {
			archiveName += ".tar.gz"
		}

		// формування шляху для архіва
		path := filepath.Join(userDir, archiveName)

		// формування шляху для архіва
		targz := archiver.TarGzArchiver{}
		targz.Compress(selectFile, path)
		fmt.Println("\n✅ Архів успішно створено:", path)
	default:
		fmt.Println("Обрано неіснуючий формат ❌")
	}
}

// функція для відображення всіх файлів у директорії
func displayFile(dir string) {
	// виклик функції
	files, err := archiver.FindFile(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	// відображення всіх файлів
	for _, f := range files {
		fmt.Printf("%d. %s\n", f.Index, f.Name)
	}
}

// функція для формування вибраних файлів користувачм
func getSelectFiles(input string, files []mod.FileEntry) []string {
	// розбиття строки
	numList := strings.Fields(input)
	// слайс для зберегання вибраних файлів
	var selectFile []string

	for _, n := range numList {
		var num int
		// перевірка на введеня коректного числа
		_, err := fmt.Sscanf(n, "%d", &num)
		if err != nil {
			fmt.Printf("Пропущене некоректне число: %s\n", n)
			continue
		}
		// додавання файлів по індексу до слайсу selectFile
		for _, f := range files {
			if f.Index == num {
				selectFile = append(selectFile, f.Path)
				break
			}
		}
	}

	return selectFile
}

// функція для відображення всіх форматів
func format() {
	fmt.Println("1. ZIP")
	fmt.Println("2. TAR")
	fmt.Println("3. TAR.GZ")
}
