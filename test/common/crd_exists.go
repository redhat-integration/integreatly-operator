package common

import (
	"testing"

	"github.com/redhat-integration/rhi-operator/test/metadata"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIntegreatlyCRDExists(t *testing.T, ctx *TestingContext) {
	_, err := ctx.ExtensionClient.ApiextensionsV1beta1().CustomResourceDefinitions().Get("rhmis.integreatly.org", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err)
		metadata.Instance.FoundCRD = false
	} else {
		metadata.Instance.FoundCRD = true
	}
}
