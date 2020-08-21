package handler

import (
	"fmt"
	//"errors"
	"time"
	//"github.com/gosimple/slug"
	"github.com/labstack/echo"
	"github.com/srinivasrdy247/Employee_assignment_go/model"
)

const (
	startdate="2000-01-01"
	dateLayout="2000-01-01"
)

type empRegisterRequest struct {
	Employee struct {
		startdate string 
		Name string `json:"User" validate:"required"`
		Email    string `json:"Email" validate:"required,email"`
		Password string `json:"Password" validate:"required,gte=7"`
		DOB	string `json:"DOB" validate:"required,DOB"`
	} `json:"Employee"`
}

func (requestingEmp *empRegisterRequest) bind(c echo.Context, employee *model.Employee) error {
	if err := c.Bind(requestingEmp); err != nil {
		fmt.Println("IN this")
		return err
	}
	/*t1,_ := time.Parse(startdate, requestingEmp.Employee.DOB)
	t1.Format("2000-01-02")
	fmt.Print("startdate:", requestingEmp.Employee.DOB)*/
	//requestingEmp.Employee.startdate=startdate
	err := c.Validate(requestingEmp)
	if  err != nil {
		return err
	}
	employee.Name = requestingEmp.Employee.Name
	employee.Email = requestingEmp.Employee.Email
	employee.Password = requestingEmp.Employee.Password
	employee.DOB,err = time.Parse("02-01-2006",requestingEmp.Employee.DOB)
	if err!=nil{
		return err
	}
	return nil
}

type empLogin struct{
	Employee struct{
		Email    string `json:"Email" validate:"required,email"`
		Password string `json:"Password" validate:"required"`
	} `json:"Employee"`
}

func (requestingEmp *empLogin) bind(c echo.Context,employee *model.Employee) error{
	if err := c.Bind(requestingEmp); err!=nil{
		return err
	}
	if err := c.Validate(requestingEmp);err!=nil{
		return err
	}
	employee.Email = requestingEmp.Employee.Email
	employee.Password = requestingEmp.Employee.Password
	return nil
}

type empUpdate struct{
	Employee struct{
		DOB string `json:"DOB" validate:"required,DOB"`
		Name string `json:"Name" validate:"required"` 
	}`json:"Employee"`
}

func (requestingEmp *empUpdate) bind(c echo.Context,employee *model.Employee)(error){
	err := c.Bind(requestingEmp);
	if  err != nil{
		return err
	}
	/*if requestingEmp.Employee.Name == ""  {
		return errors.New("Need DOB or Name to update")
	}*/
	if err := c.Validate(requestingEmp);err!=nil{
		return err
	}
	employee.DOB,err = time.Parse("02-01-2006",requestingEmp.Employee.DOB)
	if err != nil{
		return err
	}
	employee.Name = requestingEmp.Employee.Name
	return nil
}