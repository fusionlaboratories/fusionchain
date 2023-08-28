package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/qredo/fusionchain/repo"
	identitytypes "github.com/qredo/fusionchain/x/identity/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type IdentityKeeper interface {
	GetWorkspace(ctx sdk.Context, addr string) *identitytypes.Workspace
	KeyringsRepo() repo.ObjectRepo[*identitytypes.Keyring]
}
