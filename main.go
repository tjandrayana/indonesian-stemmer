package main

import (
	"fmt"
	"strings"
)

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
