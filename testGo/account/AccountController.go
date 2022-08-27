package account

import (
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", HealthCheck)
	http.HandleFunc("/accounts", GetJsonConvertedAccountDetails)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
