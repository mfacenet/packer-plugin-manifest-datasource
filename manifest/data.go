//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DataSourceOutput

package manifest

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Datasource struct {
	config Config
}

type Config struct {
	// Source manifest file
	source              string `mapstructure:"source"`
	common.PackerConfig `mapstructure:",squash"`
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
	builds        []Build `mapstructure:"builds"`
	last_run_uuid string  `mapstructure:"last_run_uuid"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}

	if len(d.config.source) < 1 {
		err := errors.New("A path must be provided for the source of the datafile")
		return err
	}

	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DataSourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	//Open File
	content, err := ioutil.ReadFile(d.config.source)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}
	output := DataSourceOutput{}
	err = json.Unmarshal(content, &output)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}
	// Build out Data structures

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), err
}
