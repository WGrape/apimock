package service

import "testing"

func TestGetUser(t *testing.T) {
	if GetUser(876352).Uid != 876352 {
		t.Fail()
		return
	}
}
