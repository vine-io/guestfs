package main

import (
	"fmt"

	"github.com/vine-io/guestfs"
)

func main() {
	g, errno := guestfs.Create()
	if errno != nil {
		panic(fmt.Sprintf("could not create handle: %s", errno))
	}
	defer g.Close()
	if err := g.Add_drive("test.img", &guestfs.OptargsAdd_drive{}); err != nil {
		panic(err)
	}
	if err := g.Launch(); err != nil {
		panic(err)
	}
	if err := g.Shutdown(); err != nil {
		panic(err)
	}
}
