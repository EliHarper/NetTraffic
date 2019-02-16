package main

import (
	"fmt"
	"time"
)

// func timer() {
// 	var f func()
// 	var t *time.Timer

// 	f = func() {
// 		main()
// 		t = time.AfterFunc(time.Duration(5)*time.Second, f)
// 	}

// 	t = time.AfterFunc(time.Duration(5)*time.Second, f)

// 	defer t.Stop()

// 	//simulate doing stuff
// 	time.Sleep(time.Minute)

/* END TIMER; BEGIN TICKER */
func ticker() {

	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			// Call main every 5 min, as defined by ticker;
			// after 15 minutes, the ticker will stop.
			main()
		}
	}()
	time.Sleep(time.Minute * 15)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
