package controller

import (
	"fmt"
	"log"
	"net/http"
	"sample/api/model"
	"sample/helper"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func Test(app *gin.Context) {
	app.JSON(http.StatusOK, gin.H{
		"message": "yes server is working",
	})
}

func Create(app *gin.Context) {
	var newPersonalInfo model.Contact
	var phoneNumberFilter []string
	if err := app.BindJSON(&newPersonalInfo); err != nil {
		return
	}
	for _, element := range newPersonalInfo.PhoneNumbers {
		if helper.CheckIfE164OrAussieNum(element) {
			phoneNumberFilter = append(phoneNumberFilter, element)
		}
	}
	newPersonalInfo.PhoneNumbers = phoneNumberFilter
	model.AddPersonalInfo(newPersonalInfo)
	app.IndentedJSON(http.StatusCreated, newPersonalInfo)
}

func Get(app *gin.Context) {
	var rows = model.GetPersonalInfo()

	var newPersonalInfos []model.Contact
	var personalInfo model.Contact
	for rows.Next() {
		if err := rows.Scan(&personalInfo.FullName, &personalInfo.Email, pq.Array(&personalInfo.PhoneNumbers)); err != nil {
			log.Fatal(err)
		}

		newPersonalInfos = append(newPersonalInfos, personalInfo)
	}
	fmt.Println(newPersonalInfos)
	app.JSON(http.StatusOK, gin.H{
		"message": newPersonalInfos,
	})

}

func Delete(app *gin.Context) {

}
