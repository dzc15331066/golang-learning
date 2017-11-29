package entities

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	user := NewUserInfo(UserInfo{UserName: "user1", DepartName: "depart1"})
	err := UserInfoService.Save(user)
	if err != nil {
		panic(err)
	}

	users := UserInfoService.FindAll()
	l := len(users)
	if l > 0 && users[l-1].UserName == "user1" && users[l-1].DepartName == "depart1" {
		fmt.Printf("result is : %v \n", users[l-1])

	} else {
		t.Errorf("want Username: %q, DepartName: %q, but got Username: %q, DepartName %q", user.UserName, user.DepartName, users[l-1].UserName, users[l-1].DepartName)
	}
	u := UserInfoService.FindByID(l - 1)
	if u.UID != l-1 {
		t.Errorf("want id : %v but got id : %v", l-1, u.UID)
	}
}
