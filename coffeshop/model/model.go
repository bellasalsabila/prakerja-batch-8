package model

import (
	"errors"
)

type Data struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Purchase int32  `json:"purchase"`
}

var dataStore = make(map[string]*Data)

func GetData(id string) (*Data, error) {
	data, exists := dataStore[id]
	if !exists {
		return nil, errors.New("Data not found")
	}
	return data, nil
}

func CreateData(data *Data) error {
	// Implement logic to create data
	dataStore[data.ID] = data
	return nil
}

func UpdateData(id string, newData *Data) error {
	// Implement logic to update data
	data, exists := dataStore[id]
	if !exists {
		return errors.New("Data not found")
	}
	data.Name = newData.Name
	// ... update field lain sesuai kebutuhan
	return nil
}

func DeleteData(id string) error {
	// Implement logic to delete data
	_, exists := dataStore[id]
	if !exists {
		return errors.New("Data not found")
	}
	delete(dataStore, id)
	return nil
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
