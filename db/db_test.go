package db

import (
	"testing"

	lua "github.com/yuin/gopher-lua"
)

func TestSqlServer(t *testing.T) {
	l := lua.NewState()
	Preload(l)
	err := l.DoFile("test/test_db_sqlserver.lua")
	if err != nil {
		t.Error(err)
	}
}
