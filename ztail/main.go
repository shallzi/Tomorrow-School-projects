package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	cCount := -1
	files := []string{}

	if args[0] == "-c" && len(args) >= 2 {
		// Ручной парсинг числа без использования strconv
		val := 0
		hasValidInt := true
		for _, r := range args[1] {
			if r < '0' || r > '9' {
				hasValidInt = false
				break
			}
			val = val*10 + int(r-'0')
		}
		
		if hasValidInt && val > 0 {
			cCount = val
		}
		files = args[2:]
	} else {
		files = args
	}

	hasError := false
	multipleFiles := len(files) > 1
	printedFilesCount := 0 // Счетчик успешно выведенных файлов для правильных отступов

	for _, filename := range files {
		content, err := os.ReadFile(filename)
		if err != nil {
			// Если файла нет, выводим только ошибку без заголовков
			fmt.Println(err)
			hasError = true
			continue
		}

		// Если это не первый успешный файл, перед его заголовком нужна пустая строка
		if multipleFiles && printedFilesCount > 0 {
			fmt.Println()
		}

		// Выводим заголовок только для успешно прочитанных файлов
		if multipleFiles {
			fmt.Printf("==> %s <==\n", filename)
		}

		// Если задан флаг -c, берем последние N байт
		if cCount >= 0 {
			if len(content) > cCount {
				content = content[len(content)-cCount:]
			}
		}

		// Выводим содержимое файла
		os.Stdout.Write(content)
		printedFilesCount++
	}

	// Если хотя бы один файл не удалось прочитать, возвращаем ненулевой статус (1)
	if hasError {
		os.Exit(1)
	}
}