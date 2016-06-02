package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
	"crypto/md5"
	"time"
	"github.com/tangtaoit/util"
	"encoding/json"
	"bytes"
)

var server *negroni.Negroni

var basesign string
var noncestr string
var timestamp string
var apikey string

func initSetting()  {

	os.Setenv("GO_ENV", "tests")

	apikey ="f99b45c3f5d747658d421c9e75c469eb"
	noncestr ="23435"
	timestamp =fmt.Sprintf("%d",time.Now().Unix())

	signStr := apikey+noncestr+timestamp
	bytes  := md5.Sum([]byte(signStr))
	basesign =fmt.Sprintf("%X",bytes)

	router :=GetRouters()
	server = negroni.Classic()
	server.UseHandler(router)
}

func TestBindUserInfo(t *testing.T) {

	initSetting()

	resource := "/users/auth"

	params :=map[string]interface{}{
		"r_id":"1",
	}

	sign := util.SignWithBaseSign(params,apikey,basesign,nil)
	paramsBytes,err := json.Marshal(params)
	util.CheckErr(err)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", resource, bytes.NewReader(paramsBytes))
	request.Header.Set("app_id","194251277981454336")
	request.Header.Set("sign",fmt.Sprintf("%s.%s",basesign,sign))
	request.Header.Set("noncestr",noncestr)
	request.Header.Set("timestamp",timestamp)

	//request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)

	var result map[string]interface{}
	err = json.Unmarshal(response.Body.Bytes(),&result)
	util.CheckErr(err)

	fmt.Println("result =%s",result)


}