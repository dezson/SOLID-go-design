package main

import "fmt"

type MySQL struct{}

func (db MySQL) Query() interface{} {
    return []string{
        "first",
        "second",
        "third",
    }
}

type PostgreSQL struct{}

func (db PostgreSQL) Query() interface{} {
    return map[string]string{
        "31c7cfcb-4c71-4ea1-abaa-77d6322401d1": "first",
        "6c4fa25c-f5db-11ea-adc1-0242ac120002": "seccond",
        "746144dc-f5db-11ea-adc1-0242ac120002": "third",
    }
}

type DBConn interface {
    Query() interface{}
}

type UserRepository struct {
    db DBConn
}

func (u UserRepository) GetUsers() []string {
    var users []string
    res := u.db.Query()

    switch res.(type) {
    case map[string]string:
        for _, u := range res.(map[string]string) {
            users = append(users, u)
        }
        return users
    case []string:
        return res.([]string)
    }

    return []string{}
}

func main() {
    mysqlDB := MySQL{}
    u := UserRepository{db: mysqlDB}
    fmt.Printf("Users: %s\n", u.GetUsers())
}
