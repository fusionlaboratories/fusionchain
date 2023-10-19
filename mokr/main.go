// Command mokr is a mocked keyring for Fusion Chain.
//
// It is used during development to test the application,
// it generates ECDSA secp256k1 keys and signatures,
// using an in-memory database.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/qredo/fusionchain/go-client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// chain configuration
var (
	fusionChainGRPCAddr = envOrDefault("FUSION_URL", "localhost:9090")
	chainID             = envOrDefault("CHAIN_ID", "qredofusiontestnet_257-1")
	derivationPath      = envOrDefault("DERIVATION_PATH", "m/44'/60'/0'/0/0")
)

// identity configuration
var (
	keyringAddr = envOrDefault("KEYRING_ADDR", "qredokeyring1ph63us46lyw56vrzgaq")
	seedPhrase  = envOrDefault("MNEMONIC", "exclude try nephew main caught favorite tone degree lottery device tissue tent ugly mouse pelican gasp lava flush pen river noise remind balcony emerge")
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
	fmt.Println(keyringAddr)

	/* keyringAddr, err := strconv.ParseUint(keyringAddr, 10, 64)
	if err != nil {
		panic(err)
	}
	*/

	identity, err := client.NewIdentityFromSeed(derivationPath, seedPhrase)
	if err != nil {
		panic(err)
	}

	keyringIdentity := KeyringIdentity{
		Identity:    identity,
		KeyringAddr: keyringAddr,
	}
	logger.Info("starting_mpc_node", "keyring_addr", keyringIdentity.KeyringAddr, "address", identity.Address.String())

	fusionConn := MustConnectFusionChain()

	queryClient := client.NewQueryClientWithConn(fusionConn)
	txClient := client.NewTxClient(identity, chainID, fusionConn, queryClient)

	keyDB := NewMemoryDB()
	keyRequestsHandler := &MockKeyRequestsHandler{
		KeyDB:       keyDB,
		QueryClient: queryClient,
		TxClient:    txClient,
		Logger:      logger.With("module", "key_requests_handler"),
	}
	signatureRequestsHandler := &MockSignatureRequestsHandler{
		KeyDB:       keyDB,
		QueryClient: queryClient,
		TxClient:    txClient,
		Logger:      logger.With("module", "signature_requests_handler"),
	}

	engine := &Engine{
		QueryClient:              queryClient,
		KeyringAddr:              keyringAddr,
		KeyRequestsHandler:       keyRequestsHandler,
		SignatureRequestsHandler: signatureRequestsHandler,
	}

	ctx := context.Background()
	engine.Loop(ctx)
}

func MustConnectFusionChain() *grpc.ClientConn {
	grpcConn, err := grpc.Dial(
		fusionChainGRPCAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return grpcConn
}

func envOrDefault(key, defaultValue string) string {
	if v, found := os.LookupEnv(key); found {
		return v
	}
	return defaultValue
}
