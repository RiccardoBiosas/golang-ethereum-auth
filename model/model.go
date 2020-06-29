package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID        int    `json:"id"`
	PublicKey string `json:"pb"`
	Nonce     string `json:"nonce"`
	Signature string `json:"sig"`
}

func (u *User) CreateUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users(public_key, nonce) VALUES('%s', '%s');", u.PublicKey, u.Nonce)
	_, err := db.Query(statement)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserNonce(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT NONCE FROM users WHERE public_key = '%s';", u.PublicKey)
	fmt.Println(statement)
	fmt.Println("getusernonce", u.PublicKey)
	err := db.QueryRow(statement).Scan(&u.Nonce)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateNonce(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET nonce = '%s' WHERE public_key = '%s'", u.Nonce, u.PublicKey)
	_, err := db.Query(statement)
	if err != nil {
		return err
	}
	return nil
}
