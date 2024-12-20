package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage_SetLocale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		locale   string
		expected string
	}{
		0: {"now", "en", "1 day after"},
		1: {"tomorrow", "zh-CN", "1 天后"},
	}

	for index, test := range tests {
		lang := NewLanguage()
		lang.SetLocale(test.locale)
		assert.Equal(test.expected, SetLanguage(lang).Parse(test.input).AddDays(1).DiffForHumans(Parse(test.input)), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLanguage_SetResources1(t *testing.T) {
	assert := assert.New(t)

	lang := NewLanguage()
	resources := map[string]string{
		"seasons":  "spring|summer|autumn|winter",
		"year":     "1 yr|%d yrs",
		"month":    "1 mo|%d mos",
		"week":     "%dw",
		"day":      "%dd",
		"hour":     "%dh",
		"minute":   "%dm",
		"second":   "%ds",
		"now":      "just now",
		"ago":      "%s ago",
		"from_now": "in %s",
		"before":   "%s before",
		"after":    "%s after",
	}
	lang.SetLocale("en").SetResources(resources)

	tests := []struct {
		input1   string
		input2   string
		expected string
	}{
		0: {"2020-08-05 13:14:15", "2020-08-05 13:14:15", "just now"},
		1: {"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 yr before"},
		2: {"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 yr after"},
		3: {"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 yrs before"},
		4: {"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 yrs after"},

		5: {"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 mo before"},
		6: {"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 mo after"},
		7: {"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 mos before"},
		8: {"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 mos after"},

		9:  {"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1d before"},
		10: {"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1d after"},
		11: {"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1w before"},
		12: {"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1w after"},

		13: {"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1h before"},
		14: {"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1h after"},
		15: {"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10h before"},
		16: {"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10h after"},

		17: {"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1m before"},
		18: {"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1m after"},
		19: {"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10m before"},
		20: {"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10m after"},

		21: {"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1s before"},
		22: {"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1s after"},
		23: {"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10s before"},
		24: {"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10s after"},
	}

	c := SetLanguage(lang)
	for index, test := range tests {
		c1 := c.Parse(test.input1)
		c2 := c.Parse(test.input2)
		assert.NoError(c1.Error)
		assert.NoError(c2.Error)
		assert.Equal(test.expected, c1.DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLanguage_SetResources2(t *testing.T) {
	assert := assert.New(t)

	lang := NewLanguage()

	resources := map[string]string{
		"xxx": "xxx",
	}
	lang.SetResources(resources)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},

		4: {"0000-00-00 00:00:00", ""},
		5: {"2021-08-05 13:14:15", ""},
	}

	c := SetLanguage(lang)

	for index, test := range tests {
		assert.Equal(test.expected, c.Parse(test.input).DiffForHumans(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, c.Parse(test.input).Constellation(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, c.Parse(test.input).Season(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, c.Parse(test.input).ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, c.Parse(test.input).ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, c.Parse(test.input).ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, c.Parse(test.input).ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}
