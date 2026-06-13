package storage

import "github.com/mrpk1906/lego/v5/certificate"

type Certificate struct {
	*certificate.Resource

	Origin string `json:"origin,omitempty"`
}
