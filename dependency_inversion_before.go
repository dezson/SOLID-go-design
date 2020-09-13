package main

import "fmt"

type MySQL struct{}

func (db MySQL) Query() []string {
	return []string{
		"first",
		"second",
		"third",
	}
}

type PostgreSQL struct{}

func (db PostgreSQL) Query() map[string]string {
	return map[string]string{
		"31c7cfcb-4c71-4ea1-abaa-77d6322401d1": "first",
		"6c4fa25c-f5db-11ea-adc1-0242ac120002": "seccond",
		"746144dc-f5db-11ea-adc1-0242ac120002": "third",
	}
}

type UserRepository struct {
	db MySQL
}

func (u UserRepository) GetUsers() []string {
	res := u.db.Query()
	return res
}

func main() {
	mysqlDB := MySQL{}
	u := UserRepository{db: mysqlDB}
	fmt.Printf("Users: %s\n", u.GetUsers())
}
