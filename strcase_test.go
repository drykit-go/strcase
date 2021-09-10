package strcase_test

import (
	"testing"

	"github.com/drykit-go/testx"

	"github.com/drykit-go/strcase"
)

func TestCamel(t *testing.T) {
	testx.Table(strcase.Camel, nil).Cases([]testx.Case{
		{In: "A", Exp: "a"},
		{In: "ABC", Exp: "abc"},
		{In: "Name", Exp: "name"},
		{In: "AAAAAa", Exp: "aaaaAa"},
		{In: "SomeName", Exp: "someName"},
		{In: "HTTPHeader", Exp: "httpHeader"},
		{In: "MyHTTPHeader", Exp: "myHTTPHeader"},
		// {In: "my_HTTP_header", Exp: "myHTTPHeader"}, // FIXME: -> myHTTPheader
	}).Run(t)
}

func TestPascal(t *testing.T) {
	testx.Table(strcase.Pascal, nil).Cases([]testx.Case{
		{In: "a", Exp: "A"},
		{In: "abc", Exp: "Abc"},
		{In: "name", Exp: "Name"},
		{In: "AAAAAa", Exp: "AAAAAa"},
		{In: "someName", Exp: "SomeName"},
		{In: "httpHeader", Exp: "HTTPHeader"},
		{In: "myHTTPHeader", Exp: "MyHTTPHeader"},
	}).Run(t)
}
