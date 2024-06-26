package main

import (
	"bytes"
	"fmt"
)	
func maskLinks(message string) []byte {
	buffer := []byte(message)
	result := bytes.Buffer{}
	i := 0

	for i < len(buffer) {
		if len(buffer)-i >= 7 && isHttp(buffer[i:i+7]) {
			result.WriteString("[LINK REMOVED]")
			i += 7 
			for i < len(buffer) && buffer[i] != ' ' {
				i++
			}
		} else {
			result.WriteByte(buffer[i])
			i++
		}
	}
	return result.Bytes()
}
func isHttp(sub []byte) bool {
	return bytes.EqualFold(sub, []byte("http://"))
}

func main() {
	message := "Here's my spammy page http://hehefouls.netHAHAHA see you."
	maskedMessage := maskLinks(message)
	fmt.Println(string(maskedMessage))
}
