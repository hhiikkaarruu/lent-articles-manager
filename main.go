package main

import "github.com/hhiikkaarruu/lent-articles-manager/controller"

func main() {
	r := controller.GetRouter()
	r.Run(":8080")
}
