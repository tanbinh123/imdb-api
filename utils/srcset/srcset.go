// Package srcset `srcset` provides a parser for the HTML5 `srcset` attribute, based on the
// [WHATWG reference algorithm](https://html.spec.whatwg.org/multipage/embedded-content.html#parse-a-srcset-attribute).
// TODO: This works, but I dislike the state manipulation.
// Use more go-like structures for reading and tokenization, like bufio.Scanner
package srcset

import (
	"regexp"
	"strconv"
)

// ImageSource is a structure that contains an image definition.
type ImageSource struct {
	URL     string
	Width   *int64
	Height  *int64
	Density *float64
}

// SourceSet is the result of parsing the value of a srcset attribute.
// A SourceSet consists of multiple ImageSource instances.
type SourceSet []ImageSource

const (
	comma            = ','
	leftParentheses  = '('
	rightParentheses = ')'
)

const (
	stateNone = iota
	stateInDescriptor
	stateInParentheses
	stateAfterDescriptor
)

var (
	regexLeadingSpaces         = regexp.MustCompile("^[ \t\n\r\u000c]+")
	regexLeadingCommasOrSpaces = regexp.MustCompile("^[, \t\n\r\u000c]+")
	regexLeadingNotSpaces      = regexp.MustCompile("^[^ \t\n\r\u000c]+")
	regexTrailingCommas        = regexp.MustCompile("[,]+$")
	regexNonNegativeInteger    = regexp.MustCompile(`^\d+$`)
	regexFloatingPoint         = regexp.MustCompile(`^-?(?:[0-9]+|[0-9]*\.[0-9]+)(?:[eE][+-]?[0-9]+)?$`)
)

func isSpace(c rune) bool {
	switch c {
	case
		'\u0020', // space
		'\u0009', // horizontal tab
		'\u000A', // new line
		'\u000C', // form feed
		'\u000D': // carriage return
		return true
	default:
		return false
	}
}

type parser struct {
	input        string
	url          string
	pos          int
	currentState int
	end          int
	candidates   SourceSet
	descriptors  []string
}

func Parse(input string) SourceSet {
	p := &parser{
		pos:          0,
		currentState: stateNone,
		end:          len(input),
		candidates:   SourceSet{},
		descriptors:  []string{},
		input:        input,
		url:          "",
	}

	return p.Parse(input)
}

// Parse takes the value of a srcset attribute and parses it.
func (p *parser) Parse(input string) SourceSet {
	for {
		p.collectChars(regexLeadingCommasOrSpaces)
		if p.pos >= p.end {
			return p.candidates
		}

		p.url = p.collectChars(regexLeadingNotSpaces)
		p.descriptors = []string{}

		if p.url[len(p.url)-1] == ',' {
			p.url = regexTrailingCommas.ReplaceAllString(p.url, "")
			p.parseDescriptors()
		} else {
			p.tokenize()
		}
	}
}

func (p *parser) collectChars(rx *regexp.Regexp) string {
	if match := rx.FindString(p.input[p.pos:]); match != "" {
		p.pos += len(match)
		return match
	}

	return ""
}

func (p *parser) parseDescriptors() {
	var (
		isErr = false
		h     *int64
		w     *int64
		d     *float64
	)

	for _, desc := range p.descriptors {
		lastIdx := len(desc) - 1
		lastChar, numericVal := desc[lastIdx], desc[:lastIdx]
		intVal, intErr := strconv.ParseInt(numericVal, 10, 64)
		floatVal, floatErr := strconv.ParseFloat(numericVal, 64)

		switch {
		case regexNonNegativeInteger.MatchString(numericVal) && lastChar == 'w':
			if w != nil || d != nil {
				isErr = true
			}
			if intErr != nil || intVal == 0 {
				isErr = true
			} else {
				w = &intVal
			}
		case regexFloatingPoint.MatchString(numericVal) && lastChar == 'x':
			if w != nil || d != nil || h != nil {
				isErr = true
			}
			if floatErr != nil || floatVal < 0 {
				isErr = true
			} else {
				d = &floatVal
			}
		case regexNonNegativeInteger.MatchString(numericVal) && lastChar == 'h':
			if h != nil || d != nil {
				isErr = true
			}
			if intErr != nil || intVal == 0 {
				isErr = true
			} else {
				h = &intVal
			}
		default:
			isErr = true
		}
	}

	if !isErr {
		p.candidates = append(p.candidates, ImageSource{
			URL:     p.url,
			Density: d,
			Width:   w,
			Height:  h,
		})
	}
}

func (p *parser) tokenize() {
	p.collectChars(regexLeadingSpaces)
	currentDescriptor := ""
	p.currentState = stateInDescriptor

	for {
		if p.pos == len(p.input) {
			if p.currentState != stateAfterDescriptor && currentDescriptor != "" {
				p.descriptors = append(p.descriptors, currentDescriptor)
			}

			p.parseDescriptors()
			return
		}

		c := rune(p.input[p.pos])

		switch p.currentState {
		case stateInDescriptor:
			switch {
			case isSpace(c):
				if currentDescriptor != "" {
					p.descriptors = append(p.descriptors, currentDescriptor)
					currentDescriptor = ""
					p.currentState = stateAfterDescriptor
				}
			case c == comma:
				p.pos++
				if currentDescriptor != "" {
					p.descriptors = append(p.descriptors, currentDescriptor)
					p.parseDescriptors()
					return
				}
			case c == leftParentheses:
				currentDescriptor += string(c)
				p.currentState = stateInParentheses
			default:
				currentDescriptor += string(c)
			}

		case stateInParentheses:
			switch c {
			case rightParentheses:
				currentDescriptor += string(c)
				p.currentState = stateInDescriptor
			default:
				currentDescriptor += string(c)
			}

		case stateAfterDescriptor:
			switch {
			case isSpace(c):
			default:
				p.currentState = stateInDescriptor
				p.pos--
			}
		}

		p.pos++
	}
}
