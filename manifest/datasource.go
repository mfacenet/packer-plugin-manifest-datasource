//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config.DatasourceOutput

package manifest

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/json"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	// Source manifest file
	source              string `mapstructure:"source"`
	common.PackerConfig `mapstructure:",squash"`
}

type Datasource struct {
	config Config
}

type File struct {
	name string
	size int
}

type Build struct {
	name            string
	builder_type    string
	build_time      int
	artifact_id     string
	packer_run_uuid string
	custom_data     string
	files           []File
}

type DataSourceOutput struct {
	builds        []Build
	last_run_uuid string
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}

	var errs *packersdk.MultiError
	errs = packersdk.MutliErrorAppend(errs, d.config.AccessConfig.Prepare(&d.config.PackerConfig)...)

	if d.config.Empty() {
		errs = packersdk.MultiErrorAppend(errs, fmt.Errorf("A source must be specified"))
	}

	if errs != nil && len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DataSourceOutput{}).Flatmapstructure().HCL2SPec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	output := DataSourceOutput{}
	emptyOutput := hcl2Helper.HCL2ValueFromConfig(output, d.OutputSpec())

	//Open File
	content, err := ioutil.ReadFile(d.config.source)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}
	var payload Data

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}
	// Build out Data structures

	return emptyOutput, err
}
