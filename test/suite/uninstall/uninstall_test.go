package uninstall_test

import (
	"testing"

	"github.com/jenkins-x/bdd-jx/test/helpers"
	"github.com/jenkins-x/jx/v2/pkg/kube"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	helpers.RunWithReporters(t, "uninstall")
}

var _ = BeforeSuite(helpers.BeforeSuiteCallback)

var _ = SynchronizedAfterSuite(func() {}, helpers.SynchronizedAfterSuiteCallback)

var _ = Describe("uninstall Jenkins X", func() {
	var (
		T                helpers.TestOptions
	)

	Context("when running jx uninstall ", func() {
		It("Jenkins X gets uninstalled and the development namespace is removed", func() {
			kuber := kube.NewKubeConfig()
			config, _, err := kuber.LoadConfig()
			Expect(err).Should(BeNil())
			currentContext := kube.CurrentContextName(config)

			args := []string{"--verbose", "-b", "uninstall", "--context", currentContext}
			T.ExpectJxExecution(T.WorkDir, helpers.TimeoutSessionWait, 0, args...)
		})
	})
})