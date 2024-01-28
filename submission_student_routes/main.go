package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// * Membuat struktur data student berdasarkan key dan value
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

// * Inisiasi Student Database
var students []Student

// * Inisiasi Web Server routes
func initRoutes() {
	e := echo.New()

	// Routing
	e.GET("/students", getStudents)
	e.GET("/students/:id", getStudent)
	e.POST("/students", createStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	e.Start(":8080")
}

// TODO Problem and Testing 1: Buat Fungsi yang memberikan seluruh informasi student
func getStudents(c echo.Context) error {
	// Tulis jawabanmu disini
}

// TODO Problem and Testing 2: Buat Fungsi yang memberikan informasi student berdasarkan ID
func getStudent(c echo.Context) error {
	// Tulis jawabanmu disini

	return c.JSON(http.StatusOK, students)
}

// TODO Problem and Testing 3: Buat Fungsi yang membuat informasi student baru
func createStudent(c echo.Context) error {
	// Tulis jawabanmu disini

	return c.JSON(http.StatusCreated, student)
}

// TODO Problem and Testing 4: Buat Fungsi yang dapat mengupdate informasi student berdasarkan ID
func updateStudent(c echo.Context) error {
	// Tulis jawabanmu disini

	return c.JSON(http.StatusOK, students)
}

// TODO Problem and Testing 5: Buat Fungsi yang dapat menghapus informasi student berdasarkan ID
func deleteStudent(c echo.Context) error {
	// Tulis jawabanmu disini

	return c.JSON(http.StatusNoContent, students)
}
