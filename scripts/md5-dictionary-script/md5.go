package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
)

var (
	lowerCaseChars                = "abcdefghijklmnopqrstuvwxyzç"
	upperCaseChars                = "ABCDEFGHIJKLMNOPQRSTUVWXYZÇ"
	specialChars                  = "!@#$%^&*()-_+=<>?/,.:;{}[]|"
	numericChars                  = "0123456789"
	allPasswordPossibleCharacters = lowerCaseChars + upperCaseChars + specialChars + numericChars
	dictionaryMD5Passwords        = make(map[int][]Password)
	mutex                         sync.Mutex
)

type Password struct {
	PlainText string `json:"plainText"`
	MD5       string `json:"md5"`
}

func generateCombinationsWithRepetition(inputString string, combinationLength int, id int) {
	generateHelper("", inputString, combinationLength, id)
}

func generateHelper(currentPassword, inputString string, combinationLength, id int) {
	if len(currentPassword) == combinationLength {
		hash := md5.Sum([]byte(currentPassword))
		md5String := hex.EncodeToString(hash[:])

		password := Password{
			PlainText: currentPassword,
			MD5:       md5String,
		}

		mutex.Lock()
		dictionaryMD5Passwords[combinationLength] = append(dictionaryMD5Passwords[combinationLength], password)

		fileName := fmt.Sprintf("./md5-dictionary-script/jsons/passwords_length_%d_worker_%d.json", combinationLength, id)
		data, err := json.MarshalIndent(dictionaryMD5Passwords[combinationLength], "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal JSON: %v", err)
		}
		err = ioutil.WriteFile(fileName, data, 0644)
		if err != nil {
			log.Fatalf("Failed to write file: %v", err)
		}
		mutex.Unlock()

		fmt.Printf("Worker ID %d => Created password: %s\n", id, currentPassword)

		return
	}

	for i := 0; i < len(inputString); i++ {
		char := inputString[i]
		newCombination := currentPassword + string(char)
		generateHelper(newCombination, inputString, combinationLength, id)
	}
}

func main() {
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	passwordLengths := []int{6, 7, 8, 9, 10, 11, 12, 13}

	if err := os.MkdirAll("./md5-dictionary-script/jsons", os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	for i := 0; i < numWorkers && i < len(passwordLengths); i++ {
		wg.Add(1)
		go func(id, length int) {
			defer wg.Done()
			fmt.Printf("Worker ID %d => generating passwords of length %d\n", id, length)
			generateCombinationsWithRepetition(allPasswordPossibleCharacters, length, id)
		}(i, passwordLengths[i])
	}

	wg.Wait()
}
