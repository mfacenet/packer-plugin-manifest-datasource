# Packer Manifest Data Source

This is a packer plugin to read a packer manifest file and
make its data available for builders.

# Usage

```hcl
packer {
  required_plugins {
    manifest-datasource= {
      version = "0.0.1"
      source = "github.com/mfacenet/manifest-datasource"
    }
  }
}
```

Initialize your packer template (it will install the plugin):

```bash
packer init your-template.pkr.hcl
```

Use this provisioner plugin from your packer template file:

```hcl
data-source "manifest-datasource" {
  source = "<path-to-manifest"
}
```
