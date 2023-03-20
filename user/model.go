package user

import (
	"database/sql"
	"log"
)

type UserModel struct {
	conn *sql.DB
}

func (um *UserModel) SetConnection(db *sql.DB) {
	um.conn = db
}

func (um *UserModel) Login(email, password string) (User, error) {
	resUser := User{}
	res := um.conn.QueryRow("SELECT id, nama, email FROM user WHERE email = ? AND password = ?AND deleted_at IS NULL", email, password)

	if res.Err() != nil {
		log.Println("look up data error", res.Err().Error())
		return User{}, res.Err()
	}

	res.Scan(&resUser.ID, &resUser.Nama, &resUser.Email)

	return resUser, nil
}

func (um *UserModel) Register(newUser User) error {
	res, err := um.conn.Exec("INSERT INTO user (nama, email, password)", newUser.Nama, newUser.Email, newUser.Password)

	if err != nil {
		log.Println("Error insert ", err.Error())
		return err
	}

	resAff, err := res.RowsAffected()

	if err != nil {
		log.Println("Error after insert ", err.Error())
		return err
	}

	if resAff <= 0 {
		log.Println("Error inserting ", err.Error())
		return err
	}

	return nil
}

func (um *UserModel) Delete(userID User) error {
	res, err := um.conn.Exec("UPDATE user SET deleted_at = current_timestamp() WHERE id = ?", userID)

	if err != nil {
		log.Println("Error delete ", err.Error())
		return err
	}

	resAff, err := res.RowsAffected()

	if err != nil {
		log.Println("Error after delete ", err.Error())
		return err
	}

	if resAff <= 0 {
		log.Println("Error deleting ", err.Error())
		return err
	}

	return nil
}
