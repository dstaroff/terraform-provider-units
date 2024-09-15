output "example" {
  size_in_pebibytes = provider::units::to_pib(42)
}