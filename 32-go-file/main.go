package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var LineSplit = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	/*read some*/
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	/*find the index of the byte '\n'
	  and find another line begin i+1
	  default token doesn't include '\n'*/
	if i := bytes.IndexByte(data, '\n'); i > 0 {
		return i + 1, dropCR(data[0:i]), nil
	}

	/*at EOF, we have a final, non-terminal line*/
	if atEOF {
		return len(data), dropCR(data), nil
	}

	/*read some more*/
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	/*drop the '\f'
	  if you don't need, you can delete it*/
	if i := bytes.IndexByte(data, '\f'); i >= 0 {
		tmp := [][]byte{data[0:i], data[(i + 1):]}
		sep := []byte("")
		data = bytes.Join(tmp, sep)
	}
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func useSplit(filename string) {
	var count int = 0

	fin, error := os.OpenFile(filename, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer fin.Close()

	sc := bufio.NewScanner(fin)
	/*Specifies the matching function, default read by lines*/
	sc.Split(LineSplit)
	/*begin scan*/
	for sc.Scan() {
		count++
		fmt.Printf("the line %d: %s\n", count, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Println("An error has hippened")
	}
}

func main() {
	fmt.Println(432000000 / 1000 / 60 / 60 / 24)
	//filename := "./32-go-file/test.rvt"
	//fin, error := os.OpenFile(filename, os.O_RDONLY, 0)
	//if error != nil {
	//	panic(error)
	//}
	//defer fin.Close()
	//
	//sc := bufio.NewScanner(fin)
	///*Specifies the matching function, default read by lines*/
	//sc.Split(LineSplit)
	///*begin scan*/
	//for i := 0; i < 20; i++ {
	//	fmt.Printf("the line %d: %s\n", i, sc.Text())
	//}
	//if err := sc.Err(); err != nil{
	//	fmt.Println("An error has hippened")
	//}
	//f, err := ioutil.ReadFile("./32-go-file/test.rvt")
	//if err != nil {
	//	fmt.Println("read fail", err)
	//}
	//r := regexp.MustCompile("(Autodesk Revit 20\\d{2})")
	//s := r.Find(f)
	//str := string(s)
	//fmt.Println(str)
}
