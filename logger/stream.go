package logger

import (
	"log"
	"os"

	"github.com/golang-collections/go-datastructures/queue"
)

const (
	// Format holds the default log format
	Format = log.Ldate | log.Ltime
	// QueueLength is the initial length of the queue
	QueueLength = int64(100)
)

// LogLines is the datatype that holds a log unit
type LogLines struct {
	// Lines is the list of strings to be logged
	Lines []string
	// Stop if true will signal the stream to stop
	Stop bool
}

// Stream handler type
type Stream struct {
	fileStream *os.File
	logStream  *log.Logger
	Queue      *queue.Queue
	Alive      bool
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
		Queue:      queue.New(QueueLength),
		Alive:      true,
	}, nil
}

// Start starts a stream
func (s *Stream) Start() {
	for true {
		// Read log from queue
		lli, err := s.Queue.Get(1)
		if err != nil {
			// Queue must have been terminated. return error
			s.close()
			return
		}

		// Convert interface oboect to log object
		ll, ok := lli[0].(LogLines)
		if !ok {
			// Ignore this log and continue loop
			continue
		}

		// Stop the stream is stop signal is set to true
		if ll.Stop {
			s.close()
			return
		}

		// Write logs into the file
		for _, line := range ll.Lines {
			s.logStream.Println(line)
		}

	}
}

func (s *Stream) close() {
	s.Alive = false
	s.fileStream.Close()
}
