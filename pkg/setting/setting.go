/*
 * @Author: Jecosine
 * @Date: 2021-01-02 05:01:49
 * @LastEditTime: 2021-01-02 05:29:44
 * @LastEditors: Jecosine
 * @Description:Settings
 */

package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	// Cfg config file
	Cfg *ini.File
	// RunMode run mode
	RunMode string
	// Name app name
	Name string
	// HTTPPort running port
	HTTPPort int
)

func Setup() {
	var e error
	Cfg, e = ini.Load("conf/app.ini")
	if e != nil {
		log.Fatalf("[ERROR] Fail to load conf/app.ini: %v", e)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	Name = Cfg.Section("app").Key("app").MustString("NAME")
	HTTPPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(8888)
	log.Printf("RUN_MODE: %v\nNAME: %v\nHTTP_PORT: %v", RunMode, Name, HTTPPort)
}
