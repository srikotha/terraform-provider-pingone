// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-terraform-plugin-framework-generator

package sso_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pingidentity/terraform-provider-pingone/internal/acctest"
	"github.com/pingidentity/terraform-provider-pingone/internal/acctest/service/sso"
	"github.com/pingidentity/terraform-provider-pingone/internal/verify"
)

func TestAccAdministratorSecurityDataSource_Get(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNoFeatureFlag(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             sso.IdentityProvider_CheckDestroy,
		ErrorCheck:               acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAdministratorSecurityDataSourceConfig_Get(resourceName),
				Check:  administratorSecurityDataSource_CheckComputedValuesComplete(resourceName),
			},
		},
	})
}

func testAccAdministratorSecurityDataSourceConfig_Get(resourceName string) string {
	return fmt.Sprintf(`
	%[1]s

resource "pingone_identity_provider" "%[2]s-idp" {
  environment_id = data.pingone_environment.general_test.id
  name           = "%[2]s-idp"
  enabled        = true

  microsoft = {
    client_id     = "dummyclientid1"
    client_secret = "dummyclientsecret1"
    tenant_id     = "dummytenantid1"
  }
}

resource "pingone_administrator_security" "%[2]s" {
  environment_id = data.pingone_environment.general_test.id
  allowed_methods = {
    email = jsonencode(
      {
        enabled = true
      }
    )
    fido2 = jsonencode(
      {
        enabled = false
      }
    )
    totp = jsonencode(
      {
        enabled = false
      }
    )
  }
  authentication_method = "HYBRID"
  mfa_status            = "ENFORCE"
  identity_provider = {
    id = pingone_identity_provider.%[2]s-idp.id
  }
  recovery = false
}

data "pingone_administrator_security" "%[2]s" {
  environment_id = data.pingone_environment.general_test.id

  depends_on = [
    pingone_administrator_security.%[2]s
  ]
}
`, acctest.GenericSandboxEnvironment(), resourceName)
}

// Validate any computed values when applying complete HCL
func administratorSecurityDataSource_CheckComputedValuesComplete(resourceName string) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "allowed_methods.email", `{"enabled":true}`),
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "allowed_methods.fido2", `{"enabled":false}`),
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "allowed_methods.totp", `{"enabled":false}`),
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "authentication_method", "HYBRID"),
		resource.TestMatchResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "environment_id", verify.P1ResourceIDRegexpFullString),
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "has_fido2_capabilities", "true"),
		resource.TestMatchResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "id", verify.P1ResourceIDRegexpFullString),
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "is_pingid_in_bom", "false"),
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "mfa_status", "ENFORCE"),
		resource.TestMatchResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "identity_provider.id", verify.P1ResourceIDRegexpFullString),
		resource.TestCheckResourceAttr(fmt.Sprintf("data.pingone_administrator_security.%s", resourceName), "recovery", "false"),
	)
}
