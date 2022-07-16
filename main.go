package main

import (
	"encoding/base64"
	"fmt"
)

func reverseBytes(decode []byte) {
	for i := 0; i < len(decode)/2; i++ {
		temp := decode[i]
		decode[i] = decode[len(decode)-1-i]
		decode[len(decode)-1-i] = temp
	}
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	reverseBytes(sd)
	whatIsIt = string(sd)
	fmt.Print(whatIsIt)
}
