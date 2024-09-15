output "example" {
  size_in_gigabytes = provider::units::to_gb(42)
}