package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if os.Args != nil && len(os.Args) > 3 {
			if (os.Args[1] == "e") {
				fmt.Println("encoding exe to txt")
				encode(os.Args[2], os.Args[3])
			} else if (os.Args[1] == "d") {
				fmt.Println("decoding txt to exe")
				decode(os.Args[2], os.Args[3])
			} else {
				fmt.Println("First argument must be 'd' (decode) or 'e' (encode)")
			}
	} else {
		fmt.Println("arg1: 'd' (decode) or 'e' (encode)")
		fmt.Println("arg2: path to exeFile")
		fmt.Println("arg3: path to txtFile")
	}
}

func encode(exeFile, base64TxtFile string) {
	f, err := os.Open(exeFile)
	if(err!= nil) {
		fmt.Print(err)
	}
	// Read entire EXE into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	f2, _ := os.Create(base64TxtFile)
	f2.Write([]byte(encoded))
}

func decode(exeFile, base64TxtFile string) {
	// Read entire base64 TXT into byte slice.
	encoded2, _ := ioutil.ReadFile(base64TxtFile)
	// Decode base64 Txt to working EXE binary
	base64TxtToExe, _ := base64.StdEncoding.DecodeString(string(encoded2))
	f3, _ := os.Create(exeFile)
	f3.Write(base64TxtToExe)
}
