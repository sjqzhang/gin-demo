
package app

import (
	"encoding/json"
	"github.com/sjqzhang/gmock/util"
	"github.com/sjqzhang/requests"
	"testing"


)


func TestPostControllerCreate(t *testing.T) {
	Reset()
	var result map[string]interface{} //gin.Context
	reqJson := `{
  "category_id": 1,
  "title": "title",
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
	Reset()
	var result map[string]interface{} //gin.Context
	reqJson := `{
  "category_id": 1,
  "title": "new",
  "head_img": "demoString dxxxsafdsf",
  "content": "demoString  asdfasdfsdf"
}`
	resp, err := requests.PutJson("http://127.0.0.1:8081/posts/86df2a8c-c840-4f2b-b075-78beda7cd849", reqJson)
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
	Reset()
	var result map[string]interface{} //gin.Context
	//reqJson := `{}`
	resp, err := requests.Delete("http://127.0.0.1:8081/posts/86df2a8c-c840-4f2b-b075-78beda7cd849")
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


