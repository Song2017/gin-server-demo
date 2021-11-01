
package main

import server "apiserver/v1/nomad-api"

func main() {
	router := server.InitRouter()

	router.Run(":8000")
}
