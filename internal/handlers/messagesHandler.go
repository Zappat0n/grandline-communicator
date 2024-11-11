package handlers

import (
	"fmt"
	"net/http"

	"albarros/grandline/communicator/internal/models"
)

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createMessage(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	user := models.FindUserByPhone(phone)

	if user.ID != 0 {
		models.CreateMessage(r.FormValue("content"), user.ID)
		fmt.Fprintf(w, "Created a message")
	}
}
