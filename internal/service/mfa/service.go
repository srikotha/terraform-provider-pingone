// Copyright © 2025 Ping Identity Corporation

package mfa

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
)

type serviceClientType struct {
	Client *pingone.Client
}

func Resources() []func() resource.Resource {
	return []func() resource.Resource{
		NewApplicationPushCredentialResource,
		NewFIDO2PolicyResource,
		NewMFADevicePolicyResource,
		NewMFASettingsResource,
	}
}

func DataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewMFADevicePoliciesDataSource,
	}
}
