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

package controller

import (
	"context"
	"fmt"
	"github.com/coreos/etcd-operator/pkg/util/huaweicloudutil/obsfactory"

	api "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	"github.com/coreos/etcd-operator/pkg/backup"
	"github.com/coreos/etcd-operator/pkg/backup/writer"
	"k8s.io/client-go/kubernetes"
)

// handleOBS saves etcd cluster's backup to specificed OBS path.
func handleOBS(ctx context.Context, kubecli kubernetes.Interface, s *api.OBSBackupSource, endpoints []string, clientTLSSecret,
	namespace string, isPeriodic bool, maxBackup int) (*api.BackupStatus, error) {
	if s.Endpoint == "" {
		s.Endpoint = "obs.cn-east-2.myhuaweicloud.com"
	}
	// TODO: controls NewClientFromSecret with ctx. This depends on upstream kubernetes to support API calls with ctx.
	cli, err := obsfactory.NewClientFromSecret(kubecli, namespace, s.Endpoint, s.OBSSecret)
	if err != nil {
		return nil, err
	}

	tlsConfig, err := generateTLSConfig(kubecli, clientTLSSecret, namespace)
	if err != nil {
		return nil, err
	}

	bm := backup.NewBackupManagerFromWriter(kubecli, writer.NewOBSWriter(cli.OBS), tlsConfig, endpoints, namespace)

	rev, etcdVersion, now, err := bm.SaveSnap(ctx, s.Path, isPeriodic)
	if err != nil {
		return nil, fmt.Errorf("failed to save snapshot (%v)", err)
	}

	if maxBackup > 0 {
		err := bm.EnsureMaxBackup(ctx, s.Path, maxBackup)
		if err != nil {
			return nil, fmt.Errorf("succeeded in saving snapshot but failed to delete old snapshot (%v)", err)
		}
	}

	return &api.BackupStatus{EtcdVersion: etcdVersion, EtcdRevision: rev, LastSuccessDate: *now}, nil
}
