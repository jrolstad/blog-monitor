data "archive_file" "cron_lambda_zip" {
  type        = "zip"
  source_file = "../../cmd/lambda/cron_monitor/main"
  output_path = "monitor_main.zip"
}

resource "aws_lambda_function" "cron_monitor" {
  function_name = "${local.service_name}_cron_monitor"

  role = aws_iam_role.lambda_exec.arn

  filename         = data.archive_file.cron_lambda_zip.output_path
  handler          = "main"
  source_code_hash = filebase64sha256(data.archive_file.cron_lambda_zip.output_path)
  runtime          = "go1.x"
  timeout          = 600

  environment {
    variables = {
      aws_region = var.aws_region
    }
  }

}

resource "aws_cloudwatch_log_group" "cron_monitor" {
  name = "/aws/lambda/${aws_lambda_function.cron_monitor.function_name}"

  retention_in_days = 30
}

resource "aws_cloudwatch_event_rule" "every_twelve_hours" {
  name                = "every-twelve-hours"
  description         = "Fires every 12 hours"
  schedule_expression = "rate(12 hours)"
}

resource "aws_cloudwatch_event_target" "load_owners_every_twelve_hours" {
  rule      = aws_cloudwatch_event_rule.every_twelve_hours.name
  target_id = "lambda"
  arn       = aws_lambda_function.cron_monitor.arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_cron_monitor" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cron_monitor.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.every_twelve_hours.arn
}