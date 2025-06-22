terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }
}

provider "aws" {
  region = "ap-south-1"
}

resource "aws_iam_user" "user_gh_action_to_ecr" {
  name = var.iam_user_name

}

resource "aws_iam_policy" "carthage-push-tag-ecr-image-policy" {
  name        = "carthage-push-tag-ecr-image-policy"
  description = "Allow pushing and tagging docker images to public ECR repositories"
  policy = jsonencode(
    {
      Version = "2012-10-17"
      Statement = [
        {
          Sid    = "AllowPushTagECRImage"
          Effect = "Allow"
          Action = [
            "ecr-public:UntagResource",
            "ecr-public:TagResource",
            "ecr-public:UploadLayerPart",
            "ecr-public:PutImage",
            "ecr-public:InitiateLayerUpload",
            "ecr-public:CompleteLayerUpload",
            "ecr-public:BatchCheckLayerAvailability"
          ]
          Resource = var.ecr_repo_arn
        },
        {
          Sid    = "AllowGetAuthorizationToken"
          Effect = "Allow"
          Action = [
            "ecr-public:GetAuthorizationToken",
            "sts:GetServiceBearerToken"
          ]
          Resource = "*"
        },
      ]
  })
}

resource "aws_iam_user_policy_attachment" "attach-ecr-policy" {
  user       = aws_iam_user.user_gh_action_to_ecr.name
  policy_arn = aws_iam_policy.carthage-push-tag-ecr-image-policy.arn
}

resource "aws_iam_access_key" "carthage-push-tag-ecr-image-key" {
  user = aws_iam_user.user_gh_action_to_ecr.name
}

output "ecr_ci_user_access_key_id" {
  description = "The access key ID for the user to push and tag images to ECR"
  value       = aws_iam_access_key.carthage-push-tag-ecr-image-key.id
  sensitive   = true
}

output "ecr_ci_user_access_key_secret" {
  description = "The access key Secret for the user to push and tag images to ECR"
  value       = aws_iam_access_key.carthage-push-tag-ecr-image-key.secret
  sensitive   = true
}
