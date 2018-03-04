# Terraform (in development)

## Config

In `conf` folder :
* YAML configuration,
* Describe your stack,
* Format file `project_region_stack.yml` or `project_region_env_stack.yml`.

### Example

Here for AWS :
```yaml
---
cloud: 'aws'
terraform: '0.10.8'
provider:
  general:
    account: 'XXXXXXXXXXXXXX'
    region: 'eu-west-1'
    env: 'prod'
  credentials:
    profile: my-profile
    role: my-role
```

You can found an other example for GCP in this folder.

## Tree

```
.
├── README.md
├── conf
│   ├── project_eu-west-1_cdn.yml
│   └── project_europe-west1_kubernetes.yml
└── project
    ├── eu-west-1
    │   └── cdn
    │       ├── features.tf
    │       ├── inputs.tf
    │       ├── main.tf
    │       └── outputs.tf
    └── europe-west1
        └── kubernetes
            ├── features.tf
            ├── inputs.tf
            ├── main.tf
            └── outputs.tf
```