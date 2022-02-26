package log

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"bufio"
	"regexp"
	"strings"
	"time"
)

// Log represents a log message with metadata
type Log struct {
	Level     string `json:"level"`
	Timestamp int64  `json:"timestamp"`
	Container string `json:"container"`
	Content   string `json:"content"`
}

// Reader reads logs line by line.
type Reader struct {
	levelExp  *regexp.Regexp
	timeExp   *regexp.Regexp
	container string
	reader    *bufio.Reader
}

// NewReader creates a reader that parses logs.
func NewReader(container string, reader *bufio.Reader) *Reader {
	levelExp, _ := regexp.Compile("(level=)[^ ]*")
	timeExp, _ := regexp.Compile("^[^ ]+")
	return &Reader{
		levelExp:  levelExp,
		timeExp:   timeExp,
		container: container,
		reader:    reader,
	}
}

// ReadLog reads a new log entry
func (r *Reader) ReadLog() (*Log, error) {
	bytes, _, err := r.reader.ReadLine()
	if err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		return nil, nil
	}

	content := string(bytes)

	level := strings.ToLower(strings.Replace(r.levelExp.FindString(content), "level=", "", 1))
	if level == "" {
		level = "info"
	}

	timestampString := r.timeExp.FindString(content)
	timestamp, err := time.Parse(time.RFC3339Nano, timestampString)
	if err != nil {
		return nil, err
	}

	return &Log{
		Level:     level,
		Timestamp: timestamp.UnixNano(),
		Container: r.container,
		Content:   strings.TrimPrefix(strings.TrimPrefix(content, timestampString), " "),
	}, nil
}
