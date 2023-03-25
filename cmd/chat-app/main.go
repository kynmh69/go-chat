package main

import (
	"fmt"
	"go-chat/controller"
	"go-chat/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	chatRoom := model.ChatRoom{
		Messages: make([]model.Message, 0),
	}

	router.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("GET")
			controller.GetChatRoom(&chatRoom, w)
		case "POST":
			log.Println("POST")
			controller.PostChatRoom(&chatRoom, r, w)
		}
	})

	port := 8080

	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf(":%d", port),
			router,
		),
	)
}
