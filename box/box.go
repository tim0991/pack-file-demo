package box

import (
	"errors"
)

var defaultBox *Box

func init() {
	defaultBox = &Box{
		files: make(map[string][]byte),
	}
}

type Box struct {
	files map[string][]byte
}

func (b Box) Get(filename string) (bt []byte, err error) {
	bt, ok := b.files[filename]
	if !ok {
		err = errors.New("file not found")
	}
	return
}

func (b Box) Add(filename string, bt []byte) {
	b.files[filename] = bt
	return
}

func GetDefault() *Box {
	return defaultBox
}

func Add(filename string, bt []byte) {
	defaultBox.Add(filename, bt)
}
