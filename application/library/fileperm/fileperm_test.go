package fileperm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webx-top/echo"
)

func TestVerify(t *testing.T) {
	u := User{
		RootDir: `/ftpdir`,
		Rules: Rules{
			NewRule(`a`, true, nil, false).SetWriteable(true),
			NewRule(`b`, true, nil, false),
			NewRule(`/c/`, true, nil, false).SetWriteable(true),
		},
	}
	allowed := u.Allowed(`.`, true)
	assert.False(t, allowed)

	allowed = u.Allowed(`a`, true)
	assert.True(t, allowed)
	allowed = u.Allowed(echo.FilePathSeparator+`a`, true)
	assert.True(t, allowed)
	allowed = u.Allowed(echo.FilePathSeparator+`a`+echo.FilePathSeparator, true)
	assert.True(t, allowed)
	allowed = u.Allowed(`a`, false)
	assert.True(t, allowed)

	allowed = u.Allowed(`b`, true)
	assert.False(t, allowed)

	allowed = u.Allowed(`c`, false)
	assert.True(t, allowed)
	allowed = u.Allowed(`c`, true)
	assert.True(t, allowed)

	allowed = u.Allowed(`/c`, false)
	assert.True(t, allowed)
	allowed = u.Allowed(`/c`, true)
	assert.True(t, allowed)

	allowed = u.Allowed(`/c/`, false)
	assert.True(t, allowed)
	allowed = u.Allowed(`/c/`, true)
	assert.True(t, allowed)
}
