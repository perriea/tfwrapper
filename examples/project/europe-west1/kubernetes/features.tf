resource "google_storage_bucket" "image-store" {
  name     = "${var.env}-image-store-${var.gcp_region}"
  location = "EU"

  website {
    main_page_suffix = "index.html"
    not_found_page   = "404.html"
  }
}