package website

import (
	"fmt"
	"strings"

	"github.com/5yr/nnew/pkg/service"
	"github.com/5yr/nnew/pkg/setting"
	"github.com/PuerkitoBio/goquery"
)

type GitHubSite struct {
	Base string `mapstructure:"base"`
}

func NewGitHub() *GitHubSite {
	v := GitHubSite{}
	err := setting.GetSiteConfig("github", &v)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &v
}

func (v GitHubSite) generateSubpageURL(pageName string) string {
	return fmt.Sprintf("%s/%s", v.Base, pageName)
}

func (g GitHubSite) list() []string {
	return []string{"trending"}
}

func (g GitHubSite) run(params ...string) (json []byte, err error) {
	jobName := params[0]

	switch jobName {
	case "trending":
		return g.FetchTrending()
	}
	return nil, fmt.Errorf("job name not exist")
}

func (v GitHubSite) FetchTrending() (json []byte, err error) {
	res := make([]string, 0)
	service.JQGet(v.generateSubpageURL("trending"), func(doc *goquery.Document) {
		doc.Find("main div.Box article").Each(func(i int, s *goquery.Selection) {
			repoLink, _ := s.Find("h1 a").Attr("href")
			repoLinkSplit := strings.Split(repoLink, `/`)
			if len(repoLinkSplit) < 3 {
				fmt.Println("error: unexpected repo link format")
			}
			// the repoLinkSplit[0] == ""
			owner := repoLinkSplit[1]
			repo := repoLinkSplit[2]
			res = append(res, fmt.Sprintf(owner, repo))
		})
	})

	resB := strings.Join(res, ",")
	return []byte(resB), nil
}
