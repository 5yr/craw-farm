package website

import "fmt"

type Quest func(params ...string) (json []byte, err error)

type siteProcessor interface {
	run(job ...string) (json []byte, err error)
	list() (proc []string)
}

type sites map[string]siteProcessor

var s sites = make(sites)

func Setup() {
	s["v2ex"] = NewV2EX()
	s["github"] = NewGitHub()
}

func Run(site string, op string, params []string) (json []byte, err error) {
	p, exist := s[site]
	if !exist {
		return nil, fmt.Errorf("site not exist")
	}

	runParams := make([]string, 0)
	runParams = append(runParams, op)
	runParams = append(runParams, params...)

	return p.run(runParams...)
}
