package pricefeed

import (
	"fmt"

	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/forbole/juno/v6/types/config"
	"github.com/spf13/cobra"

	"github.com/forbole/callisto/v4/database"
	"github.com/forbole/callisto/v4/modules/pricefeed"
	"github.com/forbole/callisto/v4/utils"
)

// priceCmd returns the Cobra command allowing to refresh token price
func priceCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "price",
		Short: "Refresh token price",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			cdc := utils.GetCodec()

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build pricefeed module
			pricefeedModule := pricefeed.NewModule(config.Cfg, cdc, db)

			err = pricefeedModule.RunAdditionalOperations()
			if err != nil {
				return fmt.Errorf("error while storing tokens: %s", err)
			}

			err = pricefeedModule.UpdatePrice()
			if err != nil {
				return fmt.Errorf("error while updating price: %s", err)
			}

			return nil
		},
	}
}
