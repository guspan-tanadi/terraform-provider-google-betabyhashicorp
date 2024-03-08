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

package vertexai

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceVertexAIFeatureGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIFeatureGroupCreate,
		Read:   resourceVertexAIFeatureGroupRead,
		Update: resourceVertexAIFeatureGroupUpdate,
		Delete: resourceVertexAIFeatureGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIFeatureGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"big_query": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Indicates that features for this group come from BigQuery Table/View. By default treats the source as a sparse time series source, which is required to have an entityId and a feature_timestamp column in the source.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"big_query_source": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `The BigQuery source URI that points to either a BigQuery Table or View.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"input_uri": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `BigQuery URI to a table, up to 2000 characters long. For example: 'bq://projectId.bqDatasetId.bqTableId.'`,
									},
								},
							},
						},
						"entity_id_columns": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Columns to construct entityId / row keys. Currently only supports 1 entity_id_column. If not provided defaults to entityId.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the FeatureGroup.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels with user-defined metadata to organize your FeatureGroup.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The resource name of the Feature Group.`,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of feature group. eg us-central1`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the FeatureGroup was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the FeatureGroup was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
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

func resourceVertexAIFeatureGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandVertexAIFeatureGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandVertexAIFeatureGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	bigQueryProp, err := expandVertexAIFeatureGroupBigQuery(d.Get("big_query"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("big_query"); !tpgresource.IsEmptyValue(reflect.ValueOf(bigQueryProp)) && (ok || !reflect.DeepEqual(v, bigQueryProp)) {
		obj["bigQuery"] = bigQueryProp
	}
	labelsProp, err := expandVertexAIFeatureGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureGroups?featureGroupId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FeatureGroup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureGroup: %s", err)
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
		return fmt.Errorf("Error creating FeatureGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating FeatureGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create FeatureGroup: %s", err)
	}

	if err := d.Set("name", flattenVertexAIFeatureGroupName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FeatureGroup %q: %#v", d.Id(), res)

	return resourceVertexAIFeatureGroupRead(d, meta)
}

func resourceVertexAIFeatureGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureGroups/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureGroup: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VertexAIFeatureGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}

	if err := d.Set("name", flattenVertexAIFeatureGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}
	if err := d.Set("create_time", flattenVertexAIFeatureGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIFeatureGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIFeatureGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}
	if err := d.Set("description", flattenVertexAIFeatureGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}
	if err := d.Set("big_query", flattenVertexAIFeatureGroupBigQuery(res["bigQuery"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}
	if err := d.Set("terraform_labels", flattenVertexAIFeatureGroupTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}
	if err := d.Set("effective_labels", flattenVertexAIFeatureGroupEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureGroup: %s", err)
	}

	return nil
}

func resourceVertexAIFeatureGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureGroup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandVertexAIFeatureGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandVertexAIFeatureGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	bigQueryProp, err := expandVertexAIFeatureGroupBigQuery(d.Get("big_query"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("big_query"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bigQueryProp)) {
		obj["bigQuery"] = bigQueryProp
	}
	labelsProp, err := expandVertexAIFeatureGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureGroups/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FeatureGroup %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("name") {
		updateMask = append(updateMask, "name")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("big_query") {
		updateMask = append(updateMask, "bigQuery")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})

		if err != nil {
			return fmt.Errorf("Error updating FeatureGroup %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating FeatureGroup %q: %#v", d.Id(), res)
		}

		err = VertexAIOperationWaitTime(
			config, res, project, "Updating FeatureGroup", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceVertexAIFeatureGroupRead(d, meta)
}

func resourceVertexAIFeatureGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureGroup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting FeatureGroup %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "FeatureGroup")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting FeatureGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting FeatureGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIFeatureGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/featureGroups/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIFeatureGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenVertexAIFeatureGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenVertexAIFeatureGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("description")
}

func flattenVertexAIFeatureGroupBigQuery(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["big_query_source"] =
		flattenVertexAIFeatureGroupBigQueryBigQuerySource(original["bigQuerySource"], d, config)
	transformed["entity_id_columns"] =
		flattenVertexAIFeatureGroupBigQueryEntityIdColumns(original["entityIdColumns"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureGroupBigQueryBigQuerySource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["input_uri"] =
		flattenVertexAIFeatureGroupBigQueryBigQuerySourceInputUri(original["inputUri"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureGroupBigQueryBigQuerySourceInputUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureGroupBigQueryEntityIdColumns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureGroupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenVertexAIFeatureGroupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAIFeatureGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureGroupBigQuery(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBigQuerySource, err := expandVertexAIFeatureGroupBigQueryBigQuerySource(original["big_query_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBigQuerySource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bigQuerySource"] = transformedBigQuerySource
	}

	transformedEntityIdColumns, err := expandVertexAIFeatureGroupBigQueryEntityIdColumns(original["entity_id_columns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEntityIdColumns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["entityIdColumns"] = transformedEntityIdColumns
	}

	return transformed, nil
}

func expandVertexAIFeatureGroupBigQueryBigQuerySource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInputUri, err := expandVertexAIFeatureGroupBigQueryBigQuerySourceInputUri(original["input_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInputUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["inputUri"] = transformedInputUri
	}

	return transformed, nil
}

func expandVertexAIFeatureGroupBigQueryBigQuerySourceInputUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureGroupBigQueryEntityIdColumns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
