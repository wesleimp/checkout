package git_test

import (
	"os"
	"testing"

	"github.com/wesleimp/checkout/internal/git"

	"github.com/stretchr/testify/require"
)

func TestGit(t *testing.T) {
	out, err := git.Run("status")
	require.NoError(t, err)
	require.NotEmpty(t, out)

	out, err = git.Run("command-that-dont-exist")
	require.Error(t, err)
	require.Empty(t, out)
	require.Equal(
		t,
		"git: 'command-that-dont-exist' is not a git command. See 'git --help'.\n",
		err.Error(),
	)
}

func TestRepo(t *testing.T) {
	require.NoError(t, os.Chdir(os.TempDir()))
	require.False(t, git.IsRepo(), os.TempDir()+" folder should be a git repo")
}
