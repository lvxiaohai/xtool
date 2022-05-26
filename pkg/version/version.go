package version

import (
	"fmt"
	"runtime"
)

// Info contains version information.
type Info struct {
	GitAuthor    string `json:"git_author"`
	GitTag       string `json:"git_tag"`
	GitCommit    string `json:"git_commit"`
	GitTreeState string `json:"git_tree_state"`
	BuildDate    string `json:"build_date"`
	GoVersion    string `json:"go_version"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitAuthor:    gitAuthor,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
