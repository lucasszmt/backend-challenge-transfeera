package vo

import "regexp"

var NameRegexp = regexp.MustCompile(`^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$`)

type Name string

func NewName(name string) (Name, error) {
	match := NameRegexp.MatchString(name)
	if !match {
		return "", ErrInvalidName
	}
	return Name(name), nil
}
