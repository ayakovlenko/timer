package timer

import (
	"fmt"
	"os"
	"sync"
	"time"
	"timer/internal/events"
)

const startMsg = "Timer set for %v minutes"

const endMsg = "Time is up! You diligently worked for %v minutes. Good job!"

func New(minutes int) error {
	if err := events.Notify(fmt.Sprintf(startMsg, minutes)); err != nil {
		return err
	}

	startTime := time.Now()

	printTicker := time.NewTicker(time.Duration(1) * time.Second)
	alertTimer := time.NewTimer(time.Duration(minutes) * time.Minute)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-printTicker.C:
				fmt.Printf("\r%v", time.Since(startTime).Round(time.Second))
			case <-alertTimer.C:
				if err := events.Alert(fmt.Sprintf(endMsg, minutes)); err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}

				return
			}
		}
	}()

	wg.Wait()

	return nil
}
