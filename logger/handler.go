package logger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type StreamHandler map[string]*Stream

func (sh *StreamHandler) Add(lf string, lp string) error {
	s, err := NewStream(lp)
	if err != nil {
		return fmt.Errorf("Error creating log: %s. Error: %v.\n", lf, err)
	}
	go s.Start()
	(*sh)[lf] = s
	return nil
}

func (sh *StreamHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Get URL parameter
		url := strings.Split(r.URL.Path, "/")
		url = removeBlanks(url)
		if len(url) != 2 {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		// Get stream for log file
		lf := url[1]
		s, ok := (*sh)[lf]
		if !ok {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		// Read logs from request body
		lb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send logs to stream
		for _, ls := range strings.Split(string(lb), "\n") {
			s.channel <- ls
		}

		// Send 201 response
		rw.WriteHeader(http.StatusCreated)
		return

	default:
		rw.WriteHeader(http.StatusNotImplemented)
	}
}
