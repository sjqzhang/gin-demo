
package app

import (
	"encoding/json"
	"github.com/sjqzhang/gmock/util"
	"github.com/sjqzhang/requests"
	"testing"


)


func TestPostControllerCreate(t *testing.T) {
	//Reset()
	var result map[string]interface{} //gin.Context
	reqJson := `{
  "category_id": 1,
  "title": "demoString",
  "head_img": "demoString",
  "content": "demoString"
}`
	resp, err := requests.PostJson("http://127.0.0.1:8081/posts", reqJson)
	if err != nil {
		t.Fail()
	}
	if resp.R.StatusCode != 200 {
		t.Fail()
	}
	err = json.Unmarshal([]byte(resp.Text()), &result)
	if err != nil {
		t.Fail()
	}
	if util.Util.Jq(result,"code").(float64) != 200 {
		t.Fail()
	}

	
}



func TestPostControllerUpdate(t *testing.T) {
	//Reset()
	var result map[string]interface{} //gin.Context
	reqJson := `{}`
	resp, err := requests.PostJson("http://127.0.0.1:8081/", reqJson)
	if err != nil {
		t.Fail()
	}
	if resp.R.StatusCode != 200 {
		t.Fail()
	}
	err = json.Unmarshal([]byte(resp.Text()), &result)
	if err != nil {
		t.Fail()
	}
	if util.Util.Jq(result,"code").(float64) != 200 {
		t.Fail()
	}

	
}



func TestPostControllerShow(t *testing.T) {
	//Reset()
	var result map[string]interface{} //gin.Context
	reqJson := `{}`
	resp, err := requests.PostJson("http://127.0.0.1:8081/", reqJson)
	if err != nil {
		t.Fail()
	}
	if resp.R.StatusCode != 200 {
		t.Fail()
	}
	err = json.Unmarshal([]byte(resp.Text()), &result)
	if err != nil {
		t.Fail()
	}
	if util.Util.Jq(result,"code").(float64) != 200 {
		t.Fail()
	}

	
}



func TestPostControllerDelete(t *testing.T) {
	//Reset()
	var result map[string]interface{} //gin.Context
	reqJson := `{}`
	resp, err := requests.PostJson("http://127.0.0.1:8081/", reqJson)
	if err != nil {
		t.Fail()
	}
	if resp.R.StatusCode != 200 {
		t.Fail()
	}
	err = json.Unmarshal([]byte(resp.Text()), &result)
	if err != nil {
		t.Fail()
	}
	if util.Util.Jq(result,"code").(float64) != 200 {
		t.Fail()
	}

	
}



func TestPostControllerPageList(t *testing.T) {
	//Reset()
	var result map[string]interface{} //gin.Context
	reqJson := `{}`
	resp, err := requests.PostJson("http://127.0.0.1:8081/", reqJson)
	if err != nil {
		t.Fail()
	}
	if resp.R.StatusCode != 200 {
		t.Fail()
	}
	err = json.Unmarshal([]byte(resp.Text()), &result)
	if err != nil {
		t.Fail()
	}
	if util.Util.Jq(result,"code").(float64) != 200 {
		t.Fail()
	}

	
}

