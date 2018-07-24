/*
vnt-to-txt converts vNote 1.1 files to text files from the command-line.

Usage:
	vnt-to-txt [path ...]

Examples

To convert all .vnt files in the current directory to .txt files:
	vnt-to-txt *.vnt
*/
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	vnt "github.com/jybp/go-vnt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	for _, m := range args {
		o := m + ".txt"
		log.Printf("converting %s to %s", m, o)
		file, err := os.Open(m)
		if err != nil {
			log.Fatal(err)
		}
		note, err := vnt.Parse(file)
		if errC := file.Close(); errC != nil {
			log.Fatal(errC)
		}
		if err != nil {
			log.Print(err)
			continue
		}
		if err := ioutil.WriteFile(o, []byte(note.Body), 0700); err != nil {
			log.Print(err)
		}
		if err := os.Chtimes(o, note.LastModified, note.LastModified); err != nil {
			log.Print(err)
		}
	}
	log.Printf("%d files processed", len(args))
}
