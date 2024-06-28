package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	users := map[string][]uint8{}

	hash, err := bcrypt.GenerateFromPassword([]byte("123"), 10)
	if err != nil {
		fmt.Println(err)
	}

	users["edu"] = hash // salva o novo usuario

	user := "edu"  // json fake
	pass := "0123" // json fake

	err = bcrypt.CompareHashAndPassword(users[user], []byte(pass))
	if err != nil {
		fmt.Println(err)
	}

	g := gin.Default()
	g.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, 10)
	})
	g.Run()
}
