package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	base64String := base64.StdEncoding.EncodeToString(key)
// 	fmt.Printf("base64String: %s\n", base64String)
// }

func main() {
	config.LoadingConfigs()
	r := router.Generate()

	fmt.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%d`, config.Port), r))

}
