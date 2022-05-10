package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k msgServer) getParentOwnerAndValidateReqToMutateMintableNFT(ctx sdk.Context, creator string, classId string, willCreate bool) (*types.ClassParentAndOwner, error) {

	// Verify class exists
	class, found := k.nftKeeper.GetClass(ctx, classId)
	if !found {
		return nil, types.ErrNftClassNotFound.Wrapf("Class id %s not found", classId)
	}

	// Verify no tokens minted under class
	totalSupply := k.nftKeeper.GetTotalSupply(ctx, class.Id)
	if totalSupply > 0 {
		return nil, types.ErrCannotUpdateClassWithMintedTokens.Wrap("Cannot update class with minted tokens")
	}

	var classData types.ClassData
	if err := k.cdc.Unmarshal(class.Data.Value, &classData); err != nil {
		return nil, types.ErrFailedToUnmarshalData.Wrapf(err.Error())
	}

	// Check max supply vs existing mintable count
	if willCreate && classData.Config.MaxSupply > 0 && classData.MintableCount >= classData.Config.MaxSupply {
		return nil, types.ErrNftNoSupply.Wrapf("NFT Class has reached its maximum supply: %d", classData.Config.MaxSupply)
	}

	// Check class parent relation is valid and current user is owner
	parentAndOwner, err := k.validateAndGetClassParentAndOwner(ctx, class.Id, &classData)
	if err != nil {
		return nil, err
	}

	userAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("%s", err.Error())
	}
	if !parentAndOwner.Owner.Equals(userAddress) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not authorized", userAddress.String())
	}

	return parentAndOwner, nil
}
