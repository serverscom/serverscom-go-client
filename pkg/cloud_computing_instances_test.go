package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCloudComputingInstancesCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.CloudComputingInstances.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudComputingInstancesCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := CloudComputingInstanceCreateInput{
		Name:     "test-instance-2",
		FlavorID: "102",
		ImageID:  "76effbf9-76e5-46d2-a21d-ee2a72cc8757",
		RegionID: 0,
		Labels:   map[string]string{"env": "test"},
	}

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.Create(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("LDdwRb1Y"))
	g.Expect(cloudInstance.Status).To(Equal("PROVISIONING"))
	g.Expect(cloudInstance.Name).To(Equal("test-instance-2"))
	g.Expect(cloudInstance.FlavorID).To(Equal("102"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("04e59a31-efe6-440c-9568-142fca1a6123"))
	g.Expect(cloudInstance.ImageID).To(Equal("76effbf9-76e5-46d2-a21d-ee2a72cc8757"))
	g.Expect(cloudInstance.PublicIPv4Address).To(BeNil())
	g.Expect(cloudInstance.PrivateIPv4Address).To(BeNil())
	g.Expect(cloudInstance.PublicIPv6Address).To(BeNil())
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:28 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:28 +0000 UTC"))
}

func TestCloudComputingInstancesGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_instances/get_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.Get(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("ACTIVE"))
	g.Expect(cloudInstance.Name).To(Equal("name37"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("f6c9c585-627a-4113-af8c-a475f5f73a21"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2").
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/cloud_instances/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newName := "some"
	newLabels := map[string]string{"env": "test"}

	cloudInstance, err := client.CloudComputingInstances.Update(ctx, "BDbDxbl2", CloudComputingInstanceUpdateInput{Name: &newName, Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("ACTIVE"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("f6c9c585-627a-4113-af8c-a475f5f73a21"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(newLabels))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesDelete(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.CloudComputingInstances.Delete(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
}

func TestCloudComputingInstancesReinstall(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/reinstall").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/reinstall_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.Reinstall(ctx, "BDbDxbl2", CloudComputingInstanceReinstallInput{ImageID: "18e1cc16-b380-4c37-8ec1-b9d306961aae"})

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("REBUILDING"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesRescue(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/rescue").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/rescue_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.Rescue(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("RESCUE"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesUnrescue(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/unrescue").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/unrescue_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.Unrescue(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("BUSY"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesUpgrade(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/upgrade").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/upgrade_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.Upgrade(ctx, "BDbDxbl2", CloudComputingInstanceUpgradeInput{FlavorID: "103"})

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("UPGRADING"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesRevertUpgrade(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/revert_upgrade").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/revert_upgrade_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.RevertUpgrade(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("ACTIVE"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesApproveUpgrade(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/approve_upgrade").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/approve_upgrade_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.ApproveUpgrade(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("ACTIVE"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesPowerOn(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/switch_power_on").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/power_on_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.PowerOn(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("SWITCHING_ON"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesPowerOff(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/switch_power_off").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/power_off_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudComputingInstances.PowerOff(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
	g.Expect(cloudInstance).ToNot(BeNil())

	g.Expect(cloudInstance.ID).To(Equal("BDbDxbl2"))
	g.Expect(cloudInstance.Status).To(Equal("SWITCHING_OFF"))
	g.Expect(cloudInstance.Name).To(Equal("some"))
	g.Expect(cloudInstance.FlavorID).To(Equal("101"))
	g.Expect(cloudInstance.OpenstackUUID).To(Equal("b9e388ff-e53b-498a-8ef4-764450236788"))
	g.Expect(cloudInstance.ImageID).To(Equal("18e1cc16-b380-4c37-8ec1-b9d306961aae"))
	g.Expect(*cloudInstance.PublicIPv4Address).To(Equal("127.0.0.1"))
	g.Expect(*cloudInstance.PrivateIPv4Address).To(Equal("127.0.0.2"))
	g.Expect(*cloudInstance.PublicIPv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudComputingInstancesCreatePTRRecord(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/xkazYeJ0/ptr_records").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/create_ptr_record_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	ttlValue := 60
	priorityValue := 3

	input := PTRRecordCreateInput{
		IP:       "100.0.0.4",
		Domain:   "ai.privateservergrid.com",
		TTL:      &ttlValue,
		Priority: &priorityValue,
	}

	ptrRecord, err := client.CloudComputingInstances.CreatePTRRecord(ctx, "xkazYeJ0", input)

	g.Expect(err).To(BeNil())
	g.Expect(ptrRecord).ToNot(BeNil())

	g.Expect(ptrRecord.ID).To(Equal("oQeZzvep"))
	g.Expect(ptrRecord.IP).To(Equal("100.0.0.4"))
	g.Expect(ptrRecord.Domain).To(Equal("ai.privateservergrid.com"))
	g.Expect(ptrRecord.Priority).To(Equal(3))
	g.Expect(ptrRecord.TTL).To(Equal(60))
}

func TestCloudComputingInstancesDeletePTRRecord(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/xkazYeJ0/ptr_records/oQeZzvep").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.CloudComputingInstances.DeletePTRRecord(ctx, "xkazYeJ0", "oQeZzvep")

	g.Expect(err).To(BeNil())
}

func TestCloudComputingInstancePTRRecordsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/a/ptr_records").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.CloudComputingInstances.PTRRecords("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}
