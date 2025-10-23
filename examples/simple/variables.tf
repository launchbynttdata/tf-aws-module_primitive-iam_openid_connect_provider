variable "url" {
  description = "(Required) URL of the identity provider, corresponding to the iss claim."
  type        = string
}

variable "client_id_list" {
  description = "(Required) List of client IDs (audiences) that this provider should accept."
  type        = list(string)
}

variable "thumbprint_list" {
  description = "(Required) List of server certificate thumbprints for the identity provider's server certificates."
  type        = list(string)
}

variable "tags" {
  description = "(Optional) A map of tags to assign to the resource."
  type        = map(string)
  default     = {}
}
