package main

import (
	"github.com/kodykantor/go-statemap/statemap"
	"time"
	"fmt"
)

func main() {
	smap := statemap.New("testmap", "myhost", "")
	smap.SetState("mymain", "starting", "", "orange", time.Now())
	smap.SetState("mymain", "done", "", "blue", time.Now())
	smap.SetState("mymain", "def done now", "", "red", time.Now())

	fmt.Println(smap.Dump())
}
