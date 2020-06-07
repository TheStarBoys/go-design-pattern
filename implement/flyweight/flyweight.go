package flyweight

import (
	"sync"
	"strings"
	"fmt"
)

var (
	once1 sync.Once
	once2 sync.Once
	factory *FlyweightFactory
	manager *SecurityManager
	db = []struct{ // 模拟数据库数据
		Username string
		SecurityEntity string
		Permit string
	}{
		{"张三", "人员列表", "查看"},
		{"李四", "人员列表", "查看"},
		{"李四", "薪资数据", "查看"},
		{"李四", "薪资数据", "修改"},
	}
)

type UserModel struct{
	Username string
	SecurityEntity string
	Permit string
}

type Flyweight interface {
	// 判断传入的安全实体和权限，是否和享元对象内部状态匹配
	Match(securityEntity, permit string) bool
}

type AuthorizationFlyweight struct {
	// 内部状态
	securityEntity string // 安全实体
	permit string // 权限
}

func NewAuthorizationFlyweight(securityEntity, permit string) *AuthorizationFlyweight {
	return &AuthorizationFlyweight{
		securityEntity: securityEntity,
		permit: permit,
	}
}

func (a AuthorizationFlyweight) GetSecurityEntity() string {
	return a.securityEntity
}

func (a AuthorizationFlyweight) GetPermit() string {
	return a.permit
}

func (a AuthorizationFlyweight) Match(securityEntity, permit string) bool {
	return a.securityEntity == securityEntity && a.permit == permit
}

type FlyweightFactory struct {
	fsMap map[string]Flyweight
}

func GetFactory() *FlyweightFactory {
	once1.Do(func() {
		factory = &FlyweightFactory{
			fsMap: map[string]Flyweight{},
		}
	})

	return factory
}

func (factory *FlyweightFactory) GetFlyweight(key string) Flyweight {
	f, ok := factory.fsMap[key]
	if !ok {
		ss := strings.Split(key, ",")
		f = NewAuthorizationFlyweight(ss[0], ss[1])
		factory.fsMap[key] = f
	}

	return f
}

// SecurityManager 安全管理
type SecurityManager struct {
	m map[string][]Flyweight
}

func GetManager() *SecurityManager {
	once2.Do(func() {
		manager = &SecurityManager{
			m: map[string][]Flyweight{},
		}
	})

	return manager
}

func (s SecurityManager) Login(user string) {
	// 登录时需要把该用户所拥有的权限，从数据库中取出来，放到缓存
	col := s.queryByUser(user)
	s.m[user] = col
}

// 判断某用户对某个安全实体是否拥有某种权限
func (s SecurityManager) HasPermit(user, securityEntity, permit string) bool {
	col, ok := s.m[user]
	if !ok {
		fmt.Printf("%s: 没有登录或没有任何权限\n", user)
		return false
	}
	for _, fm := range col {
		au, ok := fm.(*AuthorizationFlyweight)
		if ok {
			fmt.Printf("fm: %p\n", au)
		}
		if fm.Match(securityEntity, permit) {
			return true
		}
	}

	return false
}

// 从数据库中获取某人所拥有的权限
func (s SecurityManager) queryByUser(user string) []Flyweight {
	var res []Flyweight
	for _, data := range db {
		if data.Username == user {
			fm := GetFactory().GetFlyweight(fmt.Sprintf("%s,%s", data.SecurityEntity, data.Permit))
			res = append(res, fm)
		}
	}

	return res
}