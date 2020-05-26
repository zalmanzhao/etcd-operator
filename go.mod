module github.com/coreos/etcd-operator

go 1.13

require (
	cloud.google.com/go/storage v1.7.0
	contrib.go.opencensus.io/exporter/ocagent v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go v11.3.0-beta+incompatible
	github.com/Azure/go-autorest v11.9.0+incompatible
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190125095113-2b29687e15f2
	github.com/aws/aws-sdk-go v1.13.8
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/cockroachdb/cmux v0.0.0-20170110192607-30d10be49292 // indirect
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.2.13+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/go-ini/ini v1.42.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20160202185014-0b12d6b521d8 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/pborman/uuid v0.0.0-20170612153648-e790cca94e6c
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v0.9.4
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/satori/uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/zalmanzhao/huawei-obs-go-sdk v0.0.0-20200325062301-899cc93259fa
	go.etcd.io/bbolt v1.3.4 // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/api v0.24.0
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	k8s.io/api v0.0.0-20190226173710-145d52631d00
	k8s.io/apiextensions-apiserver v0.0.0-20190226180157-bd0469a053ff
	k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc
	k8s.io/client-go v0.0.0-20190226174127-78295b709ec6
	k8s.io/kube-openapi v0.0.0-20190225204428-d50a959ae76a // indirect
)

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
