package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	passUrl  = "http://pstgu.yss.su/iu9/networks/let1/getkey.php"
	quoteUrl = "http://pstgu.yss.su/iu9/networks/let1/send_from_go.php"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getPassword(client *http.Client, name string) string {
	mdName := GetMD5Hash(name)
	req, _ := http.NewRequest("POST", passUrl, nil)
	vals := req.URL.Query()
	vals.Add("hash", mdName)
	req.URL.RawQuery = vals.Encode()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(body), " ")[1]
}

func getQuote(client *http.Client, pass string) {
	req, _ := http.NewRequest("POST", quoteUrl, nil)
	vals := req.URL.Query()
	vals.Add("subject", "let1 ИУ9-32Б Павлов Иван")
	vals.Add("fio", "Павлов Иван")
	vals.Add("pass", pass)
	req.URL.RawQuery = vals.Encode()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	client := http.DefaultClient
	pass := getPassword(client, "Павлов Иван")
	getQuote(client, pass)
}
