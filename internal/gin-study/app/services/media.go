package services

import (
	"errors"
	"github.com/go-gin-template/internal/gin-study/app/common/request"
	"github.com/go-gin-template/internal/gin-study/app/models"
	"github.com/go-gin-template/internal/gin-study/global"
	"path"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type mediaService struct {
}

var MediaService = new(mediaService)

type outPut struct {
	Id   int64  `json:"id"`
	Path string `json:"path"`
	Url  string `json:"url"`
}

const mediaCacheKeyPre = "media:"

// 文件存储目录
func (mediaService *mediaService) makeFaceDir(business string) string {
	return global.App.Config.App.Env + "/" + business
}

// HashName 生成文件名称（使用 uuid）
func (mediaService *mediaService) HashName(fileName string) string {
	fileSuffix := path.Ext(fileName)
	return uuid.NewV4().String() + fileSuffix
}

// SaveFile 保存图片（公共读）
func (mediaService *mediaService) SaveFile(params request.FileUpload) (result outPut, err error) {
	file, err := params.File.Open()
	defer file.Close()
	if err != nil {
		err = errors.New("上传失败")
		return
	}
	key := mediaService.HashName(params.File.Filename)
	disk := global.App.Disk()
	err = disk.Put(key, file, params.File.Size)
	if err != nil {
		return
	}

	image := models.Media{
		DiskType: string(global.App.Config.Storage.Default),
		SrcType:  1,
		Src:      key,
	}
	err = global.App.DB.Create(&image).Error
	if err != nil {
		return
	}

	result = outPut{int64(image.ID.ID), key, disk.Url(key)}
	return
}

// GetUrlById 通过 id 获取文件 url
func (mediaService *mediaService) GetUrlById(id int64) string {
	if id == 0 {
		return ""
	}

	var url string
	cacheKey := mediaCacheKeyPre + strconv.FormatInt(id, 10)

	exist := global.App.Redis.Exists(cacheKey).Val()
	if exist == 1 {
		url = global.App.Redis.Get(cacheKey).Val()
	} else {
		media := models.Media{}
		err := global.App.DB.First(&media, id).Error
		if err != nil {
			return ""
		}
		url = global.App.Disk(media.DiskType).Url(media.Src)
		global.App.Redis.Set(cacheKey, url, time.Second*3*24*3600)
	}

	return url
}
