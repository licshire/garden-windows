package lifecycle

import (
	"github.com/cloudfoundry-incubator/garden"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container Metrics", func() {
	var gardenArgs []string

	BeforeEach(func() {
		gardenArgs = []string{}
		client = startGarden(gardenArgs...)
	})

	Describe("for a single container", func() {
		var container garden.Container

		BeforeEach(func() {
			var err error
			container, err = client.Create(garden.ContainerSpec{})
			Expect(err).ToNot(HaveOccurred())
			StreamIn(container)
		})

		AfterEach(func() {
			client.Destroy(container.Handle())
		})

		It("returns metrics", func() {
			metrics, err := container.Metrics()
			Expect(err).ToNot(HaveOccurred())

			Expect(metrics.MemoryStat.TotalRss).To(BeNumerically(">", 0), "Expected Memory Usage to be > 0")
			Expect(metrics.CPUStat.Usage).To(BeNumerically(">", 0), "Expected CPU Usage to be > 0")
			Expect(metrics.DiskStat.BytesUsed).To(BeNumerically(">", 0), "Expected Disk Usage to be > 0")
		})
	})
})