package main

const (
	CD   = "cd"
	ECHO = "echo"
	EXIT = "exit"
	PWD  = "pwd"
	TYPE = "type"
)

type LexerState int

const (
	NORMAL LexerState = iota
	QUOTE
)

func Tokenize(s string) []string {
	var results []string
	var current string
	state := NORMAL
	s += " "

	for _, c := range s {
		switch state {
		case NORMAL:
			switch c {
			case '\'':
				state = QUOTE
			case ' ':
				if current != "" {
					results = append(results, current)
				}
				current = ""
			default:
				current += string(c)
			}
		case QUOTE:
			switch c {
			case '\'':
				state = NORMAL
			default:
				current += string(c)
			}
		}
	}

	return results
}
