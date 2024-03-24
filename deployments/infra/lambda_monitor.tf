data "archive_file" "cron_lambda_zip" {
  type        = "zip"
  source_file = "../../cmd/lambda/cron_monitor/bootstrap"
  output_path = "monitor_main.zip"
}

resource "aws_lambda_function" "cron_monitor" {
  function_name = "${local.service_name}_cron_monitor"

  role = aws_iam_role.lambda_exec.arn

  filename         = data.archive_file.cron_lambda_zip.output_path
  handler          = "main"
  source_code_hash = filebase64sha256(data.archive_file.cron_lambda_zip.output_path)
  runtime          = "provided.al2"
  architectures    = ["arm64"]
  timeout          = 600

  environment {
    variables = {
      aws_region                      = var.aws_region
      email_sender                    = var.email_sender
      subscription_table_name         = aws_dynamodb_table.subscriptions.name
      notification_history_table_name = aws_dynamodb_table.notification_history.name
      secret_google_api_key           = aws_secretsmanager_secret.google_api_key.name
    }
  }

}

resource "aws_cloudwatch_log_group" "cron_monitor" {
  name = "/aws/lambda/${aws_lambda_function.cron_monitor.function_name}"

  retention_in_days = 30
}

resource "aws_cloudwatch_event_rule" "every_hour" {
  name                = "every-hour"
  description         = "Fires every 1 hours"
  schedule_expression = "rate(1 hours)"
}

resource "aws_cloudwatch_event_target" "load_owners_every_hour" {
  rule      = aws_cloudwatch_event_rule.every_hour.name
  target_id = "lambda"
  arn       = aws_lambda_function.cron_monitor.arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_cron_monitor" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cron_monitor.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.every_hour.arn
}