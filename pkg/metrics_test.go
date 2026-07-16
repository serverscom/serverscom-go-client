package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestListHostsMetrics(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/metrics/hosts").
		WithRequestMethod("GET").
		WithResponseBodyStubInline("# HELP hosts_total\nhosts_total 42\n").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	metrics, err := client.Metrics.ListHostsMetrics(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(metrics).To(Equal("# HELP hosts_total\nhosts_total 42\n"))
}

func TestListRacksMetrics(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/metrics/racks").
		WithRequestMethod("GET").
		WithResponseBodyStubInline("# HELP racks_total\nracks_total 7\n").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	metrics, err := client.Metrics.ListRacksMetrics(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(metrics).To(Equal("# HELP racks_total\nracks_total 7\n"))
}
