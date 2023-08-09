package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func SaveVideoInfo(video *entity.Video) bool {

	res := global.Datasource.Table("video").Create(&video)
	if res.Error != nil {
		global.Logger.Error(res.Error)
		return false
	}

	return true

}
