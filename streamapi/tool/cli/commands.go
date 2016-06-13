package cli

import (
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/phillp/motus/streamapi/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// AddOrdinalValuesCommand is the command line data structure for the add action of OrdinalValues
	AddOrdinalValuesCommand struct {
		// The ordinal position of the value
		Ordinal int
		// The stream for which the value is to be added
		Stream string
		// The value to be added to the stream
		Value       float64
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "add",
		Short: `add a value to a stream referencing an ordinal position`,
	}
	tmp1 := new(AddOrdinalValuesCommand)
	sub = &cobra.Command{
		Use:   `OrdinalValues [/api/add/STREAM/ORDINAL/VALUE]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

// Run makes the HTTP request corresponding to the AddOrdinalValuesCommand command.
func (cmd *AddOrdinalValuesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/add/%v/%v/%v", cmd.Ordinal, cmd.Stream, cmd.Value)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddOrdinalValues(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddOrdinalValuesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var ordinal int
	cc.Flags().IntVar(&cmd.Ordinal, "ordinal", ordinal, `The ordinal position of the value`)
	var stream string
	cc.Flags().StringVar(&cmd.Stream, "stream", stream, `The stream for which the value is to be added`)
	var value float64
	cc.Flags().Float64Var(&cmd.Value, "value", value, `The value to be added to the stream`)
}
