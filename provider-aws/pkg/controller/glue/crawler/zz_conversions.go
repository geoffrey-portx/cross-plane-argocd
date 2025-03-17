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

package crawler

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/glue"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/glue/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetCrawlerInput returns input for read
// operation.
func GenerateGetCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.GetCrawlerInput {
	res := &svcsdk.GetCrawlerInput{}

	return res
}

// GenerateCrawler returns the current state in the form of *svcapitypes.Crawler.
func GenerateCrawler(resp *svcsdk.GetCrawlerOutput) *svcapitypes.Crawler {
	cr := &svcapitypes.Crawler{}

	if resp.Crawler.Configuration != nil {
		cr.Spec.ForProvider.Configuration = resp.Crawler.Configuration
	} else {
		cr.Spec.ForProvider.Configuration = nil
	}
	if resp.Crawler.CrawlElapsedTime != nil {
		cr.Status.AtProvider.CrawlElapsedTime = resp.Crawler.CrawlElapsedTime
	} else {
		cr.Status.AtProvider.CrawlElapsedTime = nil
	}
	if resp.Crawler.CreationTime != nil {
		cr.Status.AtProvider.CreationTime = &metav1.Time{*resp.Crawler.CreationTime}
	} else {
		cr.Status.AtProvider.CreationTime = nil
	}
	if resp.Crawler.Description != nil {
		cr.Spec.ForProvider.Description = resp.Crawler.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Crawler.LakeFormationConfiguration != nil {
		f7 := &svcapitypes.LakeFormationConfiguration{}
		if resp.Crawler.LakeFormationConfiguration.AccountId != nil {
			f7.AccountID = resp.Crawler.LakeFormationConfiguration.AccountId
		}
		if resp.Crawler.LakeFormationConfiguration.UseLakeFormationCredentials != nil {
			f7.UseLakeFormationCredentials = resp.Crawler.LakeFormationConfiguration.UseLakeFormationCredentials
		}
		cr.Spec.ForProvider.LakeFormationConfiguration = f7
	} else {
		cr.Spec.ForProvider.LakeFormationConfiguration = nil
	}
	if resp.Crawler.LastCrawl != nil {
		f8 := &svcapitypes.LastCrawlInfo{}
		if resp.Crawler.LastCrawl.ErrorMessage != nil {
			f8.ErrorMessage = resp.Crawler.LastCrawl.ErrorMessage
		}
		if resp.Crawler.LastCrawl.LogGroup != nil {
			f8.LogGroup = resp.Crawler.LastCrawl.LogGroup
		}
		if resp.Crawler.LastCrawl.LogStream != nil {
			f8.LogStream = resp.Crawler.LastCrawl.LogStream
		}
		if resp.Crawler.LastCrawl.MessagePrefix != nil {
			f8.MessagePrefix = resp.Crawler.LastCrawl.MessagePrefix
		}
		if resp.Crawler.LastCrawl.StartTime != nil {
			f8.StartTime = &metav1.Time{*resp.Crawler.LastCrawl.StartTime}
		}
		if resp.Crawler.LastCrawl.Status != nil {
			f8.Status = resp.Crawler.LastCrawl.Status
		}
		cr.Status.AtProvider.LastCrawl = f8
	} else {
		cr.Status.AtProvider.LastCrawl = nil
	}
	if resp.Crawler.LastUpdated != nil {
		cr.Status.AtProvider.LastUpdated = &metav1.Time{*resp.Crawler.LastUpdated}
	} else {
		cr.Status.AtProvider.LastUpdated = nil
	}
	if resp.Crawler.LineageConfiguration != nil {
		f10 := &svcapitypes.LineageConfiguration{}
		if resp.Crawler.LineageConfiguration.CrawlerLineageSettings != nil {
			f10.CrawlerLineageSettings = resp.Crawler.LineageConfiguration.CrawlerLineageSettings
		}
		cr.Spec.ForProvider.LineageConfiguration = f10
	} else {
		cr.Spec.ForProvider.LineageConfiguration = nil
	}
	if resp.Crawler.RecrawlPolicy != nil {
		f12 := &svcapitypes.RecrawlPolicy{}
		if resp.Crawler.RecrawlPolicy.RecrawlBehavior != nil {
			f12.RecrawlBehavior = resp.Crawler.RecrawlPolicy.RecrawlBehavior
		}
		cr.Spec.ForProvider.RecrawlPolicy = f12
	} else {
		cr.Spec.ForProvider.RecrawlPolicy = nil
	}
	if resp.Crawler.Schedule != nil {
		cr.Spec.ForProvider.Schedule = resp.Crawler.Schedule.ScheduleExpression
	} else {
		cr.Spec.ForProvider.Schedule = nil
	}
	if resp.Crawler.SchemaChangePolicy != nil {
		f15 := &svcapitypes.SchemaChangePolicy{}
		if resp.Crawler.SchemaChangePolicy.DeleteBehavior != nil {
			f15.DeleteBehavior = resp.Crawler.SchemaChangePolicy.DeleteBehavior
		}
		if resp.Crawler.SchemaChangePolicy.UpdateBehavior != nil {
			f15.UpdateBehavior = resp.Crawler.SchemaChangePolicy.UpdateBehavior
		}
		cr.Spec.ForProvider.SchemaChangePolicy = f15
	} else {
		cr.Spec.ForProvider.SchemaChangePolicy = nil
	}
	if resp.Crawler.State != nil {
		cr.Status.AtProvider.State = resp.Crawler.State
	} else {
		cr.Status.AtProvider.State = nil
	}
	if resp.Crawler.TablePrefix != nil {
		cr.Spec.ForProvider.TablePrefix = resp.Crawler.TablePrefix
	} else {
		cr.Spec.ForProvider.TablePrefix = nil
	}
	if resp.Crawler.Version != nil {
		cr.Status.AtProvider.Version = resp.Crawler.Version
	} else {
		cr.Status.AtProvider.Version = nil
	}

	return cr
}

// GenerateCreateCrawlerInput returns a create input.
func GenerateCreateCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.CreateCrawlerInput {
	res := &svcsdk.CreateCrawlerInput{}

	if cr.Spec.ForProvider.Configuration != nil {
		res.SetConfiguration(*cr.Spec.ForProvider.Configuration)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.LakeFormationConfiguration != nil {
		f2 := &svcsdk.LakeFormationConfiguration{}
		if cr.Spec.ForProvider.LakeFormationConfiguration.AccountID != nil {
			f2.SetAccountId(*cr.Spec.ForProvider.LakeFormationConfiguration.AccountID)
		}
		if cr.Spec.ForProvider.LakeFormationConfiguration.UseLakeFormationCredentials != nil {
			f2.SetUseLakeFormationCredentials(*cr.Spec.ForProvider.LakeFormationConfiguration.UseLakeFormationCredentials)
		}
		res.SetLakeFormationConfiguration(f2)
	}
	if cr.Spec.ForProvider.LineageConfiguration != nil {
		f3 := &svcsdk.LineageConfiguration{}
		if cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings != nil {
			f3.SetCrawlerLineageSettings(*cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings)
		}
		res.SetLineageConfiguration(f3)
	}
	if cr.Spec.ForProvider.RecrawlPolicy != nil {
		f4 := &svcsdk.RecrawlPolicy{}
		if cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior != nil {
			f4.SetRecrawlBehavior(*cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior)
		}
		res.SetRecrawlPolicy(f4)
	}
	if cr.Spec.ForProvider.Schedule != nil {
		res.SetSchedule(*cr.Spec.ForProvider.Schedule)
	}
	if cr.Spec.ForProvider.SchemaChangePolicy != nil {
		f6 := &svcsdk.SchemaChangePolicy{}
		if cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior != nil {
			f6.SetDeleteBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior)
		}
		if cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior != nil {
			f6.SetUpdateBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior)
		}
		res.SetSchemaChangePolicy(f6)
	}
	if cr.Spec.ForProvider.TablePrefix != nil {
		res.SetTablePrefix(*cr.Spec.ForProvider.TablePrefix)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range cr.Spec.ForProvider.Tags {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		res.SetTags(f8)
	}

	return res
}

// GenerateUpdateCrawlerInput returns an update input.
func GenerateUpdateCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.UpdateCrawlerInput {
	res := &svcsdk.UpdateCrawlerInput{}

	if cr.Spec.ForProvider.Configuration != nil {
		res.SetConfiguration(*cr.Spec.ForProvider.Configuration)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.LakeFormationConfiguration != nil {
		f5 := &svcsdk.LakeFormationConfiguration{}
		if cr.Spec.ForProvider.LakeFormationConfiguration.AccountID != nil {
			f5.SetAccountId(*cr.Spec.ForProvider.LakeFormationConfiguration.AccountID)
		}
		if cr.Spec.ForProvider.LakeFormationConfiguration.UseLakeFormationCredentials != nil {
			f5.SetUseLakeFormationCredentials(*cr.Spec.ForProvider.LakeFormationConfiguration.UseLakeFormationCredentials)
		}
		res.SetLakeFormationConfiguration(f5)
	}
	if cr.Spec.ForProvider.LineageConfiguration != nil {
		f6 := &svcsdk.LineageConfiguration{}
		if cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings != nil {
			f6.SetCrawlerLineageSettings(*cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings)
		}
		res.SetLineageConfiguration(f6)
	}
	if cr.Spec.ForProvider.RecrawlPolicy != nil {
		f8 := &svcsdk.RecrawlPolicy{}
		if cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior != nil {
			f8.SetRecrawlBehavior(*cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior)
		}
		res.SetRecrawlPolicy(f8)
	}
	if cr.Spec.ForProvider.Schedule != nil {
		res.SetSchedule(*cr.Spec.ForProvider.Schedule)
	}
	if cr.Spec.ForProvider.SchemaChangePolicy != nil {
		f11 := &svcsdk.SchemaChangePolicy{}
		if cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior != nil {
			f11.SetDeleteBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior)
		}
		if cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior != nil {
			f11.SetUpdateBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior)
		}
		res.SetSchemaChangePolicy(f11)
	}
	if cr.Spec.ForProvider.TablePrefix != nil {
		res.SetTablePrefix(*cr.Spec.ForProvider.TablePrefix)
	}

	return res
}

// GenerateDeleteCrawlerInput returns a deletion input.
func GenerateDeleteCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.DeleteCrawlerInput {
	res := &svcsdk.DeleteCrawlerInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "EntityNotFoundException"
}
