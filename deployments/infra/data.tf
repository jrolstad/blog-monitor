resource "aws_dynamodb_table_item" "subscription_cliffmassweather" {
  table_name = aws_dynamodb_table.subscriptions.name
  hash_key   = aws_dynamodb_table.subscriptions.hash_key

  item = <<ITEM
{
  "id": {
    "S": "1"
  },
  "maximumLookback": {
    "N": "5"
  },
  "name": {
    "S": "Cliff Mass Weather"
  },
  "notificationMethod": {
    "S": "email"
  },
  "notificationTargets": {
    "SS": [
      "jrolstad@gmail.com"
    ]
  },
  "url": {
    "S": "https://cliffmass.blogspot.com"
  }
}
ITEM
}