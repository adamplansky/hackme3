package hackme3

import (
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

	dat, err := ioutil.ReadFile(usr.HomeDir + "/.ssh/id_rsa_yubikey.pub")
	check(err)
	fmt.Print(string(dat))

	return "Hello from func"
}

// Hello2(): Hello2 return really interesting string
func Hello2() string {
	return "Hello2 from func"
}

// Hello3(): Hello3 return really interesting string
func Hello3() string {
	return "Hello3 from func"
}
