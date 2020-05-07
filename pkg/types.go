package serverscom

import (
	"time"
)

// Location represents location
type Location struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// SSLCertificate represents ssl certificate
type SSLCertificate struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Sha1Fingerprint string     `json:"sha1_fingerprint"`
	Expires         *time.Time `json:"expires_at"`
	Created         time.Time  `json:"created_at"`
	Updated         time.Time  `json:"updated_at"`
}

// SSLCertificateCustom represents custom ssl certificate
type SSLCertificateCustom struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Sha1Fingerprint string     `json:"sha1_fingerprint"`
	Expires         *time.Time `json:"expires_at"`
	Created         time.Time  `json:"created_at"`
	Updated         time.Time  `json:"updated_at"`
}

// SSLCertificateCreateCustomInput represents custom ssl certificate create input
type SSLCertificateCreateCustomInput struct {
	Name       string `json:"name"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	ChainKey   string `json:"chain_key"`
}

// Host represents host
type Host struct {
	ID                 string     `json:"id"`
	Title              string     `json:"title"`
	LocationID         int64      `json:"location_id"`
	Status             string     `json:"status"`
	Configuration      string     `json:"configuration"`
	PrivateIPv4Address *string    `json:"private_ipv4_address"`
	PublicIPv4Address  *string    `json:"public_ipv4_address"`
	ScheduledRelease   *time.Time `json:"scheduled_release_at"`
	Created            time.Time  `json:"created_at"`
	Updated            time.Time  `json:"updated_at"`
}

// DedicatedServer represents dedicated server
type DedicatedServer struct {
	ID                 string     `json:"id"`
	Title              string     `json:"title"`
	LocationID         int64      `json:"location_id"`
	Status             string     `json:"status"`
	Configuration      string     `json:"configuration"`
	PrivateIPv4Address *string    `json:"private_ipv4_address"`
	PublicIPv4Address  *string    `json:"public_ipv4_address"`
	ScheduledRelease   *time.Time `json:"scheduled_release_at"`
	Created            time.Time  `json:"created_at"`
	Updated            time.Time  `json:"updated_at"`
}

// DedicatedServerLayoutPartitionInput represents partition for DedicatedServerLayoutInput
type DedicatedServerLayoutPartitionInput struct {
	Target string  `json:"target"`
	Size   int     `json:"size"`
	Fs     *string `json:"fs,omitempty"`
	Fill   bool    `json:"fill"`
}

// DedicatedServerLayoutInput represents layout for DedicatedServerDrivesInput
type DedicatedServerLayoutInput struct {
	SlotPositions []int                                 `json:"slot_positions"`
	Raid          *int                                  `json:"raid,omitempty"`
	Partitions    []DedicatedServerLayoutPartitionInput `json:"partitions"`
}

// DedicatedServerSlotInput represents slot for DedicatedServerDrivesInput
type DedicatedServerSlotInput struct {
	Position     int    `json:"position"`
	DriveModelID *int64 `json:"drive_model_id,omitempty"`
}

// DedicatedServerDrivesInput represents drives for DedicatedServerCreateInput
type DedicatedServerDrivesInput struct {
	Slots  []DedicatedServerSlotInput   `json:"slots"`
	Layout []DedicatedServerLayoutInput `json:"layout"`
}

// DedicatedServerPublicUplinkInput represents public uplink for DedicatedServerUplinkModelsInput
type DedicatedServerPublicUplinkInput struct {
	ID               int64  `json:"id"`
	BandwidthModelID *int64 `json:"bandwidth_model_id,omitempty"`
}

// DedicatedServerPrivateUplinkInput represents private uplink for DedicatedServerUplinkModelsInput
type DedicatedServerPrivateUplinkInput struct {
	ID int64 `json:"id"`
}

// DedicatedServerUplinkModelsInput represents uplinks for DedicatedServerCreateInput
type DedicatedServerUplinkModelsInput struct {
	Public  *DedicatedServerPublicUplinkInput `json:"public,omitempty"`
	Private DedicatedServerPrivateUplinkInput `json:"private"`
}

// DedicatedServerHostInput represents hosts for DedicatedServerCreateInput
type DedicatedServerHostInput struct {
	Hostname string `json:"hostname"`
}

// DedicatedServerCreateInput represents dedicated server create input, example:
//
//  driveModelID := int64(1)
//  osUbuntuServerID := int64(1)
//  rootFilesystem := "ext4"
//  raidLevel := 0
//
//  input := DedicatedServerCreateInput{
//    ServerModelID: int64(1),
//    LocationID: int64(1),
//    RAMSize: 32,
//    UplinkModels: DedicatedServerUplinkModelInput{
//      PublicUplink &DedicatedServerPublicUplinkInput{ID: int64(1), BandwidthModelID: int64(1)},
//      PrivateUplink: DedicatedServerPrivateUplinkInput{ID: int64(2)},
//    },
//    Drives: DedicatedServerDrivesInput{
//      Slots: []DedicatedServerSlotInput{
//        DedicatedServerSlotInput{Position: 0, DriveModelID: &driveModelID},
//        DedicatedServerSlotInput{Position: 1, DriveModelID: &driveModelID},
//      },
//      Layout: []DedicatedServerLayoutInput{
//        DedicatedServerLayoutInput{
//          SlotPositions: []int{0, 1},
//          Riad:          &raidLevel,
//          Partitions:    []DedicatedServerLayoutPartitionInput{
//            DedicatedServerLayoutPartitionInput{Target: "swap", Size: 4096, Fill: false},
//            DedicatedServerLayoutPartitionInput{Target: "/", Fs: &rootFilesystem, Size: 100000, Fill: true},
//          },
//        },
//      },
//    },
//    IPv6: true,
//    OperatingSystemID: &osUbuntuServerID,
//    SSHKeyFingerprints: []string{
//      "48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8"
//    },
//    Hosts: []DedicatedServerHostInput{
//      Hostname: "example-host",
//    },
//  }
type DedicatedServerCreateInput struct {
	ServerModelID      int64                            `json:"server_model_id"`
	LocationID         int64                            `json:"location_id"`
	RAMSize            int                              `json:"ram_size"`
	UplinkModels       DedicatedServerUplinkModelsInput `json:"uplink_models"`
	Drives             DedicatedServerDrivesInput       `json:"drives"`
	Features           []string                         `json:"features,omitempty"`
	Ipv6               bool                             `json:"ipv6"`
	Hosts              []DedicatedServerHostInput       `json:"hosts"`
	OperatingSystemID  *int64                           `json:"operating_system_id"`
	SSHKeyFingerprints []string                         `json:"ssh_key_fingerprints,omitempty"`
}

// ServerModelOption represents server model option
type ServerModelOption struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// RAMOption represents ram option
type RAMOption struct {
	RAM  int    `json:"ram"`
	Type string `json:"type"`
}

// OperatingSystemOption represents operating system option
type OperatingSystemOption struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Arch        string   `json:"arch"`
	Filesystems []string `json:"filesystems"`
}

// UplinkOption represents uplink option
type UplinkOption struct {
	ID         int64  `json:"id"`
	Type       string `json:"type"`
	Speed      int    `json:"speed"`
	Redundancy bool   `json:"redundancy"`
}

// BandwidthOption represents bandwidth option
type BandwidthOption struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	Commit *int64 `json:"commit,omitempty"`
}

// DriveModelOption represents drive model option
type DriveModelOption struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Capacity   int    `json:"capacity"`
	Interface  string `json:"interface"`
	FormFactor string `json:"form_factor"`
	MediaType  string `json:"media_type"`
}

// SSHKey represents ssh key
type SSHKey struct {
	Name        string    `json:"name"`
	Fingerprint string    `json:"fingerprint"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}

