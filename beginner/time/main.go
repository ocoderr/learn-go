package time

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%v\n", now)
	fmt.Printf("now.Year %v\n", now.Year())
	fmt.Printf("now.Month %v\n", now.Month())
	fmt.Printf("now.Day %v\n", now.Day())

	fmt.Printf("now.Hour %v\n", now.Hour())
	fmt.Printf("now.Minute %v\n", now.Minute())
	fmt.Printf("now.Second %v\n", now.Second())
	fmt.Printf("now.Nanosecond %v\n", now.Nanosecond())
	fmt.Printf("%v\n", now.Format("2006-01-02 15:04:05")) // 特定的"2006-01-02 15:04:05",

	// calculating times:
	week := 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	weekFromNow := now.Add(time.Duration(week))
	fmt.Println(weekFromNow) // Wed Dec 28 08:52:14 +0000 UTC 2011

}
