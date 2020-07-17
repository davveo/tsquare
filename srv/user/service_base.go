package user

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	s *service
	m sync.RWMutex
)

// service 服务
type service struct {
	db *gorm.DB
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func InitService() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
