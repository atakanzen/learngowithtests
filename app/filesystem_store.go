package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemPlayerStore includes a file (database) and a league, and provides read and write actions.
type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

// NewFileSystemPlayerStore creates a FileSystemPlayerStore with given file.
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not initialise db file %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("could not get league with file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		db:     json.NewEncoder(&tape{file}),
		league: league,
	}, nil
}

// GetLeague returns the league of current FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Score > f.league[j].Score
	})
	return f.league
}

// GetPlayerScore returns the given player's score from FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Score
	}

	return 0
}

// PostPlayerScore increments the given player's score at FileSystemPlayerStore
func (f *FileSystemPlayerStore) PostPlayerScore(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.db.Encode(f.league)
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("could not get info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}
