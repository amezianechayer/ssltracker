package data

import "time"

type DomainTracking struct {
	ID         int
	DomainName string
	Expires    time.Time
	Issuer     string
	Statut     string
}
