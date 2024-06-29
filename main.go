package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	go startWeb()
	startAPI()
}

func createHash() {
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
}

func startWeb() {
	pagina := http.FileServer(http.Dir("./src/index/"))
	http.Handle("/menu/", http.StripPrefix("/menu/", pagina))
	http.ListenAndServe(":9090", nil)
}

func startAPI() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/login", func(ctx *gin.Context) {
		jsonData, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			fmt.Println("erro ao recebe json")
		}
		jsonObj, err := json.Marshal(jsonData)
		if err != nil {
			fmt.Println(err)
		}

		var users user
		errr := json.Unmarshal(jsonObj, &users)
		if errr != nil {
			fmt.Println(errr)
		}

		fmt.Println(users)
	})

	router.Run()
}
