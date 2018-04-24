resource "random_string" "password" {
  length  = 16
  special = true
  number  = true
  lower   = true
  upper   = true
}

module "gke-cluster" {
  source = "github.com/google-terraform-modules/terraform-google-kubernetes-engine"

  name     = "mycluster"
  env      = "prod"
  zone     = "europe-west1-b"
  username = "admin"
  password = "${random_string.password.result}"
}
