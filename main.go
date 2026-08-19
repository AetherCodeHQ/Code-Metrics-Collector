package main

import (
	"fmt"
	"os"
)

// code_metrics_collector - Collect code metrics
func code_metrics_collector(path string) {
	fmt.Println("========================================")
	fmt.Println("  Code-Metrics-Collector")
	fmt.Println("  Collect code metrics")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	code_metrics_collector(path)
}
