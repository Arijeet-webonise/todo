// Package models contains the types for schema 'public'.
package models

import (
	"crypto/sha1"
	"database/sql"
)

func (serviceImpl *UserServiceImpl) Register(u *User) (e error) {
	h := sha1.New()
	_, err := h.Write([]byte(u.Password.String))

	has := h.Sum(nil)

	u.Password = sql.NullString{String: string(has), Valid: u.Password.Valid}

	if err != nil {
		return err
	}

	err = serviceImpl.InsertUser(u)

	if err != nil {
		return err
	}

	return nil
}
