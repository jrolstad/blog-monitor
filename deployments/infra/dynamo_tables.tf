resource "aws_dynamodb_table" "subscriptions" {
  name         = "${local.service_name}_subscriptions"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Id"

  attribute {
    name = "id"
    type = "S"
  }

}

resource "aws_dynamodb_table" "notification_history" {
  name         = "${local.service_name}_notification_history"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Id"

  attribute {
    name = "id"
    type = "S"
  }

}