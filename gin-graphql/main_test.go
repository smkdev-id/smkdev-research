// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestStudentAPI(t *testing.T) {
	e := echo.New()
	initRoutes()

	// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED
	t.Run("TestCreateStudent", func(t *testing.T) {
		reqBody := `{"name": "John Doe", "age": 20, "grade": "A"}`
		req := httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, createStudent(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})

	// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED
	t.Run("TestGetStudents", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/students", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, getStudents(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var studentsList []Student
			err := json.Unmarshal(rec.Body.Bytes(), &studentsList)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, studentsList)
			}
		}
	})

	// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED
	t.Run("TestGetStudent", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/students/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, getStudent(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED
	t.Run("TestUpdateStudent", func(t *testing.T) {

		reqBody := `{"name": "Alice", "age": 22, "grade": "B"}`
		req := httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, createStudent(c))

		reqBody = `{"name": "Alice Smith", "age": 23, "grade": "A"}`
		req = httptest.NewRequest(http.MethodPut, "/students/2", strings.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)
		if assert.NoError(t, updateStudent(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED
	t.Run("TestDeleteStudent", func(t *testing.T) {
		// Create a new student
		reqBody := `{"name": "Bob", "age": 21, "grade": "C"}`
		req := httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, createStudent(c))

		// Delete the created student
		req = httptest.NewRequest(http.MethodDelete, "/students/3", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)
		if assert.NoError(t, deleteStudent(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})
}

// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED
