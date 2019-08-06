package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

// Atom representation
type Atom struct {
	Pos      int64
	Size     int64
	AtomType string
}

var buf = make([]byte, 8)

func main() {
	mp4FilePath := os.Args[1]
	file := openFile(mp4FilePath)

	var absPosition int64
	for {
		absPosition = seekFile(file, absPosition)
		if !readFile(file, buf) {
			break
		}
		atomSize := int64(binary.BigEndian.Uint32(buf[:8/2]))
		atomType := string(buf[4:8])
		atom := Atom{absPosition + 8, atomSize, atomType}
		absPosition = absPosition + atomSize
		showChildAtoms(file, atom)
	}
}

func showChildAtoms(file *os.File, atom Atom) {
	queue := []Atom{atom}
	for len(queue) > 0 {
		childs := []Atom{}
		poppedAtom := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		printAtom(poppedAtom)
		if !leafAtom(poppedAtom.AtomType) {
			childs = getAtomChilds(file, poppedAtom)
		}
		queue = append(queue, childs...)
	}
}

func getAtomChilds(file *os.File, parentAtom Atom) []Atom {
	childs := []Atom{}
	absPosition := parentAtom.Pos
	for true {
		absPosition = seekFile(file, absPosition)
		if !readFile(file, buf) {
			break
		}
		atomSize := int64(binary.BigEndian.Uint32(buf[:8/2]))
		atomType := string(buf[4:8])

		childAtom := Atom{absPosition + 8, atomSize, atomType}
		absPosition = absPosition + atomSize

		if childAtom.Size < parentAtom.Size && childAtom.Size != 0 {
			childs = append(childs, childAtom)
		} else {
			break
		}
	}
	return childs
}

func openFile(mp4FilePath string) *os.File {
	file, err := os.Open(mp4FilePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func seekFile(file *os.File, absPosition int64) int64 {
	absPosition, seekError := file.Seek(absPosition, 0)
	if seekError != nil {
		log.Fatal(seekError)
	}
	return absPosition
}

func readFile(file *os.File, buf []byte) bool {
	b, err := file.Read(buf)
	if err != nil {
		if err == io.EOF && b == 0 {
			return false
		}
	}
	return true
}

func printAtom(atom Atom) {
	fmt.Print("Atom type: ", atom.AtomType)
	fmt.Println(" size:", atom.Size)
}
func leafAtom(lookup string) bool {
	switch lookup {
	case
		"cmvd",
		"co64",
		"vmhd",
		"dcom",
		"elst",
		"gmhd",
		"hdlr",
		"mdhd",
		"smhd",
		"stco",
		"stsc",
		"stsd",
		"stss",
		"stsz",
		"stts",
		"tkhd":
		return true
	}
	return false
}
