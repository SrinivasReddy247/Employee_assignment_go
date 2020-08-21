package handler

import (
	"github.com/srinivasrdy247/Employee_assignment_go/employee"
)

type Handler struct{
	eselect  employee.E1
}

func NewHandler(E employee.E1) *Handler{
	return &Handler{
		eselect : E,
	}
}
