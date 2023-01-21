package main

import (
	"github.com/circle/early_education/transport/routers"
)

func main() {
	r := routers.Routers()
	r.Run()
}
