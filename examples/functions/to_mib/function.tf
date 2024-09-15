output "example" {
  size_in_mebibytes = provider::units::to_mib(42)
}