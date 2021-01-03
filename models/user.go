/*
 * @Author: Jecosine
 * @Date: 2021-01-04 02:44:44
 * @LastEditTime: 2021-01-04 07:22:22
 * @LastEditors: Jecosine
 * @Description: User model
 */
package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	ID       string `gorm:"primary_key" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

// set create time before insertion
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	log.Println("[INFO] Really run? %s", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

// set modified time before update
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	// scope.SetColumn("ModifiedTime", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
func ExistUserById(id string) bool {
	var user User
	db.Select("id").Where("id=?", id).First(&user)
	if len(user.ID) != 0 {
		return true
	}
	return false
}

// GetUsers Get users from database
func GetUsers() {

}

// GetUser Get single user by specified id
func GetUser(id string) (user User) {
	db.Where("id=?", id).First(&user)
	log.Printf("[INFO] In '/model/user.go': user id %s", user.ID)
	log.Printf("[INFO] In '/model/user.go': user name %s", user.Username)
	log.Printf("[INFO] In '/model/user.go': user password %s", user.Password)
	return user
}

// AddUser add user
// func AddUser(data map[string]interface{}) bool {
// 	db.Create(&User{
// 		ID:       data["id"].(string),
// 		Username: data["username"].(string),
// 		Password: data["password"].(string),
// 		Avatar:   data["avatar"].(string),
// 		Email:    data["email"].(string),
// 	})
// 	return true
// }

func AddUser(user User) bool {
	log.Printf("[INFO] In '/model/user.go': user id %s", user.ID)
	log.Printf("[INFO] In '/model/user.go': user name %s", user.Username)
	log.Printf("[INFO] In '/model/user.go': user password %s", user.Password)
	db.Create(&user)
	return true
}

// EditUser edit user
func EditUser(id string, data interface{}) bool {
	db.Model(&User{}).Where("id=?", id).Updates(data)
	return true
}

// DeleteUser delete user
func DeleteUser(id string) bool {
	db.Where("id=?", id).Delete(User{})
	return true
}
