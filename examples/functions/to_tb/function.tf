output "example" {
  size_in_terabytes = provider::units::to_tb(42)
}