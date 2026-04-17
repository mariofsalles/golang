package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Load()
	cookies.Configure()
	utils.LoadTemplates([]string{"views/*.html", "views/templates/*.html"})
	r := router.Generate()

	fmt.Printf("Running webapp on port %d\n", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
