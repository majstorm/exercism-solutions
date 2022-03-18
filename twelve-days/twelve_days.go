//Package twelve implements 'The Twelve Days of Christmas song'
package twelve

import (
	"fmt"
)

type LyricPair struct {
	day  string
	gift string
}

var lyrics = map[int]LyricPair{
	1:  {"first", " a Partridge in a Pear Tree"},
	2:  {"second", " two Turtle Doves"},
	3:  {"third", " three French Hens"},
	4:  {"fourth", " four Calling Birds"},
	5:  {"fifth", " five Gold Rings"},
	6:  {"sixth", " six Geese-a-Laying"},
	7:  {"seventh", " seven Swans-a-Swimming"},
	8:  {"eighth", " eight Maids-a-Milking"},
	9:  {"ninth", " nine Ladies Dancing"},
	10: {"tenth", " ten Lords-a-Leaping"},
	11: {"eleventh", " eleven Pipers Piping"},
	12: {"twelfth", " twelve Drummers Drumming"},
}

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", lyrics[i].day)
	for cnt := i; cnt > 0; cnt-- {
		if cnt == 1 && i != 1 {
			verse += " and"
		}
		verse += lyrics[cnt].gift
		if cnt != 1 {
			verse += ","
		}
	}
	return verse + "."
}

func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i)
		if i != 12 {
			song += "\n"
		}
	}
	return song
}
