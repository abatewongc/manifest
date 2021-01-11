package file

import (
	"io/ioutil"
)

// ReadBytes : Beep beep!
func ReadBytes(filePath string) (data []byte, err error) {
	data, err = ioutil.ReadFile(filePath)
	return
}
