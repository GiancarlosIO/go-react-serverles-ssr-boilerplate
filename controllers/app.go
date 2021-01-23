package controllers

import (
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *httprouter.Router
}

