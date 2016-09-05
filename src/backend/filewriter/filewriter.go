package filewriter

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ARMList struct which contains an array of ARM instuction set strings
type ARMList []string

// AppendStringToFile appends a string to the code slice in the File struct
func (f ARMList) AppendStringToFile(input string) {
	f = append(f, input)
}

// writes all the strings contained in the File struct to the file
func (f ARMList) WriteToFile(fileName string) {
	file, err := os.Create(fileName)
	check(err)
	defer file.Close()
	for _, str := range f {
		_, err := file.WriteString(str)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(n3)
	}
	file.Sync()
}
