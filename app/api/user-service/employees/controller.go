package employees

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UUID uuid.UUID
type Employee struct {
	Id           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"password_hash"`

	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Position  string    `json:"position"`

	DepartmentId int `json:"department_id"`
	RoleId       int `json:"role_id"`
}

func CreateEmployee(c *gin.Context) {
	u := generateRandomEmployee()

	fmt.Println("he", u)
}
func GetEmployees(c *gin.Context) {
	u := generateRandomEmployee()

	fmt.Println("he", u)
}
func GetEmployee(c *gin.Context) {

}
func DeleteEmployee(c *gin.Context) {

}
func UpdateEmployee(c *gin.Context) {

}

func generateRandomEmployee() Employee {
	firstNames := []string{"John", "Jane", "Alex", "Emily", "Chris"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Jones", "Brown"}
	positions := []string{"Developer", "Manager", "Analyst", "Designer", "Tester"}
	genders := []string{"Male", "Female", "Other"}

	return Employee{
		Id:           uuid.New(),
		FirstName:    firstNames[rand.Intn(len(firstNames))],
		LastName:     lastNames[rand.Intn(len(lastNames))],
		Email:        fmt.Sprintf("%s.%s@example.com", firstNames[rand.Intn(len(firstNames))], lastNames[rand.Intn(len(lastNames))]),
		Phone:        fmt.Sprintf("555-01%04d", rand.Intn(10000)),
		PasswordHash: fmt.Sprintf("%x", rand.Int63()),

		BirthDate: time.Date(1980+rand.Intn(30), time.Month(1+rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC),
		Gender:    genders[rand.Intn(len(genders))],
		Position:  positions[rand.Intn(len(positions))],

		DepartmentId: rand.Intn(100) + 1,
		RoleId:       rand.Intn(10) + 1,
	}
}
