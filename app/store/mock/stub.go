package mock

import "learngowithtests/app/store"

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   store.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) PostPlayerScore(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() store.League {
	return s.League
}
