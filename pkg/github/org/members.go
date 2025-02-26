package org

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type User struct {
	Username string
	Role     string
}

type MembersArgs struct {
	Members []User
	Admins  []User
}

func createMembers(ctx *pulumi.Context, provider *github.Provider, argsMembers MembersArgs) error {
	for _, t := range argsMembers.Members {
		memberName := util.FormatResourceName(ctx, "Member")
		_, err := github.NewMembership(ctx, memberName, &github.MembershipArgs{
			Username: pulumi.String(t.Username),
			Role:     pulumi.String(t.Role),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
	}
	return nil
}
