# tf-aws-module_primitive-iam_openid_connect_provider

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_iam_openid_connect_provider.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_openid_connect_provider) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_url"></a> [url](#input\_url) | (Required) URL of the identity provider, corresponding to the iss claim. | `string` | n/a | yes |
| <a name="input_client_id_list"></a> [client\_id\_list](#input\_client\_id\_list) | (Required) List of client IDs (audiences) that this provider should accept. | `list(string)` | n/a | yes |
| <a name="input_thumbprint_list"></a> [thumbprint\_list](#input\_thumbprint\_list) | (Required) List of server certificate thumbprints for the identity provider's server certificates. | `list(string)` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | (Optional) A map of tags to assign to the resource. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the resource |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | A map of tags assigned to the resource, including those inherited from the provider default\_tags configuration block. |
| <a name="output_url"></a> [url](#output\_url) | The URL of the IAM OpenID Connect provider |
<!-- END_TF_DOCS -->
