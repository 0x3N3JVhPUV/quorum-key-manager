package models

import (
	"time"

	"github.com/consensys/quorum-key-manager/src/stores/store/entities"
)

type Key struct {
	ID               string
	PublicKey        []byte
	SigningAlgorithm string
	EllipticCurve    string
	Tags             map[string]string
	Annotations      map[string]string
	Disabled         bool
	CreatedAt        time.Time `pg:"default:now()"`
	UpdatedAt        time.Time `pg:"default:now()"`
	DeletedAt        time.Time `pg:",soft_delete"`
}

func (k *Key) ToEntity() *entities.Key {
	return &entities.Key{
		ID:        k.ID,
		PublicKey: k.PublicKey,
		Algo: &entities.Algorithm{
			Type:          entities.KeyType(k.SigningAlgorithm),
			EllipticCurve: entities.Curve(k.EllipticCurve),
		},
		Tags:        k.Tags,
		Annotations: k.Annotations,
		Metadata: &entities.Metadata{
			Disabled:  k.Disabled,
			CreatedAt: k.CreatedAt,
			UpdatedAt: k.UpdatedAt,
			DeletedAt: k.DeletedAt,
		},
	}
}
