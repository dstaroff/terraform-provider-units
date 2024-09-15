output "example" {
  size_in_tebibytes = provider::units::to_tib(42)
}