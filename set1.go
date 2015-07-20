package main

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
	"log"
)

// source hex string = 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// target b64 string = SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t


func main() {
	challenge1()
	challenge2()
}

func challenge1() {
	const hex = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64,err := HexStringToBase64(hex)
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println(b64)
}

func challenge2() {
	fmt.Println(fixedXOR("1c0111001f010100061a024b53535009181c","686974207468652062756c6c277320657965"))
}

func fixedXOR(string1 string, string2 string) (string, err) {
	str1decoded,err := hex.DecodeString(string1)
	if err != nil {
 		return "", err
	} 
	str2decoded,err := hex.DecodeString(string2)
	if err != nil {
 		return "", err
	} 
	str1 := []byte(string1)
	str2 := []byte(string2)
	for i := 0; i < len(str1)/2; i++ {
		fmt.Println(str1[i]^str2[i])
	}
	return "stuff"
}

func HexStringToBase64(instr string) (string, error) {
	a,err := hex.DecodeString(instr)
	if err != nil {
 		return "", err
	} 
	return base64.StdEncoding.EncodeToString(a), nil
}