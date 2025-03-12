#!/bin/bash

# Set environment value
ENVIRONMENT_TAG="dev"

# Function to find EC2 instances
find_vpcs() {
  echo "Finding VPCs with tag 'Environment=${ENVIRONMENT_TAG}'..."
  aws ec2 describe-vpcs --filters "Name=tag:Environment,Values=${ENVIRONMENT_TAG}" --query 'Vpcs[*].VpcId' --output text
}

# Function to find VPC endpoints
find_vpc_endpoints() {
  echo "Finding VPC Endpoints with tag 'Environment=${ENVIRONMENT_TAG}'..."
  aws ec2 describe-vpc-endpoints --filters "Name=tag:Environment,Values=${ENVIRONMENT_TAG}" --query 'VpcEndpoints[*].VpcEndpointId' --output text
}

# Function to find EKS clusters
find_eks_clusters() {
  echo "Finding EKS clusters with tag 'Environment=${ENVIRONMENT_TAG}'..."
  aws eks list-clusters --query 'clusters' --output text
}

# Function to find S3 buckets
find_s3_buckets() {
  echo "Finding S3 buckets with tag 'Environment=${ENVIRONMENT_TAG}'..."
  aws s3api list-buckets --query "Buckets[?contains(Name, '${ENVIRONMENT_TAG}')].Name" --output text
}

# Function to find CloudWatch log groups
find_cloudwatch_log_groups() {
  echo "Finding CloudWatch log groups with tag 'Environment=${ENVIRONMENT_TAG}'..."
  aws logs describe-log-groups --query 'logGroups[*].logGroupName' --output text
}

# Find resources using AWS CLI
vpcs=$(find_vpcs)
vpc_endpoints=$(find_vpc_endpoints)
eks_clusters=$(find_eks_clusters)
s3_buckets=$(find_s3_buckets)
cloudwatch_logs=$(find_cloudwatch_log_groups)

# Print found resources
echo "Found VPCs: $vpcs"
echo "Found VPC Endpoints: $vpc_endpoints"
echo "Found EKS Clusters: $eks_clusters"
echo "Found S3 Buckets: $s3_buckets"
echo "Found CloudWatch Logs: $cloudwatch_logs"

# Function to generate Crossplane YAML for VPCs
generate_vpc_yaml() {
  for vpc_id in $1; do
    cat <<EOL > "${vpc_id}-vpc.yaml"
apiVersion: ec2.aws.crossplane.io/v1alpha1
kind: VPC
metadata:
  name: ${vpc_id}
spec:
  forProvider:
    vpcId: ${vpc_id}
  providerConfigRef:
    name: aws-provider
  crossplane.io/external-name: ${vpc_id}
  managementPolicies:
    - Observe
EOL
    echo "Generated YAML for VPC: ${vpc_id}"
  done
}

# Function to generate Crossplane YAML for VPC endpoints
generate_vpc_endpoint_yaml() {
  for vpc_endpoint_id in $1; do
    cat <<EOL > "${vpc_endpoint_id}-vpc-endpoint.yaml"
apiVersion: ec2.aws.crossplane.io/v1alpha1
kind: VpcEndpoint
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
EOL
    echo "Generated YAML for VPC Endpoint: ${vpc_endpoint_id}"
  done
}

# Function to generate Crossplane YAML for EKS clusters
generate_eks_yaml() {
  for eks_cluster_name in $1; do
    cat <<EOL > "${eks_cluster_name}-eks.yaml"
apiVersion: eks.aws.crossplane.io/v1alpha1
kind: Cluster
metadata:
  name: ${eks_cluster_name}
spec:
  forProvider:
    clusterName: ${eks_cluster_name}
  providerConfigRef:
    name: aws-provider
  crossplane.io/external-name: ${eks_cluster_name}
  managementPolicies:
    - Observe
EOL
    echo "Generated YAML for EKS Cluster: ${eks_cluster_name}"
  done
}

# Function to generate Crossplane YAML for S3 buckets
generate_s3_yaml() {
  for bucket_name in $1; do
    cat <<EOL > "${bucket_name}-s3.yaml"
apiVersion: s3.aws.crossplane.io/v1alpha1
kind: Bucket
metadata:
  name: ${bucket_name}
spec:
  forProvider:
    bucketName: ${bucket_name}
  providerConfigRef:
    name: aws-provider
  crossplane.io/external-name: ${bucket_name}
  managementPolicies:
    - Observe
EOL
    echo "Generated YAML for S3 bucket: ${bucket_name}"
  done
}

# Function to generate Crossplane YAML for CloudWatch log groups
generate_cloudwatch_yaml() {
  for log_group_name in $1; do
    cat <<EOL > "${log_group_name}-cloudwatch.yaml"
apiVersion: logs.aws.crossplane.io/v1alpha1
kind: LogGroup
metadata:
  name: ${log_group_name}
spec:
  forProvider:
    logGroupName: ${log_group_name}
  providerConfigRef:
    name: aws-provider
  crossplane.io/external-name: ${log_group_name}
  managementPolicies:
    - Observe
EOL
    echo "Generated YAML for CloudWatch Log Group: ${log_group_name}"
  done
}

# Generate YAML files for the found resources
generate_vpc_yaml "$vpcs"
generate_vpc_endpoint_yaml "$vpc_endpoints"
generate_eks_yaml "$eks_clusters"
generate_s3_yaml "$s3_buckets"
generate_cloudwatch_yaml "$cloudwatch_logs"

# Helm Chart Creation

# Create Helm chart directory structure
mkdir -p crossplane-helm-chart/templates
cd crossplane-helm-chart

# Create the Chart.yaml file
cat <<EOL > Chart.yaml
apiVersion: v2
name: crossplane-aws-resources
description: A Helm chart for deploying AWS resources in Crossplane
version: 0.1.0
EOL

# Move the generated YAMLs into the templates directory
cp ../*-vpc.yaml ./templates/
cp ../*-vpc-endpoint.yaml ./templates/
cp ../*-eks.yaml ./templates/
cp ../*-s3.yaml ./templates/
cp ../*-cloudwatch.yaml ./templates/

# Package the Helm chart
helm package .
echo "Helm chart created: crossplane-aws-resources-0.1.0.tgz"
