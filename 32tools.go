package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/kwartel/ttpapi"
)

type userConfiguration struct {
	URL      string
	Username string
	Password string
}

func getUserConf() (userConfiguration, error) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	userConf := userConfiguration{}
	err := decoder.Decode(&userConf)

	return userConf, err
}

func main() {
	userConf, err := getUserConf()
	if err != nil {
		fmt.Println("error:", err)
	}

	ttp, err := ttpapi.NewTtpAPI(userConf.URL)
	if err != nil {
		log.Fatal(err)
	}
	err = ttp.Login(userConf.Username, userConf.Password)
	if err != nil {
		log.Fatal(err)
	}
	mailboxParams := url.Values{}
	mailboxParams.Set("type", "sentbox")
	mailbox, err := ttp.GetMailbox(mailboxParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mailbox)
}
