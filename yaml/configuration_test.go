package yaml_test

import (
	"os"
	"reflect"
	"testing"

	web "github.com/iwittkau/podcast-web"
	"github.com/iwittkau/podcast-web/yaml"
)

func TestReadSiteConfig(t *testing.T) {

	site := web.Site{
		Title:  "Test",
		Author: "Tester",
		Menu: []web.MenuItem{
			{
				Name: "About",
				Link: "/about",
			},
		},
		Episodes: []web.EpisodeLink{
			{
				Name:   "Episode 0001",
				Link:   "/episodes/0001",
				Number: 1,
			},
		},
	}

	err := yaml.WriteSiteConfig("config.yml", site)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll("config.yml")

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    web.Site
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "config.yml", args: args{path: "config.yml"}, want: site, wantErr: false},
		{name: "config_not_found.yml", args: args{path: "config_not_found.yml"}, want: web.Site{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yaml.ReadSiteConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSiteConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadSiteConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
