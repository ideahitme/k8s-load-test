package parser

import (
	yaml "gopkg.in/yaml.v2"
)

type (
	Header struct {
		Title string `json:"title",yaml:"title"`
	}
)

func parseHeader(header []byte) (*Header, error) {
	h := &Header{}
	err := yaml.Unmarshal(header, h)
	return h, err
}
