module github.com/qredo/fusionchain/bridge

go 1.21

replace 

replace (
    // assets dependency replacements
	github.com/ChainSafe/go-schnorrkel => github.com/ChainSafe/go-schnorrkel v0.0.0-20210222182958-bd440c890782
	github.com/btcsuite/btcd => github.com/qredo/btcd v0.21.2
	github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.0 => github.com/qredo/go-substrate-rpc-client/v4 v4.1.11
	github.com/cosmos/gogoproto => github.com/cosmos/gogoproto v1.4.10
	github.com/tclairet/cardano-go-v1 => github.com/tclairet/cardano-go v0.0.0-20220314142433-a8bfb2ff6caf
	github.com/vedhavyas/go-subkey => github.com/vedhavyas/go-subkey v1.0.2
	golang.org/x/exp => golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc
	// keyring package
	github.com/qredo/fusionchain/keyring => ../keyring
)


require (
	github.com/qredo/fusionchain/keyring v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v3 v3.0.1
	github.com/qredo/assets v1.14.1 // GOPRIVATE='github.com/qredo'
)