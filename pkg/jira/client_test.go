package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUrl(t *testing.T) {

	var actualUrl string
	
	actualUrl = buildUrl("http://server1.org/", "/api/get?id=8")
	assert.Equal(t, "http://server1.org/api/get?id=8", actualUrl)

	actualUrl = buildUrl("http://server2.org/", "http://server3.org/api/get?id=8")
	assert.Equal(t, "http://server3.org/api/get?id=8", actualUrl)

	actualUrl = buildUrl("http://server4.org", "api/get?name=John&locale=en")
	assert.Equal(t, "http://server4.org/api/get?name=John&locale=en", actualUrl)

}