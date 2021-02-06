package logger

import (
	"path"
	"strings"
	"time"

	"github.com/mohitudupa/timber/utility"
)

// Pool type to manage logfiles
type Pool struct {
	env     *utility.Environment
	streams map[string]*Stream
}

// NewPool returns a new instance of Pool struct
func NewPool(env *utility.Environment) *Pool {
	return &Pool{
		env:     env,
		streams: map[string]*Stream{},
	}
}

// Attach adds a new log Stream to the log Pool instance
func (l *Pool) Attach(name string) error {
	f := path.Join(l.env.TimberData, name)
	s, err := NewStream(f)
	if err != nil {
		return err
	}

	// Start stream as a coroutine
	go s.Start()

	// Adding stream to logger
	l.streams[name] = s
	return nil
}

// Detach closes a log stream on the log Pool instance
func (l *Pool) Detach(name string) {
	s, ok := l.streams[name]
	if !ok {
		return
	}

	s.Queue.Put(LogLines{
		Lines: []string{},
		Stop:  true,
	})

	for s.Alive {
		time.Sleep(time.Millisecond * 100)
	}

	delete(l.streams, name)
	return
}

// Log loggs a single line or multiple lines onto a specified log Stream
func (l *Pool) Log(name string, data string, multiline bool) error {
	s, ok := l.streams[name]
	if !ok {
		err := l.Attach(name)
		if err != nil {
			return err
		}
	}

	lines := []string{}

	if multiline {
		lines = strings.Split(data, "\n")
	} else {
		lines = append(lines, strings.ReplaceAll(data, "\n", " "))
	}

	s.Queue.Put(LogLines{
		Lines: lines,
		Stop:  false,
	})
	return nil
}
