package cli

import (
	"encoding/json"
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

	// PushOrdinalValuesCommand is the command line data structure for the push action of OrdinalValues
	PushOrdinalValuesCommand struct {
		Payload     string
		PrettyPrint bool
	}

	// RegisterOrdinalValuesCommand is the command line data structure for the register action of OrdinalValues
	RegisterOrdinalValuesCommand struct {
		Payload     string
		PrettyPrint bool
	}

	// StatisticsOrdinalValuesCommand is the command line data structure for the statistics action of OrdinalValues
	StatisticsOrdinalValuesCommand struct {
		Payload     string
		PrettyPrint bool
	}

	// TagOrdinalValuesCommand is the command line data structure for the tag action of OrdinalValues
	TagOrdinalValuesCommand struct {
		Payload     string
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
	command = &cobra.Command{
		Use:   "push",
		Short: `Pushes a new ordinal value onto the stream`,
	}
	tmp2 := new(PushOrdinalValuesCommand)
	sub = &cobra.Command{
		Use:   `OrdinalValues [/api/push]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "register",
		Short: `Registers a new stream`,
	}
	tmp3 := new(RegisterOrdinalValuesCommand)
	sub = &cobra.Command{
		Use:   `OrdinalValues [/api/register]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "statistics",
		Short: `Gets statistics matching search criteria`,
	}
	tmp4 := new(StatisticsOrdinalValuesCommand)
	sub = &cobra.Command{
		Use:   `OrdinalValues [/api/statistics]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "tag",
		Short: `Changes the tag assignments for a stream`,
	}
	tmp5 := new(TagOrdinalValuesCommand)
	sub = &cobra.Command{
		Use:   `OrdinalValues [/api/tag]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
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

// Run makes the HTTP request corresponding to the PushOrdinalValuesCommand command.
func (cmd *PushOrdinalValuesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/push"
	}
	var payload client.PushOrdinalValuesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.PushOrdinalValues(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *PushOrdinalValuesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the RegisterOrdinalValuesCommand command.
func (cmd *RegisterOrdinalValuesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/register"
	}
	var payload client.RegisterOrdinalValuesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RegisterOrdinalValues(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RegisterOrdinalValuesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the StatisticsOrdinalValuesCommand command.
func (cmd *StatisticsOrdinalValuesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/statistics"
	}
	var payload client.StatisticsOrdinalValuesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StatisticsOrdinalValues(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StatisticsOrdinalValuesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the TagOrdinalValuesCommand command.
func (cmd *TagOrdinalValuesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/tag"
	}
	var payload client.TagOrdinalValuesPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.TagOrdinalValues(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TagOrdinalValuesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}
