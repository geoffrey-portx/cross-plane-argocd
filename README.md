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
