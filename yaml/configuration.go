package yaml

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	web "github.com/iwittkau/podcast-web"
)

// ReadSiteConfig reads site config from yaml file
func ReadSiteConfig(path string) (web.Site, error) {
	var site web.Site
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return site, err
	}
	err = yaml.Unmarshal(data, &site)
	if err != nil {
		return site, err
	}
	return site, nil
}

// WriteSiteConfig writes a site config to a yaml file
func WriteSiteConfig(path string, conf web.Site) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
