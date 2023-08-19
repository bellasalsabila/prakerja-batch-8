package controller

import (
	"coffeshop/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetData(c echo.Context) error {
	id := c.Param("id")
	data, err := model.GetData(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve data", nil))
	}
	return c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "Data retrieved successfully", data))
}

func CreateData(c echo.Context) error {
	reqData := new(model.Data)
	if err := c.Bind(reqData); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request data", nil))
	}
	err := model.CreateData(reqData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Failed to create data", nil))
	}
	return c.JSON(http.StatusCreated, model.NewSuccessResponse(http.StatusCreated, "Data created successfully", nil))
}

func UpdateData(c echo.Context) error {
	id := c.Param("id")
	reqData := new(model.Data)
	if err := c.Bind(reqData); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request data", nil))
	}
	err := model.UpdateData(id, reqData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Failed to update data", nil))
	}
	return c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "Data updated successfully", nil))
}

func DeleteData(c echo.Context) error {
	id := c.Param("id")
	err := model.DeleteData(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Failed to delete data", nil))
	}
	return c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "Data deleted successfully", nil))
}

func Login(c echo.Context) error {
	// Implement logic for user login and JWT generation
	// You can use controller.JWTMiddleware() to protect routes
	return c.String(http.StatusOK, "Login endpoint")
}
