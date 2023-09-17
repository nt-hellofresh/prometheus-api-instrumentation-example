package main

import "observability/pkg"

func main() {
	pkg.RegisterRoutes()
	pkg.Start(":8080")
}
