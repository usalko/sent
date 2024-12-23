package sql_parser

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/usalko/prodl/internal/sql_parser"
	"github.com/usalko/prodl/internal/sql_parser/cache"
	"github.com/usalko/prodl/internal/sql_parser/dialect"
	"github.com/usalko/prodl/internal/sql_parser/mysql"
)

func TestMysqlKeywordTable(t *testing.T) {
	for _, kw := range mysql.GetKeywords() {
		lookup, ok := cache.KeywordLookup(kw.Name, dialect.MYSQL)
		require.Truef(t, ok, "keyword %q failed to match", kw.Name)
		require.Equalf(t, lookup, kw.Id, "keyword %q matched to %d (expected %d)", kw.Name, lookup, kw.Id)
	}
}

var vitessReserved = map[string]bool{
	"ESCAPE":        true,
	"NEXT":          true,
	"OFF":           true,
	"SAVEPOINT":     true,
	"SQL_NO_CACHE":  true,
	"TIMESTAMPADD":  true,
	"TIMESTAMPDIFF": true,
}

func TestMysqlCompatibility(t *testing.T) {
	file, err := os.Open(path.Join("test_data", "mysql_keywords.txt"))
	require.NoError(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	skipStep := 4
	for scanner.Scan() {
		if skipStep != 0 {
			skipStep--
			continue
		}

		afterSplit := strings.SplitN(scanner.Text(), "\t", 2)
		word, reserved := afterSplit[0], afterSplit[1] == "1"
		if reserved || vitessReserved[word] {
			word = "`" + word + "`"
		}
		sql := fmt.Sprintf("create table %s(c1 int)", word)
		_, err := sql_parser.ParseStrictDDL(sql, dialect.MYSQL)
		if err != nil {
			t.Errorf("%s is not compatible with mysql", word)
		}
	}
}
