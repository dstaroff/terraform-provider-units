output "example" {
  size_in_petabytes = provider::units::to_pb(42)
}