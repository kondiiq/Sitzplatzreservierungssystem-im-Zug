package model

import (
	"gorm.io/gorm"
	"time"
	"zugSystem/controller"
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
	distance      int
	noStops       int
	startStation  string `gorm:"type:varchar(25)"`
	finishStation string `gorm:"type:varchar(25)"`
	Trains        []Train
	stops         []routeStop `gorm:"many2many:route_routestops"`
}

type routeStop struct {
	gorm.Model
	city          string `gorm:"type:varchar(50);unique;not null"`
	arrivalTime   time.Time
	departureTime time.Time
	timeDuration  int
	routesID      []Route
}

type Train struct {
	gorm.Model
	name       string `gorm:"type:varchar(50);unique;not null"`
	number     int
	typeTrain  string `gorm:"type:varchar(25)"`
	horsePower int
	weight     int
	length     int
	topSpeed   int
	workSpeed  int
	routeID    uint `gorm:"primaryKey,AUTO_INCREMENT;unique;not null"`
}

func (rs *routeStop) onBeforeCreateRouteStop(tx *gorm.DB) (err error) {
	rs.calculateDuration()
	return nil
}

func (rs *routeStop) onBeforeUpdateRouteStop(tx *gorm.DB) (err error) {
	rs.calculateDuration()
	return nil
}

func (rs *routeStop) calculateDuration() (int, error) {
	rs.timeDuration = int(rs.departureTime.Sub(rs.arrivalTime).Minutes())
	if rs.timeDuration < 0 {
		return rs.timeDuration, controller.CustomError{Message: "Hmm it's not a time machine time is less than 0"}
	}
	return rs.timeDuration, nil
}
