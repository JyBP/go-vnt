package vnt

import (
	"strings"
	"testing"
	"time"
)

func TestSamples(t *testing.T) {
	tests := []struct {
		vnote    string
		expected Note
	}{
		{`BEGIN:VNOTE
VERSION:1.1
BODY;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:Fran=C3=A7ois-J=C3=A9r=C3=B4me : test ; test
DCREATED:20171007T161520
LAST-MODIFIED:20171108T171055
END:VNOTE`, Note{
			"François-Jérôme : test ; test",
			time.Date(2017, 10, 7, 16, 15, 20, 0, time.UTC),
			time.Date(2017, 11, 8, 17, 10, 55, 0, time.UTC),
		}},
		{`BEGIN:VNOTE
VERSION:1.1
BODY;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:` + strings.Repeat("=C3=A9", 1000) + `
DCREATED:20171007T161520
LAST-MODIFIED:20171108T171055
END:VNOTE`, Note{
			strings.Repeat("é", 1000),
			time.Date(2017, 10, 7, 16, 15, 20, 0, time.UTC),
			time.Date(2017, 11, 8, 17, 10, 55, 0, time.UTC),
		}},
	}

	for _, test := range tests {
		actual, err := Parse(strings.NewReader(test.vnote))
		if err != nil {
			t.Error(err)
		}
		if test.expected != actual {
			t.Errorf("Parse(%q) = %s, want %s", test.vnote, actual, test.expected)
		}
	}
}
