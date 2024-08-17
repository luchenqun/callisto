package source

import (
	"github.com/forbole/bdjuno/v4/modules/storage/types"
)

type Source interface {
	HeadBucket(height int64, bucketName string) (types.BucketInfo, types.BucketExtraInfo, error)
}
