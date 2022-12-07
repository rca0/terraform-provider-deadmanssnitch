# terraform-deadmanssnitch

refs to create a new terraform-provider written in Go

https://github.com/hashicorp/terraform-plugin-sdk
https://developer.hashicorp.com/terraform/plugin

## WIP - NOT READY TO PRODUCTION

Terraform provider based on deadmanssnitch API core system: https://deadmanssnitch.com/docs/api/v1

Requirements:
- Deadmanssnitch API-KEY
- Go 1.18+
- Terraform v1.0.3+
- Docker and Docker-Compose

## Terraform Provider Actions

### Post Fields

* name [string]
* interval [string]
  * opts: |
    1_minute |
    2_minute | 
    3_minute |
    5_minute | 
    10_minute |
    15_minute | 
    30_minute | 
    hourly | 
    2_hour | 
    3_hour | 
    4_hour | 
    6_hour | 
    8_hour |
    12_hour |
    daily |
    weekly |
    monthly
* alert_type [string]
  * opts: basic | smart | default: basic
* alert_mail [array|string|null]
* notes [string|null]
* tags [array|null]

### Result Fields

* name [string]
* href [string]
* token [string]
* notes [string|null]
* created_at [string] or [timestamp]
* check_in_url [string]
* checked_in_at ??
* status [string]
* tags [array|null]
* type [map[string][string]]
* interval [string]
* alert_type [string]
  * opts: basic | smart | default: basic
* alert_mail [array|string|null]

## Create a Snitch

* POST request to path `/v1/snitches`

with commandline 
```bash
curl -X POST -d '{"name":"Daily Backups","interval":"daily"}' -H "Content-Type: application/json" -u <api-key>: https://api.deadmanssnitch.com/v1/snitches
```

## Update a Snitch

* Request to path `/v1/snitches/[token]`

with command line
```bash
curl -i -X PATCH -d '{"name":"Daily Backups","interval":"daily","notes":"Postgres box at 123.213.231.132","tags": ["production", "critical"]}' -H "Content-Type: application/json" -u <api_key>: https://api.deadmanssnitch.com/v1/snitches/c2354d53d2
```

## Delete a snitch

* DELETE request to path `/v1/snitches/[token]`

```bash
curl -i -X DELETE -u <api-key>: https://api.deadmanssnitch.com/v1/snitches/c2354d53d2
```

* The response will be empty, with 204 (no contet) HTTP status code

## Filtering by Tags(s)

* GET request to `/v1/snitches`

```bash
curl -u <api-key>:
  https://api.deadmanssnitch.com/v1/snitches?tags=production,critical
```

* If no snitches match the filter, the response will be empty array

## Adding one or more tags to a snitch

* POST request to path `/v1/snitches/[token]/tags`

with command line
```bash
curl -i -X POST -d '["production"]' -H "Content-Type: application/json" -u <api-key>: https://api.deadmanssnitch.com/v1/snitches/c2354d53d2/tags
```

## Removing one or more tags to a snitch

* DELETE request to path `/v1/snitches/*c2354d53d2*/tags/*tag_name*`

with command line
```bash
curl -i -X DELETE -H "Content-Type: application/json" -u <api-key>: https://api.deadmanssnitch.com/v1/snitches/c2354d53d2/tags/critical
```

## Changing tags on snitch

* PATCH request to path `/v1/snitches/[token]`

with command line
```bash
curl -i -X PATCH -d '{"tags": ["staging", "backup"]}' -H "Content-Type: application/json" -u <api-key>: https://api.deadmanssnitch.com/v1/snitches/c2354d53d2
```

## Pausing a Snitch

* POST request to path `/v1/snitches/[token]/pause`

```bash
curl -i -X POST -u <api-key>: https://api.deadmanssnitch.com/v1/snitches/c2354d53d2/pause
```

* The response will be empty, with 204 (no content) HTTP status code 

### Error responses

* If request JSON malformed
- Return 422 (unprocessable Entity)
- type resource_invalid

```json
{
  "type": "resource_invalid",
  "error": "The requested resource attributes are not valid.",
  "validations": [
    {
      "attribute": "name",
      "message": "can't be blank"
    },
    {
      "attribute": "interval",
      "message": "can't be blank"
    }
  ]
}
```

* If request invalid interval
- Return tpe resource_invalid

```json
{
  "type": "resource_invalid",
  "error": "The requested resource attributes are not valid.",
  "validations": [
    {
      "attribute": "interval",
      "message": "must be \"15_minute\", \"30_minute\", \"hourly\", \"daily\", \"weekly\", or \"monthly\""
    }
  ]
}
```

* Max number of snitches in your plan
- Return status code 422
- type plan_limit_reached

```json
{
  "type": "plan_limit_reached",
  "error": "We could not create your snitch because you are at your plan limit of 1 snitch! Delete an unused snitch, or head over to https://deadmanssnitch.com/ to upgrade your plan."
}
```

* API Key Error Responses

incorrect api-key will return 401 (Unauthorized) HTTP status code

```json
{
  "type": "api_key_invalid",
  "error": "Access denied. Provide your API Key as the user for HTTP Basic Authentication."
}
```