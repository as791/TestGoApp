package account

import (
	"testing"
)

func TestGetAccountDetails(t *testing.T) {
	expectedAccountDetails := []AccountDetail{
		*NewAccountDetail("test@gmail.com", "1234"),
		*NewAccountDetail("test@gmail.com", "1234"),
		*NewAccountDetail("test@gmail.com", "1234"),
	}

	actualAccountDetails := GetAccountDetails()

	if len(actualAccountDetails) != len(expectedAccountDetails) {
		t.Log("length of excpected account details and actual account details don't match")
		t.Fail()
	}

	for i := 0; i < len(actualAccountDetails); i++ {
		actualDetail := actualAccountDetails[i]
		expectedDetail := expectedAccountDetails[i]
		if actualDetail.UserName != expectedDetail.UserName {
			t.Log("UserName for expected and actual account detail don't match for index i:", i)
			t.Fail()
		}

		if actualDetail.Password != expectedDetail.Password {
			t.Log("Password for expected and actual account detail don't match for index i:", i)
			t.Fail()
		}
	}
}
