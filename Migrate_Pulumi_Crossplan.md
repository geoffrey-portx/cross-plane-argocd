## Take inventory of a tenants AWS resources

aws resourcegroupstaggingapi get-resources --tag-filters Key=Environmemt,Values=dev Key=TagKey2,Values=TagValue2

## Migrate AWS resources to Crossplane manifests using Observer management policy

```plaintext
apiVersion: ec2.aws.crossplane.io/v1alpha1
kind: <Kind type> example VpcEndpoint
metadata:
  name: ${vpc_endpoint_id}
spec:
  forProvider:
    vpcEndpointId: ${vpc_endpoint_id}
  providerConfigRef:
    name: aws-provider
  crossplane.io/external-name: ${vpc_endpoint_id}
  managementPolicies:
    - Observe
```

## Create Crossplane manifests with all parameters

kubectl get 'crossplane kind' -o yaml

## Deploy Crossplane

Deploy Crossplane in a one node EKS cluster per tenant, also deploy Argocd to the EKS cluster

## Deploy Crossplane manifests to the EKS cluster

Crossplane will now be Observing the tenants eks cluster but not managing it

## Put tenant Crossplane manifests in a github repo

Put Crossplane manifests in a github repo and have argocd sync the repo/path

## Determine what resources we want to remove to Pulumi and have Crossplane manage

Change the managementPoliy from 'Observer' to [*]

## Repeat process for next tenant environment and also for all tenants

Repeat process for QA, Prod etc. for each tenant

## Create GHA pipeline that performs the following

1) A feature branch pull request is deployed, the feature gets deployed to DEV
2) The feature gets merged to main the feature deployed to dev is removed
3) A release tag is created the release gets deployed to QA
4) A quality gets is reviewed the release gets deployed to Prod




