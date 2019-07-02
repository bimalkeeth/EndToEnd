package threads

import "time"

func Sender(ch chan string, done chan bool) {
	t := time.Tick(100 * time.Microsecond)
	for {
		select {
		case <-done:
			ch <- "Sender done"
			return
		case <-t:
			ch <- "tick"
		}
	}
}
