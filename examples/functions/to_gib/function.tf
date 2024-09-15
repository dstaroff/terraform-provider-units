output "example" {
  size_in_gibibytes = provider::units::to_gib(42)
}