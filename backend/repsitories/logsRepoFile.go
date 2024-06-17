package repsitories

import (
	"bufio"
	"os"
	"sync"

	"github.com/XDBerry29/monitor-app/models"
	"github.com/XDBerry29/monitor-app/utils"
)

type LogRepoFile struct {
	file   *os.File
	writer *bufio.Writer
	mu     sync.Mutex
}

func NewLogRepoFile(file *os.File) LogRepository[*os.File] {

	return &LogRepoFile{file: file, writer: bufio.NewWriter(file)}
}

func (l *LogRepoFile) SaveLog(log *models.Log) error {
	_, err := l.writer.WriteString(utils.LogToWriteString(log))
	if err != nil {
		return err
	}
	// Flush the buffered data to the file
	if err := l.writer.Flush(); err != nil {
		return err
	}
	return nil
}

func (l *LogRepoFile) UpdateRepo(file *os.File) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.file.Close()
	l.file = file
	l.writer = bufio.NewWriter(file)
	return nil
}
