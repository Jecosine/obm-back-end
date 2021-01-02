/*
 * @Author: Jecosine
 * @Date: 2021-01-02 06:46:30
 * @LastEditTime: 2021-01-03 07:11:54
 * @LastEditors: Jecosine
 * @Description: Initialize models and connect database
 */

package models

import (
	"fmt"
	"log"

	"github.com/Jecosine/obm-back-end/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model base model
type Model struct {
	ID string `gorm:"primary_key" json:"id"`
}

type User struct {
	ID          string `gorm:"primary_key" json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CreatedTime string `json:"createdTime" gorm:"column:createdTime"`
}

func Init() {
	var (
		e                                                       error
		dbType, dbUser, dbPasswd, dbName, dbHost, dbTablePrefix string
	)
	dbSection, e := setting.Cfg.GetSection("database")
	if e != nil {
		log.Fatal(2, "Fail to get section database in app.ini: %v", e)
	}
	dbType = dbSection.Key("TYPE").String()
	dbUser = dbSection.Key("USER").String()
	dbName = dbSection.Key("NAME").String()
	dbPasswd = dbSection.Key("PASSWORD").String()
	dbHost = dbSection.Key("HOST").String()
	dbTablePrefix = dbSection.Key("TABLE_PREFIX").String()

	db, e = gorm.Open(
		dbType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser,
			dbPasswd,
			dbHost,
			dbName,
		),
	)
	if e != nil {
		log.Fatalf("[ERROR] Database connecting error: %v", e)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return dbTablePrefix + tableName
	}
	// configure database
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(12)
	db.DB().SetMaxOpenConns(100)
}

func TestDatabase() ([]*User, error) {
	var users []*User
	e := db.Select("*").Find(&users).Error
	return users, e
}

// CloseDataBase close database
func CloseDataBase() {
	defer db.Close()
}
