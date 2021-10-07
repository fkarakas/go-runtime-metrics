package main

import (
	"fmt"
	"time"

	metrics "github.com/tevjef/go-runtime-metrics"
)

// exemple program which bloats the memory and push metrics
func main() {
	metrics.DefaultConfig.CollectionInterval = time.Second
	if err := metrics.RunCollector(metrics.DefaultConfig); err != nil {
		panic(err)
	}

	arr := []string{}

	for {
		// bloat the memory
		arr = append(arr, "test")
		time.Sleep(5 * time.Second)
		fmt.Print(".")
	}
}
