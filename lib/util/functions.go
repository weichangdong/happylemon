package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	mrand "math/rand"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else if err != nil {
		panic(err)
	} else {
		return true
	}
}
func Str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func Int2str(s int) string {
	i := strconv.Itoa(s)
	return i
}

//检测Email是否正确
func FuncVerifyStringIsEmail(email string) bool {
	preg_email := regexp.MustCompile(`[a-zA-Z0-9-_.]+@[a-zA-Z0-9-.]+\.[a-zA-Z0-9]+`)
	findEmail := preg_email.FindAllString(email, -1)
	if len(findEmail) == 1 && findEmail[0] == email {
		return true
	}
	return false
}

//检测手机号码是否正确
func FuncVerifyStringIsMobile(mobile string) bool {
	preg_mobile := regexp.MustCompile(`^1([38][0-9]|14[57]|5[^4]|7[^4])\d{8}$`)
	findMobile := preg_mobile.FindAllString(mobile, -1)
	if len(findMobile) == 1 && findMobile[0] == mobile {
		return true
	}
	return false
}

//生成随机字符串
func GetRandomString(strlen int, rtype string) string {
	if rtype == "" {
		rtype = "all"
	}
	var str string
	if rtype == "all" {
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if rtype == "numb" {
		str = "0123456789"
	} else {
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bytes := []byte(str)
	result := []byte{}
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < strlen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//检测字符串中非法字符
func CheckFiterString(str string) bool {
	preg_str := regexp.MustCompile(`[(\\(128)-\\(191))|$|¥|￥|·|'|‘|、|\\，|,|。|.|～|~|！|!|＠|@|＃|#|％|%|＆|&×|（|）|(|)|｛|｝|{|}|\[|\]|【|】|/：|:|\*|＊|\+|＋|-|－|—|=|＝|<|﹤|︳|……|^|-|∕|¦|‖|︴|“|《|》|？|?|…|…|｜|：|“|《|》|=|"|;|,| |.|*]`)
	is_true := preg_str.FindAllString(str, -1)
	if len(is_true) > 0 {
		return false
	}
	isTrue := strings.Contains(str, "`")
	if isTrue {
		return false
	}
	return true
}

func Float64InArray(num float64, array []float64) bool {
	l := len(array)
	for i := 0; i < l; i++ {
		if array[i] == num {
			return true
		}
	}
	return false
}
func Utf8Len(str string) int {
	return utf8.RuneCountInString(str)
}
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	defaultIp := "127.0.0.1"
	if err != nil {
		return defaultIp
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return defaultIp
}
func CheckOs(str string) bool {
	if str != "ios" && str != "android" {
		return false
	}
	return true
}

func CheckImei(str string) bool {
	if len(str) != 32 {
		return false
	}
	return true
}

func CheckPass(str string) bool {
	if len(str) != 32 {
		return false
	}
	return true
}

func CheckSmsCode(str string) bool {
	if len(str) != 4 {
		return false
	}
	return true
}

func CheckSex(str float64) bool {
	if str != 1 && str != 2 {
		return false
	}
	return true
}

func CheckAge(str float64) bool {
	if str < 1 || str > 120 {
		return false
	}
	return true
}

func Int64ToString(from int64) (to string) {
	to = strconv.FormatInt(from, 10)
	return
}

func Float64ToString(from float64) (to string) {
	to = strconv.FormatFloat(from, 'f', -1, 64)
	return
}

func StringToFloat(v string) (d float32, err error) {
	tmp, err := strconv.ParseFloat(v, 10)
	d = float32(tmp)
	return
}

func StringToFloat64(v string) (d float64, err error) {
	d, err = strconv.ParseFloat(v, 10)
	return
}

func StringToInt64(v string) (d int64, err error) {
	d, err = strconv.ParseInt(v, 10, 64)
	return
}

func Float64ToInt(v float64) (d int, err error) {
	vStr := strconv.FormatFloat(v, 'f', -1, 64)
	d, err = strconv.Atoi(vStr)
	return
}

//通过空格，分割字符串
func SplitStrFromSpace(key string) []string {
	spacePat := "\\s+" //正则
	spacePreg, _ := regexp.Compile(spacePat)
	keyNew := spacePreg.ReplaceAllString(key, " ")
	cityList := strings.Split(keyNew, " ")
	return cityList
}

//格式话手机号码( 隐藏中间四位 )
func FormatMobile(mob string) string {
	mobArr := strings.Split(mob, "")
	formatMob := ""
	for k, v := range mobArr {
		if k > 2 && k < 7 {
			formatMob = formatMob + "*"
		} else {
			formatMob = formatMob + v
		}
	}
	return formatMob
}

func Implode(data []string, glue string) string {
	var tmp []string
	for _, item := range data {
		tmp = append(tmp, item)
	}
	return strings.Join(tmp, glue)
}

func InArray(id int64, ids []string) bool {
	for _, vstr := range ids {
		v, _ := StringToInt64(vstr)
		if v == id {
			return true
		}
	}
	return false
}

func InStrArray(str string, strs []string) bool {
	for _, vstr := range strs {
		if vstr == str {
			return true
		}
	}
	return false
}

func IpInArray(str string, strs []string) bool {
	for _, vstr := range strs {
		if vstr == str {
			return true
		}
		matched, _ := regexp.Match(`\.\*`, []byte(vstr))
		if matched {
			newIp := strings.Replace(vstr, ".*", "", -1) + "."
			if strings.Contains(str, newIp) {
				return true
			}
		}
	}
	return false
}
func CheckColorRgb(str string) bool {
	if len(str) != 9 {
		return false
	}
	reg, _ := regexp.Compile(`#[0-9a-fA-F]{4}`)
	match := reg.MatchString(str)
	return match
}

func VersionCompare(version1, version2, operator string) bool {
	var vcompare func(string, string) int
	var canonicalize func(string) string
	var special func(string, string) int

	// version compare
	vcompare = func(origV1, origV2 string) int {
		if origV1 == "" || origV2 == "" {
			if origV1 == "" && origV2 == "" {
				return 0
			} else {
				if origV1 == "" {
					return -1
				} else {
					return 1
				}
			}
		}

		ver1, ver2, compare := "", "", 0
		if origV1[0] == '#' {
			ver1 = origV1
		} else {
			ver1 = canonicalize(origV1)
		}
		if origV2[0] == '#' {
			ver2 = origV2
		} else {
			ver2 = canonicalize(origV2)
		}
		n1, n2 := 0, 0
		for {
			p1, p2 := "", ""
			n1 = strings.IndexByte(ver1, '.')
			if n1 == -1 {
				p1, ver1 = ver1, ""
			} else {
				p1, ver1 = ver1[:n1], ver1[n1+1:]
			}
			n2 = strings.IndexByte(ver2, '.')
			if n2 == -1 {
				p2, ver2 = ver2, ""
			} else {
				p2, ver2 = ver2[:n2], ver2[n2+1:]
			}
			if (p1[0] >= '0' && p1[0] <= '9') && (p2[0] >= '0' && p2[0] <= '9') { // all isdigit
				l1, _ := strconv.Atoi(p1)
				l2, _ := strconv.Atoi(p2)
				if l1 > l2 {
					compare = 1
				} else if l1 == l2 {
					compare = 0
				} else {
					compare = -1
				}
			} else if !(p1[0] >= '0' && p1[0] <= '9') && !(p2[0] >= '0' && p2[0] <= '9') { // all isndigit
				compare = special(p1, p2)
			} else { // part isdigit
				if p1[0] >= '0' && p1[0] <= '9' { // isdigit
					compare = special("#N#", p2)
				} else {
					compare = special(p1, "#N#")
				}
			}
			if compare != 0 || n1 == -1 || n2 == -1 {
				break
			}
		}

		if compare == 0 {
			if ver1 != "" {
				if ver1[0] >= '0' && ver1[0] <= '9' {
					compare = 1
				} else {
					compare = vcompare(ver1, "#N#")
				}
			} else if ver2 != "" {
				if ver2[0] >= '0' && ver2[0] <= '9' {
					compare = -1
				} else {
					compare = vcompare("#N#", ver2)
				}
			}
		}

		return compare
	}

	// canonicalize
	canonicalize = func(version string) string {
		ver := []byte(version)
		l := len(ver)
		if l == 0 {
			return ""
		}
		var buf = make([]byte, l*2)
		j := 0
		for i, v := range ver {
			next := uint8(0)
			if i+1 < l { // Have the next one
				next = ver[i+1]
			}
			if v == '-' || v == '_' || v == '+' { // repalce "-","_","+" to "."
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else if (next > 0) &&
				(!(next >= '0' && next <= '9') && (v >= '0' && v <= '9')) ||
				(!(v >= '0' && v <= '9') && (next >= '0' && next <= '9')) { // Insert '.' before and after a non-digit
				buf[j] = v
				j++
				if v != '.' && next != '.' {
					buf[j] = '.'
					j++
				}
				continue
			} else if !((v >= '0' && v <= '9') ||
				(v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')) { // Non-letters and numbers
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else {
				buf[j] = v
				j++
			}
		}

		return string(buf[:j])
	}

	//compare special version forms
	special = func(form1, form2 string) int {
		found1, found2, len1, len2 := -1, -1, len(form1), len(form2)
		// (Any string not found) < dev < alpha = a < beta = b < RC = rc < # < pl = p
		forms := map[string]int{
			"dev":   0,
			"alpha": 1,
			"a":     1,
			"beta":  2,
			"b":     2,
			"RC":    3,
			"rc":    3,
			"#":     4,
			"pl":    5,
			"p":     5,
		}

		for name, order := range forms {
			if len1 < len(name) {
				continue
			}
			if strings.Compare(form1[:len(name)], name) == 0 {
				found1 = order
				break
			}
		}
		for name, order := range forms {
			if len2 < len(name) {
				continue
			}
			if strings.Compare(form2[:len(name)], name) == 0 {
				found2 = order
				break
			}
		}

		if found1 == found2 {
			return 0
		} else if found1 > found2 {
			return 1
		} else {
			return -1
		}
	}

	compare := vcompare(version1, version2)

	switch operator {
	case "<", "lt":
		return compare == -1
	case "<=", "le":
		return compare != 1
	case ">", "gt":
		return compare == 1
	case ">=", "ge":
		return compare != -1
	case "==", "=", "eq":
		return compare == 0
	case "!=", "<>", "ne":
		return compare != 0
	default:
		panic("operator: invalid")
	}
}

//输出执行的一些信息
func EchoInfo(ok bool, str string, args ...interface{}) {
	stat := "[error]"
	if ok {
		stat = "[ok]"
	}
	line := time.Now().Format("2006-01-02 15:04:05") + "\t" + str + "\t"
	for _, arg := range args {
		switch arg.(type) {
		case int:
			line += Int2str(arg.(int)) + "\t"
		case string:
			line += arg.(string) + "\t"
		}
	}
	line += stat
	fmt.Println(line)
}

//生成区间随机数
func RandRange(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return mrand.Int63n(max-min) + min
}

func CheckImgSuffix(fileSuffix string) bool {
	if fileSuffix != ".png" && fileSuffix != ".jpg" && fileSuffix != ".jpeg" && fileSuffix != ".bmp" && fileSuffix != ".gif" {
		return false
	}
	return true
}
