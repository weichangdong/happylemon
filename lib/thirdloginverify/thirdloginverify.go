package thirdloginverify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Verify(typeThird float64, tokenThrid, uniq_id string) bool {
	switch typeThird {
	case 1:
		return wxVerify(tokenThrid, uniq_id, typeThird)
	case 2:
		return qqVerify(tokenThrid, uniq_id, typeThird)
	case 3:
		return wbVerify(tokenThrid, uniq_id, typeThird)
	default:
		return false
	}
}

//微信登录验证
func wxVerify(tokenThrid string, uniq_id string, typeThird float64) bool {
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + tokenThrid + "&openid=opPRBwyF1OeoDOuF5Kicm2yfWnhs" + uniq_id
	return verifyQequest(url, typeThird)
}

//QQ登录验证
func qqVerify(tokenThrid string, uniq_id string, typeThird float64) bool {
	url := "https://graph.qq.com/user/get_user_info?access_token=" + tokenThrid + "&oauth_consumer_key=1106802748&openid=" + uniq_id
	return verifyQequest(url, typeThird)
}

//微博登录验证
func wbVerify(tokenThrid string, uniq_id string, typeThird float64) bool {
	url := "https://api.weibo.com/2/users/show.json?access_token=" + tokenThrid + "&uid=" + uniq_id
	return verifyQequest(url, typeThird)
}

func verifyQequest(url string, typeThird float64) bool {
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	reqest.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	bodyInfo, errReq := ioutil.ReadAll(response.Body)
	if errReq != nil {
		return false
	}
	okData := map[string]interface{}{}
	errJson := json.Unmarshal(bodyInfo, &okData)
	if errJson != nil {
		return false
	}
	if typeThird == 1 {
		if okData["openid"] == nil || okData["errcode"] != nil {
			return false
		}
		return true
	} else if typeThird == 2 {
		if okData["ret"] == nil || okData["ret"].(float64) != 0 {
			return false
		}
	} else if typeThird == 3 {
		if okData["error_code"] != nil {
			return false
		}
	}
	return true
}
