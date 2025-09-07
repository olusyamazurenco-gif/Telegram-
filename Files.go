package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	filesToCreate  = 500
	dirName        = "documents"
	secretPhrase   = "This is the very big and secret key."
)

// generateRandomText створює випадковий текст для файлів
func generateRandomText(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteString("Lorem ipsum dolor sit amet, consectetur adipiscing elit. ")
	}
	return builder.String()
}

// createFiles створює 500 файлів з випадковим текстом
func createFiles() {
	fmt.Println("Створення файлів...")
	if err := os.MkdirAll(dirName, 0755); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < filesToCreate; i++ {
		fileName := fmt.Sprintf("%s/file_%03d.txt", dirName, i)
		content := generateRandomText(100)
		
		// "Секретний" текст буде захований у 10-му, 20-му, 30-му і т.д. файлах
		if i%10 == 0 && i > 0 {
			content += " " + secretPhrase
		}

		if err := ioutil.WriteFile(fileName, []byte(content), 0644); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Файли успішно створені.")
}

// findSecret шукає "секретний" текст у всіх файлах
func findSecret() string {
	fmt.Println("Пошук секретного тексту...")
	
	// Змінна для зберігання знайдених частин секрету
	var foundSecrets []string

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			filePath := filepath.Join(dirName, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Не вдалося прочитати файл %s: %v", file.Name(), err)
				continue
			}

			// Розділяємо вміст на слова
			words := strings.Fields(string(content))
			
			// Перевіряємо, чи є в файлі наше секретне слово
			for i, word := range words {
				if strings.Contains(word, "secret") {
					// Якщо слово знайдене, додаємо його до списку
					foundSecrets = append(foundSecrets, words[i:]...)
				}
			}
		}
	}

	// Збираємо усі частини секрету в один рядок
	return strings.Join(foundSecrets, " ")
}

// cleanup видаляє створену папку
func cleanup() {
	fmt.Println("Видалення тимчасових файлів...")
	if err := os.RemoveAll(dirName); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Готово.")
}

func main() {
	// Створюємо файли
	createFiles()

	// Шукаємо секрет
	secret := findSecret()

	// Виводимо знайдений секрет
	fmt.Println("Знайдений секрет:")
	fmt.Println(secret)

	// Очищуємо директорію
	cleanup()
}
