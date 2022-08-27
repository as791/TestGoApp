package account

import "github.com/google/uuid"

type AccountDetail struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewAccountDetail(userName string, password string) *AccountDetail {
	return &AccountDetail{Id: uuid.New().String(), UserName: userName, Password: password}
}
