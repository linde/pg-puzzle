package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocIsLessThanOrEqual(t *testing.T) {

	assert := assert.New(t)
	assert.NotNil(assert)

	assert.True(Loc{0, 0}.IsLessThanOrEqual(Loc{10, 0}))
	assert.True(Loc{0, 0}.IsLessThanOrEqual(Loc{0, 10}))

	assert.True(Loc{0, 0}.IsLessThanOrEqual(Loc{0, 0}))

	assert.False(Loc{10, 0}.IsLessThanOrEqual(Loc{0, 0}))
	assert.False(Loc{0, 10}.IsLessThanOrEqual(Loc{0, 0}))

}
