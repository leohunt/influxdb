package source

import (
	"fmt"

	platform "github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/http"
	"github.com/influxdata/influxdb/http/influxdb"
)

// NewBucketService creates a bucket service from a source.
func NewBucketService(s *platform.Source) (platform.BucketService, error) {
	switch s.Type {
	case platform.SelfSourceType:
		// TODO(fntlnz): this is supposed to call a bucket service directly locally,
		// we are letting it err for now since we have some refactoring to do on
		// how services are instantiated
		return nil, fmt.Errorf("self source type not implemented")
	case platform.V2SourceType:
		httpClient, err := http.NewHTTPClient(s.URL, s.Token, s.InsecureSkipVerify)
		if err != nil {
			return nil, err
		}
		return &http.BucketService{Client: httpClient}, nil
	case platform.V1SourceType:
		return &influxdb.BucketService{Source: s}, nil
	}
	return nil, fmt.Errorf("unsupported source type %s", s.Type)
}
