package manifest

import "testing"

func TestDatasource_ConfigureNil(t *testing.T) {
	datasource := Datasource{
		config: Config{
			source: "main.json",
		},
	}
	if err := datasource.Configure(nil); err == nil {
		t.Fatalf("Should error if source is blank")
	}
}
