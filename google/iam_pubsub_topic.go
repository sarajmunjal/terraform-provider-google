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

var PubsubTopicIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"topic": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type PubsubTopicIamUpdater struct {
	project string
	topic   string
	d       *schema.ResourceData
	Config  *Config
}

func PubsubTopicIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	values["project"] = project
	if v, ok := d.GetOk("topic"); ok {
		values["topic"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/topics/(?P<topic>[^/]+)", "(?P<project>[^/]+)/(?P<topic>[^/]+)", "(?P<topic>[^/]+)"}, d, config, d.Get("topic").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &PubsubTopicIamUpdater{
		project: values["project"],
		topic:   values["topic"],
		d:       d,
		Config:  config,
	}

	d.Set("project", u.project)
	d.Set("topic", u.GetResourceId())

	return u, nil
}

func PubsubTopicIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/topics/(?P<topic>[^/]+)", "(?P<project>[^/]+)/(?P<topic>[^/]+)", "(?P<topic>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &PubsubTopicIamUpdater{
		project: values["project"],
		topic:   values["topic"],
		d:       d,
		Config:  config,
	}
	d.Set("topic", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *PubsubTopicIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyTopicUrl("getIamPolicy")
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

func (u *PubsubTopicIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyTopicUrl("setIamPolicy")
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

func (u *PubsubTopicIamUpdater) qualifyTopicUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{PubsubBasePath}}%s:%s", fmt.Sprintf("projects/%s/topics/%s", u.project, u.topic), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *PubsubTopicIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/topics/%s", u.project, u.topic)
}

func (u *PubsubTopicIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-pubsub-topic-%s", u.GetResourceId())
}

func (u *PubsubTopicIamUpdater) DescribeResource() string {
	return fmt.Sprintf("pubsub topic %q", u.GetResourceId())
}
