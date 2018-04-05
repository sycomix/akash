package main

import (
	"fmt"

	"github.com/ovrclk/akash/cmd/akash/context"
	"github.com/ovrclk/akash/cmd/common"
	"github.com/ovrclk/akash/marketplace"
	"github.com/ovrclk/akash/state"
	"github.com/ovrclk/akash/types"
	"github.com/spf13/cobra"

	. "github.com/ovrclk/akash/util"
)

func marketplaceCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "marketplace",
		Short: "monitor marketplace",
		Args:  cobra.NoArgs,
		RunE:  context.WithContext(context.RequireNode(doMarketplaceMonitorCommand)),
	}

	context.AddFlagNode(cmd, cmd.PersistentFlags())

	return cmd
}

func doMarketplaceMonitorCommand(ctx context.Context, cmd *cobra.Command, args []string) error {
	handler := marketplaceMonitorHandler()
	return common.MonitorMarketplace(ctx.Log(), ctx.Client(), handler)
}

func marketplaceMonitorHandler() marketplace.Handler {
	return marketplace.NewBuilder().
		OnTxSend(func(tx *types.TxSend) {
			fmt.Printf("TRANSFER\t%v tokens from %v to %v\n", tx.GetAmount(), X(tx.From), X(tx.To))
		}).
		OnTxCreateProvider(func(tx *types.TxCreateProvider) {
			fmt.Printf("DATACENTER CREATED\t%v created by %v\n", X(state.ProviderAddress(tx.Owner, tx.Nonce)), X(tx.Owner))
		}).
		OnTxCreateDeployment(func(tx *types.TxCreateDeployment) {
			fmt.Printf("DEPLOYMENT CREATED\t%v created by %v\n", X(state.DeploymentAddress(tx.Tenant, tx.Nonce)), X(tx.Tenant))
		}).
		OnTxCreateOrder(func(tx *types.TxCreateOrder) {
			fmt.Printf("ORDER CREATED\t%v/%v/%v\n",
				X(tx.Deployment), tx.Group, tx.Seq)
		}).
		OnTxCreateFulfillment(func(tx *types.TxCreateFulfillment) {
			fmt.Printf("FULFILLMENT CREATED\t%v/%v/%v by %v [price=%v]\n",
				X(tx.Deployment), tx.Group, tx.Order,
				X(tx.Provider), tx.Price)
		}).
		OnTxCreateLease(func(tx *types.TxCreateLease) {
			fmt.Printf("LEASE CREATED\t%v/%v/%v by %x [price=%v]\n",
				X(tx.Deployment), tx.Group, tx.Order,
				X(tx.Provider), tx.Price)
		}).
		OnTxCloseDeployment(func(tx *types.TxCloseDeployment) {
			fmt.Printf("DEPLOYMENT CLOSED\t%v\n", X(tx.Deployment))
		}).
		OnTxCloseFulfillment(func(tx *types.TxCloseFulfillment) {
			fmt.Printf("FULFILLMENT CLOSED\t%v\n", X(tx.Fulfillment))
		}).
		Create()
}
