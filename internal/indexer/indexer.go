package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 512*256)

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
				textSubjet = scanner.Text()[8:]
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

func postData(records []Message) {

	// to JSON encode
	data, err := json.Marshal(records)
	check(err)

	dataString := string(data)
	base := `{ "index" : "mails", "records": %s}`
	dataBody := fmt.Sprintf(base, dataString)

	//fmt.Println(dataBody)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", strings.NewReader(string(dataBody)))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
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
			//fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		} else {
			record := readEmailFile(path)
			records = append(records, record)
		}

		return nil
	})
	check(err)
	postData(records)

}
