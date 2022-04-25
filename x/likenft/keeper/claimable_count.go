package keeper

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k Keeper) incClaimableCount(ctx sdk.Context, classId string) error {
	return k._setClaimableCount(ctx, classId, func(count uint64) uint64 {
		return count + 1
	})
}

func (k Keeper) decClaimableCount(ctx sdk.Context, classId string) error {
	return k._setClaimableCount(ctx, classId, func(count uint64) uint64 {
		return count - 1
	})
}

func (k Keeper) setClaimableCount(ctx sdk.Context, classId string, count uint64) error {
	return k._setClaimableCount(ctx, classId, func(_ uint64) uint64 {
		return count
	})
}

func (k Keeper) _setClaimableCount(ctx sdk.Context, classId string, edit func(uint64) uint64) error {
	class, found := k.nftKeeper.GetClass(ctx, classId)
	if !found {
		return types.ErrNftClassNotFound
	}
	var classData types.ClassData
	if err := k.cdc.Unmarshal(class.Data.Value, &classData); err != nil {
		return types.ErrFailedToUnmarshalData.Wrapf(err.Error())
	}
	classData.ClaimableCount = edit(classData.ClaimableCount)
	classDataInAny, err := cdctypes.NewAnyWithValue(&classData)
	if err != nil {
		return types.ErrFailedToMarshalData.Wrapf("%s", err.Error())
	}
	class.Data = classDataInAny
	if err := k.nftKeeper.UpdateClass(ctx, class); err != nil {
		return types.ErrFailedToUpdateClass.Wrapf("%s", err.Error())
	}
	return nil
}
