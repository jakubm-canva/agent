package agent

import (
	"fmt"
	"strings"
	"testing"

	"github.com/buildkite/agent/v3/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTruncateEnv(t *testing.T) {
	l := logger.NewBuffer()
	env := map[string]string{"FOO": strings.Repeat("a", 100)}
	err := truncateEnv(l, env, "FOO", 64)
	require.NoError(t, err)
	assert.Equal(t, "aaaaaaaaaaaaaaaaaaaaaaaaaa[value truncated 100 -> 59 bytes]", env["FOO"])
	assert.Equal(t, 64, len(fmt.Sprintf("FOO=%s\000", env["FOO"])))
}

func TestValidateRepository(t *testing.T) {
	pipelineRepo := "github.com/buildkite/test"
	allowedRepos := []string{}

	// empty allow list
	err := validateRepository(allowedRepos, pipelineRepo)
	require.NoError(t, err)

	// not allowed
	allowedRepos = append(allowedRepos, "^github.com/nope/.*")
	err = validateRepository(allowedRepos, pipelineRepo)
	require.Error(t, err)

	// allowed
	allowedRepos = append(allowedRepos, "^github.com/buildkite/.*")
	err = validateRepository(allowedRepos, pipelineRepo)
	require.NoError(t, err)
}
