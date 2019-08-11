package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Atom representation
type Atom struct {
	Pos      int64
	Size     int64
	AtomType string
	Depth    int
}

var buf = make([]byte, 8)

func main() {
	mp4FilePath := os.Args[1]
	file, err := os.Open(mp4FilePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	var absPosition int64
	var seekError error
	for {

		absPosition, seekError = file.Seek(absPosition, 0)
		reader.Reset(file)
		if seekError != nil {
			log.Fatal(seekError)
		}
		b, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF && b == 0 {
				fmt.Println("All atoms succefully readed.",err)
				break
			}
		}
		atomSize := int64(binary.BigEndian.Uint32(buf[:8/2]))
		atomType := string(buf[4:8])
		atom := Atom{absPosition + 8, atomSize, atomType, 0}
		absPosition = absPosition + atomSize
		showChildAtoms(file, reader, atom)
	}
}

func showChildAtoms(file *os.File, reader *bufio.Reader, atom Atom) {
	queue := []Atom{atom}
	for len(queue) > 0 {
		childs := []Atom{}
		poppedAtom := queue[0]
		queue = queue[1:len(queue)]

		printAtom(poppedAtom)
		if !leafAtom(poppedAtom.AtomType) {
			childs, _ = getAtomChilds(file, reader, poppedAtom)
		}
		queue = append(childs, queue...)
	}
}

func getAtomChilds(file *os.File, reader *bufio.Reader, parentAtom Atom) ([]Atom, error) {
	childs := []Atom{}
	var childsError error
	var seekError error

	absPosition := parentAtom.Pos
	for {
		absPosition, seekError = file.Seek(absPosition, 0)
		if seekError != nil {
			childsError = seekError
			break
		}
		reader.Reset(file)

		b, readError := reader.Read(buf)
		if readError != nil {
			if readError == io.EOF && b == 0 {
				childsError = readError
				break
			}
		}
		atomSize := int64(binary.BigEndian.Uint32(buf[:8/2]))
		atomType := string(buf[4:8])

		childAtom := Atom{absPosition + 8, atomSize, atomType, parentAtom.Depth + 1}
		absPosition = absPosition + atomSize

		if childAtom.Size < parentAtom.Size && childAtom.Size != 0 {
			childs = append(childs, childAtom)
		} else {
			break
		}
	}
	return childs, childsError
}

func printAtom(atom Atom) {
	depthTabs := strings.Repeat("\t", atom.Depth)

	fmt.Print(depthTabs+"Atom type: ", atom.AtomType)
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
