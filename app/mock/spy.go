package mock

import (
	"fmt"
	"time"
)

type ScheduledAlert struct {
	At        time.Duration
	BetAmount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.BetAmount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{
		at,
		amount,
	})

}
