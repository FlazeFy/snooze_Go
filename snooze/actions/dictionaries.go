package actions

import (
	"log"
	"net/http"
	"snooze/helpers/generator"
	"snooze/helpers/response"
	"snooze/models"

	"github.com/gobuffalo/buffalo"
)

func GetDictionaryByType(c buffalo.Context) error {
	typeParam := c.Param("type")

	var res response.Response
	tableName := models.Dictionary{}.TableName()

	data := []models.Dictionary{}
	err := models.DB.Where("dictionary_type = ?", typeParam).All(&data)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]interface{}{
			"error": "Internal Server Error",
		}))
	}
	total := len(data)

	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(tableName, total)
	if total != 0 {
		res.Data = data
	} else {
		res.Data = nil
	}

	return c.Render(http.StatusOK, r.JSON(res))
}
