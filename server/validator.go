package router

import (
	"gopkg.in/go-playground/validator.v9"
	//s "strings"
	//"strconv"
	"fmt"
	"time"
) 

func NewValidator() *Validator {
	return &Validator{
		validator : validator.New(),
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	 v.validator.RegisterValidation("DOB",ValidateDOBfunc) //mm-dd-yyyy
	//validate.RegisterValidation
	return v.validator.Struct(i)
}

func ValidateDOBfunc(f validator.FieldLevel) bool {
	
	dobString := f.Field().String()
	empDOB, err := time.Parse("02-01-2006", dobString)
    if err != nil {
		fmt.Print(err)
        return false
	}
	age := time.Since(empDOB)
	fmt.Print("age:",age)
	if age.Hours() < 175320 {
		return false
	}
	return true
}
	/*count := s.Count(dobString, "-")
	if count!=2{
		return false
	}
	spliting := s.Split(dobString, "-")
	fmt.Print("in dob")
	_, err := strconv.Atoi(spliting[0]) //day
	if err!=nil {
		return false
	}
	_, err = strconv.Atoi(spliting[1]) //month
	if err!=nil {
		return false
	}
	_, err = strconv.Atoi(spliting[2]) //year 
	if err!=nil {
		return false
	}
	//t:=time.Date(year, Month(month), day, 0, 0, 0, 0, time.UTC) // 2 jan 2001
	slice := []string{spliting[0], spliting[1], spliting[2]}
	extractedDOB := s.Join(slice,"-") */

	/*
	if year >= 1950 && year <= 2000{
		if month >= 1 && month <= 12{
			if day >=1 && day <=31{
				if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
					return true
				}
			}
			if day >= 1 && day<=30{
				if month==4 || month ==6 || month==9 || month==11{
					return true
				}
			}
			if day >=1 && day <=28{
				if month==2{
					return true
				}
			}
			if day==29 && month==2{
				if year%4==0{
					return true
				}
			}
		}
	}
	return false
	*/



