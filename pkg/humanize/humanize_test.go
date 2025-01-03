package humanize_test

import (
	"testing"

	"github.com/pourtorabehsan/myvisor/pkg/humanize"
	"github.com/stretchr/testify/assert"
)

func TestSI(t *testing.T) {
	assert.Equal(t, "-", humanize.SI(-1000))
	assert.Equal(t, "0", humanize.SI(0))
	assert.Equal(t, "999", humanize.SI(999))
	assert.Equal(t, "1K", humanize.SI(1000))
	assert.Equal(t, "1M", humanize.SI(1000000))
	assert.Equal(t, "1G", humanize.SI(1000000000))
	assert.Equal(t, "1T", humanize.SI(1000000000000))
	assert.Equal(t, "1.2K", humanize.SI(1234))
	assert.Equal(t, "1.2M", humanize.SI(1234567))
	assert.Equal(t, "1.2G", humanize.SI(1234567890))
	assert.Equal(t, "1.2T", humanize.SI(1234567890123))
}
