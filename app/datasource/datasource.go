package datasource

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB gorm.DB

type Animal struct {
    ID   int64
    Name string `gorm:"default:'galeone'"`
    Age  int64
}

func AutoInitDB() {
    InitGormDB("user", "password", "localhost", "3306", "gorm_issue", "mysql")
}

func InitGormDB(user string, password string, host string, port string, database string, driverName string) {
    DB, err := gorm.Open(driverName, user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    defer DB.Close()

    // Migrate the schema
    println("Migrating 'Animal'")
    DB.AutoMigrate(&Animal{})
}
