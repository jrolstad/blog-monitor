terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.42.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.4.2"
    }
  }

  required_version = "~> 1.0"

  backend "s3" {
    bucket = "blogmonitortfstate"
    key    = "blogmonitortfstatetfstate/terraform.tfstate"
    region = "us-west-2"
  }
}