package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"
	//"strings"
)

func main() {
	fmt.Print("challenge 1: ")
	challenge1()
	fmt.Print("challenge 2: ")
	challenge2()
	fmt.Print("challenge 3: ")
	challenge3()
	fmt.Println("\n")
}

func challenge1() {
	// source hex string = 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
	// target b64 string = SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

	const hex = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64, err := HexStringToBase64(hex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b64)
}

func challenge2() {
	xor, err := fixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", xor)
}

// interprets two input strings as hex, xors the contents and returns a hex encoding of the result
func fixedXOR(string1 string, string2 string) (string, error) {
	if len(string1) != len(string2) {
		return "", errors.New("input strings must be of equal length")
	}
	str1, _ := hex.DecodeString(string1)
	str2, _ := hex.DecodeString(string2)
	var retstr = make([]byte, len(str1))
	for i := range str1 {
		retstr[i] = str1[i] ^ str2[i]
		//fmt.Printf("\n%x xored with %x is %x\n", str1[i], str2[i], str1[i]^str2[i])
	}
	return hex.EncodeToString(retstr), nil
}

// returns a base64 encoding of a hex string
func HexStringToBase64(instr string) (string, error) {
	a, err := hex.DecodeString(instr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(a), nil
}

func challenge3() {
	const cipher = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cipherbytes, _ := hex.DecodeString(cipher)
	//letterfreqs := []float64{.0812, .0149, .0271, .0432, .1202, .023, .0203, .0592, .0731, .001, .0069, .0398, .0261, .0695, .0768, .0182, .0011, .0602, .0628, .091, .0288, .0111, .0209, .0017, .0211, .0007}
	letterfreqs := makefreqs()
	msg := ""
	max := float64(26) // max ratio 100% or 1.00 * 26 letters...
	key := byte(0)
	for b := 0; b < 256; b++ {
		candidateKey := byte(b)
		candidateMsg := tryKey(cipherbytes, candidateKey)
		score := simpleScore(candidateMsg, letterfreqs)
		if score < max {
			max = score
			msg = string(candidateMsg)
			key = candidateKey
		}
		//fmt.Printf("%v\tscore\t%v\n", candidateKey, score)
	}
	fmt.Printf("message: \"%v\", key: %v, score: %v\n", msg, key, max)
}

func tryKey(cipherbytes []byte, key byte) []byte {
	lenDecoded := len(cipherbytes)
	candidateClearHex := make([]byte, lenDecoded)

	for index := 0; index < lenDecoded; index++ {
		//fmt.Printf("index: %v", index)
		//fmt.Printf(", xor: %v; ", decoded[index] ^ candidateKey)
		candidateClearHex[index] = cipherbytes[index] ^ key
	}
	//candidateClearDecoded, _ := hex.DecodeString(candidateClearHex)
	return candidateClearHex
}

// just returns how many bytes are valid A-Z characters (case insensitive)
func simpleScore(msgbytes []byte, freqs map[byte]float64) float64 {
	histogram := make(map[byte]int)
	totalbytes := len(msgbytes)
	for i := 0; i < totalbytes; i++ {
		histogram[msgbytes[i]] = histogram[msgbytes[i]] + 1
	}
	score := float64(0)
	for key,value := range histogram {
		delta := float64(value) / float64(totalbytes) //* float64(100)
		//fmt.Printf("histogram: %v, delta: %v, totalbytes: %v\n", histogram[j], delta, totalbytes)
		targetratio := freqs[key]
		actualratio := delta
		sdelta := math.Abs(targetratio - actualratio)
		//fmt.Printf("for character: %v, target ratio: %v, actual ratio: %v, abs diff: %v", string(byte(j+97)), targetratio, actualratio, sdelta)
		score += sdelta
	}
	return score
}

func makefreqs() map[byte]float64 {
	return map[byte]float64{
		' ':  0.0000189169,
		'!':  0.000306942,
		'"':  0.00000183067,
		'#':  0.0000854313,
		'$':  0.0000970255,
		'%':  0.0000170863,
		'&':  0.0000134249,
		'\'': 0.00000122045,
		'(':  0.00000427156,
		')':  0.0000115942,
		'*':  0.000241648,
		'+':  0.0000231885,
		',':  0.0000323418,
		'-':  0.000197712,
		'.':  0.000316706,
		'/':  0.0000311214,
		'0':  0.0274381,
		'1':  0.0435053,
		'2':  0.0312312,
		'3':  0.0243339,
		'4':  0.0194265,
		'5':  0.0188577,
		'6':  0.0175647,
		'7':  0.01621,
		'8':  0.0166225,
		'9':  0.0179558,
		':':  0.00000549201,
		';':  0.0000207476,
		'<':  0.00000427156,
		'=':  0.0000140351,
		'>':  0.00000183067,
		'?':  0.0000207476,
		'@':  0.000238597,
		'A':  0.00130466,
		'B':  0.000806715,
		'C':  0.000660872,
		'D':  0.000698096,
		'E':  0.000970865,
		'F':  0.000417393,
		'G':  0.000497332,
		'H':  0.000544319,
		'I':  0.00070908,
		'J':  0.000363083,
		'K':  0.000460719,
		'L':  0.000775594,
		'M':  0.000782306,
		'N':  0.000748134,
		'O':  0.000729217,
		'P':  0.00073715,
		'Q':  0.000147064,
		'R':  0.0008476,
		'S':  0.00108132,
		'T':  0.000801223,
		'U':  0.000350268,
		'V':  0.000235546,
		'W':  0.000320367,
		'X':  0.000142182,
		'Y':  0.000255073,
		'Z':  0.000170252,
		'[':  0.000010984,
		'\\': 0.0000115942,
		']':  0.000010984,
		'^':  0.0000195272,
		'_':  0.000122655,
		'`':  0.0000115942,
		'a':  0.0752766,
		'b':  0.0229145,
		'c':  0.0257276,
		'd':  0.0276401,
		'e':  0.070925,
		'f':  0.012476,
		'g':  0.0185331,
		'h':  0.0241319,
		'i':  0.0469732,
		'j':  0.00836677,
		'k':  0.0196828,
		'l':  0.0377728,
		'm':  0.0299913,
		'n':  0.0456899,
		'o':  0.0517,
		'p':  0.0245578,
		'q':  0.00346119,
		'r':  0.0496032,
		's':  0.0461079,
		't':  0.0387388,
		'u':  0.0210191,
		'v':  0.00833626,
		'w':  0.0124492,
		'x':  0.00573305,
		'y':  0.0152483,
		'z':  0.00632558,
		'{':  0.00000122045,
		'|':  0.00000122045,
		'}':  0.000000610223,
		'~':  0.0000152556,
		'ä':  0.000000610223,
		'ï':  0.00000183067,
		'ö':  0.000000610223,
		'ü':  0.00000122045,
	}
}
