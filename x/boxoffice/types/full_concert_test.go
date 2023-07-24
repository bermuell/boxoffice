package types_test

import (
	"github.com/bermuell/boxoffice/x/boxoffice/testutil"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
)

const (
	bob = testutil.Bob
)

func GetStoredConcert1() types.StoredConcert {
	return types.StoredConcert{
		Index:  "1",
		Name:   "Awesome convert of X",
		Volume: 100,
	}
}
