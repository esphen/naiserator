package resourcecreator_test

import (
	"fmt"
	"testing"

	nais_io_v1alpha1 "github.com/nais/liberator/pkg/apis/nais.io/v1alpha1"
	"github.com/nais/naiserator/pkg/resourcecreator/resource"
	"github.com/nais/naiserator/pkg/test/goldenfile"

	"github.com/ghodss/yaml"
	"github.com/nais/naiserator/pkg/resourcecreator"
)

const (
	testDataDirectory = "testdata"
)

type applicationTestCase struct {
	ApplicationInput nais_io_v1alpha1.Application
}

func TestApplicationGoldenFile(t *testing.T) {
	goldenfile.Run(t, testDataDirectory, func(input []byte, resourceOptions resource.Options) (resource.Operations, error) {
		test := applicationTestCase{}
		err := yaml.Unmarshal(input, &test)
		if err != nil {
			return nil, err
		}

		err = test.ApplicationInput.ApplyDefaults()
		if err != nil {
			return nil, fmt.Errorf("apply default values to Application object: %s", err)
		}

		return resourcecreator.CreateApplication(&test.ApplicationInput, resourceOptions)
	})
}
