// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetworkSecurityGatewaySecurityPoliciesRule_networkSecurityGatewaySecurityPoliciesRulesBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityGatewaySecurityPoliciesRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityGatewaySecurityPoliciesRule_networkSecurityGatewaySecurityPoliciesRulesBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_gateway_security_policies_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "gateway_security_policy"},
			},
		},
	})
}

func testAccNetworkSecurityGatewaySecurityPoliciesRule_networkSecurityGatewaySecurityPoliciesRulesBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_security_gateway_security_policies" "default" {
  provider    = google-beta
  name        = "tf-test-my-gateway-security-policy%{random_suffix}"
  location    = "us-central1"
  description = "gateway security policy created to be used as reference by the rule."
}

resource "google_network_security_gateway_security_policies_rule" "default" {
  provider                = google-beta
  name                    = "tf-test-my-gateway-security-policy-rule%{random_suffix}"
  location                = "us-central1"
  gateway_security_policy = google_network_security_gateway_security_policies.default.name
  enabled                 = true  
  description             = "my description"
  priority                = 0
  session_matcher         = "host() == 'example.com'"
  basic_profile           = "ALLOW"
}
`, context)
}

func TestAccNetworkSecurityGatewaySecurityPoliciesRule_networkSecurityGatewaySecurityPoliciesRulesAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityGatewaySecurityPoliciesRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityGatewaySecurityPoliciesRule_networkSecurityGatewaySecurityPoliciesRulesAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_security_gateway_security_policies_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "gateway_security_policy"},
			},
		},
	})
}

func testAccNetworkSecurityGatewaySecurityPoliciesRule_networkSecurityGatewaySecurityPoliciesRulesAdvancedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_security_gateway_security_policies" "default" {
  provider    = google-beta
  name        = "tf-test-my-gateway-security-policy%{random_suffix}"
  location    = "us-central1"
  description = "gateway security policy created to be used as reference by the rule."
}

resource "google_network_security_gateway_security_policies_rule" "default" {
  provider                = google-beta
  name                    = "tf-test-my-gateway-security-policy-rule%{random_suffix}"
  location                = "us-central1"
  gateway_security_policy = google_network_security_gateway_security_policies.default.name
  enabled                 = true  
  description             = "my description"
  priority                = 0
  session_matcher         = "host() == 'example.com'"
  application_matcher     = "request.method == 'POST'"
  tls_inspection_enabled  = false
  basic_profile           = "ALLOW"
}
`, context)
}

func testAccCheckNetworkSecurityGatewaySecurityPoliciesRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_gateway_security_policies_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkSecurityGatewaySecurityPoliciesRule still exists at %s", url)
			}
		}

		return nil
	}
}
