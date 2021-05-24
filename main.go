package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

// GOAL: parse below samples
const (
	// ex1 = `github.name = "*f\*o\\o\"o*" AND github.owner = "*bar*"`
	ex1 = `github.name = "*foo*" AND github.owner = "*bar*"`
	// ex1 = `github.name = "*foo*"`
	// ex1 = `github.name.subField = "*foo*"`
	// ex2 = `github.name = "*foo*"`
)

/*
	Query = Requirement {`AND` Requirement}
	Requirement = Check {`OR` Check}
	Check = Path `=` Pattern | `(` Query `)`
	Path = Field {`.` Field}
	Pattern = `"` ... `"`
*/

// TODO: try codegen: participle/experimental/codegen.GenerateLexer()
var basicLexer = stateful.MustSimple([]stateful.Rule{
	// {"Comment", `(?i)rem[^\n]*`, nil},
	{"String", `"(\\"|\\\\|[^"])*"`, nil},
	// {"Number", `[-+]?(\d*\.)?\d+`, nil},
	{"Field", `[a-z][a-zA-Z0-9]*`, nil},
	{"Ident", `[a-zA-Z_]\w*`, nil},
	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	// {"EOL", `[\n\r]+`, nil},
	{"whitespace", `[ \t]+`, nil},
})

type Query struct {
	First *Requirement   `@@`
	And   []*Requirement `( "AND" @@ )*`
}

type Requirement = Check

// type Requirement struct {
// 	Left  *Check     `@@`
// 	Right []*OrCheck `@@*`
// }

// type OrCheck struct {
// 	Check *Check `"OR" @@`
// }

type Check struct {
	Cmp *Comparison `@@`
	// Subquery *Query      `| "(" @@ ")"`
}

type Comparison struct {
	Path   *Path  `@@`
	String string `"=" @String`
}

type Path struct {
	Base string   `@Field`
	Sub  []string `( "." @Field)*`
}

func main() {
	parser, err := participle.Build(&Query{}, participle.Lexer(basicLexer))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parser)
	fmt.Println(strings.Repeat("=", 30))

	var q Query
	err = parser.ParseString("", ex1, &q)
	if err != nil {
		log.Fatal(err)
	}
	buf, _ := json.MarshalIndent(q, "", "  ")
	fmt.Println("AST:", string(buf))
}
