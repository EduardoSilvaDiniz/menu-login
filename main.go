package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	go startWeb()
	startAPI()
}

//func addUser(user, password string) {
//	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
//	if err != nil {
//		panic(hash)
//	}
//
//	users[user] = hash
//}
//
//func passwordCheck(user, password string) bool {
//	err := bcrypt.CompareHashAndPassword(users[user], []byte(password))
//	if err != nil {
//		return false
//	}
//	return true
//}

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

	router.POST("/post", func(ctx *gin.Context) {
		var usuario user
		ctx.BindJSON(&usuario)
		ctx.JSON(200, usuario)
	})

	router.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
