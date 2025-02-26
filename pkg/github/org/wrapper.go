package org

import (
	p "github.com/kemadev/iac-components/pkg/github/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WrapperArgs struct {
	Provider p.ProviderArgs
	Settings SettingsArgs
}

func setDefaultArgs(args *WrapperArgs) {
	p.SetDefaults(&args.Provider)
}

func Wrapper(ctx *pulumi.Context, args WrapperArgs) error {
	setDefaultArgs(&args)
	provider, err := p.NewProvider(ctx, args.Provider)
	if err != nil {
		return err
	}
	err = createSettings(ctx, provider, args.Settings)
	if err != nil {
		return err
	}
	return nil
}
