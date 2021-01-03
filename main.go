/*
 * @Author: Jecosine
 * @Date: 2021-01-02 04:15:04
 * @LastEditTime: 2021-01-03 08:17:08
 * @LastEditors: Jecosine
 * @Description: Main Application
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/Jecosine/obm-back-end/models"
	"github.com/Jecosine/obm-back-end/pkg/setting"
	"github.com/Jecosine/obm-back-end/routes"
)

func main() {
	setting.Setup()
	models.Init()
	r := routes.Init()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: r,
		// ReadTimeout:    setting.ReadTimeout,
		// WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
