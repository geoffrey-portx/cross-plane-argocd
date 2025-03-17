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

package identitypool

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/cognitoidentity"
	svcsdk "github.com/aws/aws-sdk-go/service/cognitoidentity"
	svcsdkapi "github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/cognitoidentity/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an IdentityPool resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create IdentityPool in AWS"
	errUpdate        = "cannot update IdentityPool in AWS"
	errDescribe      = "failed to describe IdentityPool"
	errDelete        = "failed to delete IdentityPool"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.IdentityPool)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.IdentityPool)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeIdentityPoolInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeIdentityPoolWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateIdentityPool(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.IdentityPool)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateIdentityPoolInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateIdentityPoolWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.AllowClassicFlow != nil {
		cr.Spec.ForProvider.AllowClassicFlow = resp.AllowClassicFlow
	} else {
		cr.Spec.ForProvider.AllowClassicFlow = nil
	}
	if resp.AllowUnauthenticatedIdentities != nil {
		cr.Status.AtProvider.AllowUnauthenticatedIDentities = resp.AllowUnauthenticatedIdentities
	} else {
		cr.Status.AtProvider.AllowUnauthenticatedIDentities = nil
	}
	if resp.CognitoIdentityProviders != nil {
		f2 := []*svcapitypes.Provider{}
		for _, f2iter := range resp.CognitoIdentityProviders {
			f2elem := &svcapitypes.Provider{}
			if f2iter.ClientId != nil {
				f2elem.ClientID = f2iter.ClientId
			}
			if f2iter.ProviderName != nil {
				f2elem.ProviderName = f2iter.ProviderName
			}
			if f2iter.ServerSideTokenCheck != nil {
				f2elem.ServerSideTokenCheck = f2iter.ServerSideTokenCheck
			}
			f2 = append(f2, f2elem)
		}
		cr.Status.AtProvider.CognitoIdentityProviders = f2
	} else {
		cr.Status.AtProvider.CognitoIdentityProviders = nil
	}
	if resp.DeveloperProviderName != nil {
		cr.Spec.ForProvider.DeveloperProviderName = resp.DeveloperProviderName
	} else {
		cr.Spec.ForProvider.DeveloperProviderName = nil
	}
	if resp.IdentityPoolId != nil {
		cr.Status.AtProvider.IdentityPoolID = resp.IdentityPoolId
	} else {
		cr.Status.AtProvider.IdentityPoolID = nil
	}
	if resp.IdentityPoolName != nil {
		cr.Spec.ForProvider.IdentityPoolName = resp.IdentityPoolName
	} else {
		cr.Spec.ForProvider.IdentityPoolName = nil
	}
	if resp.IdentityPoolTags != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range resp.IdentityPoolTags {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		cr.Spec.ForProvider.IdentityPoolTags = f6
	} else {
		cr.Spec.ForProvider.IdentityPoolTags = nil
	}
	if resp.OpenIdConnectProviderARNs != nil {
		f7 := []*string{}
		for _, f7iter := range resp.OpenIdConnectProviderARNs {
			var f7elem string
			f7elem = *f7iter
			f7 = append(f7, &f7elem)
		}
		cr.Status.AtProvider.OpenIDConnectProviderARNs = f7
	} else {
		cr.Status.AtProvider.OpenIDConnectProviderARNs = nil
	}
	if resp.SamlProviderARNs != nil {
		f8 := []*string{}
		for _, f8iter := range resp.SamlProviderARNs {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		cr.Spec.ForProvider.SamlProviderARNs = f8
	} else {
		cr.Spec.ForProvider.SamlProviderARNs = nil
	}
	if resp.SupportedLoginProviders != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range resp.SupportedLoginProviders {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		cr.Spec.ForProvider.SupportedLoginProviders = f9
	} else {
		cr.Spec.ForProvider.SupportedLoginProviders = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.IdentityPool)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateIdentityPoolInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateIdentityPoolWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.IdentityPool)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteIdentityPoolInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteIdentityPoolWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.CognitoIdentityAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.CognitoIdentityAPI
	preObserve     func(context.Context, *svcapitypes.IdentityPool, *svcsdk.DescribeIdentityPoolInput) error
	postObserve    func(context.Context, *svcapitypes.IdentityPool, *svcsdk.IdentityPool, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.IdentityPoolParameters, *svcsdk.IdentityPool) error
	isUpToDate     func(*svcapitypes.IdentityPool, *svcsdk.IdentityPool) (bool, error)
	preCreate      func(context.Context, *svcapitypes.IdentityPool, *svcsdk.CreateIdentityPoolInput) error
	postCreate     func(context.Context, *svcapitypes.IdentityPool, *svcsdk.IdentityPool, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.IdentityPool, *svcsdk.DeleteIdentityPoolInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.IdentityPool, *svcsdk.DeleteIdentityPoolOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.IdentityPool, *svcsdk.IdentityPool) error
	postUpdate     func(context.Context, *svcapitypes.IdentityPool, *svcsdk.IdentityPool, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.IdentityPool, *svcsdk.DescribeIdentityPoolInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.IdentityPool, _ *svcsdk.IdentityPool, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.IdentityPoolParameters, *svcsdk.IdentityPool) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.IdentityPool, *svcsdk.IdentityPool) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.IdentityPool, *svcsdk.CreateIdentityPoolInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.IdentityPool, _ *svcsdk.IdentityPool, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.IdentityPool, *svcsdk.DeleteIdentityPoolInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.IdentityPool, _ *svcsdk.DeleteIdentityPoolOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.IdentityPool, *svcsdk.IdentityPool) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.IdentityPool, _ *svcsdk.IdentityPool, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
