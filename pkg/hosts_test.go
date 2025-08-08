package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	serverID  = "xkazYeJ0"
	networkID = "pen5zld7"
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
			{Hostname: "example.aa", Labels: map[string]string{"env": "test"}},
			{Hostname: "example.bb", Labels: map[string]string{"env": "test"}},
		},
	}

	ctx := context.TODO()

	dedicatedServers, err := client.Hosts.CreateDedicatedServers(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(len(dedicatedServers)).To(Equal(2))

	dedicatedServer := dedicatedServers[0]

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("init"))
	g.Expect(dedicatedServer.Configuration).To(Equal("Dell chassis-9015 / 2 GB RAM / 2 x hdd-model-404"))
	g.Expect(dedicatedServer.PrivateIPv4Address).To(BeNil())
	g.Expect(dedicatedServer.PublicIPv4Address).To(BeNil())
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
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
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
}

func TestHostsGetDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.GetDedicatedServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsUpdateDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}

	dedicatedServer, err := client.Hosts.UpdateDedicatedServer(ctx, serverID, DedicatedServerUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(newLabels))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsScheduleReleaseForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/schedule_release").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/schedule_release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.ScheduleReleaseForDedicatedServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))

	scheduledRelease := *dedicatedServer.ScheduledRelease

	g.Expect(scheduledRelease.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsAbortReleaseForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/abort_release").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/abort_release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.AbortReleaseForDedicatedServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOnDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/power_on").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.PowerOnDedicatedServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOffDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/power_off").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.PowerOffDedicatedServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerCycleDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/power_cycle").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.PowerCycleDedicatedServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(dedicatedServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(dedicatedServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsDedicatedServerPowerFeeds(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/power_feeds").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/power_feeds_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	powerFeeds, err := client.Hosts.DedicatedServerPowerFeeds(ctx, serverID)

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
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/ptr_records").
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

	ptrRecord, err := client.Hosts.CreatePTRRecordForDedicatedServer(ctx, serverID, input)

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
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/ptr_records/oQeZzvep").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.Hosts.DeletePTRRecordForDedicatedServer(ctx, serverID, "oQeZzvep")

	g.Expect(err).To(BeNil())
}

func TestHostsReinstallOperatingSystemForDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/reinstall").
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

	dedicatedServer, err := client.Hosts.ReinstallOperatingSystemForDedicatedServer(ctx, serverID, input)

	g.Expect(err).To(BeNil())
	g.Expect(dedicatedServer).ToNot(BeNil())

	g.Expect(dedicatedServer.ID).To(Equal(serverID))
	g.Expect(dedicatedServer.Title).To(Equal("example.aa"))
	g.Expect(dedicatedServer.LocationID).To(Equal(int64(1)))
	g.Expect(dedicatedServer.Status).To(Equal("active"))
	g.Expect(dedicatedServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*dedicatedServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*dedicatedServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(dedicatedServer.ScheduledRelease).To(BeNil())
	g.Expect(dedicatedServer.Labels).To(Equal(map[string]string{"env": "test"}))
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
			{Hostname: "example.aa", Labels: map[string]string{"env": "test"}},
			{Hostname: "example.bb", Labels: map[string]string{"env": "test"}},
		},
	}

	ctx := context.TODO()

	sbmServers, err := client.Hosts.CreateSBMServers(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(len(sbmServers)).To(Equal(2))

	sbmServer := sbmServers[0]

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("init"))
	g.Expect(sbmServer.Configuration).To(Equal("Dell chassis-9015 / 2 GB RAM / 2 x hdd-model-404"))
	g.Expect(sbmServer.PrivateIPv4Address).To(BeNil())
	g.Expect(sbmServer.PublicIPv4Address).To(BeNil())
	g.Expect(sbmServer.ScheduledRelease).To(BeNil())
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
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
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:04 +0000 UTC"))
}

func TestHostsGetSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/" + serverID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.GetSBMServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.ScheduledRelease).To(BeNil())
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsUpdateSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/" + serverID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}

	sbmServer, err := client.Hosts.UpdateSBMServer(ctx, serverID, SBMServerUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.ScheduledRelease).To(BeNil())
	g.Expect(sbmServer.Labels).To(Equal(newLabels))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsReleaseSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/" + serverID).
		WithRequestMethod("DELETE").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.ReleaseSBMServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOnSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/" + serverID + "/power_on").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.PowerOnSBMServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOffSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/" + serverID + "/power_off").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.PowerOffSBMServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerCycleSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/" + serverID + "/power_cycle").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.PowerCycleSBMServer(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestReinstallOperatingSystemForSBMServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/sbm_servers/" + serverID + "/reinstall").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/sbm_servers/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	sbmServer, err := client.Hosts.ReinstallOperatingSystemForSBMServer(ctx, serverID, SBMOperatingSystemReinstallInput{Hostname: "example.aa", OperatingSystemID: 1})

	g.Expect(err).To(BeNil())
	g.Expect(sbmServer).ToNot(BeNil())

	g.Expect(sbmServer.ID).To(Equal(serverID))
	g.Expect(sbmServer.Title).To(Equal("example.aa"))
	g.Expect(sbmServer.Type).To(Equal("sbm_server"))
	g.Expect(sbmServer.LocationID).To(Equal(int64(1)))
	g.Expect(sbmServer.Status).To(Equal("active"))
	g.Expect(sbmServer.Configuration).To(Equal("REMM R123"))
	g.Expect(*sbmServer.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*sbmServer.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(sbmServer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(sbmServer.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(sbmServer.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsGetKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/kubernetes_baremetal_nodes/" + serverID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/kubernetes_baremetal_nodes/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	node, err := client.Hosts.GetKubernetesBaremetalNode(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(node).ToNot(BeNil())

	g.Expect(node.ID).To(Equal(serverID))
	g.Expect(node.Title).To(Equal("example.aa"))
	g.Expect(node.LocationID).To(Equal(int64(1)))
	g.Expect(node.Status).To(Equal("active"))
	g.Expect(node.Configuration).To(Equal("REMM R123"))
	g.Expect(*node.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*node.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(node.ScheduledRelease).To(BeNil())
	g.Expect(node.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(node.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(node.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsUpdateKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/kubernetes_baremetal_nodes/" + serverID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/hosts/kubernetes_baremetal_nodes/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}
	node, err := client.Hosts.UpdateKubernetesBaremetalNode(ctx, serverID, KubernetesBaremetalNodeUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(node).ToNot(BeNil())

	g.Expect(node.ID).To(Equal(serverID))
	g.Expect(node.Title).To(Equal("example.aa"))
	g.Expect(node.LocationID).To(Equal(int64(1)))
	g.Expect(node.Status).To(Equal("active"))
	g.Expect(node.Configuration).To(Equal("REMM R123"))
	g.Expect(*node.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*node.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(node.ScheduledRelease).To(BeNil())
	g.Expect(node.Labels).To(Equal(newLabels))
	g.Expect(node.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(node.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOnKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/kubernetes_baremetal_nodes/" + serverID + "/power_on").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/kubernetes_baremetal_nodes/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	node, err := client.Hosts.PowerOnKubernetesBaremetalNode(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(node).ToNot(BeNil())

	g.Expect(node.ID).To(Equal(serverID))
	g.Expect(node.Title).To(Equal("example.aa"))
	g.Expect(node.Type).To(Equal("kubernetes_baremetal_node"))
	g.Expect(node.LocationID).To(Equal(int64(1)))
	g.Expect(node.Status).To(Equal("active"))
	g.Expect(node.Configuration).To(Equal("REMM R123"))
	g.Expect(*node.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*node.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(node.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(node.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(node.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerOffKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/kubernetes_baremetal_nodes/" + serverID + "/power_off").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/kubernetes_baremetal_nodes/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	node, err := client.Hosts.PowerOffKubernetesBaremetalNode(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(node).ToNot(BeNil())

	g.Expect(node.ID).To(Equal(serverID))
	g.Expect(node.Title).To(Equal("example.aa"))
	g.Expect(node.Type).To(Equal("kubernetes_baremetal_node"))
	g.Expect(node.LocationID).To(Equal(int64(1)))
	g.Expect(node.Status).To(Equal("active"))
	g.Expect(node.Configuration).To(Equal("REMM R123"))
	g.Expect(*node.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*node.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(node.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(node.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(node.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsPowerCycleKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/kubernetes_baremetal_nodes/" + serverID + "/power_cycle").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/kubernetes_baremetal_nodes/get_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	node, err := client.Hosts.PowerCycleKubernetesBaremetalNode(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(node).ToNot(BeNil())

	g.Expect(node.ID).To(Equal(serverID))
	g.Expect(node.Title).To(Equal("example.aa"))
	g.Expect(node.Type).To(Equal("kubernetes_baremetal_node"))
	g.Expect(node.LocationID).To(Equal(int64(1)))
	g.Expect(node.Status).To(Equal("active"))
	g.Expect(node.Configuration).To(Equal("REMM R123"))
	g.Expect(*node.PrivateIPv4Address).To(Equal("10.0.0.1"))
	g.Expect(*node.PublicIPv4Address).To(Equal("169.254.0.1"))
	g.Expect(node.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(node.Created.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
	g.Expect(node.Updated.String()).To(Equal("2020-04-22 06:22:02 +0000 UTC"))
}

func TestHostsGetDedicatedServerNetworkUsage(t *testing.T) {
	g := NewGomegaWithT(t)
	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/network_utilization").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/network_utilization.json").
		WithResponseCode(200).
		Build()
	defer ts.Close()

	ctx := context.TODO()
	usage, err := client.Hosts.GetDedicatedServerNetworkUsage(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(usage).ToNot(BeNil())
	g.Expect(usage.Type).To(Equal("traffic"))
	g.Expect(usage.Utilization).ToNot(BeNil())
	g.Expect(usage.Utilization.Value).To(Equal(int64(2000000)))
	g.Expect(usage.Utilization.Commit).To(Equal(int64(1000000)))
	g.Expect(usage.Utilization.Unit).To(Equal("KB"))
}

func TestHostsGetDedicatedServerNetwork(t *testing.T) {
	g := NewGomegaWithT(t)
	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/networks/" + networkID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_network_response.json").
		WithResponseCode(200).
		Build()
	defer ts.Close()

	ctx := context.TODO()
	network, err := client.Hosts.GetDedicatedServerNetwork(ctx, serverID, networkID)

	g.Expect(err).To(BeNil())
	g.Expect(network).ToNot(BeNil())
	g.Expect(network.ID).To(Equal(networkID))
	g.Expect(*network.Title).To(Equal("Public"))
	g.Expect(network.Status).To(Equal("active"))
	g.Expect(*network.Cidr).To(Equal("100.0.8.0/29"))
	g.Expect(network.Family).To(Equal("ipv4"))
	g.Expect(network.InterfaceType).To(Equal("public"))
	g.Expect(network.DistributionMethod).To(Equal("gateway"))
	g.Expect(network.Additional).To(Equal(false))
	g.Expect(network.Created.String()).To(Equal("2025-07-31 11:03:36 +0000 UTC"))
	g.Expect(network.Updated.String()).To(Equal("2025-07-31 11:03:36 +0000 UTC"))
}

func TestHostsAddDedicatedServerPublicNetwork(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/networks/public_ipv4").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_network_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	input := NetworkInput{
		DistributionMethod: "gateway",
		Mask:               29,
	}

	ctx := context.TODO()
	network, err := client.Hosts.AddDedicatedServerPublicIPv4Network(ctx, serverID, input)

	g.Expect(err).To(BeNil())
	g.Expect(network).ToNot(BeNil())
	g.Expect(network.ID).ToNot(BeEmpty())
}

func TestHostsAddDedicatedServerPrivateNetwork(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/networks/private_ipv4").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_network_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	input := NetworkInput{
		DistributionMethod: "gateway",
		Mask:               29,
	}

	ctx := context.TODO()
	network, err := client.Hosts.AddDedicatedServerPrivateIPv4Network(ctx, serverID, input)

	g.Expect(err).To(BeNil())
	g.Expect(network).ToNot(BeNil())
	g.Expect(network.ID).ToNot(BeEmpty())
}

func TestHostsActivateDedicatedServerPublicIPv6Network(t *testing.T) {
	g := NewGomegaWithT(t)
	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/networks/public_ipv6").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_network_response.json").
		WithResponseCode(202).
		Build()
	defer ts.Close()

	ctx := context.TODO()
	network, err := client.Hosts.ActivateDedicatedServerPubliIPv6Network(ctx, serverID)

	g.Expect(err).To(BeNil())
	g.Expect(network).ToNot(BeNil())
	g.Expect(network.ID).ToNot(BeEmpty())
}

func TestHostsDeleteDedicatedServerNetwork(t *testing.T) {
	g := NewGomegaWithT(t)
	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/" + serverID + "/networks/" + networkID).
		WithRequestMethod("DELETE").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_network_response.json").
		WithResponseCode(202).
		Build()
	defer ts.Close()

	ctx := context.TODO()
	network, err := client.Hosts.DeleteDedicatedServerNetwork(ctx, serverID, networkID)

	g.Expect(err).To(BeNil())
	g.Expect(network).ToNot(BeNil())
	g.Expect(network.ID).To(Equal(networkID))
}
