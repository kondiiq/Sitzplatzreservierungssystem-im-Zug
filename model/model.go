package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey,AUTO_INCREMENT;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Route struct {
	gorm.Model
	name          string `gorm:"type:varchar(50);unique;not null"`
	distance      int    `gorm:"type:int"`
	noStops       int    `gorm:"type:int"`
	startStation  string `gorm:"type:varchar(25)"`
	finishStation string `gorm:"type:varchar(25)"`
	Trains        []Train
}

type Train struct {
	gorm.Model
	name       string `gorm:"type:varchar(50);unique;not null"`
	number     int    `gorm:"type:int"`
	typeTrain  string `gorm:"type:varchar(25)"`
	horsePower int    `gorm:"type:int"`
	weight     int    `gorm:"type:int"`
	length     int    `gorm:"type:int"`
	topSpeed   int    `gorm:"type:int"`
	workSpeed  int    `gorm:"type:int"`
	routeID    uint   `gorm:"primaryKey,AUTO_INCREMENT;unique;not null"`
}
