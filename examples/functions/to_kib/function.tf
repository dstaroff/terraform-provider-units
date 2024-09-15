output "example" {
  size_in_kibibytes = provider::units::to_kib(42)
}