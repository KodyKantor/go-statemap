package main

import (
	"github.com/kodykantor/go-statemap/statemap"
	"time"
)

func main() {
	smap := statemap.New("testmap", "myhost", "")
	smap.SetState("mymain", "starting", "", time.Now())
	smap.SetState("mymain", "done", "", time.Now())
	smap.SetState("mymain", "def done now", "", time.Now())

	smap.Dump()
}
