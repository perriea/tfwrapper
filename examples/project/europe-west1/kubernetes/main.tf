provider "google" {
  credentials = "${file(var.gcp_credentials)}"
  project     = "${var.gcp_project}"
  region      = "${var.region}"
  version     = "1.8.0"
}

provider "random" {
  version = "1.2"
}
