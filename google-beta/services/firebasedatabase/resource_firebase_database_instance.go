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

package firebasedatabase

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func enableRTDB(config *transport_tpg.Config, d *schema.ResourceData, project string, billingProject string, userAgent string) error {
	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseDatabaseBasePath}}projects/{{project}}/locations/{{region}}/instances/{{instance_id}}:reenable")
	if err != nil {
		return err
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return fmt.Errorf("Error reenabling firebase_database_instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished reenabling firebase_database_instance %q: %#v", d.Id(), res)
	}
	return nil
}

func disableRTDB(config *transport_tpg.Config, d *schema.ResourceData, project string, billingProject string, userAgent string) error {
	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseDatabaseBasePath}}projects/{{project}}/locations/{{region}}/instances/{{instance_id}}:disable")
	if err != nil {
		return err
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return fmt.Errorf("Error disabling firebase_database_instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished disabling firebase_database_instance %q: %#v", d.Id(), res)
	}
	return nil
}

func ResourceFirebaseDatabaseInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseDatabaseInstanceCreate,
		Read:   resourceFirebaseDatabaseInstanceRead,
		Update: resourceFirebaseDatabaseInstanceUpdate,
		Delete: resourceFirebaseDatabaseInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseDatabaseInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The globally unique identifier of the Firebase Realtime Database instance.
Instance IDs cannot be reused after deletion.`,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A reference to the region where the Firebase Realtime database resides.
Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DEFAULT_DATABASE", "USER_DATABASE", ""}),
				Description: `The database type.
Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
Creating user Databases is only available for projects on the Blaze plan.
Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value: "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]`,
				Default: "USER_DATABASE",
			},
			"database_url": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
or https://{instance-id}.{region}.firebasedatabase.app in other regions.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully-qualified resource name of the Firebase Realtime Database, in the
format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
PROJECT_NUMBER: The Firebase project's ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database`,
			},
			"desired_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ACTIVE",
				Description: `The intended database state.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseDatabaseInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	typeProp, err := expandFirebaseDatabaseInstanceType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseDatabaseBasePath}}projects/{{project}}/locations/{{region}}/instances?databaseId={{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}
	if err := d.Set("name", flattenFirebaseDatabaseInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{instance_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// start of customized code
	if p, ok := d.GetOk("desired_state"); ok && p.(string) == "DISABLED" {
		if err := disableRTDB(config, d, project, billingProject, userAgent); err != nil {
			return err
		}
	}
	// end of customized code

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceFirebaseDatabaseInstanceRead(d, meta)
}

func resourceFirebaseDatabaseInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseDatabaseBasePath}}projects/{{project}}/locations/{{region}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseDatabaseInstance %q", d.Id()))
	}

	res, err = resourceFirebaseDatabaseInstanceDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing FirebaseDatabaseInstance because it no longer exists.")
		d.SetId("")
		return nil
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("desired_state"); !ok {
		if err := d.Set("desired_state", "ACTIVE"); err != nil {
			return fmt.Errorf("Error setting desired_state: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("name", flattenFirebaseDatabaseInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("database_url", flattenFirebaseDatabaseInstanceDatabaseUrl(res["databaseUrl"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("type", flattenFirebaseDatabaseInstanceType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenFirebaseDatabaseInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceFirebaseDatabaseInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	typeProp, err := expandFirebaseDatabaseInstanceType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseDatabaseBasePath}}projects/{{project}}/locations/{{region}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)

	// start of customized code
	if d.HasChange("desired_state") {
		if p, ok := d.GetOk("desired_state"); ok && p.(string) != d.Get("state").(string) {
			switch p.(string) {
			case "ACTIVE":
				if err := enableRTDB(config, d, project, billingProject, userAgent); err != nil {
					return err
				}
			case "DISABLED":
				if err := disableRTDB(config, d, project, billingProject, userAgent); err != nil {
					return err
				}
			default:
				return fmt.Errorf("Unsupported value in field `desired_state`: %v", p)
			}
		}
		// firebasedatabase does not update UpdateDatabaseInstance endpoint now.
		return nil
	}
	// end of customized code

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	return resourceFirebaseDatabaseInstanceRead(d, meta)
}

func resourceFirebaseDatabaseInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseDatabaseBasePath}}projects/{{project}}/locations/{{region}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// start of customized code
	if d.Get("state").(string) == "ACTIVE" {
		if err := disableRTDB(config, d, project, billingProject, userAgent); err != nil {
			return err
		}
	}
	if d.Get("type").(string) == "DEFAULT_DATABASE" {
		log.Printf("[WARN] Default Firebase Database Instance %q cannot be deleted, left disabled", d.Id())
		return nil
	}
	// end of customized code

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting Instance %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Instance")
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceFirebaseDatabaseInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/instances/(?P<instance_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<instance_id>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<instance_id>[^/]+)$",
		"^(?P<instance_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{instance_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("desired_state", "ACTIVE"); err != nil {
		return nil, fmt.Errorf("Error setting desired_state: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseDatabaseInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseDatabaseInstanceDatabaseUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseDatabaseInstanceType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseDatabaseInstanceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseDatabaseInstanceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceFirebaseDatabaseInstanceDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v := res["state"]; v == "DELETED" {
		return nil, nil
	}
	res["desired_state"] = res["state"]

	return res, nil
}
