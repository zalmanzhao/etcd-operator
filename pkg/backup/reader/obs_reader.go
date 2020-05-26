// Copyright 2019 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reader

import (
	"fmt"
	obs "github.com/zalmanzhao/huawei-obs-go-sdk"
	"io"

	"github.com/coreos/etcd-operator/pkg/backup/util"
)

// ensure ossReader satisfies reader interface.
var _ Reader = &obsReader{}

// ossReader provides Reader imlementation for reading a file from S3
type obsReader struct {
	client *obs.ObsClient
}

// NewOSSReader return a Reader implementation to read a file from OSS in the form of ossReader
func NewOBSReader(client *obs.ObsClient) Reader {
	return &obsReader{client: client}
}

// Open opens the file on path where path must be in the format "<oss-bucket-name>/<key>"
func (obsr *obsReader) Open(path string) (io.ReadCloser, error) {
	bk, key, err := util.ParseBucketAndKey(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse oss bucket and key: %v", err)
	}

	_, err = obsr.client.HeadBucket(bk)
	if err != nil {
		return nil,  fmt.Errorf("OSS: bucket<%s> not found", bk)
	}

	input := &obs.GetObjectInput{}
	input.Bucket = bk
	input.Key = key
	resp, err := obsr.client.GetObject(input)
	if err != nil {
		return nil, err
	}


	return resp.Body, nil
}
