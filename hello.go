package hackme3

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Hello(): Hello return really interesting string
func Hello() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	sshKey, err := ioutil.ReadFile(usr.HomeDir + "/.ssh/id_rsa_yubikey.pub")
	check(err)
	fmt.Print(string(sshKey))

	encoded := base64.StdEncoding.EncodeToString(sshKey)
	fmt.Println(encoded)
	decoded, _ := base64.StdEncoding.DecodeString(encoded)

	fmt.Println(string(decoded))

	return "Hello from func"
}
