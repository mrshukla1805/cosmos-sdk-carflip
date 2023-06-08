package keeper

import (
	"encoding/binary"

	"carflip/x/carflip/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetCarCount get the total number of car
func (k Keeper) GetCarCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CarCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCarCount set the total number of car
func (k Keeper) SetCarCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CarCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCar appends a car in the store with a new id and update the count
func (k Keeper) AppendCar(
	ctx sdk.Context,
	car types.Car,
) uint64 {
	// Create the car
	count := k.GetCarCount(ctx)

	// Set the ID of the appended value
	car.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CarKey))
	appendedValue := k.cdc.MustMarshal(&car)
	store.Set(GetCarIDBytes(car.Id), appendedValue)

	// Update car count
	k.SetCarCount(ctx, count+1)

	return count
}

// SetCar set a specific car in the store
func (k Keeper) SetCar(ctx sdk.Context, car types.Car) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CarKey))
	b := k.cdc.MustMarshal(&car)
	store.Set(GetCarIDBytes(car.Id), b)
}

// GetCar returns a car from its id
func (k Keeper) GetCar(ctx sdk.Context, id uint64) (val types.Car, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CarKey))
	b := store.Get(GetCarIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCar removes a car from the store
func (k Keeper) RemoveCar(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CarKey))
	store.Delete(GetCarIDBytes(id))
}

// GetAllCar returns all car
func (k Keeper) GetAllCar(ctx sdk.Context) (list []types.Car) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CarKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Car
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCarIDBytes returns the byte representation of the ID
func GetCarIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCarIDFromBytes returns ID in uint64 format from a byte array
func GetCarIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
