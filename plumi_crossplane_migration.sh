#!/bin/bash

# Check if multiple stacks are specified
stacks=$(pulumi stack ls --json | jq -r '.[] | .name')

if [ -z "$stacks" ]; then
    echo "No Pulumi stacks found. Please create a stack first."
    exit 1
fi

# Loop through each stack
for stack in $stacks; do
    echo "Processing stack: $stack"
    
    # Set the active stack
    pulumi stack select "$stack" > /dev/null
    
    # Get the Pulumi stack outputs in JSON format
    stack_output=$(pulumi stack output --json)

    # Extract outputs from the Pulumi stack
    vpc_id=$(echo "$stack_output" | jq -r '.vpcId')
    vpc_cidr_block=$(echo "$stack_output" | jq -r '.vpcCidrBlock')
    eks_cluster_name=$(echo "$stack_output" | jq -r '.eksClusterName')
    eks_vpc_id=$(echo "$stack_output" | jq -r '.eksVpcId')
    eks_role_arn=$(echo "$stack_output" | jq -r '.eksClusterRoleArn')
    s3_bucket_name=$(echo "$stack_output" | jq -r '.s3BucketName')
    s3_region=$(echo "$stack_output" | jq -r '.s3Region')
    endpoint_service_name=$(echo "$stack_output" | jq -r '.endpointServiceName')
    endpoint_vpc_id=$(echo "$stack_output" | jq -r '.endpointVpcId')
    cloudwatch_log_group_name=$(echo "$stack_output" | jq -r '.cloudwatchLogGroupName')
    sqs_queue_name=$(echo "$stack_output" | jq -r '.sqsQueueName')
    karpenter_version=$(echo "$stack_output" | jq -r '.karpenterVersion')

    # Create folder for the current stack
    stack_folder="$stack"
    mkdir -p "$stack_folder/templates"
    echo "Created folder: $stack_folder"

    # Define the AWS Provider configuration for Crossplane
    provider_config="apiVersion: aws.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: my-aws-provider
spec:
  credentialsSecretRef:
    name: my-aws-creds
    namespace: crossplane-system
    key: credentials.json  # Replace with your actual credentials secret
  region: us-east-1  # Default AWS region
"

    # Write ProviderConfig YAML to file
    echo "$provider_config" > "$stack_folder/provider-config.yaml"

    # Generate and save each resource YAML

    # VPC
    vpc_yaml="apiVersion: aws.crossplane.io/v1alpha1
kind: VPC
metadata:
  name: my-vpc
spec:
  providerConfigRef:
    name: my-aws-provider
  forProvider:
    cidrBlock: $vpc_cidr_block
    enableDnsHostnames: true
    enableDnsSupport: true
    instanceTenancy: default
    tags:
      Name: my-vpc
      Environment: prod
  writeConnectionSecretToRef:
    name: my-vpc-secret
    namespace: crossplane-system
"
    echo "$vpc_yaml" > "$stack_folder/templates/vpc.yaml"

    # EKS
    eks_yaml="apiVersion: eks.aws.crossplane.io/v1alpha1
kind: EKSCluster
metadata:
  name: $eks_cluster_name
spec:
  providerConfigRef:
    name: my-aws-provider
  forProvider:
    version: '1.21'
    roleArn: $eks_role_arn
    resourcesVpcConfig:
      subnetIds:
        - subnet-abc123  # Replace with your actual subnet IDs
        - subnet-def456
      securityGroupIds:
        - sg-123456  # Replace with your actual security group IDs
      endpointPublicAccess: true
      endpointPrivateAccess: true
    desiredSize: 3
    minSize: 1
    maxSize: 5
  writeConnectionSecretToRef:
    name: $eks_cluster_name-secret
    namespace: crossplane-system
"
    echo "$eks_yaml" > "$stack_folder/templates/eks.yaml"

    # S3
    s3_yaml="apiVersion: s3.aws.crossplane.io/v1alpha1
kind: Bucket
metadata:
  name: $s3_bucket_name
