package model

import (
	//"github.com/jmoiron/sqlx"
	"time"
)

type Employee struct {
    EId int `db:"id"` 
    Name  string `db:"name"`
	Email  string `db:"email"`
	Password string `db:"password"`
	Status string `db:"status"`
	DOB time.Time `db:"dob"`
}

