package model

type User struct {
	Id        string `json:"id" gorm:"column:id;"`
	Email     string `json:"email" gorm:"column:email;"`
	Password  string `json:"-" gorm:"column:password;"`
	Salt      string `json:"-" gorm:"column:salt;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Phone     string `json:"phone" gorm:"column:phone;"`
	Role      string `json:"role" gorm:"column:role;"`
}

func (User) TableName() string {
	return "users"
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}
