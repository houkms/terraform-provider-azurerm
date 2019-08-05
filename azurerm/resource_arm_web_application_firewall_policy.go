// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://github.com/Azure/magic-module-specs
//
// ----------------------------------------------------------------------------

package azurerm

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/response"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmWebApplicationFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmWebApplicationFirewallPolicyCreateUpdate,
		Read:   resourceArmWebApplicationFirewallPolicyRead,
		Update: resourceArmWebApplicationFirewallPolicyCreateUpdate,
		Delete: resourceArmWebApplicationFirewallPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"location": azure.SchemaLocation(),

			"resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

			"custom_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(network.WebApplicationFirewallActionAllow),
								string(network.WebApplicationFirewallActionBlock),
								string(network.WebApplicationFirewallActionLog),
							}, false),
							// Default: string(network.WebApplicationFirewallActionAllow),
						},
						"match_conditions": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"match_values": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"match_variables": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"variable_name": {
													Type:     schema.TypeString,
													Required: true,
													ValidateFunc: validation.StringInSlice([]string{
														string(network.RemoteAddr),
														string(network.RequestMethod),
														string(network.QueryString),
														string(network.PostArgs),
														string(network.RequestURI),
														string(network.RequestHeaders),
														string(network.RequestBody),
														string(network.RequestCookies),
													}, false),
													// Default: string(network.RemoteAddr),
												},
												"selector": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"operator": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											string(network.WebApplicationFirewallOperatorIPMatch),
											string(network.WebApplicationFirewallOperatorEqual),
											string(network.WebApplicationFirewallOperatorContains),
											string(network.WebApplicationFirewallOperatorLessThan),
											string(network.WebApplicationFirewallOperatorGreaterThan),
											string(network.WebApplicationFirewallOperatorLessThanOrEqual),
											string(network.WebApplicationFirewallOperatorGreaterThanOrEqual),
											string(network.WebApplicationFirewallOperatorBeginsWith),
											string(network.WebApplicationFirewallOperatorEndsWith),
											string(network.WebApplicationFirewallOperatorRegex),
										}, false),
										// Default: string(network.WebApplicationFirewallOperatorIPMatch),
									},
									"negation_conditon": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"transforms": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
											// Optional: true,
											// ValidateFunc: validation.StringInSlice([]string{
											// 	string(network.Lowercase),
											// 	string(network.Trim),
											// 	string(network.URLDecode),
											// 	string(network.URLEncode),
											// 	string(network.RemoveNulls),
											// 	string(network.HTMLEntityDecode),
											// }, false),
											// Default: string(network.Lowercase),
										},
									},
								},
							},
						},
						"priority": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"rule_type": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(network.WebApplicationFirewallRuleTypeMatchRule),
								string(network.WebApplicationFirewallRuleTypeInvalid),
							}, false),
							// Default: string(network.WebApplicationFirewallRuleTypeMatchRule),
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"etag": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"policy_settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled_state": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(network.Disabled),
								string(network.Enabled),
							}, false),
							Default: string(network.Disabled),
						},
						"mode": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(network.Prevention),
								string(network.Detection),
							}, false),
							Default: string(network.Prevention),
						},
					},
				},
			},

			"tags": tagsSchema(),

			"provisioning_state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"resource_state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceArmWebApplicationFirewallPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).webApplicationFirewallPoliciesClient
	ctx := meta.(*ArmClient).StopContext

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group").(string)

	if requireResourcesToBeImported {
		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Error checking for present of existing Web Application Firewall Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
			}
		}
		if !utils.ResponseWasNotFound(resp.Response) {
			return tf.ImportAsExistsError("azurerm_web_application_firewall_policy", *resp.ID)
		}
	}

	// id := d.Get("id").(string)
	location := azure.NormalizeLocation(d.Get("location").(string))
	customRules := d.Get("custom_rules").([]interface{})
	etag := d.Get("etag").(string)
	policySettings := d.Get("policy_settings").([]interface{})
	tags := d.Get("tags").(map[string]interface{})

	parameters := network.WebApplicationFirewallPolicy{
		Etag: utils.String(etag),
		// ID:       utils.String(id),
		Location: utils.String(location),
		WebApplicationFirewallPolicyPropertiesFormat: &network.WebApplicationFirewallPolicyPropertiesFormat{
			CustomRules:    expandArmWebApplicationFirewallPolicyWebApplicationFirewallCustomRule(customRules),
			PolicySettings: expandArmWebApplicationFirewallPolicyPolicySettings(policySettings),
		},
		Tags: expandTags(tags),
	}

	if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
		return fmt.Errorf("Error creating Web Application Firewall Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error retrieving Web Application Firewall Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if resp.ID == nil {
		return fmt.Errorf("Cannot read Web Application Firewall Policy %q (Resource Group %q) ID", name, resourceGroup)
	}
	d.SetId(*resp.ID)

	return resourceArmWebApplicationFirewallPolicyRead(d, meta)
}

func resourceArmWebApplicationFirewallPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).webApplicationFirewallPoliciesClient
	ctx := meta.(*ArmClient).StopContext

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["ApplicationGatewayWebApplicationFirewallPolicies"]

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Web Application Firewall Policy %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error reading Web Application Firewall Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group", resourceGroup)
	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}
	if properties := resp.WebApplicationFirewallPolicyPropertiesFormat; properties != nil {
		if err := d.Set("custom_rules", flattenArmWebApplicationFirewallPolicyWebApplicationFirewallCustomRule(properties.CustomRules)); err != nil {
			return fmt.Errorf("Error setting `custom_rules`: %+v", err)
		}
		if err := d.Set("policy_settings", flattenArmWebApplicationFirewallPolicyPolicySettings(properties.PolicySettings)); err != nil {
			return fmt.Errorf("Error setting `policy_settings`: %+v", err)
		}
		d.Set("provisioning_state", properties.ProvisioningState)
		d.Set("resource_state", string(properties.ResourceState))
	}
	d.Set("etag", resp.Etag)
	d.Set("type", resp.Type)
	flattenAndSetTags(d, resp.Tags)

	return nil
}

func resourceArmWebApplicationFirewallPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).webApplicationFirewallPoliciesClient
	ctx := meta.(*ArmClient).StopContext

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["ApplicationGatewayWebApplicationFirewallPolicies"]

	future, err := client.Delete(ctx, resourceGroup, name)
	if err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error deleting Web Application Firewall Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		if !response.WasNotFound(future.Response()) {
			return fmt.Errorf("Error waiting for deleting Web Application Firewall Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
		}
	}

	return nil
}

func expandArmWebApplicationFirewallPolicyWebApplicationFirewallCustomRule(input []interface{}) *[]network.WebApplicationFirewallCustomRule {
	results := make([]network.WebApplicationFirewallCustomRule, 0)
	for _, v := range input {
		v := v.(map[string]interface{})
		name := v["name"].(string)
		priority := v["priority"].(int)
		ruleType := v["rule_type"].(string)
		matchConditions := v["match_conditions"].([]interface{})
		action := v["action"].(string)

		item := network.WebApplicationFirewallCustomRule{
			Action:          network.WebApplicationFirewallAction(action),
			MatchConditions: expandArmWebApplicationFirewallPolicyMatchCondition(matchConditions),
			Name:            utils.String(name),
			Priority:        utils.Int32(int32(priority)),
			RuleType:        network.WebApplicationFirewallRuleType(ruleType),
		}

		results = append(results, item)
	}
	return &results
}

func expandArmWebApplicationFirewallPolicyPolicySettings(input []interface{}) *network.PolicySettings {
	if len(input) == 0 {
		return nil
	}
	v := input[0].(map[string]interface{})

	enabledState := v["enabled_state"].(string)
	mode := v["mode"].(string)

	result := network.PolicySettings{
		EnabledState: network.WebApplicationFirewallEnabledState(enabledState),
		Mode:         network.WebApplicationFirewallMode(mode),
	}
	return &result
}

