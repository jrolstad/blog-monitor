resource "aws_secretsmanager_secret" "google_api_key" {
  name = "${local.service_name}_google_apikey"
}

resource "aws_secretsmanager_secret_version" "google_api_key" {
  secret_id     = aws_secretsmanager_secret.google_api_key.id
  secret_string = var.google_api_key
}