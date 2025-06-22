resource "aws_iam_role" "iam_role_for_lambda" {
  name = "iam_role_for_lambda"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_dir  = "${path.module}/lambda"
  output_path = "${path.module}/lambda.zip"
}

resource "aws_lambda_function" "carthage_reports_lambda" {
  tags = {
    Name        = "carthage-reports-lambda"
    Environment = "Dev"
  }

  function_name = "carthage_reports_lambda"
  filename      = data.archive_file.lambda_zip.output_path

  role        = aws_iam_role.iam_role_for_lambda.arn
  timeout     = 5
  runtime     = "nodejs16.x"
  memory_size = 128
  ephemeral_storage {
    size = 512
  }

  handler = "index.handler"

  environment {
    variables = {
      ENV = "dev"
    }
  }
}
