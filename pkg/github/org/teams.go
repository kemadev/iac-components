package org

import (
	"slices"

	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TeamArgs struct {
	Name        string
	Description string
	Privacy     string
	ParentTeam  string
}

type TeamsArgs struct {
	Teams []TeamArgs
}

var (
	TeamsDefaultArgs = TeamsArgs{
		Teams: []TeamArgs{
			{
				Name:        "admins",
				Description: "Full ccess everywhere",
			},
			{
				Name:        "maintainers",
				Description: "Maintain permissions on all repositories",
			},
			{
				Name:        "developers",
				Description: "Parent team for all developers",
			},
		},
	}
)

func createTeamsSetDefaults(args *TeamsArgs) {
	if args.Teams == nil {
		args.Teams = TeamsDefaultArgs.Teams
		return
	}
	if len(args.Teams) == 0 {
		args.Teams = TeamsDefaultArgs.Teams
		return
	}
	for _, action := range TeamsDefaultArgs.Teams {
		found := slices.Contains(args.Teams, action)
		if !found {
			args.Teams = append(args.Teams, action)
		}
	}
}

func createTeams(ctx *pulumi.Context, provider *github.Provider, argsTeams TeamsArgs) error {
	for _, team := range argsTeams.Teams {
		teamName := util.FormatResourceName(ctx, team.Name)
		_, err := github.NewTeam(ctx, teamName, &github.TeamArgs{
			Name:         pulumi.String(team.Name),
			Description:  pulumi.String(team.Description),
			Privacy:      pulumi.String("closed"),
			ParentTeamId: pulumi.String(team.ParentTeam),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
	}
	return nil
}
