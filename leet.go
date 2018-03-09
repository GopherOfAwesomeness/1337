package Leet

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

const (
	everything = -1
)

type Level int32

const (
	max = Level(100)
	// HaXX0r Level replace nearly everything
	HaXX0r = Level(99)
	// N00b just replace sub-string in 15% of the caes - LAME
	N00b = Level(15)
	// Lam3 do not replace anything
	Lam3 = Level(0)
)

type Replace struct {
	src, dest string
	weight    int
}

type Replacer []Replace

func (r Replacer) Len() int           { return len(r) }
func (r Replacer) Less(i, j int) bool { return r[i].weight < r[j].weight }
func (r Replacer) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Translator contains a translate config
type Translator struct {
	rep   Replacer
	level Level
	rand  *rand.Rand
}

func defaultReplace() Replacer {
	r := Replacer{
		{
			src:  "a",
			dest: "4",
		},
		{
			src:  "e",
			dest: "3",
		},
	}
	sort.Sort(r)
	return r
}

var source = func() rand.Source { return rand.NewSource(time.Now().UnixNano()) }

// New return a new Translator
func New(l Level, r Replacer) *Translator {
	if len(r) == 0 {
		r = defaultReplace()
	}
	if l < Lam3 || l > HaXX0r {
		l = N00b
	}
	return &Translator{
		rep:   r,
		level: l,
		rand:  rand.New(source()),
	}
}

// Translate a string based on the replace rules
func (t *Translator) Translate(s string) string {
	for _, r := range t.rep {
		if Level(t.rand.Int31n(int32(max))) < t.level {
			s = strings.Replace(s, r.src, r.dest, everything)
		}
	}
	return s
}
