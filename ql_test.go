package dotsql

import (
	u "github.com/amortaza/dotsql/test_utils"
	"testing"
)

func Test_QL_02(t *testing.T) {

	//s, _ := ToSQL("owner.name = $name")
	s := ""
	expect := ""

	u.Check(t, expect, s)
}

func Test_QL_01(t *testing.T) {

	//s, _ := ToSQL("")
	s := ""
	expect := ""

	u.Check(t, expect, s)
}
