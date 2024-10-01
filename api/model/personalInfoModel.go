package model

import (
	"database/sql"
	"log"
	"sample/utils"

	"github.com/lib/pq"
)

type Contact struct {
	FullName     string   `json:"full_name" binding: "required"`
	Email        string   `json:"email" binding: "required"`
	PhoneNumbers []string `json:"phone_numbers" "binding:"required"`
}

func AddPersonalInfo(personalInfo Contact) {
	_, err := utils.Database.Query("INSERT into contact.contact_tbl (full_name,email,phone_number) VALUES ($1,$2,$3)", personalInfo.FullName, personalInfo.Email, pq.Array(personalInfo.PhoneNumbers))
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}
}

func GetPersonalInfo() *sql.Rows {
	data, err := utils.Database.Query("SELECT * FROM contact.contact_tbl")
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}

	return data
}

func DeletePersonalInfo(id string) {

}
