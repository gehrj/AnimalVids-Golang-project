package db

type Users struct{
	First string	`xorm:"first" json:"first" schema:"first"`
	Last string	`xorm:"last" json:"last" schema:"last"`
	Email string	`xorm:"email" json:"email" schema:"email"`
	Password string	`xorm:"password" json:"password" schema:"password"`
}