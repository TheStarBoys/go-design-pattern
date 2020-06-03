package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {
	normalLogin := NewNormalLogin()
	ok := normalLogin.Login(LoginModel{
		userName: "user1",
		password: "user1_pwd",
	})
	if ok == false {
		t.Errorf("except: true, got: false")
	}
	staffLogin := NewStaffLogin()
	ok = staffLogin.Login(LoginModel{
		userName: "worker1",
		password: "worker1_pwd",
	})
	if ok == false {
		t.Errorf("except: true, got: false")
	}
}
