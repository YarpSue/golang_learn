package sleep_demo

import (
	"flag"
	"time"
)

var period  = flag.Duration("period", 1*time.Second, "sleep period")
/*
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v\n", *period)
	time.Sleep(*period)
	fmt.Println()
}*/