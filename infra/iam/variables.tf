variable "iam_user_name" {
  type = string
}

variable "ecr_repo_arn" {
  type        = string
  description = "ARN of ecr repository the user can push images to"
}
