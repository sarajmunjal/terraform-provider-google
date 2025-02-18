// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var CloudFunctionsCloudFunctionIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"region": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"cloud_function": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type CloudFunctionsCloudFunctionIamUpdater struct {
	project       string
	region        string
	cloudFunction string
	d             *schema.ResourceData
	Config        *Config
}

func CloudFunctionsCloudFunctionIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	values["project"] = project
	region, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}
	values["region"] = region
	if v, ok := d.GetOk("cloud_function"); ok {
		values["cloud_function"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/functions/(?P<cloud_function>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<cloud_function>[^/]+)"}, d, config, d.Get("cloud_function").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &CloudFunctionsCloudFunctionIamUpdater{
		project:       values["project"],
		region:        values["region"],
		cloudFunction: values["cloud_function"],
		d:             d,
		Config:        config,
	}

	d.Set("project", u.project)
	d.Set("region", u.region)
	d.Set("cloud_function", u.GetResourceId())

	return u, nil
}

func CloudFunctionsCloudFunctionIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	values["project"] = project
	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	values["region"] = region

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/functions/(?P<cloud_function>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<cloud_function>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &CloudFunctionsCloudFunctionIamUpdater{
		project:       values["project"],
		region:        values["region"],
		cloudFunction: values["cloud_function"],
		d:             d,
		Config:        config,
	}
	d.Set("cloud_function", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *CloudFunctionsCloudFunctionIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyCloudFunctionUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	policy, err := sendRequest(u.Config, "GET", project, url, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *CloudFunctionsCloudFunctionIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyCloudFunctionUrl("setIamPolicy")
	if err != nil {
		return err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *CloudFunctionsCloudFunctionIamUpdater) qualifyCloudFunctionUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{CloudFunctionsBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/functions/%s", u.project, u.region, u.cloudFunction), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *CloudFunctionsCloudFunctionIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/functions/%s", u.project, u.region, u.cloudFunction)
}

func (u *CloudFunctionsCloudFunctionIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-cloudfunctions-cloudfunction-%s", u.GetResourceId())
}

func (u *CloudFunctionsCloudFunctionIamUpdater) DescribeResource() string {
	return fmt.Sprintf("cloudfunctions cloudfunction %q", u.GetResourceId())
}
