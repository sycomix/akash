package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	tmtypes "github.com/tendermint/tendermint/types"
)

func lcdCommand() *cobra.Command {

	cdc := amino.NewCodec()
	tmtypes.RegisterBlockAmino(cdc)

	return lcd.ServeCommand(cdc, lcdRegisterRoutes)
}

func lcdRegisterRoutes(rs *lcd.RestServer) {
	client.RegisterRoutes(rs.CliCtx, rs.Mux)
	authrest.RegisterTxRoutes(rs.CliCtx, rs.Mux)
}
