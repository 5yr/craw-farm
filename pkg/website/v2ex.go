package website

import (
	"fmt"

	"github.com/5yr/nnew/pkg/service"
	"github.com/5yr/nnew/pkg/setting"
	"github.com/PuerkitoBio/goquery"
)

type V2EXSite struct {
	Base string `mapstructure:"base"` //url, main page
}

func NewV2EX() *V2EXSite {
	v := V2EXSite{}
	err := setting.GetSiteConfig("v2ex", &v)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &v
}

func (v V2EXSite) generateSubpageURL(pageName string) string {
	return fmt.Sprintf("%s/?tab=%s", v.Base, pageName)
}

func (v V2EXSite) FetchPostList(name string) {
	service.JQGet(v.generateSubpageURL(name), func(doc *goquery.Document) {
		doc.Find(".content .box .cell.item").Each(func(i int, s *goquery.Selection) {
			titleNode := s.Find("span.item_title a.topic-link")
			link, _ := titleNode.Attr("href")
			title := titleNode.Text()

			topicInfoNode := s.Find("span.topic_info")
			topic := topicInfoNode.Find("a.node").Text()
			author := s.Find("strong a").Text()
			replyNum := s.Find("a.count_livid").Text()

			fmt.Println(replyNum, link, topic, author, title)
		})
	})

}
