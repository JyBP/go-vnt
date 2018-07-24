package vnt

import (
	"bufio"
	"errors"
	"io"
	"mime"
	"strings"
	"time"
)

type Note struct {
	Body         string
	Created      time.Time
	LastModified time.Time
}

var InvalidFormat error = errors.New("invalid vNote 1.1 format")

func Parse(r io.Reader) (Note, error) {
	scanner := bufio.NewScanner(r)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 6 || lines[0] != "BEGIN:VNOTE" || lines[1] != "VERSION:1.1" || lines[5] != "END:VNOTE" {
		return Note{}, InvalidFormat
	}

	idx := strings.Index(lines[2], ":")
	if idx == -1 {
		return Note{}, InvalidFormat
	}

	metadatas := strings.Split(lines[2][:idx], ";")
	if len(metadatas) != 3 || metadatas[0] != "BODY" {
		return Note{}, InvalidFormat
	}
	if metadatas[1] != "CHARSET=UTF-8" {
		return Note{}, errors.New("unsuported charset")
	}

	body := lines[2][idx+1:]
	if metadatas[2] == "ENCODING=QUOTED-PRINTABLE" {
		var err error
		if body, err = new(mime.WordDecoder).Decode("=?UTF-8?Q?" + body + "?="); err != nil {
			return Note{}, err
		}
	}

	created := strings.TrimPrefix(lines[3], "DCREATED:")
	createdAt, err := time.Parse("20060102T150405", created)
	if err != nil {
		return Note{}, errors.New("invalid created date format")
	}

	lastModified := strings.TrimPrefix(lines[4], "LAST-MODIFIED:")
	lastModifiedAt, err := time.Parse("20060102T150405", lastModified)
	if err != nil {
		return Note{}, errors.New("invalid last modified date format")
	}

	return Note{body, createdAt, lastModifiedAt}, nil
}
