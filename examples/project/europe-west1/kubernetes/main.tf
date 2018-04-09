provider "google" {
  credentials = "${file(var.gcp_credentials)}"
  project     = "${var.gcp_project}"
  region      = "${var.gcp_region}"
  version     = "1.8"
}

provider "random" {
  version = "1.2"
}
