package models

import (
	"go-web-skeleton/app/shared/database"
)

type User struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func GetUserById(id string) User {
	var u User
	database.Db.Get(&u, "SELECT * FROM public.users WHERE id = $1", id)
	return u
}
