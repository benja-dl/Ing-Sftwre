package utils

import (
	"database/sql"
	"fmt"
	"github.com/luccasniccolas/monitorWebApp/data"
)

func DeleteDomain(id int, db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM websites WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func ChangeDomainName(id int, newName string, db *sql.DB) error {
	err := db.QueryRow(`UPDATE websites SET name=$1 WHERE id=$2`, newName, id)

	if err != nil {
		return fmt.Errorf("error: modificación no exitosa")
	}

	return nil
}

func GetDomains(id int, db *sql.DB) ([]data.Domain, error) {
	// Ejecutar la consulta para obtener los dominios asociados con el ID proporcionado
	rows, err := db.Query(`SELECT id, owner_entity_id, name, url FROM websites WHERE owner_entity_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var domains []data.Domain

	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		var domain data.Domain
		if err := rows.Scan(&domain.ID, &domain.OwnerID, &domain.Name, &domain.URL); err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	// Comprobar si hubo algún error durante la iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return domains, nil
}
