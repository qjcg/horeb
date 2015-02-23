package main

type UnicodePlane struct {
	low, high int
}

/*
See:
- http:en.wikipedia.org/wiki/Plane_(Unicode)
- http:en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane
- http:en.wikipedia.org/wiki/Plane_(Unicode)#Supplementary_Multilingual_Plane
*/
var Planes = map[string]*UnicodePlane{
	"hebrew":             &UnicodePlane{0x0590, 0x05ff},
	"currency":           &UnicodePlane{0x20a0, 0x20cf},
	"letterlike":         &UnicodePlane{0x2100, 0x214f},
	"geometric":          &UnicodePlane{0x25a0, 0x25ff},
	"misc_symbols":       &UnicodePlane{0x2600, 0x26ff},
	"dingbats":           &UnicodePlane{0x2700, 0x27bf},
	"aegean_nums":        &UnicodePlane{0x10100, 0x1013f},
	"ancient_greek_nums": &UnicodePlane{0x10140, 0x1018f},
	"phaistos_disc":      &UnicodePlane{0x101d0, 0x101ff},
	"math_alnum":         &UnicodePlane{0x1d400, 0x1d7ff},
	"emoji":              &UnicodePlane{0x1f300, 0x1f5ff},
	"mahjong":            &UnicodePlane{0x1f000, 0x1f02f},
	"dominos":            &UnicodePlane{0x1f030, 0x1f09f},
	"playing_cards":      &UnicodePlane{0x1f0a0, 0x1f0ff},
}
