// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeSecurityPolicyRule_securityPolicyRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSecurityPolicyRule_securityPolicyRuleBasicExample(context),
			},
			{
				ResourceName:            "google_compute_security_policy_rule.policy_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_policy"},
			},
		},
	})
}

func testAccComputeSecurityPolicyRule_securityPolicyRuleBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_security_policy" "default" {
  name        = "policyruletest%{random_suffix}"
  description = "basic global security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_security_policy_rule" "policy_rule" {
  security_policy = google_compute_security_policy.default.name
  description     = "new rule"
  priority        = 100
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["10.10.0.0/16"]
    }
  }
  action          = "allow"
  preview         = true
}
`, context)
}

func TestAccComputeSecurityPolicyRule_securityPolicyRuleDefaultRuleExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSecurityPolicyRule_securityPolicyRuleDefaultRuleExample(context),
			},
			{
				ResourceName:            "google_compute_security_policy_rule.policy_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_policy"},
			},
		},
	})
}

func testAccComputeSecurityPolicyRule_securityPolicyRuleDefaultRuleExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_security_policy" "default" {
  name        = "policyruletest%{random_suffix}"
  description = "basic global security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_security_policy_rule" "default_rule" {
  security_policy = google_compute_security_policy.default.name
  description     = "default rule"
  action          = "deny"
  priority        = "2147483647"
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }
}

resource "google_compute_security_policy_rule" "policy_rule" {
  security_policy = google_compute_security_policy.default.name
  description     = "new rule"
  priority        = 100
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["10.10.0.0/16"]
    }
  }
  action          = "allow"
  preview         = true
}
`, context)
}

func TestAccComputeSecurityPolicyRule_securityPolicyRuleMultipleRulesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSecurityPolicyRule_securityPolicyRuleMultipleRulesExample(context),
			},
			{
				ResourceName:            "google_compute_security_policy_rule.policy_rule_one",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_policy"},
			},
		},
	})
}

func testAccComputeSecurityPolicyRule_securityPolicyRuleMultipleRulesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_security_policy" "default" {
  name        = "policywithmultiplerules%{random_suffix}"
  description = "basic global security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_security_policy_rule" "policy_rule_one" {
  security_policy = google_compute_security_policy.default.name
  description     = "new rule one"
  priority        = 100
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["10.10.0.0/16"]
    }
  }
  action          = "allow"
  preview         = true
}

resource "google_compute_security_policy_rule" "policy_rule_two" {
  security_policy = google_compute_security_policy.default.name
  description     = "new rule two"
  priority        = 101
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
    }
  }
  action          = "allow"
  preview         = true
}
`, context)
}

func testAccCheckComputeSecurityPolicyRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_security_policy_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/securityPolicies/{{security_policy}}/getRule?priority={{priority}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeSecurityPolicyRule still exists at %s", url)
			}
		}

		return nil
	}
}
