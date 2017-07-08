package main

import "regexp"

var kataDasar map[string]bool
var kataPartikel map[string]bool
var kataMilik map[string]bool
var kataFirstPrefix map[string]bool

var arrKataPartikel = []string{"kah", "lah", "pun"}
var arrKataMilik = []string{"ku", "mu", "nya"}
var arrKataFirstPrefix = []string{"meng", "meny", "men", "mem", "me", "peng", "peny", "pen", "pem", "di", "ter", "ke"}

type Parser struct {
	punc, decimal, zero, unicode, integer, special, ilegalOctalNumber, newLine *regexp.Regexp
}
