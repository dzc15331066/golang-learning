package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

// UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session := myengine.NewSession()
	defer session.Close()
	err := session.Begin()
	checkErr(err)
	_, err = session.Insert(u)
	if err != nil {
		session.Rollback()
		return err
	} else {
		session.Commit()
		return nil
	}
}

// FindAll.
func (*UserInfoAtomicService) FindAll() []UserInfo {
	all := make([]UserInfo, 0)
	err := myengine.Find(&all)
	checkErr(err)
	return all
}

// FindByID .

func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	u := &UserInfo{}
	myengine.Id(id).Get(u)
	return u
}
