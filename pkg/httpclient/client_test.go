package httpclient

import (
	"testing"
	"fmt"
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

func TestAny(t *testing.T) {
	a := func() string { return "first" } ()
	if a == "first" {
		a = func() string { return "second" } ()
		fmt.Printf("Inner a: %s", a)
	} 
	fmt.Printf("Outer a: %s", a)
}