package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	asciiByt, err := processFile("/run/media/foo/mistral-7b-instruct-v0.2.Q4_K_M.gguf", 1000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print(bytesToASCII(asciiByt))
}

func processFile(filePath string, byteSize int) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var fileData []byte

	for i := 0; i < byteSize; i++ {
		byteVal, err := reader.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("error reading byte: %v", err)
		}
		fileData = append(fileData, byteVal)
	}

	return fileData, nil
}

func bytesToASCII(file []byte) string {
	var asciiStr strings.Builder

	for i := 0; i < len(file); i += 4 {
		word1 := uint16(file[i])<<8 | uint16(file[i+1])
		word2 := uint16(file[i+2])<<8 | uint16(file[i+3])
		pair := fmt.Sprintf("%04X%04X", word1, word2)

		bytesData, err := hex.DecodeString(pair)
		if err != nil {
			fmt.Println("Error decoding hex string:", err)
			continue
		}

		for _, byteVal := range bytesData {
			if byteVal >= 32 && byteVal <= 126 {
				asciiStr.WriteByte(byteVal)
			} else {
				asciiStr.WriteByte('.')
			}
		}
	}

	return asciiStr.String()
}
