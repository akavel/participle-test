package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer/stateful"
)

// GOAL: parse below samples
const (
	ex1 = `github.name = "*foo*" AND github.owner = "*bar*"`
	ex2 = `github.name = "*foo*"`
)

/*
	Query = Requirement {`AND` Requirement}
	Requirement = Check {`OR` Check}
	Check = Field `=` Pattern | `(` Query `)`
	Field = Word {`.` Word}
	Pattern = `"` ... `"`
*/

var basicLexer = stateful.MustSimple([]stateful.Rule{
	// {"Comment", `(?i)rem[^\n]*`, nil},
	// {"String", `"(\\"|[^"])*"`, nil},
	// {"Number", `[-+]?(\d*\.)?\d+`, nil},
	// {"Ident", `[a-zA-Z_]\w*`, nil},
	// {"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	// {"EOL", `[\n\r]+`, nil},
	// {"whitespace", `[ \t]+`, nil},
})

func main() {
	fmt.Println("hello")
}
