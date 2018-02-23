# tfwrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/perriea/tfwrapper)](https://goreportcard.com/report/github.com/perriea/tfwrapper) [![Build Status](https://travis-ci.org/perriea/tfwrapper.svg?branch=master)](https://travis-ci.org/perriea/tfwrapper) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`tfwrapper` is a command created to apply best practices Terraform.

## Features

* AWS applications,
* Switching role (AWS STS),
* Possibility to use all feature of Terraform

## Requirements

* Credential AWS
* Using switch role AWS with MFA (recommanded by AWS)

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
                └── terraform.tf
```

## Roadmap

* Support Google Cloud Plateform
* Support Azure

## Licence

Source code can be found on [Github](https://github.com/georgeOsdDev/markdown-edit), licenced under [MIT](http://opensource.org/licenses/mit-license.php).

Developed by [Aurelien Perrier](http://about.me/perriea)