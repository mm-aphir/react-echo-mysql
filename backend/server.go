package main

import (
	"fmt"
	"net/http"
	
	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/middleware"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Initialize() {
	db, err := gorm.Open("mysql", "root:z2q3wd31@tcp(127.0.0.1:3306)/db_echo?charset=utf8&parseTime=True",)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected!!!")
	}

	db.AutoMigrate(&Member{}, &Gender{})
	db.Model(&Member{}).AddForeignKey("gender_id", "gender(gender_id)", "CASCADE", "CASCADE")
	db.Model(&Gender{}).AddIndex("index_gender_id_name", "gender_id", "genderName")
	h.DB = db
}

type Member struct {
	MemberID	uint   `gorm:"primary_key" json:"member_id"`
	FirstName 	string `gorm:"not null" json:"firstName,omitempty"`
	LastName  	string `gorm:"not null" json:"lastName,omitempty"`
	Age       	int    `gorm:"not null" json:"age,omitempty"`
	GenderID    int    `gorm:"not null" json:"gender_id"`
	Gender      Gender  `gorm:"foreignkey:gender_id;association_foreignkey:gender_id"`
}

type Gender struct {
	GenderID	uint   `gorm:"primary_key" json:"gender_id"`
	GenderName 	string `json:"genderName"`
}

func (h *Handler) GetAllMember(c echo.Context) error {
	member := []Member{}

	h.DB.Joins("JOIN genders on members.gender_id = genders.gender_id ").Preload("Gender").Find(&member)
	

	return c.JSON(http.StatusOK, member)
}

func (h *Handler) GetMember(c echo.Context) error {
	id := c.Param("id")
	member := Member{}

	if err := h.DB.Find(&member, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, member)
}

func (h *Handler) GetAllGender(c echo.Context) error {
	gender := []Gender{}

	h.DB.Find(&gender)

	return c.JSON(http.StatusOK, gender)
}

func (h *Handler) GetGender(c echo.Context) error {
	id := c.Param("id")
	gender := Gender{}

	if err := h.DB.Find(&gender, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, gender)
}

func (h *Handler) AddGender(c echo.Context) error {
	gender := Gender{}

	if err := c.Bind(&gender); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&gender).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, gender)
}

func (h *Handler) AddMember(c echo.Context) error {
	member := Member{}

	if err := c.Bind(&member); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&member).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, member)
}

func (h *Handler) DeleteMember(c echo.Context) error {
	id := c.Param("id")
	member := Member{}

	if err := h.DB.Find(&member, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := h.DB.Delete(&member).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) UpdateMember(c echo.Context) error {
	id := c.Param("id")
	member := Member{}

	if err := h.DB.Find(&member, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&member); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&member).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, member)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	h := Handler{}
	h.Initialize()

	e.GET("/member", h.GetAllMember)
	e.GET("/member/:id", h.GetMember)
	e.GET("/gender", h.GetAllGender)
	e.GET("/gender/:id", h.GetGender)
	e.POST("/member", h.AddMember)
	e.POST("/gender", h.AddGender)
	e.DELETE("/member/:id", h.DeleteMember)
	e.PUT("/member/:id", h.UpdateMember)

	e.Logger.Fatal(e.Start(":1323"))
}