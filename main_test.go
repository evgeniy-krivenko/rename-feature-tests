package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParsePath(t *testing.T) {
	path := "/Users/ekrivenko/work-repos/acceptance-test/api-acceptance-test/src/test/resources/features/selfcare/pz/internetSettings/SettingsInernet.feature"
	res := getDottedPath(path, "selfcare")
	require.Equal(t, "selfcare.pz.internetSettings.SettingsInernet", res)
}

func TestParsePathOther(t *testing.T) {
	path := "/Users/ekrivenko/work-repos/acceptance-test/api-acceptance-test/src/test/resources/features/selfcare/pz/tele2pay/payment/Payment.feature"
	res := getDottedPath(path, "selfcare")
	require.Equal(t, "selfcare.pz.tele2pay.payment.Payment", res)
}

func TestCutLine(t *testing.T) {
	line := "  @atest @autoTestExternalId-api.selfcare.foldet.test.Feature.1"
	res := cutLine(line, "@autoTestExternalId-api")
	require.Equal(t, "  @atest", res)
}

func TestCutLineWithoutPrefix(t *testing.T) {
	line := "  @atest\n"
	res := cutLine(line, "@autoTestExternalId-api")
	require.Equal(t, "  @atest", res)
}

func TestCutLineWithOtherTags(t *testing.T) {
	line := "  @atest @tech\n"
	res := cutLine(line, "@autoTestExternalId-api")
	require.Equal(t, "  @atest @tech", res)
}

func TestCutLineWithOtherTagsAndPrefix(t *testing.T) {
	line := "  @atest @tech @autoTestExternalId-api.selfcare.foldet.test.Feature.1\n"
	res := cutLine(line, "@autoTestExternalId-api")
	require.Equal(t, "  @atest @tech", res)
}
