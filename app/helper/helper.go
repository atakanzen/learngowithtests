package helper

import (
	"bytes"
	"io/ioutil"
	"learngowithtests/app/mock"
	"learngowithtests/app/store"
	"os"
	"reflect"
	"testing"
)

func GetLeagueBody(t testing.TB, body *bytes.Buffer) store.League {
	t.Helper()

	league, err := store.NewLeague(body)

	if err != nil {
		t.Error(err)
	}

	return league
}

func AssertLeagueBody(t testing.TB, got, want store.League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertResBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got score %q, want %q", got, want)
	}
}

func AssertResCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}

func AssertContentType(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response did not have content-type of \"application/json\", got %q", got)
	}
}

func AssertScore(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got score %d, want %d", got, want)
	}
}

func AssertErrNil(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("did not expect error, but got one %q,", got)
	}
}

func AssertScheduledAlert(t testing.TB, got, want mock.ScheduledAlert) {
	amountGot := got.BetAmount
	if amountGot != want.BetAmount {
		t.Fatalf("got amount %d, want %d", amountGot, want.BetAmount)
	}

	gotScheduledTime := got.At
	if gotScheduledTime != want.At {
		t.Fatalf("got scheduled time of %v, want %v", gotScheduledTime, want.At)
	}
}

func AssertPlayerWin(t testing.TB, playerStore *mock.StubPlayerStore, winner string) {
	t.Helper()

	if len(playerStore.WinCalls) != 1 {
		t.Errorf("got %d calls to PostPlayerScore, want %d", len(playerStore.WinCalls), 1)
	}

	if playerStore.WinCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q want %q", playerStore.WinCalls[0], winner)
	}
}

func CreateTmpFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create tmp file, %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
