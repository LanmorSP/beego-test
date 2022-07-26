package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserTable_20220726_073906 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserTable_20220726_073906{}
	m.Created = "20220726_073906"

	migration.Register("UserTable_20220726_073906", m)
}

// Run the migrations
func (m *UserTable_20220726_073906) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update)
	m.SQL(`
	CREATE TABLE public.user (
		id serial NOT NULL PRIMARY KEY,
		username character(255) NOT NULL,
		profile character(255) NOT NULL
	);`)
}

// Reverse the migrations
func (m *UserTable_20220726_073906) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE public.user")
}
