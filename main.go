package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
	Active bool `json:"active"`
}

func main() {

	employees := getEmployees()
	fmt.Println(employees)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Bienvenido a la empresa Gophers!",
		})
	})

	r.GET("/employees", func (c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	r.GET("/employees/:id", func (c *gin.Context) {
		id := c.Param("id")
		
		employee, err := FindEmployeeById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}
		c.JSON(http.StatusOK, employee)
		
	})


	//Aclaración: No es dinámico, además está en una petición GET cuando debería ser un POST.
	//Es solamente para probar y aprender.
	r.GET("/employeesparams", func (c *gin.Context) {

		employee := CreateEmployee(4, "Roberto", false)
		employees = append(employees, employee)
		c.JSON(http.StatusOK, employee)

	})

	r.GET("/employeesactive", func (c *gin.Context){
		result := FindEmployeesActive(employees, true)
		c.JSON(http.StatusOK, result)
	})

	r.Run() 

}

func getEmployees() []Employee {

	employees := []Employee{
		{1, "Francisco", true},
		{2, "Juan", false},
		{3, "Pepe", true},
	}

	return employees

}

func FindEmployeeById (id string) (Employee, error) {

	var result Employee
	var err error = nil
	var flag = 1

	idInt, err2 := strconv.Atoi(id)
		if err2 != nil {
			fmt.Println("error al convertir el id")
		}
	for _, employee := range getEmployees(){
		if employee.Id == idInt{
			result = employee
			flag = 0
		}
	}

	if flag == 1 {
		err = fmt.Errorf("No se encontró ese empleado")
	}

	return result, err

}

func CreateEmployee (id int, name string, active bool) Employee {

	employee := Employee{
		Id: id,
		Name: name,
		Active: active,
	}

	return employee

}

func FindEmployeesActive (employees []Employee, status bool) []Employee {

	var result []Employee

	for _, employee := range employees {
		if employee.Active == status{
			result = append(result, employee)
		}
	}

	return result

}