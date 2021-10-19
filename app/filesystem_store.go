package main

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	DB     *json.Encoder
	league League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)
	return &FileSystemPlayerStore{
		json.NewEncoder(&tape{file}), league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Score
	}

	return 0
}

func (f *FileSystemPlayerStore) PostPlayerScore(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.DB.Encode(f.league)
}
