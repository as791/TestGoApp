package account

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func GetAccountDetails() []AccountDetail {
	accountDetails := []AccountDetail{
		GetAccountDetail(),
		GetAccountDetail(),
		GetAccountDetail(),
	}
	return accountDetails
}

func GetJsonConvertedAccountDetails(w http.ResponseWriter, r *http.Request) {
	log.Print(GetAccountDetails())
	var flowId = uuid.New().String()
	http.SetCookie(w, &http.Cookie{Name: "accountDetailsCokkie", Value: uuid.New().String()})
	w.Header().Add("flowId", flowId)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(GetAccountDetails())
	if err != nil {
		log.Panic("[ERROR] : Error while getting account details for flow Id: ", flowId)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Service Is UP!!")
	if err != nil {
		return
	}
	log.Print("Endpoint Hit: /")
}
