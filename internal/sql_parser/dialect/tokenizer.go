/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dialect

import (
	"github.com/usalko/sent/internal/sql_parser/ast"
)

const (
	EofChar = 0x100
)

// Tokenizer is the interface used to generate SQL
// tokens for the parser.
type Tokenizer interface {
	SetParseTree(stmt ast.Statement)
	GetParseTree() ast.Statement

	SetAllowComments(allow bool)

	SetPartialDDL(node ast.Statement)
	GetPartialDDL() ast.Statement

	IncNesting()
	GetNesting() int
	DecNesting()

	SetSkipToEnd(skip bool)

	BindVar(bvar string, value struct{})
	GetBindVars() BindVars

	Scan() (int, string)

	GetIdToken() int
	GetKeywordString(token int) string

	GetLastError() error

	Cur() uint16
	Skip(count int)
	SkipBlank()
	Reset()
	SetMulti(multi bool)

	GetPos() int

	SetSkipSpecialComments(skip bool)
}