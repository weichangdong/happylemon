package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

/*
//接口请求次数
var HttpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "http request count",
	},
	[]string{"endpoint"},
)*/

//接口请求次数及响应时间
var httpRequestDuration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http request duration",
	},
	[]string{"endpoint"},
)

// http 状态码返回次数统计
var httpCode = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_code_count",
		Help: "http_code_count",
	},
	[]string{"code"},
)

//接口 json里ret返回统计次数
var jsonRet = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_ret_count",
		Help: "http_ret_count",
	},
	[]string{"ret"},
)


//注册到prometheus里
func init()  {
	//prometheus.MustRegister(HttpRequestCount)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(jsonRet)
	prometheus.MustRegister(httpCode)
}

func elapsedApi(c *gin.Context){
	if val, ok := c.Get("httpReqTime"); ok && val != nil{
		t, _ := val.(time.Time)
		elapsed := (float64)(time.Since(t) / time.Millisecond)
		httpRequestDuration.WithLabelValues(c.Request.URL.Path).Observe(elapsed)
	}
}

func Handler()  gin.HandlerFunc{
	handler := promhttp.Handler()
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func ReturnCount(c *gin.Context,ret int) {
	elapsedApi(c)
	jsonRet.WithLabelValues(strconv.Itoa(ret)).Inc()
	httpCode.WithLabelValues(strconv.Itoa(200)).Inc()
}

func HttpCodeCount(c *gin.Context,code int){
	if(code !=404){
		elapsedApi(c)
	}
	httpCode.WithLabelValues(strconv.Itoa(code)).Inc()
}
