package process

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/5yr/nnew/pkg/website"
	"github.com/spf13/viper"
)

const DefaultSep = "."

type Task struct {
	Name     string   `mapstructure:"name"`
	Sequence []string `mapstructure:"seq"`
}

func LoadTask(absPath string) (t *Task, err error) {
	var (
		f *os.File
	)
	if !path.IsAbs(absPath) {
		return nil, fmt.Errorf("need absolute path")
	}

	if f, err = os.Open(absPath); err != nil {
		return nil, fmt.Errorf("file open failed")
	}

	v := viper.New()
	v.SetConfigType("toml")
	if err = v.ReadConfig(f); err != nil {
		return nil, fmt.Errorf("read config error")
	}

	t = &Task{}
	if err = v.Unmarshal(t); err != nil {
		return nil, err
	}
	return t, nil
}

func SplitParam(taskDesc string) (site string, op string, params []string) {
	res := strings.Split(taskDesc, DefaultSep)
	frag := len(res)
	if frag < 2 {
		panic("task desc invalid")
	} else if frag == 2 {
		return res[0], res[1], make([]string, 0)
	} else {
		return res[0], res[1], res[2:]
	}
}

func (t Task) ExecSequence() {
	fmt.Println("Start Process Sequence: ", t.Name)
	for i, jobDesc := range t.Sequence {
		res, err := website.Run(SplitParam(jobDesc))
		if err != nil {
			fmt.Println("ERROR: ", err)
			continue
		}
		fmt.Printf("Job-%d: Success!\nResult: \n%s\n", i, string(res))
	}
}
