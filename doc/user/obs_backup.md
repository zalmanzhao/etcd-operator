# Backups using Huawei Cloud Object Storage Service (OBS)

Etcd backup operator backs up the data of an etcd cluster running on Kubernetes to a remote storage such as Huawei Cloud Object Storage Service (OBS). If it is not deployed yet, please follow the [instructions](walkthrough/backup-operator.md#deploy-etcd-backup-operator) to deploy it, e.g. by running

```sh
kubectl apply -f example/etcd-backup-operator/deployment.yaml
```

## Setup Huawei Cloud backup account, OBS bucket, and Secret

1. Login [Huawei Cloud Console](https://www.huaweicloud.com) and create your own AccessKeyID (AKID) and AccessKeySecret (AKS). You can optionally create an Object Storage Service bucket for backups.
2. Create a secret storing your AKID and AKS in Kubernetes.  

 ```yaml  
apiVersion: v1
  kind: Secret
  metadata:
    name: my-obs-credentials
  type: Opaque
  data:
    accessKeyID: <base64 my-access-key-id>
    accessKeySecret: <base64 my-access-key-secret>
 ```  

3. Create an `EtcdBackup` CR file `etcdbackup.yaml` which uses secret `my-oss-credentials` from the previous step.  
```yaml
apiVersion: etcd.database.coreos.com/v1beta2
kind: EtcdBackup
metadata:
  name: etcd-cluster-with-oss-backup
spec:
  backupPolicy:
    ...
  etcdEndpoints:
    - "http://example-etcd-cluster-client:2379"
  storageType: OBS
  obs:
    endpoint: obs.cn-east-2.myhuaweicloud.com
    ossSecret: my-oss-credentials
    path: my-etcd-backups-bucket/etcd.backup
```   

4. Apply yaml file to kubernetes cluster.  
```sh
kubectl apply -f etcdbackup.yaml
```
5. Check the `status` section of the `EtcdBackup` CR.
```console
$ kubectl get EtcdBackup etcd-cluster-with-oss-backup -o yaml
apiVersion: etcd.database.coreos.com/v1beta2
kind: EtcdBackup
...
spec:
  obs:
    obsSecret: my-obs-credentials
    path: my-etcd-backups-bucket/etcd.backup
    endpoint: obs.cn-east-2.myhuaweicloud.com
  etcdEndpoints:
  - http://example-etcd-cluster-client:2379
  storageType: OBS
status:
  etcdRevision: 1
  etcdVersion: 3.4.3
  succeeded: true
```

6. We should see the backup files from Huawei Cloud OBS Console.


## Restore etcd based on data from OBS.

Etcd restore operator is in charge of restoring etcd cluster from backup. If it is not deployed, please deploy by following command:

```sh
kubectl apply -f example/etcd-restore-operator/deployment.yaml
```

Now kill all the etcd pods to simulate a cluster failure:

```sh
kubectl delete pod -l app=etcd,etcd_cluster=example-etcd-cluster --force --grace-period=0
```

1. Create an EtcdRestore CR.
```yaml
apiVersion: "etcd.database.coreos.com/v1beta2"
kind: "EtcdRestore"
metadata:
  # The restore CR name must be the same as spec.etcdCluster.name
  name: example-etcd-cluster
spec:
  etcdCluster:
    # The namespace is the same as this EtcdRestore CR
    name: example-etcd-cluster
  backupStorageType: OBS
  oss:
    # The format of the path must be: "<oss-bucket-name>/<path-to-backup-file>"
    path: my-etcd-backups-bucket/etcd.backup
    ossSecret: my-obs-credentials
    endpoint: obs.cn-east-2.myhuaweicloud.com
```

2. Check the `status` section of the `EtcdRestore` CR.     
```sh
$ kubectl get etcdrestore example-etcd-cluster -o yaml
apiVersion: etcd.database.coreos.com/v1beta2
kind: EtcdRestore
...
spec:
  obs:
    obsSecret: my-obs-credentials
    path: my-etcd-backups-bucket/etcd.backup
    endpoint: obs.cn-east-2.myhuaweicloud.com
  backupStorageType: OBS
  etcdCluster:
    name: example-etcd-cluster
status:
  succeeded: true
```

3. Verify the `EtcdCluster` CR and restored pods for the restored cluster.    
```sh  
$ kubectl get etcdcluster
$ kubectl get pods -l app=etcd,etcd_cluster=example-etcd-cluster
```
