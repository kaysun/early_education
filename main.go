package main

import (
	"github.com/circle/early_education/routers"
)

func main() {
	r := routers.Routers()
	r.Run()
}
