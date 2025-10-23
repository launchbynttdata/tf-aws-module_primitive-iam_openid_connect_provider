output "arn" {
  description = "The ARN of the resource"
  value       = module.aws_iam_openid_connect_provider.arn
}

output "tags_all" {
  description = "A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block."
  value       = module.aws_iam_openid_connect_provider.tags_all
}
