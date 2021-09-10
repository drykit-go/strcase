package strcase

import "unicode"

type camelcase struct{ data []rune }

func (cc *camelcase) append(r rune, lower bool) {
	if lower {
		r = unicode.ToLower(r)
	}
	cc.data = append(cc.data, r)
}

func (cc camelcase) String() string {
	return string(cc.data)
}

func camel(in string) string {
	cc := camelcase{}
	ok := false
	for i, r := range in {
		switch {
		case r == '_':
			continue
		case i == len(in)-1:
			cc.append(r, !ok)
		case ok:
			cc.append(r, false)
		case unicode.IsLower(rune(in[i+1])):
			cc.append(r, i == 0)
			ok = true
		default:
			cc.append(r, true)
		}
	}
	return cc.String()
}
