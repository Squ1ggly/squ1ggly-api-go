package controllers

import (
	"net/http"
	"squ1ggly/squ1ggly-api-go/functions"

	"github.com/gin-gonic/gin"
)

func GetItem(context *gin.Context) {
	tasksListName := context.Param("listName")
	tasks := functions.GetRapidItems(context, &tasksListName)
	context.IndentedJSON(http.StatusOK, tasks.Value)
}
