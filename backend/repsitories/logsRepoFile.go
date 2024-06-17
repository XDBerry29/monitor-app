package repsitories

import (
	"bufio"
	"os"

	"github.com/XDBerry29/monitor-app/models"
	"github.com/XDBerry29/monitor-app/utils"
)

type logRepoFile struct {
	file   *os.File
	writer *bufio.Writer
}

func NewLogRepoFile(file *os.File) LogRepository {

	return &logRepoFile{file: file, writer: bufio.NewWriter(file)}
}

func (l *logRepoFile) SaveLog(log *models.Log) error {
	_, err := l.writer.WriteString(utils.LogToWriteString(log))
	if err != nil {
		return err
	}
	return nil
}
