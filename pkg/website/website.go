package website

type sites struct {
	v2ex *V2EXSite
}

var s sites

func Setup() {
	s.v2ex = NewV2EX()

	// test
	s.v2ex.FetchPostList("all")
}
