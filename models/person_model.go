package models

import (
	"database/sql"
	"errors"
	"github.com/Abhishek-Mali-Simform/assessments/database"
	"strconv"
)

type PersonInfo struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}

func RetrievePerson(personID int) (*PersonInfo, error) {
	if personID <= 0 {
		return nil, errors.New("no personID passed to retrieve person")
	}
	personInfo := new(PersonInfo)
	query := `
		SELECT 
            p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code
        FROM 
            person p
        JOIN 
            phone ph ON p.id = ph.person_id
        JOIN 
            address_join aj ON p.id = aj.person_id
        JOIN 
            address a ON aj.address_id = a.id
        WHERE 
            p.id =` + strconv.Itoa(personID)
	err := database.DB.QueryRow(query).Scan(&personInfo.Name, &personInfo.PhoneNumber, &personInfo.City, &personInfo.State, &personInfo.Street1, &personInfo.Street2, &personInfo.ZipCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("person not found")
		}
		return nil, err
	}
	return personInfo, nil
}

func (personInfo *PersonInfo) Save() error {
	tx, err := database.DB.Begin()
	if err != nil {
		return errors.New("failed to begin transaction: " + err.Error())
	}

	var personID int
	err = tx.QueryRow("INSERT INTO person (name, age) VALUES ($1, 0) RETURNING id", personInfo.Name).Scan(&personID)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return errors.New("failed to rollback transaction: " + err.Error())
		}
		return errors.New("failed to insert person: " + err.Error())
	}

	_, err = tx.Exec("INSERT INTO phone (number, person_id) VALUES ($1, $2)", personInfo.PhoneNumber, personID)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return errors.New("failed to rollback transaction: " + err.Error())
		}
		return errors.New("failed to insert phone number: " + err.Error())
	}

	var addressID int
	err = tx.QueryRow("INSERT INTO address (city, state, street1, street2, zip_code) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		personInfo.City, personInfo.State, personInfo.Street1, personInfo.Street2, personInfo.ZipCode).Scan(&addressID)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return errors.New("failed to rollback transaction: " + err.Error())
		}
		return errors.New("failed to insert address: " + err.Error())
	}

	_, err = tx.Exec("INSERT INTO address_join (person_id, address_id) VALUES ($1, $2)", personID, addressID)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return errors.New("failed to rollback transaction: " + err.Error())
		}
		return errors.New("failed to insert address_join: " + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("failed to commit person info")
	}
	return nil
}
