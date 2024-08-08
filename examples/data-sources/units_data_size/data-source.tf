data "units_data_size" "disk_size" {
  gigabytes = 1000
}

output "real_disk_size" {
  value = data.units_data_size.disk_size.gibibytes
}

data "units_data_size" "fs_block_size" {
  kibibytes = 4
}

output "full_fs_block_size" {
  value = data.units_data_size.fs_block_size.kibibytes
}
