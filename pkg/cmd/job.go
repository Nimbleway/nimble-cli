// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/Nimbleway/nimble-cli/internal/apiquery"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
	"github.com/Nimbleway/nimble-go"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var jobsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create Job",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-name",
			Required: true,
			BodyPath: "agent_name",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "destination",
			BodyPath: "destination",
		},
		&requestflag.Flag[*string]{
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inputs",
			BodyPath: "inputs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "schedule",
			BodyPath: "schedule",
		},
	},
	Action:          handleJobsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"destination": {
		&requestflag.InnerFlag[string]{
			Name:       "destination.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "destination.type",
			Usage:      `Allowed values: "file", "s3".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "destination.format",
			Usage:      `Allowed values: "jsonl", "csv", "parquet".`,
			InnerField: "format",
		},
	},
	"inputs": {
		&requestflag.InnerFlag[string]{
			Name:       "inputs.type",
			Usage:      `Allowed values: "s3", "inline", "file".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inputs.data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "inputs.file-path",
			InnerField: "file_path",
		},
	},
	"schedule": {
		&requestflag.InnerFlag[string]{
			Name:       "schedule.cron",
			InnerField: "cron",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "schedule.enabled",
			InnerField: "enabled",
		},
	},
})

var jobsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update Job",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "destination",
			BodyPath: "destination",
		},
		&requestflag.Flag[*string]{
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inputs",
			BodyPath: "inputs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "schedule",
			BodyPath: "schedule",
		},
	},
	Action:          handleJobsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"destination": {
		&requestflag.InnerFlag[string]{
			Name:       "destination.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "destination.type",
			Usage:      `Allowed values: "file", "s3".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "destination.format",
			Usage:      `Allowed values: "jsonl", "csv", "parquet".`,
			InnerField: "format",
		},
	},
	"inputs": {
		&requestflag.InnerFlag[string]{
			Name:       "inputs.type",
			Usage:      `Allowed values: "s3", "inline", "file".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inputs.data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "inputs.file-path",
			InnerField: "file_path",
		},
	},
	"schedule": {
		&requestflag.InnerFlag[string]{
			Name:       "schedule.cron",
			InnerField: "cron",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "schedule.enabled",
			InnerField: "enabled",
		},
	},
})

var jobsList = cli.Command{
	Name:    "list",
	Usage:   "List Jobs",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "agent-name",
			Usage:     "Filter by agent name",
			QueryPath: "agent_name",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "per-page",
			Default:   20,
			QueryPath: "per_page",
		},
		&requestflag.Flag[*string]{
			Name:      "q",
			Usage:     "Search by name or display name",
			QueryPath: "q",
		},
	},
	Action:          handleJobsList,
	HideHelpCommand: true,
}

var jobsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Job",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
	},
	Action:          handleJobsDelete,
	HideHelpCommand: true,
}

var jobsGet = cli.Command{
	Name:    "get",
	Usage:   "Get Job",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
	},
	Action:          handleJobsGet,
	HideHelpCommand: true,
}

var jobsRun = cli.Command{
	Name:    "run",
	Usage:   "Trigger Run",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
	},
	Action:          handleJobsRun,
	HideHelpCommand: true,
}

func handleJobsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomnimblewaynimblego.JobNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Jobs.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "jobs create",
		Transform:      transform,
	})
}

func handleJobsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomnimblewaynimblego.JobUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Jobs.Update(
		ctx,
		cmd.Value("job-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "jobs update",
		Transform:      transform,
	})
}

func handleJobsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomnimblewaynimblego.JobListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Jobs.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "jobs list",
		Transform:      transform,
	})
}

func handleJobsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Jobs.Delete(ctx, cmd.Value("job-id").(string), options...)
}

func handleJobsGet(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Jobs.Get(ctx, cmd.Value("job-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "jobs get",
		Transform:      transform,
	})
}

func handleJobsRun(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Jobs.Run(ctx, cmd.Value("job-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "jobs run",
		Transform:      transform,
	})
}
