module github.com/qredo/fusionchain/fusionmodel

go 1.21

replace github.com/qredo/fusionchain/keyring => ../keyring

require (
	github.com/qredo/fusionchain/keyring v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/vrischmann/envconfig v1.3.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
