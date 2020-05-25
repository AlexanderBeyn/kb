package lib

import (
	"sort"
	"strings"
)

type CommonArgs struct {
	Project    *string
	Column     *string
	FromColumn *string
	Search     *string
	More       bool
	Args       []string
}

var (
	ProjectSigil    = "%%"
	ColumnSigil     = "%"
	FromColumnSigil = "^%"
	SearchSigil     = "/"
	MoreSuffix      = "+"
)

func StripSigil(s string) (string, string) {
	sigils := []string{
		ProjectSigil,
		ColumnSigil,
		FromColumnSigil,
		SearchSigil,
	}

	sort.Slice(sigils, func(i int, j int) bool {
		return len(sigils[i]) > len(sigils[j])
	})

	for _, sigil := range sigils {
		if strings.HasPrefix(s, sigil) {
			return sigil, s[len(sigil):]
		}
	}

	return "", s
}

func ParseCommonArgs(args []string) CommonArgs {
	var out CommonArgs
	var newArgs []string

	for idx, arg := range args {
		if arg == "--" {
			if (idx + 1) < len(args) {
				newArgs = append(newArgs, args[idx+1:]...)
			}
			break
		}
		sigil, rest := StripSigil(arg)
		switch sigil {
		case ProjectSigil:
			out.Project = &rest
		case ColumnSigil:
			out.Column = &rest
		case FromColumnSigil:
			out.FromColumn = &rest
		case SearchSigil:
			out.Search = &rest
		default:
			newArgs = append(newArgs, arg)
		}
	}

	out.Args = newArgs

	if len(out.Args) > 0 && out.Args[len(out.Args)-1] == MoreSuffix {
		out.Args = out.Args[:len(out.Args)-1]
		out.More = true
	}

	return out
}
