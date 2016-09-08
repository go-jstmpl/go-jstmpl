// +build !windows

package flags

import (
	"strings"
)

const (
	defaultShortOptDelimiter = '-'
	defaultLongOptDelimiter  = "--"
	defaultNameArgDelimiter  = '='
)

func argumentIsOption(arg string) bool {
	return len(arg) > 0 && arg[0] == '-'
}

// stripOptionPrefix returns the option without the prefix and whether or
// not the option is a long option or not.
func stripOptionPrefix(optname string) (string, bool) {
	if strings.HasPrefix(optname, "--") {
		return optname[2:], true
	} else if strings.HasPrefix(optname, "-") {
		return optname[1:], false
	}

	return optname, false
}

// splitOption attempts to split the passed option into a name and an argument.
// When there is no argument specified, nil will be returned for it.
func splitOption(option string) (string, *string) {
	pos := strings.Index(option, "=")
	if pos >= 0 {
		rest := option[pos+1:]
		return option[:pos], &rest
	}

	return option, nil
}

// newHelpGroup returns a new group that contains default help parameters.
func newHelpGroup(showHelp func() error) *Group {
	var help struct {
		ShowHelp func() error `short:"h" long:"help" description:"Show this help message"`
	}
	help.ShowHelp = showHelp
	return NewGroup("Help Options", &help)
}