func expandArmWebApplicationFirewallPolicyMatchCondition(input []interface{}) *[]network.MatchCondition {
	results := make([]network.MatchCondition, 0)
	for _, v := range input {
		v := v.(map[string]interface{})
		matchVariables := v["match_variables"].([]interface{})
		operator := v["operator"].(string)
		negationConditon := v["negation_conditon"].(bool)
		matchValues := v["match_values"].([]interface{})
		transforms := v["transforms"].([]interface{})

		item := network.MatchCondition{
			MatchValues:      expandArmWebApplicationFirewallPolicy(matchValues),
			MatchVariables:   expandArmWebApplicationFirewallPolicyMatchVariable(matchVariables),
			NegationConditon: utils.Bool(negationConditon),
			Operator:         network.WebApplicationFirewallOperator(operator),
			// Transforms:       network.WebApplicationFirewallTransform(transforms),
			Transforms: expandArmWebApplicationFirewallTransform(transforms),
		}

		results = append(results, item)
	}
	return &results
}

func expandArmWebApplicationFirewallPolicy(input []interface{}) *[]string {
	results := make([]string, 0)
	for _, v := range input {
		results = append(results, v.(string))
	}
	return &results
}

func expandArmWebApplicationFirewallPolicyMatchVariable(input []interface{}) *[]network.MatchVariable {
	results := make([]network.MatchVariable, 0)
	for _, v := range input {
		v := v.(map[string]interface{})
		variableName := v["variable_name"].(string)
		selector := v["selector"].(string)

		item := network.MatchVariable{
			Selector:     utils.String(selector),
			VariableName: network.WebApplicationFirewallMatchVariable(variableName),
		}

		results = append(results, item)
	}
	return &results
}

func expandArmWebApplicationFirewallTransform(input []interface{}) *[]network.WebApplicationFirewallTransform {
	results := make([]network.WebApplicationFirewallTransform, 0)
	for _, v := range input {
		results = append(results, v.(network.WebApplicationFirewallTransform))
	}
	return &results
}

func flattenArmWebApplicationFirewallPolicyWebApplicationFirewallCustomRule(input *[]network.WebApplicationFirewallCustomRule) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	for _, item := range *input {
		v := make(map[string]interface{})

		v["action"] = string(item.Action)
		v["match_conditions"] = flattenArmWebApplicationFirewallPolicyMatchCondition(item.MatchConditions)
		if priority := item.Priority; priority != nil {
			v["priority"] = int(*priority)
		}
		v["rule_type"] = string(item.RuleType)

		results = append(results, v)
	}

	return results
}

func flattenArmWebApplicationFirewallPolicyPolicySettings(input *network.PolicySettings) []interface{} {
	if input == nil {
		return make([]interface{}, 0)
	}

	result := make(map[string]interface{})

	result["enabled_state"] = string(input.EnabledState)
	result["mode"] = string(input.Mode)

	return []interface{}{result}
}

func flattenArmWebApplicationFirewallPolicyMatchCondition(input *[]network.MatchCondition) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	for _, item := range *input {
		v := make(map[string]interface{})

		v["match_values"] = flattenArmWebApplicationFirewallPolicy(item.MatchValues)
		v["match_variables"] = flattenArmWebApplicationFirewallPolicyMatchVariable(item.MatchVariables)
		if negationConditon := item.NegationConditon; negationConditon != nil {
			v["negation_conditon"] = *negationConditon
		}
		v["operator"] = string(item.Operator)
		v["transforms"] = flattenArmWebApplicationFirewallTransform(item.Transforms)

		results = append(results, v)
	}

	return results
}

func flattenArmWebApplicationFirewallPolicy(input *[]string) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	for _, item := range *input {
		results = append(results, item)
	}

	return results
}

func flattenArmWebApplicationFirewallPolicyMatchVariable(input *[]network.MatchVariable) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	for _, item := range *input {
		v := make(map[string]interface{})

		if selector := item.Selector; selector != nil {
			v["selector"] = *selector
		}
		v["variable_name"] = string(item.VariableName)

		results = append(results, v)
	}

	return results
}

func flattenArmWebApplicationFirewallTransform(input *[]network.WebApplicationFirewallTransform) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	for _, item := range *input {
		results = append(results, item)
	}

	return results
}
