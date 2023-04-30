package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readEmailFile(path string) Message {

	// open file
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	// scan file
	scanner := bufio.NewScanner(f)

	// declare storages for data
	textFrom := ""
	textTo := ""
	textSubjet := ""
	textBody := ""
	findBody := false // to find body (after last field)

	// get data
	for scanner.Scan() {
		if findBody {
			textBody = textBody + scanner.Text()
		}
		if strings.Contains(scanner.Text(), "Subject:") {
			if textSubjet == "" {
				textSubjet = scanner.Text()[9:]
			}
		}
		if strings.Contains(scanner.Text(), "To:") {
			if textTo == "" {
				textTo = scanner.Text()[4:]
			}
		}
		if strings.Contains(scanner.Text(), "From:") {
			if textFrom == "" {
				textFrom = scanner.Text()[6:]
			}
		}
		if strings.Contains(scanner.Text(), "X-FileName:") {
			findBody = true
		}

	}
	check(scanner.Err())
	m := Message{From: textFrom, To: textTo, Subject: textSubjet, Body: textBody}
	return m

}

func main() {

	records := make([]Message, 0)
	// take first command line argument, path
	pathArg := os.Args[1]

	// get directory list
	innerPath := "maildir"
	dirPath := filepath.Join(pathArg, innerPath)
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		check(err)
		if info.IsDir() {
			fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		} else {
			record := readEmailFile(path)
			records = append(records, record)
		}

		return nil
	})
	check(err)

	// to JSON encode
	data, err := json.Marshal(records)
	check(err)
	fmt.Println(string(data))
}
