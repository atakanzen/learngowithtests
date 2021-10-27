package mock

import "time"

type alert struct {
	scheduledAt time.Duration
	betAmount   int
}

type SpyBlindAlerter struct {
	Alerts []alert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, alert{
		duration,
		amount,
	})

}
