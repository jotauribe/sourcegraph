package command

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
)

// Logger tracks command invocations and stores the command's output and
// error stream values.
type Logger struct {
	redactedValues []string
	entries        []LogEntry
}

type LogEntry struct {
	Command []string
	Out     string
}

// NewLogger creates a new logger instance with the given redacted values.
// When the log messages are serialized, any occurrence of the values are
// replaced with a canned string.
func NewLogger(redactedValues ...string) *Logger {
	return &Logger{
		redactedValues: redactedValues,
	}
}

// RecordCommand pushes a new command invocation into the logger. The given
// output and error stream readers are read concurrently until completion.
// This method blocks.
func (l *Logger) RecordCommand(command []string, stdout, stderr io.Reader) {
	out := &bytes.Buffer{}
	var m sync.Mutex
	var wg sync.WaitGroup

	readIntoBuf := func(prefix string, r io.Reader) {
		defer wg.Done()

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			m.Lock()
			fmt.Fprintf(out, "%s: %s\n", prefix, scanner.Text())
			m.Unlock()
		}
	}

	wg.Add(2)
	go readIntoBuf("stdout", stdout)
	go readIntoBuf("stderr", stderr)
	wg.Wait()

	payload := out.String()

	for _, v := range l.redactedValues {
		payload = strings.Replace(payload, v, "******", -1)
	}

	l.entries = append(l.entries, LogEntry{
		Command: command,
		Out:     payload,
	})
}

func (l *Logger) Entries() (entries []LogEntry) {
	for _, entry := range l.entries {
		entries = append(entries, entry)
	}

	return entries
}
