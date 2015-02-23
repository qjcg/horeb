package main

type UnicodeBlock struct {
	low, high int
}

/*
See:
- http:en.wikipedia.org/wiki/Plane_(Unicode)
- http:en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane
- http:en.wikipedia.org/wiki/Plane_(Unicode)#Supplementary_Multilingual_Plane
*/
var Blocks = map[string]*UnicodeBlock{
	// Basic Multilingual Plane (0000-ffff)
	"hebrew":       &UnicodeBlock{0x0590, 0x05ff},
	"currency":     &UnicodeBlock{0x20a0, 0x20cf},
	"letterlike":   &UnicodeBlock{0x2100, 0x214f},
	"geometric":    &UnicodeBlock{0x25a0, 0x25ff},
	"misc_symbols": &UnicodeBlock{0x2600, 0x26ff},
	"dingbats":     &UnicodeBlock{0x2700, 0x27bf},
	// Supplementary Multilingual Plane (10000-1ffff)
	"aegean_nums":        &UnicodeBlock{0x10100, 0x1013f},
	"ancient_greek_nums": &UnicodeBlock{0x10140, 0x1018f},
	"phaistos_disc":      &UnicodeBlock{0x101d0, 0x101ff},
	"math_alnum":         &UnicodeBlock{0x1d400, 0x1d7ff},
	"emoji":              &UnicodeBlock{0x1f300, 0x1f5ff},
	"mahjong":            &UnicodeBlock{0x1f000, 0x1f02f},
	"dominos":            &UnicodeBlock{0x1f030, 0x1f09f},
	"playing_cards":      &UnicodeBlock{0x1f0a0, 0x1f0ff},
}
