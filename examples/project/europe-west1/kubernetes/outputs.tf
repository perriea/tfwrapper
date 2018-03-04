output "bucket_state_self_link" {
  value = "${google_storage_bucket.image-store.self_link}"
}

output "bucket_state_url" {
  value = "${google_storage_bucket.image-store.url}"
}