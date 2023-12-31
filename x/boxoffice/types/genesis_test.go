package types_test

import (
	"testing"

	"github.com/bermuell/boxoffice/x/boxoffice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				SystemInfo: types.SystemInfo{
					NextId: 67,
				},
				StoredConcertList: []types.StoredConcert{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedConcert",
			genState: &types.GenesisState{
				StoredConcertList: []types.StoredConcert{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

// Basic test to check initial SystemInfo
func TestDefaultGenesisState_CheckInitialNextId(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			StoredConcertList: []types.StoredConcert{},
			SystemInfo:        types.SystemInfo{uint64(1)},
		},
		types.DefaultGenesis(),
	)
}
