package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/asdine/storm/q"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserInfo : Detailed info about a user.
type UserInfo struct {
	Pk             int    `storm:"id,increment,index"`
	Email          string `storm:"unique"`
	DisplayName    string `storm:"unique"`
	Password       []byte
	Token          [16]byte
	TokenExpires   time.Time
	AccountCreated time.Time
	FailedAttempts int
	Locked         bool
	LockTime       time.Time
	Active         bool
}

type UserPermission struct {
	Pk         int `storm:"id,increment,index"`
	Permission int
	User       string
	Active     bool
}

type Permission struct {
	Pk     int `storm:"id,increment,index"`
	Name   string
	Active bool
}

// Identity : A user's ID
type Identity struct {
	Email       string    `json:"email"`
	DisplayName string    `json:"displayname"`
	Token       [16]byte  `json:"token"`
	Expires     time.Time `json:"expires"`
	Password    string    `json:"password"`
}

// Login logs a user into the database
func Login(email string, pw string) (Identity, error) {
	var user UserInfo

	err := db.One("Email", email, &user)
	if err != nil {
		return Identity{}, err
	}

	if user.Locked {
		if time.Now().After(user.LockTime) {
			user.Locked = false
		} else {
			return Identity{}, fmt.Errorf("this account is locked")
		}
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(pw))

	if err != nil {
		user.FailedAttempts++
		if user.FailedAttempts >= 7 {
			user.Locked = true
			user.LockTime = time.Now().Add(time.Hour * time.Duration(24))
		}
		return Identity{}, err
	}

	uid := uuid.NewV4()

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
	var user UserInfo

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
	var perm UserPermission

	err := db.Select(q.And(q.Eq("User", email), q.Eq("Pk", permission))).First(&perm)
	if err != nil {
		return err
	}

	return nil
}

// RegisterAccount : Creates a new account.
func RegisterAccount(id *Identity) error {
	var user UserInfo

	uid := uuid.NewV4()
	user.Token, id.Token = uid, uid

	expires := time.Now().Add(time.Hour * 120)
	user.TokenExpires, id.Expires = expires, expires

	pwHash, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(id.Password)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	id.Email = strings.ToLower(strings.TrimSpace(id.Email))

	user.Password = pwHash
	user.Email = id.Email
	user.DisplayName = id.DisplayName
	user.Active = true
	user.FailedAttempts = 0
	user.Locked = false
	user.LockTime = time.Time{}
	user.AccountCreated = time.Now()

	err = db.Save(&user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAccount() error {

}

func UpdateUserPermission() error {

}

func UpdatePermission() error {

}
