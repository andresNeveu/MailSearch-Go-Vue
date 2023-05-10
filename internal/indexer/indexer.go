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
	"sync"
)

// Struct to format the email
type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

func main() {

	const MAX = 8

	var wg sync.WaitGroup

	// take first command line argument, path
	pathArg := os.Args[1]

	// get directory list
	innerPath := "maildir"
	dirPath := filepath.Join(pathArg, innerPath, "allen-p")
	files, err := os.ReadDir(dirPath)
	check(err)

	sem := make(chan int, MAX)
	for _, file := range files {
		wg.Add(1)
		sem <- 1
		subDirPath := filepath.Join(dirPath, file.Name())
		go func(subDirPath string) {
			records := make([]Message, 0)
			defer wg.Done()
			err := filepath.Walk(subDirPath, func(path string, info os.FileInfo, err error) error {
				check(err)
				if !info.IsDir() {
					record := readEmailFile(path)
					records = append(records, record)
				}
				return nil
			})
			check(err)
			postData(records)
			<-sem
		}(subDirPath)
	}

	wg.Wait()

	fmt.Println("Successfull")
}

// Basic error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Read the file to search Subject, To, From, Body of the email
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
	found := make([]bool, 4) // 0 = from, 1 = to, 2 = subject, 3 = x-file

	// get data
	for scanner.Scan() {
		if found[3] {
			textBody = textBody + scanner.Text() + "\u000a"
			continue
		}
		if !found[0] {
			if strings.Contains(scanner.Text(), "From:") {
				textFrom = scanner.Text()[6:]
				found[0] = true
				continue
			}
		}
		if !found[1] {
			if strings.Contains(scanner.Text(), "To:") {
				textTo = scanner.Text()[4:]
				found[1] = true
				continue
			}
		}
		if !found[2] {
			if strings.Contains(scanner.Text(), "Subject:") {
				textSubjet = scanner.Text()[8:]
				found[2] = true
				continue
			} else if found[1] {
				textTo = textTo + scanner.Text()
				continue
			}
		}
		if !found[3] {
			if strings.Contains(scanner.Text(), "X-FileName:") {
				found[3] = true
				continue
			}
		}
	}
	check(scanner.Err())

	m := Message{From: textFrom, To: textTo, Subject: textSubjet, Body: textBody}
	return m

}

// Post request to ZincSearch
func postData(records []Message) {

	// to JSON encode
	data, err := json.Marshal(records)
	check(err)

	dataString := string(data)
	base := `{ "index" : "mails", "records":`
	end := `}`
	dataBody := base + dataString + end

	//fmt.Println(dataBody)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", strings.NewReader(string(dataBody)))
	check(err)

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	check(err)

	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	check(err)

	log.Println(string(body))
}
