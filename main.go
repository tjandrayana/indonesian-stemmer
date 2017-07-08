package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var kataDasar map[string]bool
var kataPartikel map[string]bool
var kataMilik map[string]bool
var kataFirstPrefix map[string]bool

var arrKataPartikel = []string{"kah", "lah", "pun"}
var arrKataMilik = []string{"ku", "mu", "nya"}
var arrKataFirstPrefix = []string{"meng", "meny", "men", "mem", "me", "peng", "peny", "pen", "pem", "di", "ter", "ke"}

func LoadKata() {
	kataPartikel = make(map[string]bool)
	kataMilik = make(map[string]bool)
	kataFirstPrefix = make(map[string]bool)

	for i := range arrKataPartikel {
		kataPartikel[arrKataPartikel[i]] = true
	}

	for i := range arrKataMilik {
		kataMilik[arrKataMilik[i]] = true
	}

	for i := range arrKataFirstPrefix {
		kataFirstPrefix[arrKataFirstPrefix[i]] = true
	}

}

func main() {
	kataDasar = make(map[string]bool)

	p := New()

	text := "buku miliknya yoshua tertinggal di rumahku dimana ibuku tinggal bersama ayahmu"
	fmt.Println(text)
	text = strings.ToLower(text)
	text = p.Normalize(text)

	// fmt.Println(text)

	textArray := strings.Split(text, " ")

	var tempText string

	setStopAddress("files/kata-dasar.csv")
	LoadKata()

	for i := range textArray {
		if textArray[i] == "" {
			continue
		}
		//fmt.Println(textArray[i])

		if kataDasar[textArray[i]] {
			tempText = tempText + " " + textArray[i]
		} else {
			kata := removePartikel(textArray[i])
			if kataDasar[kata] {
				tempText = tempText + " " + kata
				continue
			}
			kata = removeKepemilikan(kata)
			if kataDasar[kata] {
				tempText = tempText + " " + kata
				continue
			}

			tempText = tempText + " " + kata
		}
	}

	fmt.Println("temp = ", tempText)

}

func setStopAddress(path string) error {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Println("[Error] Set Stopwords from ", path)
		return err
	}

	csvFile := csv.NewReader(f)
	content, err := csvFile.ReadAll()
	if err != nil {
		log.Println(err)
	}

	csvData := content

	// Create a new reader.
	for i := range csvData {
		for j := range csvData[i] {
			kataDasar[csvData[i][j]] = true
		}
		//fmt.Println("data = ", kataDasar["zulu"])
	}
	return nil
}

func (p *Parser) Normalize(text string) string {

	text = p.RemovePunc(text)

	return text
}

func removePartikel(kata string) string {

	var substring2, substring3 string
	size := len(kata)

	if size >= 2 {
		substring2 = kata[size-2 : size]

	}
	if size >= 3 {
		substring3 = kata[size-3 : size]

	}

	if kataPartikel[substring2] && substring2 != "" {
		kata = kata[0 : size-2]
	} else if kataPartikel[substring3] && substring3 != "" {
		kata = kata[0 : size-3]
	}

	return kata
}

func removeKepemilikan(kata string) string {
	var substring2, substring3 string
	size := len(kata)

	if size >= 2 {
		substring2 = kata[size-2 : size]

	}
	if size >= 3 {
		substring3 = kata[size-3 : size]

	}

	if kataMilik[substring2] && substring2 != "" {
		kata = kata[0 : size-2]
	} else if kataMilik[substring3] && substring3 != "" {
		kata = kata[0 : size-3]
	}

	return kata
}

func removeFirstPrefix(kata string) string {
	var substring2, substring3 string
	size := len(kata)

	if size >= 2 {
		substring2 = kata[size-2 : size]

	}
	if size >= 3 {
		substring3 = kata[size-3 : size]

	}

	if kataMilik[substring2] && substring2 != "" {
		kata = kata[0 : size-2]
	} else if kataMilik[substring3] && substring3 != "" {
		kata = kata[0 : size-3]
	}

	return kata
}

type Parser struct {
	punc, decimal, zero, unicode, integer, special, ilegalOctalNumber, newLine *regexp.Regexp
}

func New() *Parser {
	parser := &Parser{
		punc:              regexp.MustCompile(`\W`),
		decimal:           regexp.MustCompile(`\D`),
		zero:              regexp.MustCompile(`^0`),
		unicode:           regexp.MustCompile(`[\pL\p{Mc}\p{Mn}-_']+`),
		integer:           regexp.MustCompile(`^(?:[-+]?(?:0|[1-9][0-9]*))$`),
		special:           regexp.MustCompile(`[_+-.:,!@#$%^&*();\\/|<>"'?=]`),
		ilegalOctalNumber: regexp.MustCompile(` 0+`),
		newLine:           regexp.MustCompile("\\[.*?]"),
	}

	return parser
}
func (p *Parser) RemovePunc(str string) string {
	return p.punc.ReplaceAllString(str, " ")
}
