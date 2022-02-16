package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	nft "github.com/likecoin/likechain/backport/cosmos-sdk/v0.46.0-alpha2/x/nft"
	iscntypes "github.com/likecoin/likechain/x/iscn/types"
)

type IscnKeeper interface {
	// Methods imported from iscn should be defined here
	GetContentIdRecord(ctx sdk.Context, iscnIdPrefix iscntypes.IscnIdPrefix) *iscntypes.ContentIdRecord
	GetIscnIdSequence(ctx sdk.Context, iscnId iscntypes.IscnId) uint64
	GetStoreRecord(ctx sdk.Context, seq uint64) *iscntypes.StoreRecord
}

type NftKeeper interface {
	// Methods imported from nft should be defined here
	SaveClass(ctx sdk.Context, class nft.Class) error
	GetClass(ctx sdk.Context, classID string) (nft.Class, bool)
}

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
