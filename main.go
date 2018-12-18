/*
 *  @Author Muhammad Edwin (muhammad.edwin@ruma.co.id)
 *
 *
 */

package main

import (
	_ "testbdd/controllers"
	_ "testbdd/models"
	_ "testbdd/routers"

	"github.com/astaxie/beego"
)

func init() {
}

func main() {
	beego.Run()
}
