package db

import "errors"

const (
	DB_NAME    = "0-directory-db"
	DB_SESSION = "0-directory/dbconnection"
)

var (
	//ErrDuplicate indicates the entry is invalid because of a primary key violation
	ErrDuplicate = errors.New("Duplicate")
)
