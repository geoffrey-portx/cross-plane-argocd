/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package service

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/servicediscovery"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/servicediscovery/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetServiceInput returns input for read
// operation.
func GenerateGetServiceInput(cr *svcapitypes.Service) *svcsdk.GetServiceInput {
	res := &svcsdk.GetServiceInput{}

	if cr.Status.AtProvider.ID != nil {
		res.SetId(*cr.Status.AtProvider.ID)
	}

	return res
}

// GenerateService returns the current state in the form of *svcapitypes.Service.
func GenerateService(resp *svcsdk.GetServiceOutput) *svcapitypes.Service {
	cr := &svcapitypes.Service{}

	if resp.Service.Arn != nil {
		cr.Status.AtProvider.ARN = resp.Service.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.Service.CreateDate != nil {
		cr.Status.AtProvider.CreateDate = &metav1.Time{*resp.Service.CreateDate}
	} else {
		cr.Status.AtProvider.CreateDate = nil
	}
	if resp.Service.CreatorRequestId != nil {
		cr.Spec.ForProvider.CreatorRequestID = resp.Service.CreatorRequestId
	} else {
		cr.Spec.ForProvider.CreatorRequestID = nil
	}
	if resp.Service.Description != nil {
		cr.Spec.ForProvider.Description = resp.Service.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Service.DnsConfig != nil {
		f4 := &svcapitypes.DNSConfig{}
		if resp.Service.DnsConfig.DnsRecords != nil {
			f4f0 := []*svcapitypes.DNSRecord{}
			for _, f4f0iter := range resp.Service.DnsConfig.DnsRecords {
				f4f0elem := &svcapitypes.DNSRecord{}
				if f4f0iter.TTL != nil {
					f4f0elem.TTL = f4f0iter.TTL
				}
				if f4f0iter.Type != nil {
					f4f0elem.Type = f4f0iter.Type
				}
				f4f0 = append(f4f0, f4f0elem)
			}
			f4.DNSRecords = f4f0
		}
		if resp.Service.DnsConfig.NamespaceId != nil {
			f4.NamespaceID = resp.Service.DnsConfig.NamespaceId
		}
		if resp.Service.DnsConfig.RoutingPolicy != nil {
			f4.RoutingPolicy = resp.Service.DnsConfig.RoutingPolicy
		}
		cr.Spec.ForProvider.DNSConfig = f4
	} else {
		cr.Spec.ForProvider.DNSConfig = nil
	}
	if resp.Service.HealthCheckConfig != nil {
		f5 := &svcapitypes.HealthCheckConfig{}
		if resp.Service.HealthCheckConfig.FailureThreshold != nil {
			f5.FailureThreshold = resp.Service.HealthCheckConfig.FailureThreshold
		}
		if resp.Service.HealthCheckConfig.ResourcePath != nil {
			f5.ResourcePath = resp.Service.HealthCheckConfig.ResourcePath
		}
		if resp.Service.HealthCheckConfig.Type != nil {
			f5.Type = resp.Service.HealthCheckConfig.Type
		}
		cr.Spec.ForProvider.HealthCheckConfig = f5
	} else {
		cr.Spec.ForProvider.HealthCheckConfig = nil
	}
	if resp.Service.HealthCheckCustomConfig != nil {
		f6 := &svcapitypes.HealthCheckCustomConfig{}
		if resp.Service.HealthCheckCustomConfig.FailureThreshold != nil {
			f6.FailureThreshold = resp.Service.HealthCheckCustomConfig.FailureThreshold
		}
		cr.Spec.ForProvider.HealthCheckCustomConfig = f6
	} else {
		cr.Spec.ForProvider.HealthCheckCustomConfig = nil
	}
	if resp.Service.Id != nil {
		cr.Status.AtProvider.ID = resp.Service.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.Service.InstanceCount != nil {
		cr.Status.AtProvider.InstanceCount = resp.Service.InstanceCount
	} else {
		cr.Status.AtProvider.InstanceCount = nil
	}
	if resp.Service.Name != nil {
		cr.Spec.ForProvider.Name = resp.Service.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.Service.NamespaceId != nil {
		cr.Spec.ForProvider.NamespaceID = resp.Service.NamespaceId
	} else {
		cr.Spec.ForProvider.NamespaceID = nil
	}
	if resp.Service.Type != nil {
		cr.Spec.ForProvider.Type = resp.Service.Type
	} else {
		cr.Spec.ForProvider.Type = nil
	}

	return cr
}

// GenerateCreateServiceInput returns a create input.
func GenerateCreateServiceInput(cr *svcapitypes.Service) *svcsdk.CreateServiceInput {
	res := &svcsdk.CreateServiceInput{}

	if cr.Spec.ForProvider.CreatorRequestID != nil {
		res.SetCreatorRequestId(*cr.Spec.ForProvider.CreatorRequestID)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.DNSConfig != nil {
		f2 := &svcsdk.DnsConfig{}
		if cr.Spec.ForProvider.DNSConfig.DNSRecords != nil {
			f2f0 := []*svcsdk.DnsRecord{}
			for _, f2f0iter := range cr.Spec.ForProvider.DNSConfig.DNSRecords {
				f2f0elem := &svcsdk.DnsRecord{}
				if f2f0iter.TTL != nil {
					f2f0elem.SetTTL(*f2f0iter.TTL)
				}
				if f2f0iter.Type != nil {
					f2f0elem.SetType(*f2f0iter.Type)
				}
				f2f0 = append(f2f0, f2f0elem)
			}
			f2.SetDnsRecords(f2f0)
		}
		if cr.Spec.ForProvider.DNSConfig.NamespaceID != nil {
			f2.SetNamespaceId(*cr.Spec.ForProvider.DNSConfig.NamespaceID)
		}
		if cr.Spec.ForProvider.DNSConfig.RoutingPolicy != nil {
			f2.SetRoutingPolicy(*cr.Spec.ForProvider.DNSConfig.RoutingPolicy)
		}
		res.SetDnsConfig(f2)
	}
	if cr.Spec.ForProvider.HealthCheckConfig != nil {
		f3 := &svcsdk.HealthCheckConfig{}
		if cr.Spec.ForProvider.HealthCheckConfig.FailureThreshold != nil {
			f3.SetFailureThreshold(*cr.Spec.ForProvider.HealthCheckConfig.FailureThreshold)
		}
		if cr.Spec.ForProvider.HealthCheckConfig.ResourcePath != nil {
			f3.SetResourcePath(*cr.Spec.ForProvider.HealthCheckConfig.ResourcePath)
		}
		if cr.Spec.ForProvider.HealthCheckConfig.Type != nil {
			f3.SetType(*cr.Spec.ForProvider.HealthCheckConfig.Type)
		}
		res.SetHealthCheckConfig(f3)
	}
	if cr.Spec.ForProvider.HealthCheckCustomConfig != nil {
		f4 := &svcsdk.HealthCheckCustomConfig{}
		if cr.Spec.ForProvider.HealthCheckCustomConfig.FailureThreshold != nil {
			f4.SetFailureThreshold(*cr.Spec.ForProvider.HealthCheckCustomConfig.FailureThreshold)
		}
		res.SetHealthCheckCustomConfig(f4)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}
	if cr.Spec.ForProvider.NamespaceID != nil {
		res.SetNamespaceId(*cr.Spec.ForProvider.NamespaceID)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f7 := []*svcsdk.Tag{}
		for _, f7iter := range cr.Spec.ForProvider.Tags {
			f7elem := &svcsdk.Tag{}
			if f7iter.Key != nil {
				f7elem.SetKey(*f7iter.Key)
			}
			if f7iter.Value != nil {
				f7elem.SetValue(*f7iter.Value)
			}
			f7 = append(f7, f7elem)
		}
		res.SetTags(f7)
	}
	if cr.Spec.ForProvider.Type != nil {
		res.SetType(*cr.Spec.ForProvider.Type)
	}

	return res
}

// GenerateUpdateServiceInput returns an update input.
func GenerateUpdateServiceInput(cr *svcapitypes.Service) *svcsdk.UpdateServiceInput {
	res := &svcsdk.UpdateServiceInput{}

	if cr.Status.AtProvider.ID != nil {
		res.SetId(*cr.Status.AtProvider.ID)
	}

	return res
}

// GenerateDeleteServiceInput returns a deletion input.
func GenerateDeleteServiceInput(cr *svcapitypes.Service) *svcsdk.DeleteServiceInput {
	res := &svcsdk.DeleteServiceInput{}

	if cr.Status.AtProvider.ID != nil {
		res.SetId(*cr.Status.AtProvider.ID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ServiceNotFound"
}
