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
	DOUBLE_QUOTE
	LITERAL
)

func Tokenize(s string) []string {
	var results []string
	var current string
	var isInsideQuote bool
	state := NORMAL
	s += " "

	for _, c := range s {
		switch state {
		case NORMAL:
			switch c {
			case '\\':
				state = LITERAL
			case '\'':
				state = QUOTE
			case '"':
				state = DOUBLE_QUOTE
				isInsideQuote = true
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
		case DOUBLE_QUOTE:
			switch c {
			case '"':
				state = NORMAL
			case '\\':
				state = LITERAL
				isInsideQuote = true
			default:
				current += string(c)
			}
		case LITERAL:
			current += string(c)
			if isInsideQuote {
				state = DOUBLE_QUOTE
				isInsideQuote = false
			} else {
				state = NORMAL
			}
		}
	}

	return results
}
