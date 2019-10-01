package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*Gender{
		"gender_id": &Gender{3, "Bi"},
	}
	genderJSON = `{"genderName":"Bi"}`
)

// var (
// 	mockDB = map[string]*Member{
// 		"member_id": &Member{13, "Memtest","aaa",22,1,1},
// 	}
// 	memberJSON = `{"firstname":"Memtest","lastname":"aaa","age":22,"gender_id":1}`
// )

func TestCreateGender(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/gender", strings.NewReader(genderJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &main{mockDB}

	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, genderJSON, rec.Body.String())
	}
}

// func TestGetMember(t *testing.T) {
// 	r := setupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/major", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotNil(t, w.Body)
// }