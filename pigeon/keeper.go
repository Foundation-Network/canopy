go
package pigeon

import (
    "://github.com"
    "://github.com"
)

// Keeper handles reading and writing pigeon data to the blockchain
type Keeper struct {
    storeKey types.StoreKey
}

func NewKeeper(key types.StoreKey) Keeper {
    return Keeper{storeKey: key}
}

// SetPigeon saves a bird's record permanently to the Foundation L1
func (k Keeper) SetPigeon(ctx types.Context, pigeon Pigeon) {
    store := ctx.KVStore(k.storeKey)
    bz := types.MustMarshal(&pigeon)
    store.Set([]byte(pigeon.ID), bz)
}

// GetPigeon retrieves a bird's pedigree from the chain
func (k Keeper) GetPigeon(ctx types.Context, id string) (Pigeon, bool) {
    store := ctx.KVStore(k.storeKey)
    bz := store.Get([]byte(id))
    if bz == nil {
        return Pigeon{}, false
    }
    var pigeon Pigeon
    types.MustUnmarshal(bz, &pigeon)
    return pigeon, true
}
