package git

import (
	"strings"

	"github.com/arl/gitstatus"
)

var status *gitstatus.Status

func init() {
	status, _ = gitstatus.New()
}

func CurrentBranch() string {
	if status == nil {
		return ""
	}

	if status.IsDetached {
		return "(detached " + status.HEAD + ")"
	}

	return status.LocalBranch
}

func RemoteBranch() string {
	if status == nil {
		return ""
	}
	return status.RemoteBranch
}

func Status(separator string, statuses ...func() string) func() string {
	if status == nil {
		return func() string {
			return ""
		}
	}

	result := make([]string, len(statuses))
	for i := range statuses {
		result[i] = statuses[i]()
	}

	return func() string {
		return strings.Join(result, separator)
	}

}

func HasStashed(symbol string) func() string {
	return func() string {
		if status == nil || status.NumStashed == 0 {
			return ""
		}
		return symbol
	}
}

func HasModified(symbol string) func() string {
	return func() string {
		if status == nil || status.NumModified == 0 {
			return ""
		}
		return symbol
	}
}

func HasStaged(symbol string) func() string {
	return func() string {
		if status == nil || status.NumStaged == 0 {
			return ""
		}
		return symbol
	}
}

func HasUntracked(symbol string) func() string {
	return func() string {
		if status == nil || status.NumUntracked == 0 {
			return ""
		}
		return symbol
	}
}
