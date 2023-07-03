package model

import (
	"context"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/wejectchen/ginblog/utils"
	"github.com/wejectchen/ginblog/utils/errmsg"
)

var Zone = utils.Zone
var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuSever

// UpLoadFile 上传文件函数
func UpLoadFile(file multipart.File, fileSize int64, fileHeader *multipart.FileHeader) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := setConfig()

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + "/" + ret.Key

	// oss
	// client, err := oss.New(utils.OssEndpoint, utils.OssAccessKeyID, utils.OssAccessKeySecret)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(-1)
	// }
	// bucket, err := client.Bucket(utils.OssBucket)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(-1)
	// }
	// err = bucket.PutObject(fileHeader.Filename, file)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(-1)
	// }
	// url, err := bucket.SignURL(fileHeader.Filename, oss.HTTPGet, 20)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(-1)
	// }
	return url, errmsg.SUCCSE
}

func setConfig() storage.Config {
	cfg := storage.Config{
		Zone:          selectZone(Zone),
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	return cfg
}

func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuabei
	case 3:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}
