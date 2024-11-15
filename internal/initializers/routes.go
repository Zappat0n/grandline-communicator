package initializers

import (
	"fmt"
	"net/http"

	"albarros/grandline/communicator/internal/handlers"
)

func Routes() {
	http.HandleFunc("/message", handlers.MessagesHandler)
	http.HandleFunc("/", handlers.HealthCheckHandler)

	fmt.Println("Starting server at port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
