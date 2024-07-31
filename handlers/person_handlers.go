package handlers

import (
	"fmt"
	"github.com/Abhishek-Mali-Simform/assessments/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func RetrievePersonInfo(ginContext *gin.Context) {
	personIDStr := ginContext.Param("person_id")
	if strings.TrimSpace(personIDStr) == "" {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Person ID required"})
		return
	}
	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		fmt.Println("Error Log: ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Person ID"})
		return
	}
	person, err := models.RetrievePerson(personID)
	if err != nil {
		fmt.Println("Error Log: ", err)
		ginContext.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ginContext.JSON(http.StatusOK, person)
}

func CreatePersonInfo(ginContext *gin.Context) {
	personInfo := new(models.PersonInfo)
	if err := ginContext.ShouldBindJSON(personInfo); err != nil {
		fmt.Println("Error Log: ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := personInfo.Save(); err != nil {
		fmt.Println("Error Log: ", err)
		ginContext.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "unable to save person's information"})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "person's information saved successfully"})
}
