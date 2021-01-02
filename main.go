/*
 * @Author: Jecosine
 * @Date: 2021-01-02 04:15:04
 * @LastEditTime: 2021-01-03 06:52:54
 * @LastEditors: Jecosine
 * @Description: Main Application
 */
package main

import (
	"fmt"

	"github.com/Jecosine/obm-back-end/models"
	"github.com/Jecosine/obm-back-end/pkg/setting"
)

func main() {
	setting.Setup()
	models.Init()
	users, e := models.TestDatabase()
	if e != nil {
		fmt.Printf("error: %v", e)
		return
	}
	for _, user := range users {
		fmt.Println(user)
	}

}
