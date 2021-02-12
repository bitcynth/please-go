package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type passwdEntry struct {
	name   string
	passwd string
	uid    int
	gid    int
	gecos  string
	dir    string
	shell  string
}

func parsePasswd() ([]passwdEntry, error) {
	passwdBytes, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		return nil, err
	}

	passwdLines := strings.Split(string(passwdBytes), "\n")

	var entries []passwdEntry
	for _, line := range passwdLines {
		parts := strings.Split(line, ":")
		if len(parts) != 7 {
			continue
		}

		uid, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}
		gid, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, err
		}

		entry := passwdEntry{
			name:   parts[0],
			passwd: parts[1],
			uid:    uid,
			gid:    gid,
			gecos:  parts[4],
			dir:    parts[5],
			shell:  parts[6],
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func getPwNam(name string) *passwdEntry {
	entries, err := parsePasswd()
	if err != nil {
		os.Exit(-1)
	}

	for _, entry := range entries {
		if entry.name == name {
			return &entry
		}
	}

	return nil
}
