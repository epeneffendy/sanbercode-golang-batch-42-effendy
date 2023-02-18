package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/database"
	"practice/repository"
	"practice/stucts"
	"strconv"
)

func GetAllPerson(c *gin.Context) {
	var (
		result gin.H
	)

	persons, err := repository.GetAllPerson(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": persons,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person stucts.Person

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success INsert Person",
	})
}

func UpdatePerson(c *gin.Context) {
	var person stucts.Person
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}

	person.ID = int64(id)

	err = repository.UpdatePerson(database.DBConnection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Person",
	})
}

func DeletePerson(c *gin.Context) {
	var person stucts.Person
	id, err := strconv.Atoi(c.Param("id"))

	err = repository.DeletePerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Person",
	})
}
