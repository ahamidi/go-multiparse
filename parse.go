// Package multiparse provides tools to perform basic type detection and
// parsing on strings.
package multiparse

// MultiParse is the general interface all multiparse parsers implement.
type MultiParse interface {
	Parse(string) (interface{}, error)
}

// Parse a string to determine whether it represents a numeric or time value.
// This is a convenience function that is equivalent to calling the
// ParseType method on a general purpose parser instance returned by
// MakeGeneralParser.
func Parse(s string) (*Parsed, error) {
	p := MakeGeneralParser()
	return p.parse(s)
}

// ParseType determines whether a numeric or time representation
// according to the initialzed parsers.
func ParseType(s string) (*Parsed, error) {
	p := MakeGeneralParser()
	return p.ParseType(s)
}

// ParseInt reports whether the string parses to an integer according
// to the parser rules.
func ParseInt(s string) (int, error) {
	p := MakeGeneralParser()
	return p.ParseInt(s)
}

// ParseFloat reports whether the string parses to a float according
// to the parser rules.
func ParseFloat(s string) (float64, error) {
	p := MakeGeneralParser()
	return p.ParseFloat(s)
}

// ParseMoneyreports whether the string parses to a moneytary value
// according to the parser rules.
func ParseMoney(s string) (*Money, error) {
	p := MakeGeneralParser()
	return p.ParseMoney(s)
}

// ParseTime determines whether the input string parses for any of
// a number of common layouts. Calling this function is equivalent to
// constructing a general time parser with MakeGeneralTimeParser
// and invoking its ParseTime method.
func ParseTime(s string) (*Time, error) {
	p := MakeGeneralParser()
	return p.ParseTime(s)
}