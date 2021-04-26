package font

import (
	"io/ioutil"
	"testing"

	"github.com/golang/freetype/truetype"
	"github.com/tdewolff/test"
	"golang.org/x/image/font/sfnt"
)

func TestParseTTF(t *testing.T) {
	b, err := ioutil.ReadFile("../resources/DejaVuSerif.ttf")
	test.Error(t, err)

	sfnt, err := ParseFont(b, 0)
	test.Error(t, err)
	test.T(t, sfnt.Head.UnitsPerEm, uint16(2048))
}

func TestParseOTF(t *testing.T) {
	b, err := ioutil.ReadFile("../resources/EBGaramond12-Regular.otf")
	test.Error(t, err)

	sfnt, err := ParseFont(b, 0)
	test.Error(t, err)
	test.T(t, sfnt.Head.UnitsPerEm, uint16(1000))
}

//func TestParseOTF_CFF2(t *testing.T) {
//	b, err := ioutil.ReadFile("../resources/AdobeVFPrototype.otf") // TODO: CFF2
//	test.Error(t, err)
//
//	sfnt, err := ParseFont(b, 0)
//	test.Error(t, err)
//	test.T(t, sfnt.Head.UnitsPerEm, uint16(1000))
//}

func TestParseWOFF(t *testing.T) {
	b, err := ioutil.ReadFile("../resources/DejaVuSerif.woff")
	test.Error(t, err)

	sfnt, err := ParseFont(b, 0)
	test.Error(t, err)
	test.T(t, sfnt.Head.UnitsPerEm, uint16(2048))
}

func TestParseWOFF2(t *testing.T) {
	b, err := ioutil.ReadFile("../resources/DejaVuSerif.woff2")
	test.Error(t, err)

	sfnt, err := ParseFont(b, 0)
	test.Error(t, err)
	test.T(t, sfnt.Head.UnitsPerEm, uint16(2048))
}

func TestParseEOT(t *testing.T) {
	b, err := ioutil.ReadFile("../resources/DejaVuSerif.eot")
	test.Error(t, err)

	sfnt, err := ParseFont(b, 0)
	test.Error(t, err)
	test.T(t, sfnt.Head.UnitsPerEm, uint16(2048))
}

func TestFromGoFreetype(t *testing.T) {
	b, err := ioutil.ReadFile("../resources/DejaVuSerif.ttf")
	test.Error(t, err)

	font, err := truetype.Parse(b)
	test.Error(t, err)

	buf := FromGoFreetype(font)
	sfnt, err := ParseSFNT(buf, 0)
	test.Error(t, err)
	test.T(t, sfnt.Head.UnitsPerEm, uint16(2048))
}

func TestFromGoSFNT(t *testing.T) {
	b, err := ioutil.ReadFile("../resources/DejaVuSerif.ttf")
	test.Error(t, err)

	font, err := sfnt.Parse(b)
	test.Error(t, err)

	buf := FromGoSFNT(font)
	sfnt, err := ParseSFNT(buf, 0)
	test.Error(t, err)
	test.T(t, sfnt.Head.UnitsPerEm, uint16(2048))
}
