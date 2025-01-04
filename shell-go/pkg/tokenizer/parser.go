package tokenizer

import (
	"strings"
)

type Tokenizer struct {
	Command string
	Args    []string
}

func NewTokenizer(text string) *Tokenizer {
	tokens := tokenize(text)
	if len(tokens) == 0 {
		return &Tokenizer{}
	}

	return &Tokenizer{
		Command: tokens[0],
		Args:    tokens[1:],
	}
}

func tokenize(text string) []string {
	var tokens []string
	var sb strings.Builder
	var inSingleQuotes, inDoubleQuotes bool

	i := 0
	for i < len(text) {
		char := text[i]

		if char == '\'' && !inDoubleQuotes {
			inSingleQuotes = !inSingleQuotes
			i++
			continue
		}

		if char == '"' && !inSingleQuotes {
			inDoubleQuotes = !inDoubleQuotes
			i++
			continue
		}

		if inSingleQuotes {
			sb.WriteByte(char)
			i++
			continue
		}

		if inDoubleQuotes {
			if char == '\\' {
				if i+1 < len(text) {
					next := text[i+1]
					switch next {
					case '\\', '"', '$':
						sb.WriteByte(next)
						i += 2
						continue
					case '\n':
						i += 2
						continue
					default:
						sb.WriteByte('\\')
						i++
						continue
					}
				} else {
					sb.WriteByte('\\')
					i++
					continue
				}
			} else {
				sb.WriteByte(char)
				i++
				continue
			}
		}

		if char == '\\' {
			if i+1 < len(text) {
				sb.WriteByte(text[i+1])
				i += 2
			} else {
				sb.WriteByte('\\')
				i++
			}
			continue
		}

		if char == ' ' || char == '\t' || char == '\n' {
			if sb.Len() > 0 {
				tokens = append(tokens, sb.String())
				sb.Reset()
			}
			i++
			continue
		}

		sb.WriteByte(char)
		i++
	}

	if sb.Len() > 0 {
		tokens = append(tokens, sb.String())
	}

	return tokens
}
