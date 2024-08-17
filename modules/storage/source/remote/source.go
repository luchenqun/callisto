package remote

import (
	types "github.com/forbole/bdjuno/v4/modules/storage/types"
	"github.com/forbole/juno/v5/node/remote"

	storagesource "github.com/forbole/bdjuno/v4/modules/storage/source"
)

var (
	_ storagesource.Source = &Source{}
)

// Source implements storage.Source using a remote node
type Source struct {
	*remote.Source
	storageClient types.QueryClient
}

// NewSource returns a new Source implementation
func NewSource(source *remote.Source, storageClient types.QueryClient) *Source {
	return &Source{
		Source:        source,
		storageClient: storageClient,
	}
}

// HeadBucket implements storage source.Source
func (s Source) HeadBucket(height int64, bucketName string) (types.BucketInfo, types.BucketExtraInfo, error) {
	res, err := s.storageClient.HeadBucket(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadBucketRequest{BucketName: bucketName},
	)
	if err != nil {
		return types.BucketInfo{}, types.BucketExtraInfo{}, err
	}

	return *res.BucketInfo, *res.ExtraInfo, err
}
