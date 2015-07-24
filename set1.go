package main

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
	"log"
	"errors"
)

func main() {
	challenge1()
	challenge2()
}

func challenge1() {
	// source hex string = 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
	// target b64 string = SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

	const hex = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64,err := HexStringToBase64(hex)
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println(b64)
}

func challenge2() {
	xor,err := fixedXOR("1c0111001f010100061a024b53535009181c","686974207468652062756c6c277320657965")
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println(xor)
}

// interprets two input strings as hex, xors the contents and returns a hex encoding of the result
func fixedXOR(string1 string, string2 string) (string, error) {
	if len(string1) != len(string2) {
		return "", errors.New("input strings must be of equal length")
	}
	str1,_ := hex.DecodeString(string1)
	str2,_ := hex.DecodeString(string2)
	var retstr = make([]byte, len(str1))
	for i := 0; i < len(str1); i++ {
		retstr[i] = str1[i]^str2[i]
		//fmt.Printf("\n%x xored with %x is %x\n", str1[i], str2[i], str1[i]^str2[i])
	}
	return hex.EncodeToString(retstr),nil
}

// returns a base64 encoding of a hex string
func HexStringToBase64(instr string) (string, error) {
	a,err := hex.DecodeString(instr)
	if err != nil {
 		return "", err
	} 
	return base64.StdEncoding.EncodeToString(a), nil
}