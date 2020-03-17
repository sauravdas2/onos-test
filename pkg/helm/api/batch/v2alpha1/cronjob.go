// Code generated by onit-generate. DO NOT EDIT.

package v2alpha1

import (
	"github.com/onosproject/onos-test/pkg/helm/api/resource"
	batchv2alpha1 "k8s.io/api/batch/v2alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var CronJobKind = resource.Kind{
	Group:   "batch",
	Version: "v2alpha1",
	Kind:    "CronJob",
}

var CronJobResource = resource.Type{
	Kind: CronJobKind,
	Name: "cronjobs",
}

func NewCronJob(cronJob *batchv2alpha1.CronJob, client resource.Client) *CronJob {
	return &CronJob{
		Resource: resource.NewResource(cronJob.ObjectMeta, CronJobKind, client),
		Object:   cronJob,
	}
}

type CronJob struct {
	*resource.Resource
	Object *batchv2alpha1.CronJob
}

func (r *CronJob) Delete() error {
	return r.Clientset().
		BatchV2alpha1().
		RESTClient().
		Delete().
		Namespace(r.Namespace).
		Resource(CronJobResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}