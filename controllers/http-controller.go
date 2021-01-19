package controllers

import (
	"fmt"

	"aot-technologies.com/saltsecurities/echo/models"
)

// PostHTTPRequest is ..
func PostHTTPRequest(req models.HTTPRequest) string {

	fmt.Println(req)
	return "OK"

}
