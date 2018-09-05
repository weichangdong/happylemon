package controllers

import (
	"encoding/json"
	"happylemon/conf"
	"happylemon/lib/encrypt"
	"runtime"

	"happylemon/lib/prometheus"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (self *BaseController) Response(c *gin.Context, ret int, data interface{}) {

	prometheus.ReturnCount(c, ret)

	mode := conf.Config.Server.Mode
	if mode != "online" {
		_, file, line, _ := runtime.Caller(1)
		if ret != conf.CodeOk {
			c.JSON(200, gin.H{"ret": ret, "line": line, "file": file})
			return
		}
	} else {
		if ret != conf.CodeOk {
			c.JSON(200, gin.H{"ret": ret})
			return
		}
	}

	if ret == conf.CodeOk && data == nil {
		c.JSON(200, gin.H{"ret": conf.CodeOk})
		return
	}
	if self.isEncry(c.Query("aes")) {
		dataStr, _ := json.Marshal(data)
		ok, okData := encrypt.Encryper(dataStr)
		if !ok {
			c.JSON(200, gin.H{"ret": conf.CodeAesErr})
			return
		}
		c.JSON(200, gin.H{"ret": conf.CodeOk, "dat": okData})
	} else {
		c.JSON(200, gin.H{"ret": conf.CodeOk, "dat": data})
	}
}

//post 方法统一接收数据,没有做json的解析
func (self *BaseController) Request(c *gin.Context) (int, []byte) {
	postData, err := c.GetRawData()
	if err != nil {
		return conf.CodeParaErr, nil
	}
	if self.isEncry(c.Query("aes")) {
		ok, okData := encrypt.Decryper(postData)
		if !ok {
			return conf.CodeAesErr, nil
		}
		return conf.CodeOk, okData
	} else {
		return conf.CodeOk, postData
	}
}

//接收特殊场景的数据
func (self *BaseController) SpRequest(c *gin.Context) (int, []byte) {
	tmpData := c.PostForm("data")
	postData := []byte(tmpData)
	if self.isEncry(c.Query("aes")) {
		ok, okData := encrypt.Decryper(postData)
		if !ok {
			return conf.CodeAesErr, nil
		}
		return conf.CodeOk, okData
	} else {
		return conf.CodeOk, postData
	}
}

//判断是否有这个参数字段
func (self *BaseController) ParamMustExist(paraTmp interface{}, key string) bool {
	switch paraTmp.(type) {
	case map[string]float64:
		para := paraTmp.(map[string]float64)
		if _, ok := para[key]; !ok {
			return false
		}
		return true
	case map[string]string:
		para := paraTmp.(map[string]float64)
		if _, ok := para[key]; !ok {
			return false
		}
		return true
	case map[string]interface{}:
		para := paraTmp.(map[string]interface{})
		if _, ok := para[key]; !ok {
			return false
		}
		return true
	}

	return true
}

func (self *BaseController) SpPostInterface(c *gin.Context) (code int, okData map[string]interface{}) {
	code, data := self.SpRequest(c)
	if code != conf.CodeOk {
		return code, nil
	}
	err := json.Unmarshal(data, &okData)
	if err != nil {
		return conf.CodeParaErr, nil
	}
	return
}

func (self *BaseController) SpPostString(c *gin.Context) (code int, okData map[string]string) {
	code, data := self.SpRequest(c)
	if code != conf.CodeOk {
		return code, nil
	}
	err := json.Unmarshal(data, &okData)
	if err != nil {
		return conf.CodeParaErr, nil
	}
	return
}

func (self *BaseController) SpPostFloat64(c *gin.Context) (code int, okData map[string]float64) {
	code, data := self.SpRequest(c)
	if code != conf.CodeOk {
		return code, nil
	}
	err := json.Unmarshal(data, &okData)
	if err != nil {
		return conf.CodeParaErr, nil
	}
	return
}

// 验证post body中的数据
func (self *BaseController) PostInterface(c *gin.Context) (code int, okData map[string]interface{}) {
	code, data := self.Request(c)
	if code != conf.CodeOk {
		return code, nil
	}
	err := json.Unmarshal(data, &okData)
	if err != nil {
		return conf.CodeParaErr, nil
	}
	return
}

// 验证post body中的数据
func (self *BaseController) PostString(c *gin.Context) (code int, okData map[string]string) {
	code, data := self.Request(c)
	if code != conf.CodeOk {
		return code, nil
	}
	err := json.Unmarshal(data, &okData)
	if err != nil {
		return conf.CodeParaErr, nil
	}
	return
}

// 验证post body中的数据
func (self *BaseController) PostFloat64(c *gin.Context) (code int, okData map[string]float64) {
	code, data := self.Request(c)
	if code != conf.CodeOk {
		return code, nil
	}
	err := json.Unmarshal(data, &okData)
	if err != nil {
		return conf.CodeParaErr, nil
	}
	return
}

func (self *BaseController) CheckImgSize(size int64) bool {
	if size > 10240 {
		return true
	}
	return true
}

func (*BaseController) isEncry(aes string) bool {
	mode := conf.Config.Server.Mode
	if aes == "1" && mode != "online" {
		return false
	} else {
		return true
	}
}
