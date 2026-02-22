go
package pigeon

import (
    "time"
)

// Pigeon defines the immutable record of a racing bird
type Pigeon struct {
    ID           string    `json:"id"`            // e.g., FNDN-2026-001
    OwnerAddress string    `json:"owner_address"` // Your 0xA62... address
    Strain       string    `json:"strain"`        // Bloodline (e.g., Jan Aarden)
    HatchDate    time.Time `json:"hatch_date"`
    DNAHash      string    `json:"dna_hash"`      // Proof of biological identity
    ParentIDs    []string  `json:"parent_ids"`    // Link for pedigree tracking
}
