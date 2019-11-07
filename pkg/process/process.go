package process

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

type Task struct {
	Name     string   `mapstruct:"name"`
	Sequence []string `mapstruct:"seq"`
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
	if err = v.ReadConfig(f); err != nil {
		return nil, fmt.Errorf("read config error")
	}

	t = &Task{}
	if err = v.Unmarshal(t); err != nil {
		return nil, err
	}
	return t, nil
}
