package store

import (
	"errors"
    "strconv"

    //"log"
    "fmt"
	"github.com/jmoiron/sqlx"
    "github.com/srinivasrdy247/Employee_assignment_go/model"
    "github.com/srinivasrdy247/Employee_assignment_go/util"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Connecting struct{
    db *sqlx.DB
}

func NewCon(d *sqlx.DB) *Connecting{
    return &Connecting{
        db : d,
    }
}

func (connect *Connecting) CreateUser(employee *model.Employee) (error){
    encryptPass := util.Encrypt([]byte(employee.Password),"e_pass")

//    fmt.Println("error:",pass)
    err := connect.db.Get(&employee.EId,"INSERT INTO Employee(email,name,password,status,dob) VALUES ($1,$2,$3,$4,$5) returning id",employee.Email,employee.Name,encryptPass,true,employee.DOB)
    
    if err != nil{
        return errors.New("Not inserted into database")
    }
   // connect.db.Get(&employee.EId,"SELECT id from employee where email=$1;", employee.Email)
    return nil
}

    /*var n int //4 
    //err:= d.db.Get(&n,"SELECT count(*) from Employee where Email=$1 and Name=$2;",u.Email,u.Name)
    fmt.Println("n:",err)
    if err!=nil{
        return false,err
    }
    fmt.Println("n:",n)
    if n==0{
        err=d.db.Get(&n,"SELECT count(*) from Employee;")
        n++
        fmt.Println("inside:",n)
        u.Status="active"
        u.EId=n
       
        fmt.Print("result:",r)
    return true,err
    }
    return false,nil*/


func (connect *Connecting) EmpByID(id string) (*model.Employee, error){
    employee := model.Employee{}
    fmt.Print("Id")
    err := connect.db.Get(&employee,"SELECT * from Employee where id=$1;", id)
    if err != nil{
        return nil, errors.New("Invalid user id in param")
    }
    if employee.EId != 0 {
        if employee.Status == "true" {
            return &employee, nil
        }
        return nil, errors.New("user not present")
    }
    return nil,errors.New("while fetching param")
}

func (connect *Connecting) EmpByName(name string) (*model.Employee, error){
    employee := model.Employee{}
    fmt.Print("name")
    err := connect.db.Get(&employee,"SELECT * from Employee where name=$1;", name)
    if err != nil{
        return nil,errors.New("No employee present with name:"+name)
    }
    if employee.EId != 0 {
        if employee.Status == "true" {
            return &employee,nil
        }
        return nil,errors.New("User not present")
    }
    return nil,errors.New("While fetching param")
}

func (connect *Connecting) EmpList() ([]model.Employee,error){
    employee := []model.Employee{}
    //status := "true"
    err := connect.db.Select(&employee,"SELECT * FROM employee WHERE status=$1;",true)
    if err != nil{
        return employee, errors.New("Database")
    }
    return employee, nil
}

func (connect *Connecting) EmpAuth(email,password string) (error){
    var pass []byte
    err := connect.db.Get(&pass,"SELECT password from employee where email=$1 and status=$2;",email,true)
    //fmt.Print("s:",status)
    if err != nil{
        return errors.New("No such user is present")
    }
    if password == string(util.Decrypt(pass,"e_pass")) {
        return nil
    }
    return errors.New("Unauthorised")
}

func (connect *Connecting) EmpDelete(ID string) (error){
    var status string
    err := connect.db.Get(&status,"SELECT status from Employee where id=$1;",ID)
    if err != nil{
        return errors.New("No user present to delete")
    }
    if status == "true"{
       connect.db.MustExec("UPDATE EMPLOYEE SET STATUS=$1 WHERE ID=$2;",false,ID)
        return nil
    }
    return errors.New("No user present to delete")
    
}

func(connect *Connecting) EmpUpdate(employee *model.Employee, ID string)(error){
    var count int
    err := connect.db.Get(&count,"SELECT count(*) from Employee where id=$1 and status=$2;",ID,true)
    if err != nil{
        return errors.New("While fetching data from database")
    }
    if count == 1 {
        connect.db.Get(&count,"SELECT count(*) from Employee where name=$1;",employee.Name)
        if count == 0{
            connect.db.MustExec("UPDATE EMPLOYEE SET name=$1, dob=$2 WHERE ID=$3;",employee.Name,employee.DOB,ID)
            return nil
        } 
        return errors.New("Name already present")
    }
    return errors.New("User not present")
      // return errors.New("Need any DOB or Name to update")
}

func (d *Connecting) Makefile() {
    employee := []model.Employee{}
    err := d.db.Select(&employee,"SELECT * FROM employee WHERE status=$1;",true)
    if err != nil{
       fmt.Println(err.Error())
       return
    }
    f := excelize.NewFile()
    index:=f.NewSheet("sheet2")
    f.SetCellValue("sheet2","A1","Employee_id")
    f.SetCellValue("sheet2","B1","Name")
    f.SetCellValue("sheet2","C1","Email")
    f.SetCellValue("sheet2","D1","Status")
    f.SetCellValue("sheet2","E1","DOB")
    for i := 0;i<len(employee);i++{
        s:=strconv.Itoa(i+2)
        f.SetCellValue("sheet2","A"+s,employee[i].EId)
        f.SetCellValue("sheet2","B"+s,employee[i].Name)
        f.SetCellValue("sheet2","C"+s,employee[i].Email)
        f.SetCellValue("sheet2","D"+s,employee[i].Status)
        f.SetCellValue("sheet2","E"+s,employee[i].DOB)
    }
    f.SetActiveSheet(index)
    if err := f.SaveAs("Book2.xlsx"); err != nil {
        println(err.Error())
    }

}
/*func (d *Connecting) Empbymail(s string) (*model.Employee){
    var u model.Employee
    err:=d.db.Get(&u.Status,"SELECT id from Employee where Email=$1;",u.Email)
    if err!=nil{
        return nil
    }
    if u.Status=="active"{
        d.db.
    }

    //doubt
    v:=model.Employee{}
            r,_:=d.db.Queryx("SELECT * from Employee where Email=$1 and Password=$2;",u.Email,u.Password)
            for r.Next(){
               err:= r.StructScan(&v)
               
                fmt.Print(r)
                fmt.Printf("v: %#v\n",v)
                if err!=nil{
                    log.Fatalln(err)
                }
            }

}*/



