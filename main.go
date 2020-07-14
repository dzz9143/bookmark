package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var bookmark *Bookmark
var dataFile *os.File

func init() {
	bookmark = NewBookmark()
	homePath := os.Getenv("HOME")
	if homePath == "" {
		homePath = os.Getenv("HOMEPATH")
		if homePath == "" {
			log.Fatalln("can not find home directory, exit")
			os.Exit(1)
		}
	}

	dataPath := filepath.Join(homePath, ".bookmark")
	f, err := os.OpenFile(dataPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)

	dataFile = f

	if err != nil {
		log.Fatalf("fail to open bookmark open file: %s\n", err.Error())
	}

	scanner := bufio.NewScanner(dataFile)

	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		ss := strings.Split(s, " ")
		l := len(ss)
		if l < 2 {
			continue
		}

		key, value := ss[0], ss[1]
		var tags []string
		if l > 2 {
			tags = ss[2:]
		}

		bookmark.Add(key, value, tags...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("fail to load bookmark data: %s\n", err.Error())
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage:
1. To add: bm -a <key> <value> <tag1> <tag2> ... <tagN>
2. To query: bm <keyword1> ... <keywordN>
`)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		usage()
		return
	}

	cmd := args[0]

	if cmd == "-a" {
		// handle add to bookmark
		if len(args) < 3 {
			fmt.Fprintf(os.Stderr, "when add to bookmark, <key> <value> must be specified\n\n")
			usage()
			return
		}

		dataFile.WriteString(strings.Join(args[1:], " ") + "\n")
		defer dataFile.Close()
		return
	}

	// handle query
	result := bookmark.Query(args[:])

	if len(result) < 1 {
		return
	}

	maxLen := 0

	for _, kv := range result {
		if l := len(kv.Key()); l > maxLen {
			maxLen = l
		}
	}

	for _, kv := range result {
		k := kv.Key()
		format := "%-" + strconv.Itoa(maxLen) + "s | %s\n"
		fmt.Printf(format, k, kv.Value())
	}
}
