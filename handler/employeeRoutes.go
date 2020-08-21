package handler

import(
	//"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/srinivasrdy247/Employee_assignment_go/handler/middle"
)

func (h *Handler) Register(v *echo.Group){
	guestEmpGroup := v.Group("")
	guestEmpGroup.POST("/signup", h.SignIn)
	//fmt.Print("b:",b)
	//basicauth:=middleware1.Basic()
	empGroup := v.Group("/employee")
	empGroup.Use(middleware.BasicAuth( func (username, password string, c echo.Context) (bool, error){
		err := h.eselect.EmpAuth(username, password)
		if err != nil{
			return false, err
		}
		return true, nil
	}))
	empGroup.GET("", h.EmployeeList)
	empGroup.GET("/:id", h.GetByID)
	empGroup.GET("/name/:username", h.GetByName)
	empGroup.DELETE("/:id", h.DeleteEmployee)
	empGroup.PUT("/:id", h.UpdateEmployee) // api/id/
}