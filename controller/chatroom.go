package controller

import (
	"encoding/json"
	"go-chat/model"
	"net/http"
)

func GetChatRoom(chatRoom *model.ChatRoom, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(chatRoom)
}

func PostChatRoom(chatRoom *model.ChatRoom, r *http.Request, w http.ResponseWriter) {
	var message model.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chatRoom.Messages = append(chatRoom.Messages, message)

	w.WriteHeader(http.StatusCreated)
}
