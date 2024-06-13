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

func RestHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("rest/index.plush.html"))
}

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
	err := models.DB.Where("end_at IS NULL OR is_finished = 0").All(&data)
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

func PutActiveRest(c buffalo.Context) error {
	var res response.Response
	tableName := models.Rest{}.TableName()

	id := c.Param("id")

	endAt := c.Request().FormValue("end_at")
	if endAt == "" {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"error": "Missing end_at parameter",
		}))
	}

	err := models.DB.RawQuery("UPDATE rests SET end_at = ?, is_finished = 1 WHERE id = ?", endAt, id).Exec()
	if err != nil {
		log.Printf("Error updating data: %v", err)
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]interface{}{
			"error": "Internal Server Error",
		}))
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg("Update", tableName, 1)

	return c.Render(http.StatusOK, r.JSON(res))
}

func PostRest(c buffalo.Context) error {
	var res response.Response
	tableName := models.Rest{}.TableName()
	id, _ := generator.GenerateUUID(16)
	createdAt := generator.GenerateTimeNow("timestamp")

	restNotes := c.Request().FormValue("rest_notes")
	if restNotes == "" {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"error": "Missing rest_notes parameter",
		}))
	}
	restCat := c.Request().FormValue("rest_category")
	if restCat == "" {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"error": "Missing rest_category parameter",
		}))
	}
	startAt := c.Request().FormValue("started_at")
	if startAt == "" {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"error": "Missing started_at parameter",
		}))
	}
	endAt := c.Request().FormValue("end_at")

	var query string
	var params []interface{}

	if endAt != "" {
		query = "INSERT INTO rests(id, rest_notes, rest_category, started_at, end_at, is_finished, created_at, created_by) VALUES (?,?,?,?,?,?,?,?)"
		params = []interface{}{id, restNotes, restCat, startAt, endAt, 0, createdAt, "123"}
	} else {
		query = "INSERT INTO rests(id, rest_notes, rest_category, started_at, is_finished, created_at, created_by) VALUES (?,?,?,?,?,?,?)"
		params = []interface{}{id, restNotes, restCat, startAt, 0, createdAt, "123"}
	}

	err := models.DB.RawQuery(query, params...).Exec()
	if err != nil {
		log.Printf("Error create data: %v", err)
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]interface{}{
			"error": err.Error(),
		}))
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg("Create", tableName, 1)

	return c.Render(http.StatusOK, r.JSON(res))
}
