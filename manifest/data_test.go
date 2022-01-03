package manifest

import "testing"

func TestDatasource_ConfigureNil(t *testing.T) {
	datasource := Datasource{
		config: Config{
			source: "",
		},
	}
	if err := datasource.Configure(nil); err == nil {
		t.Fatalf("Should error if source is blank")
	}
}

func TestDatasource_ConfigureFile(t *testing.T) {
	datasource := Datasource{
		config: Config{
			source: "test-data/manifest.json",
		},
	}
	if err := datasource.Configure(nil); err != nil {
		t.Fatalf("Error parsing config")
	}
	if datasource.config.source != "test-data/manifest.json" {
		t.Fatalf("Values do not match")
	}
}
