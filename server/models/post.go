package models

import (
	"github.com/go-bongo/bongo"
)

type (
	Post struct {
		bongo.DocumentBase `bson:",inline"`
	}

	ActivityPubNote struct {
	}
)
