package ingress_test

import (
	"testing"

	nais_io_v1 "github.com/nais/liberator/pkg/apis/nais.io/v1"
	"github.com/nais/naiserator/pkg/generators"
	"github.com/nais/naiserator/pkg/resourcecreator/ingress"
	"github.com/nais/naiserator/pkg/resourcecreator/resource"
	"github.com/nais/naiserator/pkg/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestIngress(t *testing.T) {
	t.Run("invalid ingress URLs are rejected", func(t *testing.T) {
		for _, i := range []nais_io_v1.Ingress{"crap", "htp:/foo", "http://valid.fqdn/foo", "ftp://test"} {
			app := fixtures.MinimalApplication()
			app.Spec.Ingresses = []nais_io_v1.Ingress{i}
			ast := resource.NewAst()
			err := app.ApplyDefaults()
			assert.NoError(t, err)

			opts := &generators.Options{}
			opts.Config.Features.Linkerd = false
			err = ingress.Create(app, ast, opts)

			assert.NotNil(t, err)
			assert.Len(t, ast.Operations, 0)
		}
	})
}
