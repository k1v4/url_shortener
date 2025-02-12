package DataBase

import "errors"

var (
	ErrUrlExists   = errors.New("url already exists")
	ErrUrlNotFound = errors.New("url not found")
)
