package controllers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func test() {
	//users := map[string][]byte {
	//"edu": bcrypt.GenerateFromPassword([]byte("123"), 10),
	//}
	fmt.Println(bcrypt.GenerateFromPassword([]byte("123"), 10))
}
