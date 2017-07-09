package parser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HeaderTestSuite struct {
	suite.Suite
	headerYAMLgood string
	headerJSONgood string
}

func (suite *HeaderTestSuite) SetupTest() {
	suite.headerYAMLgood = `title: Something`
	suite.headerJSONgood = `{"title": "Something"}`
}

func (suite *HeaderTestSuite) TestGoodHeaders() {
	h, err := parseHeader([]byte(suite.headerJSONgood))
	suite.NoError(err, "should not be any error for valid JSON header")
	suite.Equal(Header{Title: "Something"}, *h, "header should be successfully parsed")

	h, err = parseHeader([]byte(suite.headerYAMLgood))
	suite.NoError(err, "should not be any error for valid JSON header")
	suite.Equal(Header{Title: "Something"}, *h, "header should be successfully parsed")
}

func TestHeader(t *testing.T) {
	suite.Run(t, new(HeaderTestSuite))
}
