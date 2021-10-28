package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(time.Duration, int)
}

type BlindAlerterFunc func(time.Duration, int)

func (b BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, betAmount int) {
	b(duration, betAmount)
}

func StdOutAlerter(duration time.Duration, betAmount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind bet is now %d\n", betAmount)
	})
}
