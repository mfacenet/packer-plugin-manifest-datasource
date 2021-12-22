//go:generate packer-sdc mapstructure-to-hcl2 -type Config.DatasourceOutput
package manifest

import (
	"errors"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	// Source manifest file
	source string `mapstructure:"source"`
}

type Datasource struct {
	config Config
}

/* Json Format used to model output on
{
  "builds": [
    {
      "name": "virtualbox_zero_config",
      "builder_type": "virtualbox-iso",
      "build_time": 1639951415,
      "files": [
        {
          "name": "..\\builds\\zero-config-disk001.vmdk",
          "size": 4894903296
        },
        {
          "name": "..\\builds\\zero-config.ovf",
          "size": 7010
        }
      ],
      "artifact_id": "VM",
      "packer_run_uuid": "d53141f2-4e9c-e14c-67aa-06b3e4cb25dc",
      "custom_data": null
    }
  ],
  "last_run_uuid": "d53141f2-4e9c-e14c-67aa-06b3e4cb25dc"
}
*/

type DataSourceOutput struct {
	latestBuild string `mapstructure:"latestBuild"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.Flatmapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}
	if d.config.source == "" {
		d.config.source = "manifest.json"
	}
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DataSourceOutput{}).Flatmapstructure().HCL2SPec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	output := DataSourceOutput{}
	emptyOutput := hcl2Helper.HCL2ValueFromConfig(output, d.OutputSpec())
	err := errors.New("This functionality hasn't been implemented yet")

	//Open File

	//Parse File

	return (emptyOutput, err)
}
