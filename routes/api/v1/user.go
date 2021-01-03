package v1

import (
	"log"
	"net/http"

	"github.com/Jecosine/obm-back-end/models"
	"github.com/Jecosine/obm-back-end/pkg/e"
	"github.com/gin-gonic/gin"
)

// GetUsers Get users from database
func GetUsers(c *gin.Context) {

}

// GetUser Get single user by specified id
func GetUser(c *gin.Context) {
	id := string(c.Param("id"))
	log.Printf("[INFO] In 'routes/api/v1/user.go': Get url param id %v", id)
	status := e.ERROR
	var data interface{}
	if models.ExistUserById(id) {
		data = models.GetUser(id)
		status = e.SUCCESS
	} else {
		status = e.NOTFOUND
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    e.GetMsg(status),
		"data":   data,
	})
}

// AddUser add user
func AddUser(c *gin.Context) {

}

// EditUser edit user
func EditUser(c *gin.Context) {

}

// DeleteUser delete user
func DeleteUser(c *gin.Context) {

}
