package models

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int
	Username string
	Profile  string
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(u User) (i int64) {
	o := orm.NewOrm()
	var user User
	if u.Username != "" {
		user.Username = u.Username
	}
	if u.Profile != "" {
		user.Profile = u.Profile
	}
	id, err := o.Insert(&user)

	if err == nil {
		return id
	}
	return 0
}

func GetUser(id int) (u *User, err error) {
	o := orm.NewOrm()
	var user User
	user = User{Id: id}

	err = o.Read(&user)

	if err == orm.ErrNoRows {
		return nil, errors.New("User not exists")
	} else if err == orm.ErrMissPK {
		return nil, errors.New("User not exists")
	} else {
		return &user, nil
	}
}

func GetAllUsers(offset int, limit int) (uu []*User, err error) {
	o := orm.NewOrm()
	var users []*User
	qs := o.QueryTable("user")
	//set limit
	num, _ := qs.Limit(limit, offset).All(&users)
	if num < 1 {
		return []*User{}, errors.New("User not exists")
	}
	return users, nil
}

func GetAllUsersCount() (c int64, err error) {
	o := orm.NewOrm()

	cnt, err := o.QueryTable("user").Count()

	return cnt, err
}

func UpdateUser(id int, uu *User) (a *User, err error) {
	o := orm.NewOrm()
	var user User
	user = User{Id: id}
	err = o.Read(&user)

	if err == orm.ErrNoRows {
		return nil, errors.New("User not exists")
	} else if err == orm.ErrMissPK {
		return nil, errors.New("User not exists")
	} else {
		if uu.Username != "" {
			user.Username = uu.Username
		}
		if uu.Profile != "" {
			user.Profile = uu.Profile
		}
		o.Update(&user)
		return &user, nil
	}
}

func DeleteUser(id int) (s string, err error) {
	o := orm.NewOrm()
	var user User
	user = User{Id: id}
	err = o.Read(&user)

	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return "", errors.New("User not exists")
	} else {
		o.Delete(&user)
		return "ok", nil
	}
}
