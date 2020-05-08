package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCloudInstancesCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := CloudInstanceCreateInput{
		Name:     "test-instance-2",
		FlavorID: "102",
		ImageID:  "76effbf9-76e5-46d2-a21d-ee2a72cc8757",
		RegionID: 0,
	}

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.Create(ctx, input)

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
	g.Expect(cloudInstance.PublicIpv6Address).To(BeNil())
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:28 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:28 +0000 UTC"))
}

func TestCloudInstancesGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_instances/get_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.Get(ctx, "BDbDxbl2")

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesUpdate(t *testing.T) {
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

	cloudInstance, err := client.CloudInstances.Update(ctx, "BDbDxbl2", CloudInstanceUpdateInput{Name: &newName})

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesDelete(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.CloudInstances.Delete(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
}

func TestCloudInstancesReinstall(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/reinstall").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/reinstall_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.Reinstall(ctx, "BDbDxbl2", CloudInstanceReinstallInput{ImageID: "18e1cc16-b380-4c37-8ec1-b9d306961aae"})

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesRescue(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/rescue").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/rescue_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.Rescue(ctx, "BDbDxbl2")

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesUnrescue(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/unrescue").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/unrescue_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.Unrescue(ctx, "BDbDxbl2")

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesUpgrade(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/upgrade").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/upgrade_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.Upgrade(ctx, "BDbDxbl2", CloudInstanceUpgradeInput{FlavorID: "103"})

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesRevertUpgrade(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/revert_upgrade").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/revert_upgrade_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.RevertUpgrade(ctx, "BDbDxbl2")

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesApproveUpgrade(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/approve_upgrade").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/approve_upgrade_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.ApproveUpgrade(ctx, "BDbDxbl2")

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesPowerOn(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/switch_power_on").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/power_on_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.PowerOn(ctx, "BDbDxbl2")

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesPowerOff(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/BDbDxbl2/switch_power_off").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_instances/power_off_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudInstance, err := client.CloudInstances.PowerOff(ctx, "BDbDxbl2")

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
	g.Expect(*cloudInstance.PublicIpv6Address).To(Equal("::1"))
	g.Expect(cloudInstance.Created.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
	g.Expect(cloudInstance.Updated.String()).To(Equal("2020-04-22 06:22:32 +0000 UTC"))
}

func TestCloudInstancesCreatePTRRecord(t *testing.T) {
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

	ptrRecord, err := client.CloudInstances.CreatePTRRecord(ctx, "xkazYeJ0", input)

	g.Expect(err).To(BeNil())
	g.Expect(ptrRecord).ToNot(BeNil())

	g.Expect(ptrRecord.ID).To(Equal("oQeZzvep"))
	g.Expect(ptrRecord.IP).To(Equal("100.0.0.4"))
	g.Expect(ptrRecord.Domain).To(Equal("ai.privateservergrid.com"))
	g.Expect(ptrRecord.Priority).To(Equal(3))
	g.Expect(ptrRecord.TTL).To(Equal(60))
}

func TestCloudInstancesDeletePTRRecord(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances/xkazYeJ0/ptr_records/oQeZzvep").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.CloudInstances.DeletePTRRecord(ctx, "xkazYeJ0", "oQeZzvep")

	g.Expect(err).To(BeNil())
}
