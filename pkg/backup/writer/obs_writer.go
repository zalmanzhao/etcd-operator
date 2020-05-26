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

package writer

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/coreos/etcd-operator/pkg/backup/util"

	obs "github.com/zalmanzhao/huawei-obs-go-sdk"
)

type obsWriter struct {
	obs *obs.ObsClient
}

// NewOBSWriter creates a obs writer.
func NewOBSWriter(obs *obs.ObsClient) Writer {
	return &obsWriter{obs: obs}
}

// Write writes the backup file to the given obs path, "<obs-bucket-name>/<key>".
func (obsw *obsWriter) Write(ctx context.Context, path string, r io.Reader) (int64, error) {
	// TODO: support context.
	bk, key, err := util.ParseBucketAndKey(path)
	if err != nil {
		return 0, err
	}

	// If bucket doesn't exist, we create it.
	_, err = obsw.obs.GetBucketStorageInfo(bk)
	if err != nil {
		input := &obs.CreateBucketInput{}
		input.Bucket = bk
		if _, err = obsw.obs.CreateBucket(input); err != nil {
			return 0, fmt.Errorf("failed to create bucket, error: %v", err)
		}
	}


	input := &obs.PutObjectInput{}
	input.Bucket = bk
	input.Key = key
	input.Body = r
	if _, err := obsw.obs.PutObject(input); err != nil {
		return 0, err
	}

	i := &obs.GetObjectInput{}
	i.Bucket = bk
	i.Key = key
	rc, err := obsw.obs.GetObject(i)
	if err != nil {
		return 0, fmt.Errorf("failed to get obs object: %v", err)
	}

	return  rc.ContentLength, nil
}

func (obsw *obsWriter) List(ctx context.Context, basePath string) ([]string, error) {
	// TODO: support context.
	bk, key, err := util.ParseBucketAndKey(basePath)
	if err != nil {
		return nil, err
	}

	i := &obs.ListObjectsInput{}
	i.Marker = ""
	i.Prefix = key
	i.MaxKeys = 1000

	var objKeys []string
	for {

		resp, err := obsw.obs.ListObjects(i)
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %v", err)
		}

		for _, obj := range resp.Contents {
			objKeys = append(objKeys, path.Join(bk, obj.Key))
		}

		i.Prefix = resp.Prefix
		i.Marker = resp.NextMarker

		if !resp.IsTruncated {
			break
		}
	}

	return objKeys, nil
}

func (obsw *obsWriter) Delete(ctx context.Context, path string) error {
	// TODO: support context.
	bk, key, err := util.ParseBucketAndKey(path)
	if err != nil {
		return err
	}

	i := &obs.DeleteObjectInput{}
	i.Bucket = bk
	i.Key = key
	if _, err = obsw.obs.DeleteObject(i); err != nil {
		return err
	}
	return nil
}
