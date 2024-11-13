package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/luccasniccolas/monitorWebApp/data"
)

func CheckEmailExist(email string, db *sql.DB) bool {
	row := db.QueryRow("SELECT email FROM users WHERE email=$1", email)

	if errors.Is(row.Scan(), sql.ErrNoRows) {
		return false
	}

	return true
}

func GetUserByEmail(email string, db *sql.DB) (*data.User, error) {
	user := new(data.User)

	row := db.QueryRow("SELECT id, entity_id, first_name, last_name, email, password_hash FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.EntityID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err
	}

	return user, nil
}

func GetUserByID(id int, db *sql.DB) (*data.User, error) {
	user := new(data.User)

	row := db.QueryRow("SELECT id, entity_id, first_name, last_name, email, password_hash FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.EntityID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err
	}

	return user, nil
}

func GetAccountType(id int, db *sql.DB) (string, error) {
	var accType string
	err := db.QueryRow("SELECT entity_type FROM entities WHERE id=$1", id).Scan(&accType)
	if err != nil {
		return "", fmt.Errorf("problemas en la base de datos")
	}

	return accType, nil
}
