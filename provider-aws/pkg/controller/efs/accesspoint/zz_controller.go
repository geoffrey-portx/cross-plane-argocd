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

package accesspoint

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/efs"
	svcsdk "github.com/aws/aws-sdk-go/service/efs"
	svcsdkapi "github.com/aws/aws-sdk-go/service/efs/efsiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/efs/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an AccessPoint resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create AccessPoint in AWS"
	errUpdate        = "cannot update AccessPoint in AWS"
	errDescribe      = "failed to describe AccessPoint"
	errDelete        = "failed to delete AccessPoint"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.AccessPoint) (managed.TypedExternalClient[*svcapitypes.AccessPoint], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.AccessPoint) (managed.ExternalObservation, error) {
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeAccessPointsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeAccessPointsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.AccessPoints) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateAccessPoint(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, cr *svcapitypes.AccessPoint) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateAccessPointInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateAccessPointWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.AccessPointArn != nil {
		cr.Status.AtProvider.AccessPointARN = resp.AccessPointArn
	} else {
		cr.Status.AtProvider.AccessPointARN = nil
	}
	if resp.AccessPointId != nil {
		cr.Status.AtProvider.AccessPointID = resp.AccessPointId
	} else {
		cr.Status.AtProvider.AccessPointID = nil
	}
	if resp.ClientToken != nil {
		cr.Status.AtProvider.ClientToken = resp.ClientToken
	} else {
		cr.Status.AtProvider.ClientToken = nil
	}
	if resp.FileSystemId != nil {
		cr.Status.AtProvider.FileSystemID = resp.FileSystemId
	} else {
		cr.Status.AtProvider.FileSystemID = nil
	}
	if resp.LifeCycleState != nil {
		cr.Status.AtProvider.LifeCycleState = resp.LifeCycleState
	} else {
		cr.Status.AtProvider.LifeCycleState = nil
	}
	if resp.Name != nil {
		cr.Status.AtProvider.Name = resp.Name
	} else {
		cr.Status.AtProvider.Name = nil
	}
	if resp.OwnerId != nil {
		cr.Status.AtProvider.OwnerID = resp.OwnerId
	} else {
		cr.Status.AtProvider.OwnerID = nil
	}
	if resp.PosixUser != nil {
		f7 := &svcapitypes.PosixUser{}
		if resp.PosixUser.Gid != nil {
			f7.Gid = resp.PosixUser.Gid
		}
		if resp.PosixUser.SecondaryGids != nil {
			f7f1 := []*int64{}
			for _, f7f1iter := range resp.PosixUser.SecondaryGids {
				var f7f1elem int64
				f7f1elem = *f7f1iter
				f7f1 = append(f7f1, &f7f1elem)
			}
			f7.SecondaryGids = f7f1
		}
		if resp.PosixUser.Uid != nil {
			f7.Uid = resp.PosixUser.Uid
		}
		cr.Spec.ForProvider.PosixUser = f7
	} else {
		cr.Spec.ForProvider.PosixUser = nil
	}
	if resp.RootDirectory != nil {
		f8 := &svcapitypes.RootDirectory{}
		if resp.RootDirectory.CreationInfo != nil {
			f8f0 := &svcapitypes.CreationInfo{}
			if resp.RootDirectory.CreationInfo.OwnerGid != nil {
				f8f0.OwnerGid = resp.RootDirectory.CreationInfo.OwnerGid
			}
			if resp.RootDirectory.CreationInfo.OwnerUid != nil {
				f8f0.OwnerUid = resp.RootDirectory.CreationInfo.OwnerUid
			}
			if resp.RootDirectory.CreationInfo.Permissions != nil {
				f8f0.Permissions = resp.RootDirectory.CreationInfo.Permissions
			}
			f8.CreationInfo = f8f0
		}
		if resp.RootDirectory.Path != nil {
			f8.Path = resp.RootDirectory.Path
		}
		cr.Spec.ForProvider.RootDirectory = f8
	} else {
		cr.Spec.ForProvider.RootDirectory = nil
	}
	if resp.Tags != nil {
		f9 := []*svcapitypes.Tag{}
		for _, f9iter := range resp.Tags {
			f9elem := &svcapitypes.Tag{}
			if f9iter.Key != nil {
				f9elem.Key = f9iter.Key
			}
			if f9iter.Value != nil {
				f9elem.Value = f9iter.Value
			}
			f9 = append(f9, f9elem)
		}
		cr.Spec.ForProvider.Tags = f9
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.AccessPoint) (managed.ExternalUpdate, error) {
	return e.update(ctx, cr)

}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.AccessPoint) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteAccessPointInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteAccessPointWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EFSAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		update:         nopUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EFSAPI
	preObserve     func(context.Context, *svcapitypes.AccessPoint, *svcsdk.DescribeAccessPointsInput) error
	postObserve    func(context.Context, *svcapitypes.AccessPoint, *svcsdk.DescribeAccessPointsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.AccessPoint, *svcsdk.DescribeAccessPointsOutput) *svcsdk.DescribeAccessPointsOutput
	lateInitialize func(*svcapitypes.AccessPointParameters, *svcsdk.DescribeAccessPointsOutput) error
	isUpToDate     func(context.Context, *svcapitypes.AccessPoint, *svcsdk.DescribeAccessPointsOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.AccessPoint, *svcsdk.CreateAccessPointInput) error
	postCreate     func(context.Context, *svcapitypes.AccessPoint, *svcsdk.CreateAccessPointOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.AccessPoint, *svcsdk.DeleteAccessPointInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.AccessPoint, *svcsdk.DeleteAccessPointOutput, error) (managed.ExternalDelete, error)
	update         func(context.Context, *svcapitypes.AccessPoint) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.AccessPoint, *svcsdk.DescribeAccessPointsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.AccessPoint, _ *svcsdk.DescribeAccessPointsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.AccessPoint, list *svcsdk.DescribeAccessPointsOutput) *svcsdk.DescribeAccessPointsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.AccessPointParameters, *svcsdk.DescribeAccessPointsOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.AccessPoint, *svcsdk.DescribeAccessPointsOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.AccessPoint, *svcsdk.CreateAccessPointInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.AccessPoint, _ *svcsdk.CreateAccessPointOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.AccessPoint, *svcsdk.DeleteAccessPointInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.AccessPoint, _ *svcsdk.DeleteAccessPointOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopUpdate(context.Context, *svcapitypes.AccessPoint) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
