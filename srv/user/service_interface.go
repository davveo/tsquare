package user

// Service 用户服务类
type Service interface {
	CreateUser(username, password string) (*UserModel, error)
	SetPassword(user *UserModel, password string) error
	UpdateUsername(user *UserModel, username string) error
	AuthUser(username, thePassword string) (*UserModel, error)
	UserExists(username string) bool
	QueryUserByUserId(id int64) (*UserModel, error)
	QueryUserByPhone(phone string) (*UserModel, error)
	QueryUserByName(username string) (*UserModel, error)
	QueryUserList(page, size int) (*[]UserModel, error)
}
