package test

import (
	"Gee/geeorm/geeEngine"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	//engine, err := geeEngine.NewEngine("sqlite3", "Gee.db")
	engine, _ := geeEngine.NewEngine("sqlite3", "Gee.db")

	//if err != nil {
	//
	//}
	s := engine.NewSession()
	s.Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
