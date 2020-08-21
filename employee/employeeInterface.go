package employee

import (
	"github.com/srinivasrdy247/Employee_assignment_go/model"
)

type E1 interface{
	CreateUser(*model.Employee) (error)
	EmpList() ([]model.Employee, error)
	EmpAuth(string, string) (error)
	EmpByID(string) (*model.Employee, error)
	EmpByName(string) (*model.Employee, error)
	EmpDelete(string) (error)
	EmpUpdate(*model.Employee, string) (error)
}