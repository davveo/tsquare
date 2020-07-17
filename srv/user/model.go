package user

import (
	"database/sql"
	"time"
)

type BaseGormModel struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type TimestampModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserModel struct {
	BaseGormModel
	UserName string         `sql:"type:varchar(254);unique;not null"`
	Password sql.NullString `sql:"type:varchar(60)"`
}

func (u *UserModel) TableName() string {
	return "tb_user"
}
