package sonarqube

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/golang"
	"github.com/trustacks/trustacks/pkg/actions/javascript"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	SonarProjectPropertiesExists = engine.NewFact()
)

var SonarProjectPropertiesExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "sonar-project.properties")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = SonarProjectPropertiesExists
	return fact, nil
}

func init() {
	engine.AddToRuleset(&javascript.PackageJSONExistsRule, &SonarProjectPropertiesExistsRule)
	engine.AddToRuleset(&golang.GoModExistsRule, &SonarProjectPropertiesExistsRule)
}
