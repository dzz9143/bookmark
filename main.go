package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var bookmark *Bookmark
var indexFile *os.File

func init() {
	bookmark = NewBookmark()
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("HOMEPATH")
		if home == "" {
			log.Fatalln("can not spot home directory, exit...")
			os.Exit(1)
		}
	}
	bmPath := filepath.Join(home, ".bookmark")
	err := os.MkdirAll(bmPath, 0755)
	if err != nil {
		log.Fatalf("can not create .bookmark directory: %s\n", err.Error())
	}

	indexPath := filepath.Join(bmPath, "index")
	indexFile, err := os.OpenFile(indexPath, os.O_CREATE|os.O_RDWR, 0755)

	scanner := bufio.NewScanner(indexFile)

	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		ss := strings.Split(s, " ")
		bookmark.index.Add(ss[0], ss[1:]...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("fail to read from bookmark index: %s\n", err.Error())
	}

}

func main() {
	// b := NewBookmark()
	// b.Add("google", "www.google.com", "search")
	// b.Add("baidu", "www.baidu.com", "search", "china")
	// b.Add("bing", "www.bing.com", "search", "microsoft")
	// b.Add("cnbing", "cn.bing.com", "search", "microsoft", "china")

	// r := b.Query(os.Args[1:])
	// if len(r) > 0 {
	// 	fmt.Println("Search result:")
	// }
	// for i, kv := range r {
	// 	fmt.Printf("[#%d] %s: %s\n", i, kv.Key(), kv.Value())
	// }
}
