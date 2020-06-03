package template_method

import (
	"fmt"
	"crypto/sha256"
)

type LoginMethod interface {
	Login(model LoginModel) bool // 模板方法
	findLoginUser(username string) *LoginModel // 强制子类实现的方法
	encryptPwd(pwd string) string // 提供了缺省实现，即不加密
	match(lm1, lm2 LoginModel) bool // 提供了缺省实现
}

type LoginModel struct {
	userName string
	password string
}

type LoginTemplate struct {
	LoginMethod
}

func NewLoginTemplate(method LoginMethod) *LoginTemplate {
	return &LoginTemplate{
		LoginMethod: method,
	}
}

func (t *LoginTemplate) Login(model LoginModel) bool {
	lm2 := t.LoginMethod.findLoginUser(model.userName)
	// 数据库中有该用户
	if lm2 != nil {
		// 加密用户密码
		model.password = t.LoginMethod.encryptPwd(model.password)
		// 比对是否匹配
		return t.LoginMethod.match(model, *lm2)
	}

	return false
}

// findLoginUser 根据用户名查找登录模型
func (t *LoginTemplate) findLoginUser(username string) *LoginModel {
	// 等待子类实现
	panic("need to implement")
	return &LoginModel{}
}

// encryptPwd 加密密码
func (t *LoginTemplate) encryptPwd(pwd string) string {
	// 如果需要，子类可以覆盖
	return pwd
}

// match 比对两个LoginModel是否相等
func (t *LoginTemplate) match(lm1, lm2 LoginModel) bool {
	return lm1.userName == lm2.userName && lm1.password == lm2.password
}

// 普通用户登录
type NormalLogin struct {
	*LoginTemplate
}

func NewNormalLogin() *NormalLogin {
	l := &NormalLogin{}
	l.LoginTemplate = NewLoginTemplate(l)

	return l
}

func (n *NormalLogin) findLoginUser(username string) *LoginModel {
	// 不实际获取信息，而是提供一个默认的用户信息
	return &LoginModel{
		userName: "user1",
		password: "user1_pwd",
	}
}

// 员工登录
type StaffLogin struct {
	*LoginTemplate
}

func NewStaffLogin() *StaffLogin {
	s := &StaffLogin{}
	s.LoginTemplate = NewLoginTemplate(s)

	return s
}

func (s *StaffLogin) findLoginUser(username string) *LoginModel {
	// 不实际获取信息，而是提供一个默认的用户信息
	return &LoginModel{
		userName: "worker1",
		password: "9ea9c3a837014564cef1e29a18794f8bf440f248cf7e78effb22b3c6792b1c49", // worker1_pwd 进行SHA256的结果
	}
}

// 员工登录需要对密码加密
func (s *StaffLogin) encryptPwd(pwd string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(pwd)))
}