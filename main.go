package main

import (
	"net/http"

	router "menu-login/src/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// go startWeb()
	startAPI()
}

func startWeb() {
	pagina := http.FileServer(http.Dir("./src/index/"))
	http.Handle("/menu/", http.StripPrefix("/menu/", pagina))
	http.ListenAndServe(":9090", nil)
}

func startAPI() {
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.AppRoutes(app)

	app.Run()
}
