package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/likecoin/likechain/x/likenft/types"
)

// SetClassRevealQueueEntry set a specific classRevealQueueEntry in the store from its index
func (k Keeper) SetClassRevealQueueEntry(ctx sdk.Context, classRevealQueueEntry types.ClassRevealQueueEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassRevealQueueKeyPrefix))
	b := k.cdc.MustMarshal(&classRevealQueueEntry)
	store.Set(types.ClassRevealQueueKey(
		classRevealQueueEntry.RevealTime,
		classRevealQueueEntry.ClassId,
	), b)
}

// RemoveClassRevealQueueEntry removes a classRevealQueueEntry from the store
func (k Keeper) RemoveClassRevealQueueEntry(
	ctx sdk.Context,
	revealTime time.Time,
	classId string,

) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassRevealQueueKeyPrefix))
	if !store.Has(types.ClassRevealQueueKey(revealTime, classId)) {
		return sdkerrors.ErrKeyNotFound.Wrapf("classRevealQueueEntry not found: classId: %s, revealTime: %s", classId, revealTime)
	}
	store.Delete(types.ClassRevealQueueKey(
		revealTime,
		classId,
	))
	return nil
}

// UpdateClassRevealQueueEntry updates a classRevealQueueEntry in the store
func (k Keeper) UpdateClassRevealQueueEntry(ctx sdk.Context, originalRevealTime time.Time, classId string, updatedRevealTime time.Time) error {
	err := k.RemoveClassRevealQueueEntry(ctx, originalRevealTime, classId)
	if err != nil {
		return err
	}
	k.SetClassRevealQueueEntry(ctx, types.ClassRevealQueueEntry{
		RevealTime: updatedRevealTime,
		ClassId:    classId,
	})
	return nil
}

func (k Keeper) ClassRevealQueueIterator(ctx sdk.Context) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassRevealQueueKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	return iterator
}

// IterateClassRevealQueue iterates over all classRevealQueueEntry
func (k Keeper) IterateClassRevealQueue(ctx sdk.Context, cb func(val types.ClassRevealQueueEntry) (stop bool)) {
	iterator := k.ClassRevealQueueIterator(ctx)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClassRevealQueueEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if cb(val) {
			break
		}
	}

	return
}

// GetClassRevealQueue returns all classRevealQueueEntry
func (k Keeper) GetClassRevealQueue(ctx sdk.Context) (list []types.ClassRevealQueueEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassRevealQueueKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClassRevealQueueEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}