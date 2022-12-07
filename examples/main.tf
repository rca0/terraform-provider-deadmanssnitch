provider "deadmanssnitch" {
  // ENV VARS: DEADMANSSNITCH_APIKEY DEADMANSSNITCH_BASEURL 
  // baseurl = "https://api.deadmanssnitch.com/v1/snitches" (default value)
  apikey = var.deadmanssnitch_api_key
}

resource "deadmanssnitch_snitch" "foo" {
  name       = "foo"
  interval   = "2_minute"
  alert_type = "basic"
  alert_mail = "me@ruanc.dev"
  notes      = "The sample example of snitch creation with terraform"
  tags = [
    "basic",
    "production",
    "critical",
    "terraform"
  ]
}

resource "deadmanssnitch_snitch" "bar" {
  name       = "bar"
  interval   = "weekly"
  alert_type = "smart"
  alert_mail = "me@ruanc.dev"
  notes      = "The sample example of snitch creation with terraform"
  tags = [
    "smart",
    "production",
    "critical",
    "terraform"
  ]
}

terraform {
  required_providers {
    deadmanssnitch = {
      source  = "github.com/rca0/terraform-provider-deadmanssnitch"
      version = "0.1.0"
    }
  }
}