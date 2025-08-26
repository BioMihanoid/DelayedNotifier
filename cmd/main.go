package main

import (
	"BioMihanoid/DelayedNotifier/internal/config"
	"fmt"
)

func main() {
	conf := config.NewConfig()
	fmt.Println(conf)
}
