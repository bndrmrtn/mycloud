package config

import "github.com/bndrmrtn/my-cloud/utils"

func Containers() (int64, int64) {
	var (
		fileSizeMBPerContainer = 1024
		fileLimitPerContainer  = 100
	)

	return int64(fileSizeMBPerContainer * utils.MB), int64(fileLimitPerContainer)
}
