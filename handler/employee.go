package handler

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/srinivasrdy247/Employee_assignment_go/model"
	"github.com/srinivasrdy247/Employee_assignment_go/util"
)

func (h *Handler) SignIn(c echo.Context) error{
	var employeeDetails model.Employee
	req := &empRegisterRequest{}
	if err := req.bind(c, &employeeDetails); err != nil{
		fmt.Println("Bind")
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	err := h.eselect.CreateUser(&employeeDetails);
	if err != nil{
		fmt.Println("Cannot create")
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&employeeDetails))
}

func(h *Handler) UpdateEmployee(c echo.Context) error{
	ID := c.Param("id")
	_, err := h.eselect.EmpByID(ID)
	if err != nil{
		return c.JSON(http.StatusBadRequest, util.NewError(err))
	}
	var employee model.Employee
	req := &empUpdate{}
	err = req.bind(c, &employee)
	if err != nil{
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	err = h.eselect.EmpUpdate(&employee, ID)
	if err != nil{
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	return c.NoContent(http.StatusNoContent)
}

func(h *Handler) EmployeeList(c echo.Context) error{
	v, err := h.eselect.EmpList()
	//fmt.Print("n:", n)
	if err!=nil{
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	return c.JSON(http.StatusOK, elist(v))
}

func(h *Handler) GetByID(c echo.Context) error{
	id := c.Param("id")
	u, err := h.eselect.EmpByID(id)
	if err!=nil{
		return c.JSON(http.StatusBadRequest, util.NewError(err))
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}

func(h *Handler) GetByName(c echo.Context) error{
	user := c.Param("username")
	employeeDetails, err := h.eselect.EmpByName(user)
	if err != nil{
		return c.JSON(http.StatusBadRequest, util.NewError(err))
	}
	return c.JSON(http.StatusOK, newUserResponse(employeeDetails))
}

func(h *Handler) DeleteEmployee(c echo.Context) error{
	id := c.Param("id")
	err := h.eselect.EmpDelete(id)
	if err!=nil{
		return c.JSON(http.StatusBadRequest, util.NewError(err))
	}
	return c.NoContent(http.StatusNoContent)
}
