package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var ErrNoRows = errors.New("no rows in result set")
var ErrDuplicate = errors.New("duplicate")
var ErrBadRequest = errors.New("bad request")
var ErrUnauthorized = errors.New("unauthorized")

var (
	OK           = map[string]string{"message": "success"}
	NotFound     = map[string]string{"message": "not found"}
	Duplicate    = map[string]string{"message": "duplicate"}
	BadRequest   = map[string]string{"message": "bad request"}
	InternalErr  = map[string]string{"message": "internal server error"}
	Unauthorized = map[string]string{"message": "unauthorized"}
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id,omitempty" gorm:"not null"`
	Title       string    `json:"title" gorm:"size:256;not null"`
	Description string    `json:"description" gorm:"size:256"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:NOW()"`
	DoneAt      time.Time `json:"done_at,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
}

func (t *Task) Validate() bool {
	if t.UserID <= 0 {
		return false
	}
	if len(t.Title) < 3 || len(t.Title) > 256 {
		return false
	}
	if len(t.Description) > 256 {
		return false
	}
	return true
}

type User struct {
	ID       int    `json:"id,omitempty" gorm:"primaryKey"`
	Username string `json:"username,omitempty" gorm:"unique;size:32;not null"`
	Password string `json:"password,omitempty" gorm:"size:128;not null;->:false;<-"`
}

func (u *User) Validate() bool {
	if len(u.Username) < 3 || len(u.Username) > 128 {
		return false
	}
	if len(u.Password) < 4 || len(u.Password) > 32 {
		return false
	}
	return true
}

type JWTClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func ReplyError(c *gin.Context, err error) {
	switch err {
	case ErrUnauthorized:
		c.JSON(http.StatusUnauthorized, Unauthorized)
	case ErrBadRequest:
		c.JSON(http.StatusBadRequest, BadRequest)
	case ErrNoRows:
		c.JSON(http.StatusNotFound, NotFound)
	case ErrDuplicate:
		c.JSON(http.StatusNotAcceptable, Duplicate)
	default:
		c.JSON(http.StatusInternalServerError, InternalErr)
	}
	return
}

type Settings struct {
	AppParams      Params           `json:"app"`
}

type Params struct {
	ServerURL         string `json:"server_url"`
	Login             string `json:"login"`
	OfferLink         string `json:"offer_link"`
	ServerName        string `json:"server_name"`
	AppVersion        string `json:"app_version"`
	PortRun           string `json:"port_run"`
	LogInfo           string `json:"log_info"`
	LogError          string `json:"log_error"`
	LogDebug          string `json:"log_debug"`
	LogWarning        string `json:"log_warning"`
	LogMachineHWID    string `json:"log_machine_hw_id"`
	LogMaxSize        int    `json:"log_max_size"`
	LogMaxBackups     int    `json:"log_max_backups"`
	LogMaxAge         int    `json:"log_max_age"`
	LogCompress       bool   `json:"log_compress"`
	AuthServiceURL    string `json:"auth_service_url"`
	SecretKey         string `json:"secret_key"`
	PaymentServiceURL string `json:"payment_service_url"`
}