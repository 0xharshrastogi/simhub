package resourcecontroller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/lib/go/common"
	"github.com/harshrastogiexe/lib/go/proxy"
)

func getRealmFromId(ctx *gin.Context) (common.RealmType, error) {
	id, err := strconv.Atoi(ctx.Param("realm"))
	if err != nil {
		return 0, err
	}
	return common.RealmType(id), nil
}

func loadRealm(t common.RealmType) (*common.Realm, error) {
	switch t {
	case common.Magnates:
		return &proxy.Magnates, nil
	case common.Entrepreneurs:
		return nil, fmt.Errorf("entrepreneurs realm didn't developed")
	default:
		return nil, fmt.Errorf("invalid realm type %d", t)
	}
}

func GetAllResources(ctx *gin.Context) {
	rType, err := getRealmFromId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = loadRealm(rType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	file, err := os.OpenFile(path.Join("data", fmt.Sprintf("resources_%d.json", rType)), os.O_RDONLY, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()
	io.Copy(ctx.Writer, file)
}
