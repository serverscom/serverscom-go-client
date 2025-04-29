package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLocationsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetLocation(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	location, err := client.Locations.GetLocation(ctx, 1108)

	g.Expect(err).To(BeNil())
	g.Expect(location).ToNot(BeNil())

	g.Expect(location.ID).To(Equal(int64(1108)))
	g.Expect(location.Name).To(Equal("location2251"))
	g.Expect(location.Status).To(Equal("active"))
	g.Expect(location.Code).To(Equal("location2251"))
	g.Expect(location.SupportedFeatures).To(Equal([]string{
		"disaggregated_public_ports",
		"disaggregated_private_ports",
		"no_public_network",
		"no_private_ip",
		"no_public_ip_address",
		"host_rescue_mode",
		"oob_public_access",
	}))
	g.Expect(location.L2SegmentsEnabled).To(BeFalse())
	g.Expect(location.PrivateRacksEnabled).To(BeFalse())
	g.Expect(location.LoadBalancersEnabled).To(BeFalse())
}

func TestServerModelOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.ServerModelOptions(int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetServerModelOption(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108/order_options/server_models/231").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/server_model_option_get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	serverModel, err := client.Locations.GetServerModelOption(ctx, 1108, 231)

	g.Expect(err).To(BeNil())
	g.Expect(serverModel).ToNot(BeNil())

	g.Expect(serverModel.ID).To(Equal(int64(231)))
	g.Expect(serverModel.Name).To(Equal("server-model-469"))
	g.Expect(serverModel.CPUName).To(Equal("cpu-model-model-468"))
	g.Expect(serverModel.CPUCount).To(Equal(2))
	g.Expect(serverModel.CPUCoresCount).To(Equal(16))
	g.Expect(serverModel.CPUFrequency).To(Equal(1970))
	g.Expect(serverModel.RAM).To(Equal(32))
	g.Expect(serverModel.RAMType).To(Equal("DDR3"))
	g.Expect(serverModel.MaxRAM).To(Equal(384))
	g.Expect(serverModel.HasRAIDController).To(BeTrue())
	g.Expect(serverModel.RAIDControllerName).To(Equal("AHCI controller"))
	g.Expect(serverModel.DriveSlotsCount).To(Equal(8))

	g.Expect(serverModel.DriveSlots).To(HaveLen(2))

	g.Expect(serverModel.DriveSlots[0].Position).To(Equal(0))
	g.Expect(serverModel.DriveSlots[0].Interface).To(Equal("SAS"))
	g.Expect(serverModel.DriveSlots[0].FormFactor).To(Equal("2_5"))
	g.Expect(serverModel.DriveSlots[0].DriveModelID).To(Equal(int64(234)))
	g.Expect(serverModel.DriveSlots[0].HotSwappable).To(BeFalse())

	g.Expect(serverModel.DriveSlots[1].Position).To(Equal(1))
	g.Expect(serverModel.DriveSlots[1].Interface).To(Equal("SAS"))
	g.Expect(serverModel.DriveSlots[1].FormFactor).To(Equal("2_5"))
	g.Expect(serverModel.DriveSlots[1].DriveModelID).To(Equal(int64(235)))
	g.Expect(serverModel.DriveSlots[1].HotSwappable).To(BeFalse())
}

func TestRAMOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/2/ram").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.RAMOptions(int64(1), int64(2))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestOperatingSystemOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/2/operating_systems").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.OperatingSystemOptions(int64(1), int64(2))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetOperatingSystemOption(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108/order_options/server_models/231/operating_systems/50").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/operating_system_option_get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	osOption, err := client.Locations.GetOperatingSystemOption(ctx, 1108, 231, 50)

	g.Expect(err).To(BeNil())
	g.Expect(osOption).ToNot(BeNil())

	g.Expect(osOption.ID).To(Equal(int64(50)))
	g.Expect(osOption.FullName).To(Equal("Ubuntu 18.04-server x86_64"))
	g.Expect(osOption.Name).To(Equal("Ubuntu"))
	g.Expect(osOption.Version).To(Equal("18.04-server"))
	g.Expect(osOption.Arch).To(Equal("x86_64"))
	g.Expect(osOption.Filesystems).To(ContainElements("ext2", "ext4", "swap", "xfs", "reiser"))
}

func TestDriveModelOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/2/drive_models").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.DriveModelOptions(int64(1), int64(2))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetDriveModel(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108/order_options/server_models/231/drive_models/369").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/drive_model_get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	driveModel, err := client.Locations.GetDriveModelOption(ctx, 1108, 231, 369)

	g.Expect(err).To(BeNil())
	g.Expect(driveModel).ToNot(BeNil())

	g.Expect(driveModel.ID).To(Equal(int64(369)))
	g.Expect(driveModel.Name).To(Equal("ssd-model-504"))
	g.Expect(driveModel.Capacity).To(Equal(100))
	g.Expect(driveModel.Interface).To(Equal("SATA3"))
	g.Expect(driveModel.FormFactor).To(Equal("2.5"))
	g.Expect(driveModel.MediaType).To(Equal("SSD"))
}

func TestUplinkOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/2/uplink_models").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.UplinkOptions(int64(1), int64(2))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetUplinkOption(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108/order_options/server_models/231/uplink_models/294").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/uplink_option_get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	uplinkOption, err := client.Locations.GetUplinkOption(ctx, 1108, 231, 294)

	g.Expect(err).To(BeNil())
	g.Expect(uplinkOption).ToNot(BeNil())

	g.Expect(uplinkOption.ID).To(Equal(int64(294)))
	g.Expect(uplinkOption.Name).To(Equal("Private 1 Gbps without redundancy"))
	g.Expect(uplinkOption.Type).To(Equal("private"))
	g.Expect(uplinkOption.Speed).To(Equal(1000))
	g.Expect(uplinkOption.Redundancy).To(BeFalse())
}

func TestBandwidthOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/2/uplink_models/3/bandwidth").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.BandwidthOptions(int64(1), int64(2), int64(3))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetBandwidthOption(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108/order_options/server_models/231/uplink_models/294/bandwidth/348").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/bandwidth_option_get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	bandwidthOption, err := client.Locations.GetBandwidthOption(ctx, 1108, 231, 294, 348)

	g.Expect(err).To(BeNil())
	g.Expect(bandwidthOption).ToNot(BeNil())

	g.Expect(bandwidthOption.ID).To(Equal(int64(348)))
	g.Expect(bandwidthOption.Name).To(Equal("20002 GB"))
	g.Expect(bandwidthOption.Type).To(Equal("bytes"))
	g.Expect(*bandwidthOption.Commit).To(Equal(int64(20002000000)))
}

func TestSBMFlavorOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/sbm_flavor_models").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.SBMFlavorOptions(int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetSBMFlavor(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108/order_options/sbm_flavor_models/119").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/sbm_flavor_get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmFlavor, err := client.Locations.GetSBMFlavorOption(ctx, 1108, 119)

	g.Expect(err).To(BeNil())
	g.Expect(sbmFlavor).ToNot(BeNil())

	g.Expect(sbmFlavor.ID).To(Equal(int64(119)))
	g.Expect(sbmFlavor.Name).To(Equal("P-101"))
	g.Expect(sbmFlavor.CPUName).To(Equal("cpu_name"))
	g.Expect(sbmFlavor.CPUCount).To(Equal(1))
	g.Expect(sbmFlavor.CPUCoresCount).To(Equal(2))
	g.Expect(sbmFlavor.CPUFrequency).To(Equal("3.8"))
	g.Expect(sbmFlavor.RAMSize).To(Equal(4096))
	g.Expect(sbmFlavor.DrivesConfiguration).To(Equal(""))
	g.Expect(sbmFlavor.PublicUplinkModelID).To(Equal(117))
	g.Expect(sbmFlavor.PublicUplinkModelName).To(Equal("uplink-model-name-36"))
	g.Expect(sbmFlavor.PrivateUplinkModelID).To(Equal(116))
	g.Expect(sbmFlavor.PrivateUplinkModelName).To(Equal("uplink-model-name-35"))
	g.Expect(sbmFlavor.BandwidthID).To(Equal(118))
	g.Expect(sbmFlavor.BandwidthName).To(Equal("public-bandwidth-model-21"))
}

func TestSBMOperatingSystemOptionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/sbm_flavor_models/1/operating_systems").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Locations.SBMOperatingSystemOptions(int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsGetSBMOperatingSystemOption(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1108/order_options/sbm_flavor_models/119/operating_systems/50").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/locations/operating_system_option_get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	osOption, err := client.Locations.GetSBMOperatingSystemOption(ctx, 1108, 119, 50)

	g.Expect(err).To(BeNil())
	g.Expect(osOption).ToNot(BeNil())

	g.Expect(osOption.ID).To(Equal(int64(50)))
	g.Expect(osOption.FullName).To(Equal("Ubuntu 18.04-server x86_64"))
	g.Expect(osOption.Name).To(Equal("Ubuntu"))
	g.Expect(osOption.Version).To(Equal("18.04-server"))
	g.Expect(osOption.Arch).To(Equal("x86_64"))
	g.Expect(osOption.Filesystems).To(ContainElements("ext2", "ext4", "swap", "xfs", "reiser"))
}
