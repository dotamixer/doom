package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Plugin struct {
	Name string `yaml:"name"`
	//Output string `yaml:"output"`
	Flags string `yaml:"flags"`
	Type  string `yaml:"type"`
}

type GoOptions struct {
	Modifiers map[string]string `yaml:"modifiers"`
}

type Generate struct {
	GoOptions GoOptions `yaml:"go_options"`
	Plugins   []Plugin  `yaml:"plugins"`
	Output    string    `yaml:"output"`
	Modifier  string    `yaml:"modifier"`
}

type Config struct {
	Module     string   `yaml:"module"`
	ImportPath string   `yaml:"import_path"`
	Protos     []string `yaml:"protos"`
	Includes   []string `yaml:"includes"`
	Generate   Generate `yaml:"generate"`
}

func NewConfig() (c *Config) {
	c = &Config{}
	return
}

func (c *Config) Output() (err error) {
	var (
		file *os.File
	)
	file, err = os.Create("idl.yaml")
	if err != nil {
		err = errors.WithMessagef(err, "create idl.yaml failed")
		return
	}

	_, err = file.WriteString(tmpl)
	if err != nil {
		err = errors.WithMessagef(err, "write idl.yaml failed")
	}
	return
}

func (c *Config) Load() (err error) {
	f, err := ioutil.ReadFile("idl.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, c)
	if err != nil {
		return err
	}

	//支持环境变量
	var absPath []string
	for _, itr := range c.Includes {
		tmp := os.ExpandEnv(itr)
		absPath = append(absPath, tmp)
	}

	c.Includes = absPath
	c.ImportPath = os.ExpandEnv(c.ImportPath)
	c.Includes = append(c.Includes, c.ImportPath)
	for _, itr := range c.Includes {
		log.Println("include path:", itr)
	}
	return nil
}
