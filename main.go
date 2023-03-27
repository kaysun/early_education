package main

import (
	"github.com/circle/early_education/global"
	"github.com/circle/early_education/transport/routers"
)

func main() {
	global.InitMySQL()

	r := routers.Routers()
	r.Run()
}
