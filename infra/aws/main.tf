terraform {
  backend "s3" {
    bucket = "carthage-infra-tf-state"
    key    = "infra/terraform.tfstate"
    region = "ap-south-1"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "ap-south-1"
}

# VARIABLES
variable "s3_bucket_name" {
  description = "The name of the S3 bucket"
  type        = string
}

variable "fifo_queue_name" {
  description = "The name of the FIFO queue"
  type        = string
}

variable "fifo_sns_topic_name" {
  description = "The name of the FIFO SNS topic"
  type        = string
}

resource "aws_s3_bucket" "carthage-private-assets" {
  bucket        = var.s3_bucket_name
  force_destroy = true
}

resource "aws_sqs_queue" "carthage_fifo_queue" {
  tags = {
    Name        = "CarthageQueue"
    Environment = "Dev"
  }

  name = var.fifo_queue_name

  fifo_queue                  = true
  content_based_deduplication = true
  deduplication_scope         = "messageGroup"
  fifo_throughput_limit       = "perMessageGroupId"

  delay_seconds              = 0     # time in seconds, after which message is made available for consumers
  max_message_size           = 2048  # max size of a single message body in bytes
  message_retention_seconds  = 86400 # time in seconds, for which message is stored in queue, after the retention period, message is deleted
  receive_wait_time_seconds  = 20    # time in seconds, for which the consumer long polls the queue in case of empty queue. This reduces number of API calls to SQS
  visibility_timeout_seconds = 60    # time in seconds, for which a consumer can process message before the message is made available for other consumers (deadline for consuming)
}

resource "aws_sns_topic" "carthage_reports_fifo" {
  name                        = var.fifo_sns_topic_name
  fifo_topic                  = true
  content_based_deduplication = true

  tags = {
    Name        = "CarthageReportsTopic"
    Environment = "Dev"
  }
}

data "aws_iam_policy_document" "carthage_policy_sns_sqs" {
  statement {
    effect = "Allow"
    sid    = "Allow-SNS-SQS"

    principals {
      type        = "Service"
      identifiers = ["sns.amazonaws.com"]
    }

    actions   = ["sqs:SendMessage"]
    resources = [aws_sqs_queue.carthage_fifo_queue.arn]

    condition {
      test     = "ArnEquals"
      variable = "aws:SourceArn"
      values   = [aws_sns_topic.carthage_reports_fifo.arn]
    }
  }
}

resource "aws_sqs_queue_policy" "carthage_queue_allow_sns" {
  queue_url = aws_sqs_queue.carthage_fifo_queue.id
  policy    = data.aws_iam_policy_document.carthage_policy_sns_sqs.json
}

resource "aws_sns_topic_subscription" "carthage_queue_subs" {
  topic_arn = aws_sns_topic.carthage_reports_fifo.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.carthage_fifo_queue.arn
}