// SSHKeyCreateInput represents ssh key create input
type SSHKeyCreateInput struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

// SSHKeyUpdateInput represents ssh key update input
type SSHKeyUpdateInput struct {
	Name string `json:"name"`
}

// CloudInstance represents cloud instance
type CloudInstance struct {
	Name               string    `json:"name"`
	ID                 string    `json:"id"`
	OpenstackUUID      string    `json:"openstack_uuid"`
	Status             string    `json:"status"`
	FlavorID           string    `json:"flavor_id"`
	ImageID            string    `json:"image_id"`
	PublicIPv4Address  *string   `json:"public_ipv4_address"`
	PrivateIPv4Address *string   `json:"private_ipv4_address"`
	PublicIpv6Address  *string   `json:"public_ipv6_address"`
	Created            time.Time `json:"created_at"`
	Updated            time.Time `json:"updated_at"`
}

// CloudInstanceCreateInput represents cloud instance create input
type CloudInstanceCreateInput struct {
	Name              string  `json:"name"`
	RegionID          int     `json:"region_id"`
	FlavorID          string  `json:"flavor_id"`
	ImageID           string  `json:"image_id"`
	GpnEnabled        *bool   `json:"gpn_enabled,omitempty"`
	Ipv6Enabled       *bool   `json:"ipv6_enabled,omitempty"`
	SSHKeyFingerprint *string `json:"ssh_key_fingerprint,omitempty"`
	BackupCopies      *int    `json:"backup_copies,omitempty"`
}

// CloudInstanceUpdateInput represents cloud instance update input
type CloudInstanceUpdateInput struct {
	Name         *string `json:"name"`
	BackupCopies *int    `json:"backup_copies"`
	GpnEnabled   *bool   `json:"gpn_enabled,omitempty"`
	Ipv6Enabled  *bool   `json:"ipv6_enabled,omitempty"`
}

// L2Segment represents l2 segment
type L2Segment struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	LocationGroupID int64     `json:"location_group_id"`
	Created         time.Time `json:"created_at"`
	Updated         time.Time `json:"updated_at"`
}

// L2SegmentMemberInput represents l2 segment member input for L2SegmentCreateInput and L2SegmentUpdateInput
type L2SegmentMemberInput struct {
	ID   string `json:"id"`
	Mode string `json:"mode"`
}

// L2SegmentCreateInput represents l2 segment create input
type L2SegmentCreateInput struct {
	Name            *string                `json:"name,omitempty"`
	Type            string                 `json:"type"`
	LocationGroupID int64                  `json:"location_group_id"`
	Members         []L2SegmentMemberInput `json:"members"`
}

// L2SegmentUpdateInput represents l2 segment update input
type L2SegmentUpdateInput struct {
	Name    *string                `json:"name,omitempty"`
	Members []L2SegmentMemberInput `json:"members,omitempty"`
}

// L2Member respresents l2 segment member
type L2Member struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Mode    string    `json:"mode"`
	Vlan    *int      `json:"vlan"`
	Status  string    `json:"status"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

// Network represents network
type Network struct {
	ID                 string    `json:"id"`
	Title              *string   `json:"title,omitempty"`
	State              string    `json:"state"`
	Cidr               *string   `json:"cidr,omitempty"`
	Family             string    `json:"family"`
	InterfaceType      string    `json:"interface_type"`
	DistributionMethod string    `json:"distribution_method"`
	Additional         bool      `json:"additional"`
	Created            time.Time `json:"created_at"`
	Updated            time.Time `json:"updated_at"`
}
