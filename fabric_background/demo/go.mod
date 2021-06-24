module demo

go 1.13

replace git.huawei.com/poissonsearch/wienerchain/proto => ../proto/go

replace git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk => ../wienerchain-go-sdk

replace gmssl => ../thirdparty/GmSSL/gmssl

require (
	git.huawei.com/poissonsearch/wienerchain/proto v0.0.0
	git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk v0.0.0
	github.com/golang/protobuf v1.4.2
	github.com/spf13/cobra v0.0.5
)
