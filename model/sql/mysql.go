package sql

import (
	"crypto/md5"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetTx() (*sql.Tx, error) {
	db, err := sql.Open("mysql", "root:15094129@tcp(127.0.0.1:3306)/goweb?charset=utf8")
	if err != nil {
		fmt.Println("Error calling Open", err.Error())
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error calling Begin", err.Error())
		return nil, err
	}
	return tx, nil
}

func Encrypt(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}

func CloseTx(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		fmt.Println("Error calling Commit", err.Error())
		tx.Rollback()
	}
}
