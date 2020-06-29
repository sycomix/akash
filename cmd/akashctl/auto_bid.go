package main

import (
	"context"
	"fmt"
	"os"

	ccontext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ovrclk/akash/cmd/common"
	"github.com/ovrclk/akash/events"
	"github.com/ovrclk/akash/pubsub"
	mquery "github.com/ovrclk/akash/x/market/query"
	mtypes "github.com/ovrclk/akash/x/market/types"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// OrderOut is the struct format of order event
type OrderOut struct {
	ID mtypes.OrderID `json:ID`
}

// autoBidCmd prints out events in real time
func autoBidCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auto-bid",
		Short: "Automatically creates bid on order creation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return common.RunForever(func(ctx context.Context) error {
				return getEvents(ctx, cdc, cmd, args)
			})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

func getEvents(ctx context.Context, cdc *codec.Codec, _ *cobra.Command, _ []string) error {
	cctx := ccontext.NewCLIContext().WithCodec(cdc)

	if err := cctx.Client.Start(); err != nil {
		return err
	}

	bus := pubsub.NewBus()
	defer bus.Close()

	group, ctx := errgroup.WithContext(ctx)

	subscriber, err := bus.Subscribe()

	if err != nil {
		return err
	}

	group.Go(func() error {
		return events.Publish(ctx, cctx.Client, "akash-cli", bus)
	})

	group.Go(func() error {
		for {
			select {
			case <-subscriber.Done():
				return nil
			case ev := <-subscriber.Events():
				switch evt := ev.(type) {
				case mtypes.EventOrderCreated:
					_ = createBid(cctx, cdc, evt.ID)
				}
			}
		}
	})

	return group.Wait()
}

func createBid(ctx ccontext.CLIContext, cdc *codec.Codec, orderID mtypes.OrderID) error {
	order, err := mquery.NewClient(ctx, mtypes.StoreKey).Order(orderID)
	if err != nil {
		return err
	}

	if order.State != mtypes.OrderOpen {
		return nil
	}

	coins := sdk.NewCoin(order.Spec.Price().Denom, sdk.NewInt(0))

	bldr := auth.NewTxBuilderFromCLI(os.Stdin).WithTxEncoder(utils.GetTxEncoder(cdc))
	msg := mtypes.MsgCreateBid{
		Order:    orderID,
		Provider: ctx.GetFromAddress(),
		Price:    coins,
	}

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	fmt.Printf("Msg...%v\n", msg)
	return utils.GenerateOrBroadcastMsgs(ctx, bldr, []sdk.Msg{msg})
}
