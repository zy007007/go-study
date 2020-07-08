package main

import (
	"gin_demo/http"
	"gin_demo/manager"
	"gin_demo/controller"
)

main() {
	manager.Start()
	controller.Init()
	http.Start()
}
