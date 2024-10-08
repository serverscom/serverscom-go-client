package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestHostsEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Hosts.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostsCreateDedicatedServers(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	driveModelID := int64(1)
	osUbuntuServerID := int64(1)
	rootFilesystem := "ext4"
	raidLevel := 0

	input := DedicatedServerCreateInput{
		ServerModelID: int64(1),
		LocationID:    int64(1),
		RAMSize:       32,
		UplinkModels: DedicatedServerUplinkModelsInput{
			Public:  &DedicatedServerPublicUplinkInput{ID: int64(1), BandwidthModelID: int64(1)},
			Private: DedicatedServerPrivateUplinkInput{ID: int64(2)},
		},
		Drives: DedicatedServerDrivesInput{
			Slots: []DedicatedServerSlotInput{
				{Position: 0, DriveModelID: &driveModelID},
				{Position: 1, DriveModelID: &driveModelID},
			},
			Layout: []DedicatedServerLayoutInput{
				{
					SlotPositions: []int{0, 1},
					Raid:          &raidLevel,
					Partitions: []DedicatedServerLayoutPartitionInput{
						{Target: "swap", Size: 4096, Fill: false},
						{Target: "/", Fs: &rootFilesystem, Size: 100000, Fill: true},
					},
				},
			},
		},
		IPv6:              true,
		OperatingSystemID: &osUbuntuServerID,
		SSHKeyFingerprints: []string{
			"48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8",
		},
		Hosts: []DedicatedServerHostInput{
			{Hostname: "example.aa"},
			{Hostname: "example.bb"},
		},
	}

	ctx := context.TODO()

	dedicatedServers, err := client.Hosts.CreateDedicatedServers(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(len(dedicatedServers)).To(Equal(2))

	dedicatedServer := dedicatedServers[0]

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("init"))
	g.Expect(dedicatedServer.Configuration).To(Equal("Dell chassis-9015 / 2 GB RAM / 2 x hdd-model-404"))
	g.Expect(dedicatedServer.PrivateIPv4Address).To(BeNil())
	g.Expect(dedicatedServer.PublicIPv4Address).To(BeNil())
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))

	dedicatedServer = dedicatedServers[1]

	g.Expect(dedicatedServer.ID).To(Equal("w9aAOdvM"))
	g.Expect(dedicatedServer.Title).To(Equal("example.bb"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("init"))
	g.Expect(dedicatedServer.Configuration).To(Equal("Dell chassis-9015 / 2 GB RAM / 2 x hdd-model-404"))
	g.Expect(dedicatedServer.PrivateIPv4Address).To(BeNil())
	g.Expect(dedicatedServer.PublicIPv4Address).To(BeNil())
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
}

func TestHostsGetDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.GetDedicatedServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsScheduleReleaseForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/schedule_release").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/schedule_release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.ScheduleReleaseForDedicatedServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))

	scheduledRelease := *dedicatedServer.ScheduledRelease

	g.Expect(scheduledRelease.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsAbortReleaseForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/abort_release").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/abort_release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.AbortReleaseForDedicatedServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOnDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/power_on").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.PowerOnDedicatedServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOffDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/power_off").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.PowerOffDedicatedServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerCycleDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/power_cycle").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.PowerCycleDedicatedServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsDedicatedServerPowerFeeds(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/power_feeds").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/power_feeds_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	powerFeeds, err := client.Hosts.DedicatedServerPowerFeeds(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(len(powerFeeds)).To(Equal(2))

	powerFeed := powerFeeds[0]

	g.Expect(powerFeed.Name).To(Equal("Power 2"))
	g.Expect(powerFeed.Status).To(Equal("on"))

	powerFeed = powerFeeds[1]

	g.Expect(powerFeed.Name).To(Equal("Power 1"))
	g.Expect(powerFeed.Status).To(Equal("on"))
}

func TestHostsCreatePTRRecordForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/ptr_records").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/ptr_record_create_response.json").
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

	ptrRecord, err := client.Hosts.CreatePTRRecordForDedicatedServer(ctx, "xkazYeJ0", input)

	g.Expect(err).To(BeNil())
	g.Expect(ptrRecord).ToNot(BeNil())

	g.Expect(ptrRecord.ID).To(Equal("oQeZzvep"))
	g.Expect(ptrRecord.IP).To(Equal("100.0.0.4"))
	g.Expect(ptrRecord.Domain).To(Equal("ai.privateservergrid.com"))
	g.Expect(ptrRecord.Priority).To(Equal(3))
	g.Expect(ptrRecord.TTL).To(Equal(60))
}

func TestHostsDeletePTRRecordForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/ptr_records/oQeZzvep").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.Hosts.DeletePTRRecordForDedicatedServer(ctx, "xkazYeJ0", "oQeZzvep")

	g.Expect(err).To(BeNil())
}

func TestHostsReinstallOperatingSystemForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/reinstall").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	osUbuntuServerID := int64(1)
	rootFilesystem := "ext4"
	raidLevel := 0

	input := OperatingSystemReinstallInput{
		Hostname: "new-hostname",
		Drives: OperatingSystemReinstallDrivesInput{
			Layout: []OperatingSystemReinstallLayoutInput{
				{
					SlotPositions: []int{0, 1},
					Raid:          &raidLevel,
					Partitions: []OperatingSystemReinstallPartitionInput{
						{Target: "swap", Size: 4096, Fill: false},
						{Target: "/", Fs: &rootFilesystem, Size: 100000, Fill: true},
					},
				},
			},
		},
		OperatingSystemID:  &osUbuntuServerID,
		SSHKeyFingerprints: []string{"48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8"},
	}

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.ReinstallOperatingSystemForDedicatedServer(ctx, "xkazYeJ0", input)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestDedicatedServerConnectionsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/connections").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Hosts.DedicatedServerConnections("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDedicatedServerNetworksCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/networks").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Hosts.DedicatedServerNetworks("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDedicatedServerDriveSlotsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Hosts.DedicatedServerDriveSlots("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDedicatedServerPTRRecordsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/ptr_records").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Hosts.DedicatedServerPTRRecords("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostsCreateSBMServers(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := SBMServerCreateInput{
		LocationID:    int64(1),
		FlavorModelID: int64(1),
		Hosts: []SBMServerHostInput{
			{Hostname: "example.aa"},
			{Hostname: "example.bb"},
		},
	}

	ctx := context.TODO()

	sbmServers, err := client.Hosts.CreateSBMServers(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(len(sbmServers)).To(Equal(2))

	sbmServer := sbmServers[0]

	g.Expect(sbmServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("init"))
	g.Expect(sbmServer.Configuration).To(Equal("Dell chassis-9015 / 2 GB RAM / 2 x hdd-model-404"))
	g.Expect(sbmServer.PrivateIPv4Address).To(BeNil())
	g.Expect(sbmServer.PublicIPv4Address).To(BeNil())
	g.Expect(sbmServer.ScheduledRelease).To(BeNil())
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))

	sbmServer = sbmServers[1]

	g.Expect(sbmServer.ID).To(Equal("w9aAOdvM"))
	g.Expect(sbmServer.Title).To(Equal("example.bb"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("init"))
	g.Expect(sbmServer.Configuration).To(Equal("Dell chassis-9015 / 2 GB RAM / 2 x hdd-model-404"))
	g.Expect(sbmServer.PrivateIPv4Address).To(BeNil())
	g.Expect(sbmServer.PublicIPv4Address).To(BeNil())
	g.Expect(sbmServer.ScheduledRelease).To(BeNil())
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
}

func TestHostsGetSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/xkazYeJ0").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.GetSBMServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.ScheduledRelease).To(BeNil())
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsReleaseSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/xkazYeJ0").
		WithRequestMethod("DELETE").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.ReleaseSBMServer(ctx, "xkazYeJ0")

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal("xkazYeJ0"))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}
