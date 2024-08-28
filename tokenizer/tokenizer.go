package tokenizer

import (
	"fmt"
)

type Tokenizer struct {
	input string

	curr int
	next int

	ch byte

	results []string
}


func (t * Tokenizer) PrintTokens() {
	for _, t := range t.results {
		fmt.Println(t)
	}
}

type QueryToken struct {
	start   int
	end     int
	value   string
	isMedia bool
}

func (qt *QueryToken) Value() string {
	return qt.value
}
func (qt *QueryToken) String() string {
	return fmt.Sprintf("start_pos %d\nend_pos %d\nmedia_query %v\nvalue %s", qt.start, qt.end, qt.isMedia, qt.value)
}

func New(input string) *Tokenizer {
	t := &Tokenizer{
		input: input,
	}

	t.readChar()

	return t
}

func (t *Tokenizer) Tokenize() error {
	for t.next < len(t.input) {
		if t.ch == '@' {
			t.readQuery(false)

		}
		t.readChar() // reading char
	}

	return nil
}

func (t *Tokenizer) readQuery(hasParent bool) string {
	var temp int = t.curr
	media := t.isMedia()

	var str string = t.input[temp:t.next]

	var cnt int = 0

	query := &QueryToken{
		start:   temp,
		isMedia: media,
	}

	for t.next < len(t.input) {
		if t.ch == '@' {
			cnt++
			str += t.readQuery(true)
		}

		str += string(t.ch)

		if t.ch == '}' {
			cnt--
			if cnt == 0 {
				break
			}
		}

		t.readChar()
	}
	query.end = t.curr
	query.value = str

	fmt.Println(query)

	if hasParent {
		//t.results = append(t.results, query.Value())
		return query.Value()
	}
	t.results = append(t.results, query.String())

	return query.String()
}

func (t *Tokenizer) isMedia() bool {
	var res string = ""
	for t.ch != ' ' && t.next < len(t.input) {
		res += string(t.ch)
		t.readChar()
	}

	return res == "@media"
}

func (t *Tokenizer) readChar() {
	if t.next >= len(t.input) {
		t.ch = 0
	} else {
		t.ch = t.input[t.next]
	}
	t.curr = t.next
	t.next++
}
