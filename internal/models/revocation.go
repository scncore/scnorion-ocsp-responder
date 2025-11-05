package models

import (
	"context"

	scncore_ent "github.com/scncore/ent"
	"github.com/scncore/ent/revocation"
)

func (m *Model) GetRevoked(serial int64) (*scncore_ent.Revocation, error) {
	return m.Client.Revocation.Query().Where(revocation.ID(serial)).Only(context.Background())
}
