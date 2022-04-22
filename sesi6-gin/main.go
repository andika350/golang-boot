package main

import "sesi6-gin/routers"

func main() {
	var PORT = ":8081"

	routers.StartServer().Run(PORT)
}