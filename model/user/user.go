package user

import (
	"github.com/chitchat/model/sql"
	"fmt"
)

type User struct {
	UserID   int
	UserName string
	Email    string
	PassWord string
}

//CreateUser ...
func CreateUser(user *User) (int, error) {
	tx, err := sql.GetTx()
	if err != nil {
		fmt.Println("CreateUser GetTx:", err.Error())
		return -1, err
	}
	defer sql.CloseTx(tx)
	sqlInsert := fmt.Sprintf(`INSERT INTO users (username,email, password) VALUES (?,?,?);`)
	result, err := tx.Exec(sqlInsert, user.UserName, user.Email, user.PassWord)

	if err != nil {
		fmt.Println("CreateUser QueryRow:", err.Error())
		return -1, err
	}

	id, err := result.LastInsertId()
	return int(id), nil
}

// DeleteUser ...
func DeleteUser(user *User) (err error) {
	tx, err := sql.GetTx()
	if err != nil {
		fmt.Println("CreateUser GetTx:", err.Error())
		return err
	}
	defer sql.CloseTx(tx)
	statement := fmt.Sprintf(`delete from users where userid =?;`)

	_, err = tx.Exec(statement, user.UserID)
	if err != nil {
		fmt.Println("DeleteUser Exec:", err.Error())
		return err
	}
	return nil
}

//UpdateUser ...
func UpdateUser(user *User) (err error) {
	tx, err := sql.GetTx()
	if err != nil {
		fmt.Println("CreateUser GetTx:", err.Error())
		return err
	}
	defer sql.CloseTx(tx)
	statement := fmt.Sprintf(`update users set username = ?, email = ? where userid = ?;`)

	_, err = tx.Exec(statement, user.UserID, user.UserName, user.Email)
	if err != nil {
		fmt.Println("Update Exec:", err.Error())
		return err
	}
	return nil
}

//SelectUser ...
func SelectUser(email string) (User, error) {
	user := User{}
	tx, err := sql.GetTx()
	if err != nil {
		fmt.Println("CreateUser GetTx:", err.Error())
		return user, err
	}
	defer sql.CloseTx(tx)
	statement := fmt.Sprintf(`SELECT userid, username, email, password FROM users WHERE email = ?;`)

	err = tx.QueryRow(statement, email).
		Scan(&user.UserID, &user.UserName, &user.Email, &user.PassWord)
	if err != nil {
		fmt.Println("SelectUser :", err.Error())
		return user, err
	}
	return user, nil
}
