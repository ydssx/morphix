module github.com/ydssx/morphix

go 1.20

require (
	github.com/Gre-Z/common v0.0.0-20191024025434-2dbc6bd196f9
	github.com/IBM/sarama v1.42.1
	github.com/axiaoxin-com/goutils v1.0.39
	github.com/bsm/redislock v0.9.4
	github.com/bufbuild/connect-go v1.10.0
	github.com/casbin/casbin/v2 v2.79.0
	github.com/casbin/gorm-adapter/v3 v3.20.0
	github.com/cheggaaa/pb/v3 v3.1.4
	github.com/chenjiandongx/ginprom v0.0.0-20210617023641-6c809602c38a
	github.com/cloudevents/sdk-go/protocol/nats/v2 v2.14.0
	github.com/cloudevents/sdk-go/v2 v2.14.0
	github.com/dapr/go-sdk v1.9.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dtm-labs/driver-kratos v0.0.9
	github.com/emicklei/proto v1.12.2
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/fatih/color v1.16.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20231201055217-d13071e45f8d
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20231201055217-d13071e45f8d
	github.com/go-kratos/kratos/v2 v2.7.1
	github.com/go-kratos/swagger-api v1.0.1
	github.com/go-playground/validator/v10 v10.16.0
	github.com/go-redis/redis_rate/v10 v10.0.1
	github.com/go-resty/resty/v2 v2.10.0
	github.com/google/uuid v1.4.0
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus v1.0.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1
	github.com/hibiken/asynq v0.24.1
	github.com/nats-io/nats.go v1.31.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.17.0
	github.com/redis/go-redis/v9 v9.3.0
	github.com/ydssx/api-gen v0.0.0-20231014024148-dcbd019c3dcf
	go.etcd.io/etcd/client/v3 v3.5.10
	go.mongodb.org/mongo-driver v1.13.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.1
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.44.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/sdk/metric v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/automaxprocs v1.5.3
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.26.0
	golang.org/x/sync v0.5.0
	google.golang.org/genproto/googleapis/api v0.0.0-20231127180814-3a041ad873d4
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/antlabs/strsim v0.0.2 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/casbin/govaluate v1.1.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/dapr/dapr v1.12.0-rc.4 // indirect
	github.com/dave/dst v0.27.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dtm-labs/dtmdriver v0.0.6 // indirect
	github.com/dtm-labs/dtmdriver-kratos v0.0.9 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.4.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.20.3 // indirect
	github.com/glebarez/sqlite v1.7.0 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-kratos/grpc-gateway/v2 v2.5.1-0.20210811062259-c92d36e434b1 // indirect
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220414054820-d0b704b8f38d // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/consul/api v1.24.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/microsoft/go-mssqldb v1.5.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.7.0 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/speps/go-hashids v2.0.0+incompatible // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.11.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.etcd.io/etcd/api/v3 v3.5.10 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.10 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.15.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/postgres v1.4.4 // indirect
	gorm.io/driver/sqlserver v1.4.1 // indirect
	gorm.io/plugin/dbresolver v1.3.0 // indirect
	modernc.org/libc v1.24.1 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.6.0 // indirect
	modernc.org/sqlite v1.25.0 // indirect
)
