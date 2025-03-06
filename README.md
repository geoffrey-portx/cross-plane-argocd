# AWS Infrastructure with Crossplane, Helm, ArgoCD, and GitHub Actions

This repository provides a comprehensive setup to deploy AWS infrastructure using **Crossplane**, **Helm**, **ArgoCD**, and **GitHub Actions**. The infrastructure includes various AWS resources such as **EKS (Elastic Kubernetes Service)**, **S3** buckets, **ECR (Elastic Container Registry)**, **CloudWatch** logging, **VPC**, and **VPC endpoints**. The deployment pipeline is fully automated with GitHub Actions for CI/CD, ensuring seamless deployments across **dev**, **qa**, and **prod** environments.

## Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Prerequisites](#prerequisites)
4. [Directory Structure](#directory-structure)
5. [Configuration](#configuration)
   - [Crossplane Configuration](#crossplane-configuration)
   - [GitHub Actions Pipeline Configuration](#github-actions-pipeline-configuration)
6. [Deployment Process](#deployment-process)
7. [Environment Specific Configuration](#environment-specific-configuration)
8. [Quality Gates](#quality-gates)
9. [Deploying with ArgoCD](#deploying-with-argocd)
10. [References](#references)

## Overview

This repository automates the deployment of AWS resources using **Crossplane**. The resources include EKS clusters, S3 buckets, ECR repositories, VPC, VPC Endpoints, and CloudWatch logs. The deployment pipeline leverages **Helm** for packaging and managing infrastructure as code, **ArgoCD** for GitOps-based continuous delivery, and **GitHub Actions** for CI/CD automation.

## Architecture

The architecture consists of several key components:

- **Crossplane**: Used for managing AWS infrastructure declaratively. Crossplane manifests are used to create and manage AWS resources such as EKS clusters, VPC, S3 buckets, ECR repositories, and CloudWatch logs.
  
- **Helm**: The Crossplane resources are packaged as a Helm chart to enable versioning, templating, and consistent deployments.

- **ArgoCD**: A GitOps tool that ensures continuous synchronization of Kubernetes resources. ArgoCD automatically deploys infrastructure changes to Kubernetes clusters when changes are made to the Git repository.

- **GitHub Actions**: Automates the CI/CD pipeline, including packaging the Helm chart, applying the manifests, and deploying to different environments (dev, qa, prod).

## Prerequisites

- **Crossplane** installed and configured on your Kubernetes cluster.
- **Helm** installed and configured to work with your Kubernetes cluster.
- **ArgoCD** installed and configured on your Kubernetes cluster.
- **GitHub Actions** enabled for your repository.
- AWS access credentials with the necessary permissions for deploying the infrastructure (use AWS IAM roles and secrets for security).
- A Kubernetes cluster with the necessary namespaces and permissions for ArgoCD.

## Directory Structure

The directory structure of the repository is as follows:

```plaintext
.
├── crossplane-helm-chart/                 # Contains Helm chart for Crossplane resources
│   ├── templates/                         # Crossplane resource templates
│   │   ├── crossplane-cloudwatch.yaml     # CloudWatch Log Group and Stream
│   │   ├── crossplane-vpc.yaml           # VPC resources (VPC, Subnet, RouteTable)
│   │   ├── crossplane-vpc-endpoints.yaml # VPC endpoints
│   │   ├── crossplane-eks.yaml           # EKS cluster
│   │   ├── crossplane-s3.yaml            # S3 Bucket
│   │   ├── crossplane-ecr.yaml           # ECR Repository
│   │   └── crossplane-eks-iam-roles.yaml # EKS IAM roles
│   ├── values-dev.yaml                   # Dev environment values
│   ├── values-qa.yaml                    # QA environment values
│   ├── values-prod.yaml                  # Production environment values
├── .github/
│   └── workflows/
│       └── deploy.yaml                   # GitHub Actions pipeline configuration
├── argocd-applications/                  # ArgoCD Application manifests
│   └── argocd-application.yaml           # ArgoCD App for the Crossplane resources
└── README.md                             # This file


```
## Configuration
Crossplane Configuration
The Helm chart (crossplane-helm-chart) defines all the Crossplane-managed AWS resources, including EKS, VPC, S3, ECR, CloudWatch, and VPC Endpoints. These resources are templated using values-<env>.yaml files, which define environment-specific configurations (e.g., CIDR block for VPC, region, S3 bucket names, etc.).

The values-<env>.yaml files contain:

AWS region and account-specific configurations.
VPC configuration (CIDR, subnet IDs, and route tables).
CloudWatch logging configurations.
S3, ECR, and EKS resource configurations.
These configurations allow the creation of these resources in AWS, with different parameters for each environment (dev, qa, and prod).

GitHub Actions Pipeline Configuration
The GitHub Actions pipeline (deploy.yaml) automates the deployment of Crossplane resources. The pipeline is triggered on pull requests or when a release tag is created. It performs the following tasks:

Setup: Checks out the code, sets up AWS credentials, and packages the Helm chart.
Deployment: Deploys resources to the target environment (dev, qa, prod) using ArgoCD. Each environment has quality gates to ensure proper validation before moving to the next stage.
GitHub Actions Workflow Overview:
Feature Branch PR: The pipeline runs on a pull request to a feature branch. It packages the Helm chart for dev environment and pushes the packaged chart to GitHub Container Registry (GHCR).
Release Tag: When a release tag is created (e.g., v1.0.0), the pipeline runs for all environments, deploying to dev, then qa, and finally prod, with quality gates in between.
## Environment Specific Configuration
The environment-specific configuration for dev, qa, and prod is stored in values-dev.yaml, values-qa.yaml, and values-prod.yaml, respectively. These files specify the following:

AWS region.
VPC CIDR blocks, subnet IDs, route tables.
EKS cluster configuration (including IAM roles).
CloudWatch log group and stream.
S3 bucket and ECR repository.

## Quality Gates
Quality gates are implemented as part of the GitHub Actions pipeline to ensure each environment is properly validated before moving to the next. If the validation fails (e.g., the application doesn't deploy successfully), the pipeline stops, preventing the deployment to the next environment.

## Deployment Process
Clone this repository:
Clone the repository to your local machine or directly on GitHub.

Configure AWS Credentials:
Ensure that your AWS credentials are properly configured on your local machine or GitHub repository secrets for use with the GitHub Actions pipeline.

Install and Configure Crossplane:
Install Crossplane on your Kubernetes cluster, configure it to use AWS as a provider, and apply the necessary provider configurations (like AWS credentials).

Set up ArgoCD:
Ensure ArgoCD is installed on your Kubernetes cluster. Create an ArgoCD application (argocd-application.yaml) to synchronize the Crossplane resources with your Kubernetes cluster.

Run GitHub Actions:

Create a pull request with changes, or tag a release.
The GitHub Actions pipeline will package the Helm chart and deploy it to the target environment.
After deployment, ArgoCD will sync and deploy the resources in the Kubernetes cluster using Crossplane.

## Deploying with ArgoCD
ArgoCD will be responsible for applying the manifests defined in the Helm chart. It syncs the desired state from the Git repository to the Kubernetes cluster and ensures the infrastructure resources (EKS, S3, ECR, CloudWatch, etc.) are created according to the manifests.

ArgoCD Configuration
The ArgoCD application is defined in the argocd-application.yaml. The application will:

Monitor the repository for changes.
Deploy resources from the crossplane-helm-chart directory.
Automatically synchronize the resources whenever a change is made to the repository.

## References
Crossplane Documentation
ArgoCD Documentation
GitHub Actions Documentation