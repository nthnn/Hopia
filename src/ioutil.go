package main

import "github.com/magiconair/properties"

func loadPropertiesFile(fileName string) *properties.Properties {
	return properties.MustLoadFile(fileName, properties.UTF8)
}
