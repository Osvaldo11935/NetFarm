package istorages

import (
	"NetFarm/shared/models/common"
	"github.com/gin-gonic/gin"
)

type IFileManagerStorage interface {
	Upload(c *gin.Context) (string, *common.ErrorResponse)
}
