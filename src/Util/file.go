package util

/*
File TODO  :
	read Text file
	Store the data
	other deatures will come soon

Example :
	f := File {
		"file_path",
		"Data"
	}
*/
import (
	_ "embed"
	"fmt"
	"io/ioutil"

	config "github.com/Totoies/Totoies/src/Config"
)

/*
FilePath
*/
type File struct {
	FilePath string
	Data     string
}

/*
TODO : read the file and store the data inside Data attribute
Return ---------
Error - if any error occure while reading the file
*/
func (f *File) GetFile() *string {

	LogPrintln(config.Enviourment, "Reading the file : ", f.FilePath)
	// Read the file

	data, err := ioutil.ReadFile(f.FilePath)
	if err != nil {
		LogPrintln(config.Enviourment, "Error reading file: ", err.Error())
		return nil
	}
	// Convert the byte sice to a string
	f.Data = string(data)
	LogPrintln(config.Enviourment, "Data of the file : ", f.Data)

	return &f.Data
}

func (f *File) GetFileRuntime() *string {
	LogPrintln(config.Enviourment, "Reading the file : ", f.FilePath)
	// Read the file

	data, err := ioutil.ReadFile(f.FilePath)
	if err != nil {
		LogPrintln(config.Enviourment, "Error reading file: ", err.Error())
		return nil
	}
	// Convert the byte sice to a string
	f.Data = string(data)
	LogPrintln(config.Enviourment, "Data of the file : ", f.Data)

	return &f.Data
}

// Print the data of thefile
func (f *File) PrintFile() {
	fmt.Println(f.Data)
}
