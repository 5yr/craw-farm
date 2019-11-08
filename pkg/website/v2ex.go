package website

import (
	"fmt"
	"strings"

	"github.com/5yr/nnew/pkg/service"
	"github.com/5yr/nnew/pkg/setting"
	"github.com/PuerkitoBio/goquery"
)

type V2EXSiteConfig struct {
	Base string `mapstructure:"base"` //url, main page
}

type V2EX struct {
	Cfg *V2EXSiteConfig
	Op  map[string]func(v *V2EX, params ...string) (json []byte, err error)
}

func NewV2EX() *V2EX {
	var err error

	vcfg := V2EXSiteConfig{}
	if err = setting.GetSiteConfig("v2ex", &vcfg); err != nil {
		fmt.Println(err)
		return nil
	}

	v := V2EX{
		Cfg: &vcfg,
		Op: map[string]func(v *V2EX, params ...string) (json []byte, err error){
			"post": FetchPostList,
		},
	}
	return &v
}

func (v V2EX) list() (ops []string) {
	ops = make([]string, 0)
	for k := range v.Op {
		ops = append(ops, k)
	}
	return ops
}

func (v V2EX) run(job ...string) (b []byte, err error) {
	questName := job[0]
	q, exist := v.Op[questName]
	if !exist {
		return nil, fmt.Errorf("quest not exist")
	}
	return q(&v, job[1:]...)
}

func generateSubpageURL(base string, pageName string) string {
	return fmt.Sprintf("%s/?tab=%s", base, pageName)
}

func FetchPostList(v *V2EX, params ...string) ([]byte, error) {
	name := params[0]

	res := make([]string, 0)
	service.JQGet(generateSubpageURL(v.Cfg.Base, name), func(doc *goquery.Document) {
		doc.Find(".content .box .cell.item").Each(func(i int, s *goquery.Selection) {
			titleNode := s.Find("span.item_title a.topic-link")
			link, _ := titleNode.Attr("href")
			title := titleNode.Text()

			topicInfoNode := s.Find("span.topic_info")
			topic := topicInfoNode.Find("a.node").Text()
			author := s.Find("strong a").Text()
			replyNum := s.Find("a.count_livid").Text()

			res = append(res, fmt.Sprintf(replyNum, link, topic, author, title))
		})
	})
	resB := strings.Join(res, ", ")
	return []byte(resB), nil
}
