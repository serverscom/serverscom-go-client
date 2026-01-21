package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCloudComputingRegionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/regions").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_regions/list_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.CloudComputingRegions.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(HaveLen(2))
	g.Expect(list[0].ID).To(Equal(int64(1)))
	g.Expect(list[0].Name).To(Equal("Region 1"))
	g.Expect(list[0].Code).To(Equal("region1"))
	g.Expect(list[1].ID).To(Equal(int64(2)))
	g.Expect(list[1].Name).To(Equal("Region 2"))
	g.Expect(list[1].Code).To(Equal("region2"))
}

func TestCloudComputingRegionsImages(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/regions/123/images").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_regions/images_list_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.CloudComputingRegions.Images(123)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(HaveLen(2))
	g.Expect(list[0].ID).To(Equal("image1"))
	g.Expect(list[0].Name).To(Equal("Ubuntu 20.04"))
	g.Expect(list[1].ID).To(Equal("image2"))
	g.Expect(list[1].Name).To(Equal("CentOS 7"))
}

func TestCloudComputingRegionsFlavors(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/regions/456/flavors").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_regions/flavors_list_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.CloudComputingRegions.Flavors(456)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(HaveLen(2))
	g.Expect(list[0].ID).To(Equal("flavor1"))
	g.Expect(list[0].Name).To(Equal("Small"))
	g.Expect(list[1].ID).To(Equal("flavor2"))
	g.Expect(list[1].Name).To(Equal("Medium"))
}

func TestCloudComputingRegionsCredentials(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/regions/789/credentials").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_regions/credentials_list_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	creds, err := client.CloudComputingRegions.Credentials(ctx, 789)

	g.Expect(err).To(BeNil())
	g.Expect(creds.Password).To(Equal("secret"))
	g.Expect(creds.TenantName).To(Equal(int64(123)))
	g.Expect(creds.URL).To(Equal("https://example.com"))
	g.Expect(creds.Username).To(Equal(int64(456)))
}
