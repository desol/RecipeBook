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
	Token          []byte
	TokenExpires   int64
	AccountCreated time.Time
	FailedAttempts int
	Locked         bool
	LockTime       time.Time
	Active         bool
}

// UserPermission : A user's mapping to a site permission
type UserPermission struct {
	Pk         int `storm:"id,increment,index"`
	Permission int
	User       int
	Active     bool
}

// Permission : A site permission
type Permission struct {
	Pk     int `storm:"id,increment,index"`
	Name   string
	Active bool
}

// Identity : A user's ID
type Identity struct {
	Pk          int    `json:"pk"`
	Email       string `json:"email"`
	DisplayName string `json:"displayname"`
	Token       string `json:"token"`
	Expires     int64  `json:"expires"`
	Password    string `json:"password"`
}

// Login logs a user into the database
func Login(email string, pw string) (*Identity, error) {
	var user UserInfo

	err := db.One("Email", email, &user)
	if err != nil {
		return new(Identity), err
	}

	if user.Locked {
		if time.Now().After(user.LockTime) {
			user.Locked = false
		} else {
			return new(Identity), fmt.Errorf("this account is locked")
		}
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(pw))

	if err != nil {
		user.FailedAttempts++
		if user.FailedAttempts >= 7 {
			user.Locked = true
			user.LockTime = time.Now().Add(time.Hour * time.Duration(24))
		}
		return new(Identity), err
	}

	uid := uuid.NewV4()

	user.Token, err = uid.MarshalText()
	if err != nil {
		return new(Identity), err
	}
	user.TokenExpires = time.Now().Add(time.Hour * 120).UTC().Unix()

	err = db.Update(&user)
	if err != nil {
		return new(Identity), err
	}

	return &Identity{Pk: user.Pk, Email: user.Email, DisplayName: user.DisplayName, Token: string(user.Token), Expires: user.TokenExpires}, nil
}

// CheckToken : Checks if a user's session is valid
func CheckToken(frontid *Identity) error {
	id := *frontid
	var user UserInfo

	err := db.One("Pk", id.Pk, &user)
	if err != nil {
		return err
	}

	if id.Token != string(user.Token) || time.Now().UTC().Unix() > user.TokenExpires || id.Expires != user.TokenExpires {
		return fmt.Errorf("Valid Token Required")
	}

	id.Expires = time.Now().Add(time.Hour * 120).UTC().Unix()
	user.TokenExpires = id.Expires
	err = db.Update(&user)
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
	var err error

	uid := uuid.NewV4()
	user.Token, err = uid.MarshalText()
	if err != nil {
		return err
	}
	id.Token = string(user.Token)

	expires := time.Now().Add(time.Hour * 120).UTC().Unix()
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

// UpdateAccount : Updates a user's account
func UpdateAccount(user *UserInfo) error {
	err := db.Update(user)
	return err
}

// AddUserPermission : Creates a user and permission relationship
func AddUserPermission(user, perm int) error {
	up := UserPermission{User: user, Permission: perm, Active: true}
	err := db.Save(&up)
	return err
}

// UpdateUserPermission : Updates a user's permission
func UpdateUserPermission(up *UserPermission) error {
	err := db.Update(up)
	return err
}

// UpdatePermission Updates a permission
func UpdatePermission(perm *Permission) error {
	err := db.Update(perm)
	return err
}
