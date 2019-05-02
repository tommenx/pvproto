#generate new protocol buffer
protoc ./proto/coordinator.proto --go_out=plugins=grpc:pkg
protoc ./proto/executor.proto --go_out=plugins=grpc:pkg
#remove old
rm ./pkg/proto/coordinatorpb/coordinator.pb.go
rm ./pkg/proto/executorpb/executor.pb.go
#mv new
mv ./pkg/proto/coordinator.pb.go ./pkg/proto/coordinatorpb/
mv ./pkg/proto/executor.pb.go ./pkg/proto/executorpb/

# update go vendor
cd $GOPATH/src/github.com/tommenx/coordinator
govendor remove github.com/tommenx/pvproto
rm -rf vendor/github.com/tommenx/pvproto
govendor add +e