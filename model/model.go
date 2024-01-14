package model

//type disconutType struct {
//	employeeDiscount  bool
//	studentDiscount   bool
//	pensionerDiscount bool
//	blindnessDiscount bool
//}
//
//type journeyClass struct {
//	firstClass   bool
//	economyClass bool
//}
//
//type BaseModel struct {
//	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
//	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
//	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
//}
//
//type User struct {
//	ID        uint      `gorm:"primary_key;AUTO_INCREMENT;"`
//	login     string    `gorm:"type:varchar(30);" json:"login"`
//	password  string    `gorm:"type:varchar(30);" json:"password"`
//	name      string    `gorm:""`
//	lastName  string    `gorm:""`
//	birthDate time.Time `gorm:""`
//	age       int       `gorm:""`
//}
//
//type Route struct {
//	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
//	startRoute  string `gorm:"type:varchar(30);" json:"startRoute"`
//	destination string `gorm:"type:varchar(50);" json:"destination"`
//	trainStops  string `gorm:"type:varchar(30);" json:"trainStops"`
//	routeLength int    `gorm:"type:int;" json:"routeLength"`
//	train       Train  `gorm:"foreign_key" json:"train"`
//	eatingCar   bool   `gorm:"type:bool" json:"eatingCar"`
//}
//
//type Train struct {
//	ID               uint   `gorm:"primary_key;AUTO_INCREMENT"`
//	name             string `gorm:"type:varchar(30);" json:"name"`
//	number           int    `gorm:"type:uint;" json:"number"`
//	destinationSpeed int    `gorm:"type:uint;" json:"destinationSpeed"`
//	isDelayed        bool   `gorm:"type:bool;" json:"isDelayed"`
//	minDelayed       int    `gorm:"type:int;" json:"minDelayed"`
//}
//
//type Ticket struct {
//	ID           uint         `gorm:"primary_key;AUTO_INCREMENT"`
//	owner        User         `gorm:""`
//	route        Route        `gorm:""`
//	trainNumber  int          `gorm:""`
//	date         time.Time    `gorm:""`
//	class        journeyClass `gorm:""`
//	discount     bool         `gorm:""`
//	discountType disconutType `gorm:""`
//	noAdventures int          `gorm:""`
//	timeStart    time.Time    `gorm:""`
//}

type User struct {
	ID       uint   `gorm:"type:uint;primary_key;unique" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	ID      uint   `gorm:"type:uint;primary_key;unique'" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}
