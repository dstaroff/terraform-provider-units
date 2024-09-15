output "example" {
  size_in_kilobytes = provider::units::to_kb(42)
}