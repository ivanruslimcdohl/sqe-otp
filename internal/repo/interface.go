package repo

import (
	repodb "github.com/ivanruslimcdohl/sqe-otp/internal/repo/db"
)

type Repo struct {
	DB repodb.DB
}
