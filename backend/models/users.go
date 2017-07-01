package models

import (
	"fmt"
	"time"

	"github.com/asdine/storm/q"
	"github.com/satori/go.uuid"
)

type userInfo struct {
	Pk             int    `storm:"id,increment,index"`
	Email          string `storm:"unique"`
	DisplayName    string `storm:"unique"`
	Password       [60]byte
	Token          [16]byte
	TokenExpires   time.Time
	AccountCreated time.Time
	Active         bool
}

type userPermission struct {
	Pk   int `storm:"id,increment,index"`
	User int
}

type permission struct {
	Pk     int `storm:"id,increment,index"`
	Name   string
	Active bool
}

// Identity : A user's ID
type Identity struct {
	Email       string
	DisplayName string
	Token       [16]byte
	Expires     time.Time
}

// Login logs a user into the database
func Login(email string, hash []byte) (Identity, error) {

	var user userInfo

	err := db.One("Email", email, &user)
	if err != nil {
		return Identity{}, err
	}

	uid := uuid.NewV4()
	uid = uuid.NewV5(uid, email)

	user.Token = uid
	user.TokenExpires = time.Now().Add(time.Hour * 120)

	err = db.Update(&user)
	if err != nil {
		return Identity{}, err
	}

	return Identity{Email: user.Email, DisplayName: user.DisplayName, Token: uid, Expires: user.TokenExpires}, nil
}

// CheckToken : Checks if a user's session is valid
func CheckToken(frontid *Identity) error {
	id := *frontid
	var user userInfo

	err := db.One("Email", id.Email, &user)
	if err != nil {
		return err
	}

	if id.Token != user.Token || time.Now().After(user.TokenExpires) || !id.Expires.Equal(user.TokenExpires) {
		return fmt.Errorf("Valid Token Required")
	}

	id.Expires = time.Now().Add(time.Hour * 120)
	err = db.Update(&id)
	if err != nil {
		return err
	}

	return nil
}

// CheckPermission : Checks to see if a user is permitted to do an action
func CheckPermission(email string, permission int) error {
	var user userInfo
	var perm userPermission

	err := db.One("Email", email, &user)
	if err != nil {
		return err
	}

	err = db.Select(q.And(q.Eq("User", user.Pk), q.Eq("Pk", permission))).First(&perm)
	if err != nil {
		return err
	}

	return nil
}
