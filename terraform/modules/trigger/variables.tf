# ------------------------------------------------------------------------------
# Variables
# ------------------------------------------------------------------------------
variable "name_prefix" {
  description = "Prefix used for resource names."
}

variable "lambda_arn" {
  description = "ARN of the Github Lambda."
}

variable "team_config" {
  description = "Valid JSON representation of a Team (see Go code)."
}
