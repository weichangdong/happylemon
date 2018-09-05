package upload

import (
	"happylemon/conf"
	"os/exec"

	"net/http"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadTos3(localFile string, s3File string, s3Url string) (bool, string, interface{}) {
	var info = "ok"
	// ======upload img to s3 start======
	cmd := exec.Command("/usr/local/bin/aws", "s3", "cp", localFile, s3File)
	stdout, stderr := cmd.Output()
	if stderr != nil {
		info = string(stdout) + " upload2s3-error"
		return false, info, stderr
	}
	cmd.Run()

	checkRes, err := http.Head(s3Url)
	if 200 != checkRes.StatusCode {
		return false, "check-error", err
	}
	// ======upload img to s3 end======

	return true, info, nil
}
func InitOss() *oss.Bucket {
	endpoint := conf.Config.Oss.Endpoint
	mybucket := conf.Config.Oss.OssBucket
	accessKeyId := conf.Config.Ots.AccessKeyId
	accessKeySecret := conf.Config.Ots.AccessKeySecret
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(mybucket)
	if err != nil {
		panic(err)
	}
	return bucket
}

// 文件传到阿里云的oss
func UploadToOss(fileName string, localFileName string) (bool, string) {
	bucket := InitOss()
	err := bucket.PutObjectFromFile(fileName, localFileName)
	if err != nil {
		return false, " upload oss error-1 " + err.Error()
	}
	return true, ""
}
