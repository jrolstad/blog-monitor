variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-west-2"
}

variable "environment" {
  description = "Environment the infrasructure is for."

  type    = string
  default = "prd"
}

variable "google_api_key" {
  description = "API Key to use when calling Google services"

  type = string
}

variable "email_sender" {
  description = "Email address to send from when notifying via email."

  type = string
}