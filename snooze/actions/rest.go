package actions

import (
	"log"
	"math"
	"net/http"
	"snooze/helpers/generator"
	"snooze/helpers/pagination"
	"snooze/helpers/response"
	"snooze/models"
	"strconv"

	"github.com/gobuffalo/buffalo"
)

func GetAllRest(c buffalo.Context) error {
	page, _ := strconv.Atoi(c.Params().Get("page"))

	var res response.Response
	tableName := models.Rest{}.TableName()
	pageSize := 10

	data := []models.Rest{}
	err := models.DB.All(&data)
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

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, "api/v1/rest")

	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(tableName, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           data,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return c.Render(http.StatusOK, r.JSON(res))
}

func GetActiveRest(c buffalo.Context) error {
	var res response.Response
	tableName := models.Rest{}.TableName()

	data := []models.Rest{}
	err := models.DB.Where("end_at IS NULL").All(&data)
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
