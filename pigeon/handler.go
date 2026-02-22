go
package pigeon

import (
    "://github.com"
)

// HandleRegisterBird processes the incoming registration request
func HandleRegisterBird(ctx types.Context, k Keeper, msg MsgRegisterBird) types.Result {
    // 1. Check if the bird ID already exists to prevent fraud
    if k.HasPigeon(ctx, msg.ID) {
        return types.Result{Log: "Pigeon ID already exists"}
    }

    // 2. Create the new Pigeon record
    newPigeon := Pigeon{
        ID:           msg.ID,
        OwnerAddress: msg.Owner,
        Strain:       msg.Strain,
        DNAHash:      msg.DNA,
    }

    // 3. Use the Keeper to save it permanently to the Foundation L1
    k.SetPigeon(ctx, newPigeon)
    
    return types.Result{Log: "Bird successfully registered on-chain"}
}
