data "manifest-datasource" "datasource" {
  source = "manifest.json"
}

locals {
  latest_build = data.manifest-datasource.datasource.latest_build
}

source "null" "test" {
  communicator = "none"
}

build {
  sources = [
    "source.null.test"
  ]

  provisioner "shell-local" {
    inline = [
      "echo latest_build: ${local.latest_build}"
    ]
  }
}