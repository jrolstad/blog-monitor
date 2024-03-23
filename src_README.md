# Blog Monitor
The components in Blog Monitor poll subscribed blogs and notify when a new post is detected. Components are primarily written in Golang and use AWS resources for processing, data storage, and notification capabilities.

# Requirements
* golang 1.21 or higher

# Components
## Applications

|Name|Location|Purpose|
|---|---|---|

## Libraries

|Name|Location|Purpose|
|---|---|---|

# Infrastructure
The infrastructure used to implement this solution is defined with Terraform HCL.  See the [deployments/infra](deployments/infra/) directory for more details.

A sample deployment pipelien for this can be seen at [terraform-deploy-prd](.github/workflows/deploy-prd.yml)

# Pipelines
The following pipelines are used by this repository to build, test, and deploy a sample instance.  They are:
|Name|Location|Purpose|
|---|---|---|
|application-continuous-integration|[.github/workflows/application-ci.yml](.github/workflows/application-ci.yml)|Application code (golang) CI build|

# How to Use
To start using Blog Monitor, use the following steps

## 1. Create Cloud Resources
Using the infrastructure defined via HCL code in the [deployments/infra](deployments/infra/) folder, run this on an AWS Account that you own so the resources are properly configured

# Development Environment Setup
To set up your development environment, follow the steps above in the _How to Use_ section.
Once complete, you should be able to open this solution in an editor of your choice and start making changes / running on your own.