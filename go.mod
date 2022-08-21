module github.com/kanengo/go-actor

go 1.18

replace github.com/kanengo/goutil => ../goutil

require (
	github.com/kanengo/goutil v0.0.0
	github.com/orcaman/concurrent-map/v2 v2.0.0
	google.golang.org/protobuf v1.28.1
)

require github.com/spaolacci/murmur3 v1.1.0

require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
