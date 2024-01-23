package keeper

import (
	"testing"

	cbftdb "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"

	identitytypes "github.com/qredo/fusionchain/x/identity/keeper"
	policytypes "github.com/qredo/fusionchain/x/policy/keeper"
	qassetstypes "github.com/qredo/fusionchain/x/qassets/keeper"
	treasurytypes "github.com/qredo/fusionchain/x/treasury/keeper"
)

type KeeperTest struct {
	Ctx            sdk.Context
	PolicyKeeper   *policytypes.Keeper
	IdentityKeeper *identitytypes.Keeper
	TreasuryKeeper *treasurytypes.Keeper
	QassetsKeeper  *qassetstypes.Keeper
}

func NewTest(t testing.TB) *KeeperTest {
	db := cbftdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	policyKeeper, ctx := PolicyKeeper(t, db, stateStore)
	identityKeeper, _ := IdentityKeeper(t, policyKeeper, db, stateStore)
	treasuryKeeper, _ := TreasuryKeeper(t, policyKeeper, identityKeeper, db, stateStore)
	qassetsKeeper, _ := QassetsKeeper(t, identityKeeper, treasuryKeeper, db, stateStore)

	return &KeeperTest{
		Ctx:            ctx,
		PolicyKeeper:   policyKeeper,
		IdentityKeeper: identityKeeper,
		TreasuryKeeper: treasuryKeeper,
		QassetsKeeper:  qassetsKeeper,
	}
}
