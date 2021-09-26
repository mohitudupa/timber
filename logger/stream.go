package logger

import (
	"log"
	"os"
)

const (
	// Format holds the default log format
	Format = log.Ldate | log.Ltime
)

// Stream handler type
type Stream struct {
	fileStream *os.File
	logStream  *log.Logger
	channel    chan string
}

// NewStream returns a new instance of Stream struct
func NewStream(file string) (*Stream, error) {
	var fs, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &Stream{
		fileStream: fs,
		logStream:  log.New(fs, "", log.Ldate|log.Ltime),
		channel:    make(chan string),
	}, nil
}

// Start starts a stream
func (s *Stream) Start() {
	for l := range s.channel {
		s.logStream.Println(l)
	}
	s.fileStream.Close()
}

// Stop the stream
func (s *Stream) Stop() {
	close(s.channel)
}
