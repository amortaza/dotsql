package parser

import (
	u "github.com/amortaza/dotsql/test_utils"
	"testing"
)

func Test_Token_10(t *testing.T) {

	s, e := tokenize(" h $abc.$xyz jkl.xY")
	expectedLen := 4

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "h", s[0])
	u.Check(t, "$abc.", s[1])
	u.Check(t, "$xyz", s[2])
	u.Check(t, "jkl.xY", s[3])
}

func Test_Token_09(t *testing.T) {

	s, e := tokenize(" h $abc$xyz jkl.xY")
	expectedLen := 4

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "h", s[0])
	u.Check(t, "$abc", s[1])
	u.Check(t, "$xyz", s[2])
	u.Check(t, "jkl.xY", s[3])
}

func Test_Token_08(t *testing.T) {

	s, e := tokenize(" h $abc  def    jkl.xY")
	expectedLen := 4

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "h", s[0])
	u.Check(t, "$abc", s[1])
	u.Check(t, "def", s[2])
	u.Check(t, "jkl.xY", s[3])
}

func Test_Token_07(t *testing.T) {

	s, e := tokenize("(age= 3.14) or (name=$name) and (owner.manager.location = $loc)")
	expectedLen := 17

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "(", s[0])
	u.Check(t, "age", s[1])
	u.Check(t, "=", s[2])
	u.Check(t, "3.14", s[3])
	u.Check(t, ")", s[4])
	u.Check(t, "or", s[5])
	u.Check(t, "(", s[6])
	u.Check(t, "name", s[7])
	u.Check(t, "=", s[8])
	u.Check(t, "$name", s[9])
	u.Check(t, ")", s[10])
	u.Check(t, "and", s[11])
	u.Check(t, "(", s[12])
	u.Check(t, "owner.manager.location", s[13])
	u.Check(t, "=", s[14])
	u.Check(t, "$loc", s[15])
	u.Check(t, ")", s[16])
}

func Test_Token_06(t *testing.T) {

	s, e := tokenize("age = 3.14 name<>$name (and)")
	expectedLen := 9

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "age", s[0])
	u.Check(t, "=", s[1])
	u.Check(t, "3.14", s[2])
	u.Check(t, "name", s[3])
	u.Check(t, "<>", s[4])
	u.Check(t, "$name", s[5])
	u.Check(t, "(", s[6])
	u.Check(t, "and", s[7])
	u.Check(t, ")", s[8])
}

func Test_Token_05(t *testing.T) {

	s, e := tokenize(" h abc    1jkl.xY")
	expectedLen := 4

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "h", s[0])
	u.Check(t, "abc", s[1])
	u.Check(t, "1", s[2])
	u.Check(t, "jkl.xY", s[3])
}

func Test_Token_04(t *testing.T) {

	s, e := tokenize(" h abc  def    jkl.xY")
	expectedLen := 4

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "h", s[0])
	u.Check(t, "abc", s[1])
	u.Check(t, "def", s[2])
	u.Check(t, "jkl.xY", s[3])
}

func Test_Token_03(t *testing.T) {

	s, e := tokenize(" h abc  def    jkl    ")
	expectedLen := 4

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "h", s[0])
	u.Check(t, "abc", s[1])
	u.Check(t, "def", s[2])
	u.Check(t, "jkl", s[3])
}

func Test_Token_02(t *testing.T) {

	s, e := tokenize(" hi ")
	expectedLen := 1

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "hi", s[0])
}

func Test_Token_01(t *testing.T) {

	s, e := tokenize("")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_StringRemover_01(t *testing.T) {

	s, e := removeStringConstant("")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
}

func Test_StringRemover_02(t *testing.T) {

	s, e := removeStringConstant(" hello 23.23 my {} m,. ")
	expectedLen := 23

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, " hello 23.23 my {} m,. ", s)
}

func Test_StringRemover_03(t *testing.T) {

	s, e := removeStringConstant("\"\"")
	expectedLen := 0

	u.NoError(t, e)
	u.Checki(t, expectedLen, len(s))
	u.Check(t, "", s)
}

func Test_StringRemover_04(t *testing.T) {

	s, e := removeStringConstant(" abc \"gdfed 678 \" def ")

	u.NoError(t, e)
	u.Check(t, " abc  def ", s)
}

func Test_StringRemover_05(t *testing.T) {

	_, e := removeStringConstant(" abc 'gd\"fed\\\\'678 ' def ")

	u.YesError(t, e)
}

func Test_StringRemover_06(t *testing.T) {

	_, e := removeStringConstant(" abc 'def ")
	u.YesError(t, e)
}

func Test_StringRemover_07(t *testing.T) {

	_, e := removeStringConstant("\"")

	u.YesError(t, e)
}
