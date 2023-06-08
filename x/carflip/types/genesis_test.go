package types_test

import (
	"testing"

	"carflip/x/carflip/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
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

				CarList: []types.Car{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CarCount: 2,
				RequestList: []types.Request{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				RequestCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated car",
			genState: &types.GenesisState{
				CarList: []types.Car{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid car count",
			genState: &types.GenesisState{
				CarList: []types.Car{
					{
						Id: 1,
					},
				},
				CarCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated request",
			genState: &types.GenesisState{
				RequestList: []types.Request{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid request count",
			genState: &types.GenesisState{
				RequestList: []types.Request{
					{
						Id: 1,
					},
				},
				RequestCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
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
