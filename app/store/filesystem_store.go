package store

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemPlayerStore includes a file (database) and a league, and provides read and write actions.
type FileSystemPlayerStore struct {
	DB     *json.Encoder
	League League
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
		DB:     json.NewEncoder(&Tape{file}),
		League: league,
	}, nil
}

func NewFileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file %s, %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	playerStore, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create file system player store, %v", err)
	}

	return playerStore, closeFunc, nil
}

// GetLeague returns the league of current FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.League, func(i, j int) bool {
		return f.League[i].Score > f.League[j].Score
	})
	return f.League
}

// GetPlayerScore returns the given player's score from FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.League.Find(name)

	if player != nil {
		return player.Score
	}

	return 0
}

// PostPlayerScore increments the given player's score at FileSystemPlayerStore
func (f *FileSystemPlayerStore) PostPlayerScore(name string) {

	player := f.League.Find(name)

	if player != nil {
		player.Score++
	} else {
		f.League = append(f.League, Player{Name: name, Score: 1})
	}

	f.DB.Encode(f.League)
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
