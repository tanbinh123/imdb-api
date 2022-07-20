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

// TODO: refactor to decrease function complicity
// Parse takes the value of a srcset attribute and parses it.
func Parse(input string) SourceSet {
	var (
		url          string
		pos          = 0
		currentState = stateNone
		end          = len(input)
		candidates   = SourceSet{}
		descriptors  = []string{}
	)

	collectChars := func(rx *regexp.Regexp) string {
		if match := rx.FindString(input[pos:]); match != "" {
			pos += len(match)
			return match
		}

		return ""
	}

	parseDescriptors := func() {
		var (
			isErr = false
			h     *int64
			w     *int64
			d     *float64
		)

		for _, desc := range descriptors {
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
			candidates = append(candidates, ImageSource{
				URL:     url,
				Density: d,
				Width:   w,
				Height:  h,
			})
		}
	}

	tokenize := func() {
		collectChars(regexLeadingSpaces)
		currentDescriptor := ""
		currentState = stateInDescriptor

		for {
			if pos == len(input) {
				if currentState != stateAfterDescriptor && currentDescriptor != "" {
					descriptors = append(descriptors, currentDescriptor)
				}

				parseDescriptors()
				return
			}

			c := rune(input[pos])

			switch currentState {
			case stateInDescriptor:
				switch {
				case isSpace(c):
					if currentDescriptor != "" {
						descriptors = append(descriptors, currentDescriptor)
						currentDescriptor = ""
						currentState = stateAfterDescriptor
					}
				case c == comma:
					pos++
					if currentDescriptor != "" {
						descriptors = append(descriptors, currentDescriptor)
						parseDescriptors()
						return
					}
				case c == leftParentheses:
					currentDescriptor += string(c)
					currentState = stateInParentheses
				default:
					currentDescriptor += string(c)
				}

			case stateInParentheses:
				switch c {
				case rightParentheses:
					currentDescriptor += string(c)
					currentState = stateInDescriptor
				default:
					currentDescriptor += string(c)
				}

			case stateAfterDescriptor:
				switch {
				case isSpace(c):
				default:
					currentState = stateInDescriptor
					pos--
				}
			}

			pos++
		}
	}

	for {
		collectChars(regexLeadingCommasOrSpaces)
		if pos >= end {
			return candidates
		}

		url = collectChars(regexLeadingNotSpaces)
		descriptors = []string{}

		if url[len(url)-1] == ',' {
			url = regexTrailingCommas.ReplaceAllString(url, "")
			parseDescriptors()
		} else {
			tokenize()
		}
	}
}
