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
  },
  "type": {
      "S": "blogger"
  }
}
ITEM
}

resource "aws_dynamodb_table_item" "subscription_wasmoke" {
  table_name = aws_dynamodb_table.subscriptions.name
  hash_key   = aws_dynamodb_table.subscriptions.hash_key

  item = <<ITEM
{
  "id": {
    "S": "2"
  },
  "maximumLookback": {
    "N": "5"
  },
  "name": {
    "S": "Washington Smoke information"
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
    "S": "https://wasmoke.blogspot.com/"
  },
  "type": {
      "S": "blogger"
  }
}
ITEM
}

resource "aws_dynamodb_table_item" "subscription_nws_seattle_area_forecast" {
  table_name = aws_dynamodb_table.subscriptions.name
  hash_key   = aws_dynamodb_table.subscriptions.hash_key

  item = <<ITEM
{
  "id": {
    "S": "3"
  },
  "maximumLookback": {
    "N": "5"
  },
  "name": {
    "S": "Seattle Area Forecast Discussion"
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
    "S": "https://api.weather.gov/products/types/AFD/locations/SEW"
  },
  "type": {
      "S": "national-weather-service"
  }
}
ITEM
}