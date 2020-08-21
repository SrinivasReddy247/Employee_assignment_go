package handler

import (
	//"encoding/json"
	"github.com/srinivasrdy247/Employee_assignment_go/model"
	//"fmt"
	"time"
)

type userResponse struct{
	Employee struct{
		EId int `json:"Id"` 
		Name string `json:"Name"`
		Email string `json:"Email"`
		//Status string `json:"Status"`
		DOB time.Time `json:"DOB"`
	}`json:"Employee"`
}

type userlist struct{
	Employee struct{
		EId int `json:"Id"`
		Name string `json:"Name"`
		Email string `json:"Email"`
		DOB time.Time `json:"DOB"`
	}`json:"Employee"`
}

func newUserResponse(u *model.Employee) *userResponse{
	r := new(userResponse)
	r.Employee.EId = u.EId
	r.Employee.Email = u.Email
	r.Employee.Name = u.Name
//	r.Employee.Status = u.Status
	r.Employee.DOB = u.DOB
	return r
}

func elist(List []model.Employee) *[]userlist{
	len := len(List)
	r := make([]userlist,len)
	for i := 0; i < len; i++{
		r[i].Employee.EId = List[i].EId
		r[i].Employee.Name = List[i].Name
		r[i].Employee.Email = List[i].Email
		r[i].Employee.DOB = List[i].DOB
		//fmt.Print("r[i]:",r[i])
	}
	return &r
}