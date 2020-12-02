package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type info struct {
	db  *sql.DB
	err error
}

type Result struct {
	studentId   string
	studentName string
}

func InitDB(dbName, dbSrcName string) (db *sql.DB, err error) {
	db, err = sql.Open(dbName, dbSrcName)
	if err != nil {
		return nil, errors.Wrap(err, "Scan error")
	}
	return
}

func DoWork(db *sql.DB, sql string) (int, error) {
	var info Result
	errScan := db.QueryRow(sql).Scan(&info.studentId, &info.studentName)
	if errScan != nil {
		return -1, errors.Wrap(errScan, "Scan error")
	}

	defer db.Close()

	return 1, nil
}
