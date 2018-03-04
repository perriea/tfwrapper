# tfwrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/perriea/tfwrapper)](https://goreportcard.com/report/github.com/perriea/tfwrapper) [![Build Status](https://travis-ci.org/perriea/tfwrapper.svg?branch=master)](https://travis-ci.org/perriea/tfwrapper) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`tfwrapper` is a command created to apply best practices Terraform & Cloud Providers.

## Features

* AWS applications,
* Switching role (AWS STS),
* Autoswitch terraform version (based on file `terraform.tf`),
* Possibility to use all feature of Terraform.

## Requirements

* Credential AWS,
* Using AWS STS service with MFA (recommanded).

### Terraform architecture

**Example 1 :**
```
client
├── conf
│   └── client_project_region_stack.yml
└── project
    └── region
        └── stack
            ├── features.tf
            ├── inputs.tf
            ├── main.tf
            ├── outputs.tf
            └── terraform.tf
```

**Example 2 :**
```
client
├── conf
│   └── project_env_region_stack.yml
└── project
    └── env
        └── region
            └── stack
                ├── features.tf
                ├── inputs.tf
                ├── main.tf
                ├── outputs.tf
                └── terraform.tf
```

## Roadmap

* Support Google Cloud Plateform,
* Support Azure,
* Windows Support.

## Licence

Licence [MIT](http://opensource.org/licenses/mit-license.php)   
Developed by [Aurelien Perrier](http://about.me/perriea)