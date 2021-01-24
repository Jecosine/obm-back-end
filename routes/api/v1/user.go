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
	if models.ExistUserByID(id) {
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
	var user models.User
	status := e.UNAUTHORIZED
	log.Printf("[INFO] Try query %s", c.Param("id"))
	if err := c.ShouldBindJSON(&user); err == nil {
		log.Printf("[INFO] In '/api/v1/user.go': user id %s", user.ID)
		log.Printf("[INFO] In '/api/v1/user.go': user name %s", user.Username)
		log.Printf("[INFO] In '/api/v1/user.go': user password %s", user.Password)
		models.AddUser(user)
	} else {
		log.Printf("[ERROR] Error binding model when adding user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": status,
			"msg":    e.GetMsg(status),
		})
	}

}

// EditUser edit user
func EditUser(c *gin.Context) {
	var user models.User
	status := e.UNAUTHORIZED
	id := c.Param("id")
	log.Printf("[INFO] Try modifiying %s", c.Param("id"))
	if err := c.ShouldBindJSON(&user); err == nil {
		log.Printf("[INFO] In '/api/v1/user.go': user id %s", user.ID)
		log.Printf("[INFO] In '/api/v1/user.go': user name %s", user.Username)
		log.Printf("[INFO] In '/api/v1/user.go': user password %s", user.Password)
		models.EditUser(id, user)
	} else {
		log.Printf("[ERROR] Error binding model when adding user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": status,
			"msg":    e.GetMsg(status),
		})
	}
}

// DeleteUser delete user
func DeleteUser(c *gin.Context) {
	var user models.User
	status := e.UNAUTHORIZED
	id := c.Param("id")
	log.Printf("[INFO] Try modifiying %s", c.Param("id"))
	if len(id) != 0 {
		log.Printf("[INFO] In '/api/v1/user.go': user id %s", user.ID)
		log.Printf("[INFO] In '/api/v1/user.go': user name %s", user.Username)
		log.Printf("[INFO] In '/api/v1/user.go': user password %s", user.Password)
		models.DeleteUser(id)
	} else {
		log.Printf("[ERROR] Error binding model when adding user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": status,
			"msg":    e.GetMsg(status),
		})
	}
}
