output "arn" {
  description = "The ARN of the resource"
  value       = aws_iam_openid_connect_provider.this.arn
}

output "tags_all" {
  description = "A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block."
  value       = aws_iam_openid_connect_provider.this.tags_all
}

output "url" {
  description = "The URL of the IAM OpenID Connect provider"
  # AWS stores OIDC URLs without the https:// prefix, so we normalize it here for consistency
  value = format("https://%s", trimprefix(aws_iam_openid_connect_provider.this.url, "https://"))
}
