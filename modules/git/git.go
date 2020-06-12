package git

import (
	"os"
	"strings"

	gg "github.com/go-git/go-git/v5"
)

var repo *gg.Repository

const (
	untracked uint8 = iota
	modified
	added
	deleted
)

func init() {
	repo, _ = gg.PlainOpenWithOptions(os.Getenv("PWD"), &gg.PlainOpenOptions{
		DetectDotGit: true,
	})
}

func CurrentBranch() string {
	if repo == nil {
		return ""
	}

	head, err := repo.Head()
	if err != nil {
		return ""
	}

	return head.Name().Short()
}

func Status() string {
	if repo == nil {
		return ""
	}

	wt, err := repo.Worktree()
	if err != nil {
		return ""
	}

	status, err := wt.Status()
	if err != nil {
		return ""
	}

	flags := []string{"", "", "", ""}

	for _, s := range status {
		switch s.Worktree {
		case gg.Untracked:
			flags[untracked] = "?"
		case gg.Modified:
			flags[modified] = "!"
		case gg.Added:
			flags[added] = "+"
		case gg.Deleted:
			flags[deleted] = "-"
		}

	}
	return strings.Join(flags, "")
}
