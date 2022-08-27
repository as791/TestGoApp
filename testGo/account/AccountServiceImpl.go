package account

import "github.com/google/uuid"

func (accountDetail *AccountDetail) CreateAccountDetail(userName string, password string) {
	accountDetail.Id = uuid.New().String()
	accountDetail.UserName = userName
	accountDetail.Password = password
}

func GetAccountDetail() AccountDetail {
	return *NewAccountDetail("test@gmail.com", "1234")
}
