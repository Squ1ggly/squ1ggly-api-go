package controllers

import (
	"net/http"
	"squ1ggly/squ1ggly-api-go/functions"

	"github.com/gin-gonic/gin"
)

func GetItem(context *gin.Context) {
	tasksListName := context.Param("listName")
	filter := context.Query("$filter")
	tasks := functions.GetRapidItems(context, &tasksListName, &filter)
	context.IndentedJSON(http.StatusOK, tasks.Value)
}
