package dob

import (
	"os"
	"time"
)

// FileWatcherTarget Implements Target
type FileWatcherTarget struct {
	Filename string
	LastEdit time.Time

	deps []Target
}

func (t *FileWatcherTarget) GetDeps() []Target {
	return t.deps
}
func (t *FileWatcherTarget) AddDependencies(toAdd []Target) {
	t.deps = append(t.deps, toAdd...)
}

// A file watcher cannot ever fail -- it only triggers a rebuild
func (t *FileWatcherTarget) Build() bool {
	return true
}

// check if the file changed
func (t *FileWatcherTarget) NeedsRebuild() bool {
	// get new file info
	fileinfo, err := os.Stat(t.Filename)
	if err != nil {
		return false
	}
	if fileinfo.ModTime().After(t.LastEdit) {
		return true
	}
	return false
}

func NewFileWatcherTarget(filename string) (*FileWatcherTarget, error) {

	// make sure the file exists
	fileinfo, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	return &FileWatcherTarget{
		Filename: filename,
		LastEdit: fileinfo.ModTime(),
	}, nil
}
