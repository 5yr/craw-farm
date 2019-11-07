package website

type sites struct {
	v2ex   *V2EXSite
	github *GitHubSite
}

var s sites

func Setup() {
	s.v2ex = NewV2EX()
	s.github = NewGitHub()
	// test
	s.github.FetchTrending()
}
