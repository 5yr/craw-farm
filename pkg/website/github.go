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

func (v GitHubSite) FetchTrending() {
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
			fmt.Println(owner, repo)
		})
	})
}
