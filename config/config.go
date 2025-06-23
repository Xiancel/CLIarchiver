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
	fmt.Println("Введіть путь к діректорії:")
	userDir := getInput("> ")

	files, _ := archiver.FindFile(userDir)

	fmt.Println("\nФайли у дерикторіі")
	displayFile(userDir)

	fmt.Println("\nВведіть номера файлів які бажаете добавити(через пробіл):")
	number := getInput("> ")

	selectFile := getSelectFiles(number, files)
	if len(selectFile) == 0 {
		fmt.Println("Жодного файлу не вибрано")
	}

	fmt.Println("\nОберіть формат:")
	format()
	format := getIntInput("> ")

	fmt.Println("\nВведіть назву архіву:")
	archiveName := getInput("> ")

	if !strings.HasSuffix(archiveName, ".zip") {
		archiveName += ".zip"
	}

	path := filepath.Join(userDir, archiveName)

	switch format {
	case 1:
		zip := archiver.ZipArchiver{}
		zip.Compress(selectFile, path)
		fmt.Println("\n✅ Архів успішно створено:", path)
	default:
		fmt.Println("Обрано неіснуючий формат ❌")
	}
}

func displayFile(dir string) {
	files, _ := archiver.FindFile(dir)

	for _, f := range files {
		fmt.Printf("%d. %s\n", f.Index, f.Name)
	}
}

func getSelectFiles(input string, files []mod.FileEntry) []string {
	numList := strings.Fields(input)
	var selectFile []string

	for _, n := range numList {
		var num int
		_, err := fmt.Sscanf(n, "%d", &num)
		if err != nil {
			fmt.Printf("Пропущене некоректне число: %s\n", n)
			continue
		}

		for _, f := range files {
			if f.Index == num {
				selectFile = append(selectFile, f.Path)
				break
			}
		}
	}

	return selectFile
}
func format() {
	fmt.Println("1. ZIP")
}
