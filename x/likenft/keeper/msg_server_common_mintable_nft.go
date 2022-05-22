package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/likecoin/likechain/backport/cosmos-sdk/v0.46.0-alpha2/x/nft"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k msgServer) validateReqToMutateMintableNFT(ctx sdk.Context, creator string, class nft.Class, classData types.ClassData, parent types.ClassParentWithOwner, willCreate bool) error {

	// Verify no tokens minted under class
	totalSupply := k.nftKeeper.GetTotalSupply(ctx, class.Id)
	if totalSupply > 0 {
		return types.ErrCannotUpdateClassWithMintedTokens.Wrap("Cannot update class with minted tokens")
	}

	// Check max supply vs existing mintable count
	if willCreate && classData.Config.MaxSupply > 0 && classData.MintableCount >= classData.Config.MaxSupply {
		return types.ErrNftNoSupply.Wrapf("NFT Class has reached its maximum supply: %d", classData.Config.MaxSupply)
	}

	// Check class parent relation is valid and current user is owner
	userAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("%s", err.Error())
	}
	if !parent.Owner.Equals(userAddress) {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not authorized", userAddress.String())
	}

	return nil
}
