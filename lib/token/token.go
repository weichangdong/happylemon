package token

import (
	"happylemon/lib/mytime"
	"happylemon/lib/util"
)

func GenNewToken() string {
	naMiao := mytime.Nanosecond()
	naMiaoStr := util.Int64ToString(naMiao)
	randStr := util.GetRandomString(8, "")
	oktoken := util.Md5(naMiaoStr) + randStr
	return oktoken
}
