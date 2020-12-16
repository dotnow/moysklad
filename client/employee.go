package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// EmployeeRequest структура для запросов сущности 'employee'
type EmployeeRequest struct{ *Request }

// Employee устанавливает нужный endpoint
func (api *APIClient) Employee(params Params) *EmployeeRequest {
	return &EmployeeRequest{api.newRequest("entity/employee", params)}
}

// GetByUUID возвращает сущность по UUID
func (client *EmployeeRequest) GetByUUID(uuid string) (employee *entities.Employee, err error) {

	response, err := client.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// GetList возвращает список сущностей
func (client *EmployeeRequest) GetList() (employees []entities.Employee, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}
