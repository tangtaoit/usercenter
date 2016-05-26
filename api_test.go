package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"strings"
)

func Test(t *testing.T) {
	TestingT(t)
}
var _ = Suite(&MiddlewaresTestSuite{})
var server *negroni.Negroni

type MiddlewaresTestSuite struct{}

func (s *MiddlewaresTestSuite) SetUpTest(c *C) {

	//router := NewRouter()
	//server = negroni.Classic()
	//server.UseHandler(router)
}

func (self *MiddlewaresTestSuite) TestQueryUserInfo(c *C) {

	//_,err :=QueryUserInfo("test","1");
	//
	//if err!=nil{
	//	c.Error(err)
	//}

}

func  (self *MiddlewaresTestSuite) TestBindUserInfo(c *C) {

	resource := "/users/auth/test/1"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", resource, strings.NewReader("{\"r_id\":\"1\"}"))
	server.ServeHTTP(response, request)

	assert.Equal(c, response.Code, http.StatusOK)

}



func (self *MiddlewaresTestSuite) TestGetUserInfo(c *C)   {

	resource := "/users/auth/test/1?r_id=1"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	server.ServeHTTP(response, request)

	assert.Equal(c, response.Code, http.StatusOK)

}
