package conf

import (
	"encoding/hex"

	"github.com/BurntSushi/toml"
)

var Config Conf

type Conf struct {
	Server struct {
		Port        string //http服务端口
		LogPath     string //日志落地目录
		LogLevel    int    //0 关闭日志输出 1 debug 2 info 3 error
		Aeskey      string //aes加密秘钥
		AeskeyBytes []byte //同上，二进制状态，初始访问会对其赋值，进行Cache
		WhiteIp     string //input接口IP白名单
		Mode        string
		Ginmode     string
		RootPath    string
	}
	Redis struct {
		Host string //地址
		Port string //端口
		Auth string //口令
		Db   int    //select db
	}
	Mysql struct {
		Host         string //地址
		Port         string //端口
		UserName     string //用户名
		Password     string //密码
		Database     string //数据库
		MaxOpenConns int    //最大连接数
		MaxIdleConns int    //最少启动连接数
	}
	RedisKey struct {
		UserPrefix   string
		TokenPrefix  string
		TokenUidKey  string
		CronListKey  string
		QueueListKey string
	}

	Grpc struct {
		ApiPort    string
		ApiHost    string
		GrpcSwitch bool
	}
	//发送短信配置
	Sms struct {
		AccessKeyId     string
		AccessKeySecret string
		Signname        string
	}
	Upload struct {
		QrLogoPath      string // 二维码本地logo地址
		SaveDir         string
		AvatarSaveDir   string
		GroupSaveDir    string
		OssBucket       string
		QrImgSaveDir    string
		IdCardSaveDir   string
		FeedsSaveDir    string
		FeedsGifSaveDir string
	}
	Oss struct {
		Endpoint  string
		OssBucket string
		OssCdn    string
	}
	IpWhiteList struct {
		Ips []string
	}
}

func ReadCfg(path string) Conf {
	if Config.Mysql.Host != "" {
		return Config
	}
	if _, err := toml.DecodeFile(path, &Config); err != nil {
		panic("Parse File Error:" + err.Error())
	}
	if aeskey, err := hex.DecodeString(Config.Server.Aeskey); err != nil {
		panic("aes key error:" + err.Error())
	} else {
		Config.Server.AeskeyBytes = aeskey
	}
	return Config
}