spec:
  providerConfigRef:
    name: my-aws-provider
  forProvider:
    acl: private
    region: $s3_region
    versioning:
      enabled: true
    lifecycleRules:
      - enabled: true
        id: 'expire-old-objects'
        expiration:
          days: 365
    tags:
      Name: $s3_bucket_name
      Environment: dev
  writeConnectionSecretToRef:
    name: $s3_bucket_name-secret
    namespace: crossplane-system
"
    echo "$s3_yaml" > "$stack_folder/templates/s3.yaml"

    # Endpoint
    endpoint_yaml="apiVersion: vpc.aws.crossplane.io/v1alpha1
kind: VPCPrivateLink
metadata:
  name: my-endpoint-service
spec:
  providerConfigRef:
    name: my-aws-provider
  forProvider:
    serviceName: $endpoint_service_name
    vpcId: $endpoint_vpc_id
    subnetIds:
      - subnet-abc123  # Replace with your actual subnet IDs
    securityGroupIds:
      - sg-123456  # Replace with your actual security group IDs
    privateDnsEnabled: true
    tags:
      Name: my-endpoint-service
      Environment: prod
  writeConnectionSecretToRef:
    name: my-endpoint-service-secret
    namespace: crossplane-system
"
    echo "$endpoint_yaml" > "$stack_folder/templates/endpoint.yaml"

    # CloudWatch
    cloudwatch_yaml="apiVersion: logs.aws.crossplane.io/v1alpha1
kind: LogGroup
metadata:
  name: $cloudwatch_log_group_name
spec:
  providerConfigRef:
    name: my-aws-provider
  forProvider:
    logGroupName: $cloudwatch_log_group_name
    retentionInDays: 30  # Retention period for the logs
    tags:
      Name: $cloudwatch_log_group_name
      Environment: prod
  writeConnectionSecretToRef:
    name: $cloudwatch_log_group_name-secret
    namespace: crossplane-system
"
    echo "$cloudwatch_yaml" > "$stack_folder/templates/cloudwatch-log-group.yaml"

    # SQS
    sqs_yaml="apiVersion: sqs.aws.crossplane.io/v1alpha1
kind: Queue
metadata:
  name: $sqs_queue_name
spec:
  providerConfigRef:
    name: my-aws-provider
  forProvider:
    queueName: $sqs_queue_name
    delaySeconds: 0  # Default delay for the queue
    maximumMessageSize: 262144  # Max message size
    messageRetentionSeconds: 345600  # Retention period
    tags:
      Name: $sqs_queue_name
      Environment: prod
  writeConnectionSecretToRef:
    name: $sqs_queue_name-secret
    namespace: crossplane-system
"
    echo "$sqs_yaml" > "$stack_folder/templates/sqs-queue.yaml"

    # Karpenter
    karpenter_yaml="apiVersion: karpenter.sh/v1alpha5
kind: Provisioner
metadata:
  name: karpenter-provisioner
spec:
  provider:
    aws:
      capacityTypes: [on-demand]  # Available capacity types (on-demand, spot)
  requirements:
    - key: "kubernetes.io/arch"
      operator: In
      values: ["amd64"]
    - key: "kubernetes.io/os"
      operator: In
      values: ["linux"]
  limits:
    resources:
      cpu: 1000  # CPU resources limit
      memory: 1000Gi  # Memory resources limit
  ttlSecondsAfterEmpty: 30  # TTL for provisioners when idle
"
    echo "$karpenter_yaml" > "$stack_folder/templates/karpenter-provisioner.yaml"

    # Create Helm chart structure
    mkdir -p "$stack_folder/templates"
    cat <<EOF > "$stack_folder/Chart.yaml"
apiVersion: v2
name: $stack
description: A Crossplane configuration for AWS resources.
version: 0.1.0
EOF

    cat <<EOF > "$stack_folder/values.yaml"
# Define values for this stack's configuration
EOF

    echo "Helm chart created in $stack_folder"
done

echo "All stacks processed and Helm charts created successfully."

