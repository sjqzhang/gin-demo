
package app

import (
	"testing"
	"encoding/json"
	"github.com/sjqzhang/requests"


)


func TestCategoryControllerCreate(t *testing.T) {
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
	if result["code"].(float64) != 200 {
		t.Fail()
	}

	
}



func TestCategoryControllerUpdate(t *testing.T) {
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
	if result["code"].(float64) != 200 {
		t.Fail()
	}

	
}



func TestCategoryControllerShow(t *testing.T) {
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
	if result["code"].(float64) != 200 {
		t.Fail()
	}

	
}



func TestCategoryControllerDelete(t *testing.T) {
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
	if result["code"].(float64) != 200 {
		t.Fail()
	}

	
}

