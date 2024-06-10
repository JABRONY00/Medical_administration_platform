package query

import (
	"context"
	"fmt"
	"strings"

	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/models"
	_ "github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/models"
)

func InsertClient(client *models.ClientWithPassword) error {
	_, err := DB.Exec(context.Background(),
		"INSERT INTO clients VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		client.ID,
		client.FirstName,
		client.LastName,
		client.Phone,
		client.Email,
		client.PasswordHash,
		client.BirthDate,
		client.Gender,
		client.Address,
	)
	return err
}

func GetClient(clientID *string) (*models.ClientInfo, error) {
	row := DB.QueryRow(context.Background(), "SELECT * FROM clients WHERE id = $1", clientID)
	var client models.ClientInfo
	err := row.Scan(&client.ID,
		&client.FirstName,
		&client.LastName,
		&client.Phone,
		&client.Email,
		&client.BirthDate,
		&client.Gender,
		&client.Address,
	)
	return &client, err
}

func GetAllClients() ([]models.ClientInfo, error) {
	rows, err := DB.Query(context.Background(), "SELECT * FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.ClientInfo
	for rows.Next() {
		var client models.ClientInfo
		err := rows.Scan(&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Phone,
			&client.Email,
			&client.BirthDate,
			&client.Gender,
			&client.Address,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func GetClientsByIDs(ids []string) ([]models.ClientInfo, error) {
	var sqlScript strings.Builder
	sqlScript.WriteString("SELECT * FROM clients WHERE ")

	for i := 0; i < len(ids); i++ {
		sqlScript.WriteString(fmt.Sprintf("id = '%s', ", ids[i]))

	}

	rows, err := DB.Query(context.Background(), sqlScript.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.ClientInfo
	for rows.Next() {
		var client models.ClientInfo
		err := rows.Scan(&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Phone,
			&client.Email,
			&client.BirthDate,
			&client.Gender,
			&client.Address,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func DeleteClient(id *string) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM clients WHERE id=$1", id)
	return err
}

func UpdateClient(client *models.ClientInfo) error {
	_, err := DB.Exec(context.Background(), "UPDATE clients SET first_name = $1, last_name = $2, phone = $3, email = $4, birth_date = $5, gender = $6, address = $7 WHERE id =$8",
		client.FirstName,
		client.LastName,
		client.Phone,
		client.Email,
		client.BirthDate,
		client.Gender,
		client.Address,
		client.ID,
	)
	return err
}
