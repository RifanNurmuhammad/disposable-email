package disposableemail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDomain(t *testing.T) {
	t.Run("Test Validate Domain", func(t *testing.T) {
		boolTrue := IsDisabledEmail("temp@getnada.com")
		assert.True(t, boolTrue)

		boolFalse := IsDisabledEmail("temp@google.com")
		assert.False(t, boolFalse)

		boolFalse = IsDisabledEmail("google.com")
		assert.False(t, boolFalse)

		boolTrue = IsDisabledDomain("getnada.com")
		assert.True(t, boolTrue)

		boolFalse = IsDisabledDomain("gmail.com")
		assert.False(t, boolFalse)
	})
}
