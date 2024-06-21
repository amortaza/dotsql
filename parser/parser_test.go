package parser

import (
	u "github.com/amortaza/dotsql/test_utils"
	"testing"
)

func Test_Parser_07(t *testing.T) {

	s, e := ParseDotted("(age= 3.14) or (name=\"ace.man\")")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_Parser_06(t *testing.T) {

	s, e := ParseDotted("age = 3.14 name<>$name (and)")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_Parser_05(t *testing.T) {

	s, e := ParseDotted(" h abc  def    jkl. xY")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_Parser_05p1(t *testing.T) {

	s, e := ParseDotted("jkl.")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_Parser_05p2(t *testing.T) {

	s, e := ParseDotted("jkl.a")
	expectedLen := 1

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "jkl.a", s[0])
}

func Test_Parser_04(t *testing.T) {

	s, e := ParseDotted(" h abc  def    jkl.xY")
	expectedLen := 1

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "jkl.xY", s[0])
}

func Test_Parser_03(t *testing.T) {

	s, e := ParseDotted(" h abc  def    jkl    ")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_Parser_02(t *testing.T) {

	s, e := ParseDotted(" hi ")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_Parser_01(t *testing.T) {

	s, e := ParseDotted("")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}
