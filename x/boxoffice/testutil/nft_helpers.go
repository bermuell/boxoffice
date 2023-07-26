package testutil

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func (nft *MockNftKeeper) ExpectAny(context context.Context) {
	//f-sig: Mint(ctx sdk.Context, token nft.NFT, receiver sdk.AccAddress) error
	nft.EXPECT().Mint(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any()).AnyTimes()

	//f-sig: GetClass(ctx sdk.Context, classID string) (nft.Class, bool)
	nft.EXPECT().GetClass(sdk.UnwrapSDKContext(context), gomock.Any()).AnyTimes()

	//f-sig: SaveClass(ctx sdk.Context, class nft.Class) error
	nft.EXPECT().SaveClass(sdk.UnwrapSDKContext(context), gomock.Any()).AnyTimes()

}
