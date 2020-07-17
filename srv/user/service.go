package user

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/RichardKnop/uuid"
	"github.com/jinzhu/gorm"
	"github.com/zbrechave/tsquare/util"
	pass "github.com/zbrechave/tsquare/util/password"
)

var (
	MinPasswordLength         = 6
	ErrUserNotFound           = errors.New("User not found")
	ErrUsernameTaken          = errors.New("Username taken")
	ErrCannotSetEmptyUsername = errors.New("Cannot set empty username")
	ErrPasswordTooShort       = fmt.Errorf("Password must be at least %d characters long", MinPasswordLength)
)

// 创建用户
func (s *service) CreateUser(username, password string) (*UserModel, error) {
	return s.createUserCommon(s.db, username, password)
}

// 更新用户
func (s *service) UpdateUsername(user *UserModel, username string) error {
	if username == "" {
		return ErrCannotSetEmptyUsername
	}
	return s.updateUsernameCommon(s.db, user, username)
}

func (s *service) SetPassword(user *UserModel, password string) error {
	return s.setPasswordCommon(s.db, user, password)
}

func (s *service) AuthUser(username, thePassword string) (*UserModel, error) {

}

func (s *service) UserExists(username string) bool {
	_, err := s.QueryUserByName(username)
	return err == nil
}

func (s *service) QueryUserByName(username string) (*UserModel, error) {
	user := new(UserModel)
	notFound := s.db.Where("username = LOWER(?)", username).
		First(user).RecordNotFound()
	if notFound {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *service) createUserCommon(db *gorm.DB, username, password string) (*UserModel, error) {
	user := &UserModel{
		BaseGormModel: BaseGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		UserName: strings.ToLower(username),
		Password: util.StringOrNull(""),
	}
	if password != "" {
		if len(password) < MinPasswordLength {
			return nil, ErrPasswordTooShort
		}
		passwordhash, err := pass.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.Password = util.StringOrNull(string(passwordhash))
	}

	if s.UserExists(user.UserName) {
		return nil, ErrUsernameTaken
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) updateUsernameCommon(db *gorm.DB, user *UserModel, username string) error {
	if username == "" {
		return ErrCannotSetEmptyUsername
	}
	return db.Model(user).UpdateColumn("username", strings.ToLower(username)).Error
}

func (s *service) setPasswordCommon(db *gorm.DB, user *UserModel, password string) error {
	if len(password) < MinPasswordLength {
		return ErrPasswordTooShort
	}
	passwordHash, err := pass.HashPassword(password)
	if err != nil {
		return err
	}

	return db.Model(user).UpdateColumn(UserModel{
		Password:      util.StringOrNull(string(passwordHash)),
		BaseGormModel: BaseGormModel{UpdatedAt: time.Now().UTC()},
	}).Error
}
