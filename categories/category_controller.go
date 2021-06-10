package categories

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/leandrorochaadm/time-register/database"
	"net/http"
	"strconv"
)

type Category struct {
	ID          uint64 `json:"id"`
	Description string `json:"description"`
}

type Categories []Category

var ciclePrimary bool

func GetCategories(c echo.Context) error {
	var categoryList Categories
	rows, err := database.DBClient.
		Query("SELECT id,  description FROM categories;")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			"Não foi possível listar os categorias \n"+err.Error())
	}
	for rows.Next() {
		var singleCategory Category
		if err := rows.Scan(&singleCategory.ID, &singleCategory.Description); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity,
				"Não foi possível deserializer os categorias \n"+err.Error())
		}

		categoryList = append(categoryList, singleCategory)
		ciclePrimary = true
	}
	return c.JSON(http.StatusOK, categoryList)
	// return c.JSON(http.StatusOK, categoryList)
}

func CreateCategory(c echo.Context) error {
	category := Category{}
	err := c.Bind(&category)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	res, err := database.DBClient.Exec("INSERT INTO categories ( description) VALUES (?);",
		category.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			"Não foi possível gravar esse categoria: \n"+err.Error())
	}
	id, _ := res.LastInsertId()

	// categoryList = append(categoryList, category)
	return c.JSON(http.StatusCreated, id)
}

func DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusUnprocessableEntity,
			"É necessário informar a categoria que deseja deletar\n "+err.Error())
	}
	fmt.Print(id)

	_, err = database.DBClient.Exec("delete from categories where id = ?", id)
	if err != nil {
		echo.NewHTTPError(http.StatusUnprocessableEntity,
			"Não foi possível deletar essa categoria \n "+err.Error())
	}
	return c.String(http.StatusNoContent, "")
}
