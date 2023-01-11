package puzzle

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFunc(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	strOut := Map([]string{"a", "b", "c"}, strings.ToUpper)
	assert.Equal(strOut, []string{"A", "B", "C"})

	f64Out := Map([]float64{1, -2, 3, -4}, math.Abs)
	assert.Equal(f64Out, []float64{1, 2, 3, 4})

	str2Out := Map([]string{"a", "b", "c"}, func(s string) string { return s + s })
	assert.Equal(str2Out, []string{"aa", "bb", "cc"})

}
