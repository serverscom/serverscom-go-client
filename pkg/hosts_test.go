package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestDedicatedServersCreate(t *testing.T) {
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
	bandwidthID := int64(1)
	rootFilesystem := "ext4"
	raidLevel := 0

	input := DedicatedServerCreateInput{
		ServerModelID: int64(1),
		LocationID:    int64(1),
		RAMSize:       32,
		UplinkModels: DedicatedServerUplinkModelsInput{
			Public:  &DedicatedServerPublicUplinkInput{ID: int64(1), BandwidthModelID: &bandwidthID},
			Private: DedicatedServerPrivateUplinkInput{ID: int64(2)},
		},
		Drives: DedicatedServerDrivesInput{
			Slots: []DedicatedServerSlotInput{
				DedicatedServerSlotInput{Position: 0, DriveModelID: &driveModelID},
				DedicatedServerSlotInput{Position: 1, DriveModelID: &driveModelID},
			},
			Layout: []DedicatedServerLayoutInput{
				DedicatedServerLayoutInput{
					SlotPositions: []int{0, 1},
					Raid:          &raidLevel,
					Partitions: []DedicatedServerLayoutPartitionInput{
						DedicatedServerLayoutPartitionInput{Target: "swap", Size: 4096, Fill: false},
						DedicatedServerLayoutPartitionInput{Target: "/", Fs: &rootFilesystem, Size: 100000, Fill: true},
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
			DedicatedServerHostInput{Hostname: "example.aa"},
			DedicatedServerHostInput{Hostname: "example.bb"},
		},
	}

	ctx := context.TODO()

	dedicatedServers, err := client.Hosts.DedicatedServersCreate(ctx, input)

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

func TestDedicatedServersGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.DedicatedServerGet(ctx, "xkazYeJ0")

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

func TestDedicatedServersScheduleRelease(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/schedule_release").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/schedule_release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.DedicatedServerScheduleRelease(ctx, "xkazYeJ0")

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

func TestDedicatedServersAbortRelease(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/xkazYeJ0/abort_release").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/hosts/dedicated_servers/abort_release_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dedicatedServer, err := client.Hosts.DedicatedServerAbortRelease(ctx, "xkazYeJ0")

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
