package main

import (
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)



func main() {

	fmt.Println("Please navigate to localhost:3000/users to check if the server is up and running")

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}