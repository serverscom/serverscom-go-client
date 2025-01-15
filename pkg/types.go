package serverscom

import (
	"time"
)

// Location represents location
type Location struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// SSLCertificate represents ssl certificate
type SSLCertificate struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Sha1Fingerprint string            `json:"sha1_fingerprint"`
	Labels          map[string]string `json:"labels"`
	Expires         *time.Time        `json:"expires_at"`
	Created         time.Time         `json:"created_at"`
	Updated         time.Time         `json:"updated_at"`
}

// SSLCertificateCustom represents custom ssl certificate
type SSLCertificateCustom struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Type            string            `json:"type"`
	Issuer          *string           `json:"issuer"`
	Subject         string            `json:"subject"`
	DomainNames     []string          `json:"domain_names"`
	Sha1Fingerprint string            `json:"sha1_fingerprint"`
	Labels          map[string]string `json:"labels"`
	Expires         *time.Time        `json:"expires_at"`
	Created         time.Time         `json:"created_at"`
	Updated         time.Time         `json:"updated_at"`
}

// SSLCertificateLE represents let's encrypt ssl certificate
type SSLCertificateLE struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Issuer      *string           `json:"issuer"`
	Subject     string            `json:"subject"`
	DomainNames []string          `json:"domain_names"`
	Labels      map[string]string `json:"labels"`
	Expires     *time.Time        `json:"expires_at"`
	Created     time.Time         `json:"created_at"`
	Updated     time.Time         `json:"updated_at"`
}

// SSLCertificateCreateCustomInput represents custom ssl certificate create input
type SSLCertificateCreateCustomInput struct {
	Name       string            `json:"name"`
	PublicKey  string            `json:"public_key"`
	PrivateKey string            `json:"private_key"`
	ChainKey   string            `json:"chain_key,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

// SSLCertificateUpdateCustomInput represents custom ssl certificate update input
type SSLCertificateUpdateCustomInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// SSLCertificateUpdateLEInput represents let's encrypt ssl certificate update input
type SSLCertificateUpdateLEInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// Host represents host
type Host struct {
	ID                 string     `json:"id"`
	Type               string     `json:"type"`
	Title              string     `json:"title"`
	LocationID         int64      `json:"location_id"`
	LocationCode       string     `json:"location_code"`
	Status             string     `json:"status"`
	OperationalStatus  string     `json:"operational_status"`
	PowerStatus        string     `json:"power_status"`
	Configuration      string     `json:"configuration"`
	PrivateIPv4Address *string    `json:"private_ipv4_address"`
	PublicIPv4Address  *string    `json:"public_ipv4_address"`
	ScheduledRelease   *time.Time `json:"scheduled_release_at"`
	Created            time.Time  `json:"created_at"`
	Updated            time.Time  `json:"updated_at"`
}

// ConfigurationDetails represents host configuration details
type ConfigurationDetails struct {
	RAMSize                 int     `json:"ram_size"`
	ServerModelID           *int64  `json:"server_model_id"`
	ServerModelName         *string `json:"server_model_name"`
	PublicUplinkID          *int64  `json:"public_uplink_id"`
	PublicUplinkName        *string `json:"public_uplink_name"`
	PrivateUplinkID         *int64  `json:"private_uplink_id"`
	PrivateUplinkName       *string `json:"private_uplink_name"`
	BandwidthID             *int64  `json:"bandwidth_id"`
	BandwidthName           *string `json:"bandwidth_name"`
	OperatingSystemID       *int64  `json:"operating_system_id"`
	OperatingSystemFullName *string `json:"operating_system_full_name"`
}

// DedicatedServer represents dedicated server
type DedicatedServer struct {
	ID                   string               `json:"id"`
	Type                 string               `json:"type"`
	Title                string               `json:"title"`
	LocationID           int64                `json:"location_id"`
	LocationCode         string               `json:"location_code"`
	Status               string               `json:"status"`
	OperationalStatus    string               `json:"operational_status"`
	PowerStatus          string               `json:"power_status"`
	Configuration        string               `json:"configuration"`
	PrivateIPv4Address   *string              `json:"private_ipv4_address"`
	PublicIPv4Address    *string              `json:"public_ipv4_address"`
	ScheduledRelease     *time.Time           `json:"scheduled_release_at"`
	ConfigurationDetails ConfigurationDetails `json:"configuration_details"`
	Labels               map[string]string    `json:"labels"`
	Created              time.Time            `json:"created_at"`
	Updated              time.Time            `json:"updated_at"`
}

// KubernetesBaremetalNode represents kubernetes baremetal node
type KubernetesBaremetalNode struct {
	ID                   string               `json:"id"`
	Type                 string               `json:"type"`
	Title                string               `json:"title"`
	LocationID           int64                `json:"location_id"`
	LocationCode         string               `json:"location_code"`
	Status               string               `json:"status"`
	OperationalStatus    string               `json:"operational_status"`
	PowerStatus          string               `json:"power_status"`
	Configuration        string               `json:"configuration"`
	PrivateIPv4Address   *string              `json:"private_ipv4_address"`
	PublicIPv4Address    *string              `json:"public_ipv4_address"`
	ScheduledRelease     *time.Time           `json:"scheduled_release_at"`
	ConfigurationDetails ConfigurationDetails `json:"configuration_details"`
	Labels               map[string]string    `json:"labels"`
	Created              time.Time            `json:"created_at"`
	Updated              time.Time            `json:"updated_at"`
}

// SBMServer represents scalable baremetal server
type SBMServer struct {
	ID                   string               `json:"id"`
	Type                 string               `json:"type"`
	Title                string               `json:"title"`
	LocationID           int64                `json:"location_id"`
	LocationCode         string               `json:"location_code"`
	Status               string               `json:"status"`
	OperationalStatus    string               `json:"operational_status"`
	PowerStatus          string               `json:"power_status"`
	Configuration        string               `json:"configuration"`
	PrivateIPv4Address   *string              `json:"private_ipv4_address"`
	PublicIPv4Address    *string              `json:"public_ipv4_address"`
	ScheduledRelease     *time.Time           `json:"scheduled_release_at"`
	ConfigurationDetails ConfigurationDetails `json:"configuration_details"`
	Labels               map[string]string    `json:"labels"`
	Created              time.Time            `json:"created_at"`
	Updated              time.Time            `json:"updated_at"`
}

// DedicatedServerLayoutPartitionInput represents partition for DedicatedServerLayoutInput
type DedicatedServerLayoutPartitionInput struct {
	Target string  `json:"target"`
	Size   int     `json:"size"`
	Fs     *string `json:"fs,omitempty"`
	Fill   bool    `json:"fill,omitempty"`
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
	ID               int64 `json:"id"`
	BandwidthModelID int64 `json:"bandwidth_model_id"`
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
	Hostname             string            `json:"hostname"`
	PublicIPv4NetworkID  *string           `json:"public_ipv4_network_id,omitempty"`
	PrivateIPv4NetworkID *string           `json:"private_ipv4_network_id,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
}

// DedicatedServerCreateInput represents dedicated server create input, example:
//
//	driveModelID := int64(1)
//	osUbuntuServerID := int64(1)
//	rootFilesystem := "ext4"
//	raidLevel := 0
//
//	input := DedicatedServerCreateInput{
//	  ServerModelID: int64(1),
//	  LocationID: int64(1),
//	  RAMSize: 32,
//	  UplinkModels: DedicatedServerUplinkModelInput{
//	    PublicUplink &DedicatedServerPublicUplinkInput{ID: int64(1), BandwidthModelID: int64(1)},
//	    PrivateUplink: DedicatedServerPrivateUplinkInput{ID: int64(2)},
//	  },
//	  Drives: DedicatedServerDrivesInput{
//	    Slots: []DedicatedServerSlotInput{
//	      DedicatedServerSlotInput{Position: 0, DriveModelID: &driveModelID},
//	      DedicatedServerSlotInput{Position: 1, DriveModelID: &driveModelID},
//	    },
//	    Layout: []DedicatedServerLayoutInput{
//	      DedicatedServerLayoutInput{
//	        SlotPositions: []int{0, 1},
//	        Riad:          &raidLevel,
//	        Partitions:    []DedicatedServerLayoutPartitionInput{
//	          DedicatedServerLayoutPartitionInput{Target: "swap", Size: 4096, Fill: false},
//	          DedicatedServerLayoutPartitionInput{Target: "/", Fs: &rootFilesystem, Size: 100000, Fill: true},
//	        },
//	      },
//	    },
//	  },
//	  IPv6: true,
//	  OperatingSystemID: &osUbuntuServerID,
//	  SSHKeyFingerprints: []string{
//	    "48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8"
//	  },
//	  Hosts: []DedicatedServerHostInput{
//		{
//	 	   Hostname: "example-host",
//		},
//	  },
//	}
type DedicatedServerCreateInput struct {
	ServerModelID      int64                            `json:"server_model_id"`
	LocationID         int64                            `json:"location_id"`
	RAMSize            int                              `json:"ram_size"`
	UplinkModels       DedicatedServerUplinkModelsInput `json:"uplink_models"`
	Drives             DedicatedServerDrivesInput       `json:"drives"`
	Features           []string                         `json:"features,omitempty"`
	IPv6               bool                             `json:"ipv6"`
	Hosts              []DedicatedServerHostInput       `json:"hosts"`
	OperatingSystemID  *int64                           `json:"operating_system_id,omitempty"`
	SSHKeyFingerprints []string                         `json:"ssh_key_fingerprints,omitempty"`
	UserData           *string                          `json:"user_data,omitempty"`
}

// SBMServerHostInput represents hosts for SBMServerCreateInput
type SBMServerHostInput struct {
	Hostname             string            `json:"hostname"`
	PublicIPv4NetworkID  *string           `json:"public_ipv4_network_id,omitempty"`
	PrivateIPv4NetworkID *string           `json:"private_ipv4_network_id,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
}

// SBMServerCreateInput represents SBM server create input, example:
//
//	input := &SBMServerCreateInput{
//		SBMFlavorModelID: 1,
//		LocationID: 2,
//		Hosts: []SBMServerHostInput{
//			{
//				Hostname: "test",
//			},
//		},
//	}
type SBMServerCreateInput struct {
	FlavorModelID      int64                `json:"sbm_flavor_model_id"`
	LocationID         int64                `json:"location_id"`
	Hosts              []SBMServerHostInput `json:"hosts"`
	OperatingSystemID  *int64               `json:"operating_system_id,omitempty"`
	SSHKeyFingerprints []string             `json:"ssh_key_fingerprints,omitempty"`
	UserData           *string              `json:"user_data,omitempty"`
}

// ServerModelOption represents server model option
type ServerModelOption struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	RAM  int    `json:"ram"`
}

// RAMOption represents ram option
type RAMOption struct {
	RAM  int    `json:"ram"`
	Type string `json:"type"`
}

// OperatingSystemOption represents operating system option
type OperatingSystemOption struct {
	ID          int64    `json:"id"`
	FullName    string   `json:"full_name"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Arch        string   `json:"arch"`
	Filesystems []string `json:"filesystems"`
}

// UplinkOption represents uplink option
type UplinkOption struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Speed      int    `json:"speed"`
	Redundancy bool   `json:"redundancy"`
}

// BandwidthOption represents bandwidth option
type BandwidthOption struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Commit *int64 `json:"commit,omitempty"`
}

// DriveModel represents drive model
type DriveModel struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Capacity   int    `json:"capacity"`
	Interface  string `json:"interface"`
	FormFactor string `json:"form_factor"`
	MediaType  string `json:"media_type"`
}

// SSHKey represents ssh key
type SSHKey struct {
	Name        string            `json:"name"`
	Fingerprint string            `json:"fingerprint"`
	Labels      map[string]string `json:"labels"`
	Created     time.Time         `json:"created_at"`
	Updated     time.Time         `json:"updated_at"`
}

// SSHKeyCreateInput represents ssh key create input
type SSHKeyCreateInput struct {
	Name      string            `json:"name"`
	PublicKey string            `json:"public_key"`
	Labels    map[string]string `json:"labels,omitempty"`
}

// SSHKeyUpdateInput represents ssh key update input
type SSHKeyUpdateInput struct {
	Name   string            `json:"name,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

// CloudComputingInstance represents cloud instance
type CloudComputingInstance struct {
	Name               string            `json:"name"`
	ID                 string            `json:"id"`
	RegionID           int64             `json:"region_id"`
	RegionCode         string            `json:"region_code"`
	OpenstackUUID      string            `json:"openstack_uuid"`
	Status             string            `json:"status"`
	FlavorID           string            `json:"flavor_id"`
	FlavorName         string            `json:"flavor_name"`
	ImageID            string            `json:"image_id"`
	ImageName          *string           `json:"image_name"`
	PublicIPv4Address  *string           `json:"public_ipv4_address"`
	PrivateIPv4Address *string           `json:"private_ipv4_address"`
	PublicIPv6Address  *string           `json:"public_ipv6_address"`
	GPNEnabled         bool              `json:"gpn_enabled"`
	IPv6Enabled        bool              `json:"ipv6_enabled"`
	Labels             map[string]string `json:"labels"`
	Created            time.Time         `json:"created_at"`
	Updated            time.Time         `json:"updated_at"`
}

// CloudComputingInstanceCreateInput represents cloud instance create input
type CloudComputingInstanceCreateInput struct {
	Name              string            `json:"name"`
	RegionID          int64             `json:"region_id"`
	FlavorID          string            `json:"flavor_id"`
	ImageID           string            `json:"image_id"`
	GPNEnabled        *bool             `json:"gpn_enabled,omitempty"`
	IPv6Enabled       *bool             `json:"ipv6_enabled,omitempty"`
	SSHKeyFingerprint *string           `json:"ssh_key_fingerprint,omitempty"`
	BackupCopies      *int              `json:"backup_copies,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
}

// CloudComputingInstanceUpdateInput represents cloud instance update input
type CloudComputingInstanceUpdateInput struct {
	Name         *string           `json:"name,omitempty"`
	BackupCopies *int              `json:"backup_copies,omitempty"`
	GPNEnabled   *bool             `json:"gpn_enabled,omitempty"`
	IPv6Enabled  *bool             `json:"ipv6_enabled,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
}

// CloudComputingInstanceReinstallInput represents cloud instance reinstall input
type CloudComputingInstanceReinstallInput struct {
	ImageID string `json:"image_id"`
}

// CloudComputingInstanceUpgradeInput represents cloud instance upgrade input
type CloudComputingInstanceUpgradeInput struct {
	FlavorID string `json:"flavor_id"`
}

// CloudComputingRegion represents cloud computing region
type CloudComputingRegion struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// CloudComputingImage represents cloud computing image
type CloudComputingImage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CloudComputingFlavor represents cloud computing flavor
type CloudComputingFlavor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// L2Segment represents l2 segment
type L2Segment struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	Status            string            `json:"status"`
	LocationGroupID   int64             `json:"location_group_id"`
	LocationGroupCode string            `json:"location_group_code"`
	Labels            map[string]string `json:"labels"`
	Created           time.Time         `json:"created_at"`
	Updated           time.Time         `json:"updated_at"`
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
	Labels          map[string]string      `json:"labels,omitempty"`
}

// L2SegmentUpdateInput represents l2 segment update input
type L2SegmentUpdateInput struct {
	Name    *string                `json:"name,omitempty"`
	Members []L2SegmentMemberInput `json:"members,omitempty"`
	Labels  map[string]string      `json:"labels,omitempty"`
}

// L2Member respresents l2 segment member
type L2Member struct {
	ID      string            `json:"id"`
	Title   string            `json:"title"`
	Mode    string            `json:"mode"`
	Vlan    *int              `json:"vlan"`
	Status  string            `json:"status"`
	Labels  map[string]string `json:"labels"`
	Created time.Time         `json:"created_at"`
	Updated time.Time         `json:"updated_at"`
}

// L2SegmentCreateNetworksInput represents input to create networks for L2SegmentChangeNetworksInput
type L2SegmentCreateNetworksInput struct {
	Mask               int    `json:"mask"`
	DistributionMethod string `json:"distribution_method"`
}

// L2SegmentChangeNetworksInput represents input to change networks
type L2SegmentChangeNetworksInput struct {
	Create []L2SegmentCreateNetworksInput `json:"create,omitempty"`
	Delete []string                       `json:"delete,omitempty"`
}

// Network represents network
type Network struct {
	ID                 string    `json:"id"`
	Title              *string   `json:"title,omitempty"`
	Status             string    `json:"status"`
	Cidr               *string   `json:"cidr,omitempty"`
	Family             string    `json:"family"`
	InterfaceType      string    `json:"interface_type"`
	DistributionMethod string    `json:"distribution_method"`
	Additional         bool      `json:"additional"`
	Created            time.Time `json:"created_at"`
	Updated            time.Time `json:"updated_at"`

	// DEPRECATED: should be replaced by Statu
	State string `json:"state"`
}

// L2LocationGroup represents l2 location groups
type L2LocationGroup struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	GroupType   string  `json:"group_type"`
	LocationIDs []int64 `json:"location_ids"`
}

// HostPowerFeed represents feed status
type HostPowerFeed struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// HostConnection represents host connection
type HostConnection struct {
	Port       string  `json:"port"`
	Type       string  `json:"type"`
	MACAddress *string `json:"macaddr"`
}

// PTRRecord represents ptr record
type PTRRecord struct {
	ID       string `json:"id"`
	IP       string `json:"ip"`
	Domain   string `json:"domain"`
	Priority int    `json:"priority"`
	TTL      int    `json:"ttl"`
}

// PTRRecordCreateInput represents ptr record create input
type PTRRecordCreateInput struct {
	IP       string `json:"ip"`
	Domain   string `json:"domain"`
	Priority *int   `json:"priority"`
	TTL      *int   `json:"ttl"`
}

// OperatingSystemReinstallPartitionInput represents partition for os reinstallation layout input
type OperatingSystemReinstallPartitionInput struct {
	Target string  `json:"target"`
	Size   int     `json:"size"`
	Fs     *string `json:"fs,omitempty"`
	Fill   bool    `json:"fill,omitempty"`
}

// OperatingSystemReinstallLayoutInput represents layout for os reinstallation drives input
type OperatingSystemReinstallLayoutInput struct {
	SlotPositions []int                                    `json:"slot_positions"`
	Raid          *int                                     `json:"raid,omitempty"`
	Ignore        *bool                                    `json:"ignore,omitempty"`
	Partitions    []OperatingSystemReinstallPartitionInput `json:"partitions,omitempty"`
}

// OperatingSystemReinstallDrivesInput represents drives for os reinstallation input
type OperatingSystemReinstallDrivesInput struct {
	Layout []OperatingSystemReinstallLayoutInput `json:"layout,omitempty"`
}

// OperatingSystemReinstallInput represents os reinstallation input
type OperatingSystemReinstallInput struct {
	Hostname           string                              `json:"hostname"`
	Drives             OperatingSystemReinstallDrivesInput `json:"drives"`
	OperatingSystemID  *int64                              `json:"operating_system_id,omitempty"`
	SSHKeyFingerprints []string                            `json:"ssh_key_fingerprints,omitempty"`
}

// OperatingSystemReinstallInput represents os reinstallation input
type SBMOperatingSystemReinstallInput struct {
	Hostname           string   `json:"hostname"`
	OperatingSystemID  int64    `json:"operating_system_id"`
	SSHKeyFingerprints []string `json:"ssh_key_fingerprints,omitempty"`
	UserData           *string  `json:"user_data,omitempty"`
}

// HostDriveSlot represents host drive slot
type HostDriveSlot struct {
	Position   int         `json:"position"`
	Interface  string      `json:"interface"`
	FormFactor string      `json:"form_factor"`
	DriveModel *DriveModel `json:"drive_model"`
}

// NetworkPool represents network pool
type NetworkPool struct {
	ID          string            `json:"id"`
	Title       *string           `json:"title"`
	CIDR        string            `json:"cidr"`
	Type        string            `json:"type"`
	LocationIDs []int             `json:"location_ids"`
	Labels      map[string]string `json:"labels"`
	Created     time.Time         `json:"created_at"`
	Updated     time.Time         `json:"updated_at"`
}

// NetworkPoolInput represents network pool input
type NetworkPoolInput struct {
	Title  *string           `json:"title"`
	Labels map[string]string `json:"labels,omitempty"`
}

// Subnetwork represents subnetwork
type Subnetwork struct {
	ID            string    `json:"id"`
	NetworkPoolID string    `json:"network_pool_id"`
	Title         *string   `json:"title"`
	CIDR          string    `json:"cidr"`
	Attached      bool      `json:"attached"`
	InterfaceType string    `json:"interface_type"`
	Created       time.Time `json:"created_at"`
	Updated       time.Time `json:"updated_at"`
}

// SubnetworkUpdateInput respresents subnetwork update input
type SubnetworkUpdateInput struct {
	Title *string `json:"title"`
}

// SubnetworkCreateInput represents subnetwork create input
type SubnetworkCreateInput struct {
	Title *string `json:"title"`
	CIDR  *string `json:"cidr,omitempty"`
	Mask  *int    `json:"mask,omitempty"`
}

// LoadBalancer represents load balancer
type LoadBalancer struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	Status            string            `json:"status"`
	ExternalAddresses []string          `json:"external_addresses"`
	LocationID        int64             `json:"location_id"`
	ClusterID         *string           `json:"cluster_id"`
	Labels            map[string]string `json:"labels"`
	Created           time.Time         `json:"created_at"`
	Updated           time.Time         `json:"updated_at"`
}

// L4LoadBalancer represents l4 load balancer
type L4LoadBalancer struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	Status            string            `json:"status"`
	ExternalAddresses []string          `json:"external_addresses"`
	LocationID        int64             `json:"location_id"`
	StoreLogs         bool              `json:"store_logs"`
	ClusterID         *string           `json:"cluster_id"`
	Labels            map[string]string `json:"labels"`
	Created           time.Time         `json:"created_at"`
	Updated           time.Time         `json:"updated_at"`
}

// L4VHostZoneInput represents l4 vhost zone input
type L4VHostZoneInput struct {
	ID            string  `json:"id"`
	UDP           bool    `json:"udp"`
	ProxyProtocol bool    `json:"proxy_protocol"`
	Ports         []int32 `json:"ports"`
	Description   *string `json:"description"`
	UpstreamID    string  `json:"upstream_id"`
}

// L4UpstreamInput represents l4 upstream input
type L4UpstreamInput struct {
	IP     string `json:"ip"`
	Port   int32  `json:"port"`
	Weight int32  `json:"weight"`
}

// L4UpstreamZoneInput represents l4 upstream zone input
type L4UpstreamZoneInput struct {
	ID         string            `json:"id"`
	Method     *string           `json:"method,omitempty"`
	UDP        bool              `json:"udp"`
	HCInterval *int              `json:"hc_interval,omitempty"`
	HCJitter   *int              `json:"hc_jitter,omitempty"`
	Upstreams  []L4UpstreamInput `json:"upstreams"`
}

// L4LoadBalancerUpdateInput represents l4 load balancer update input
type L4LoadBalancerUpdateInput struct {
	Name          *string               `json:"name,omitempty"`
	StoreLogs     *bool                 `json:"store_logs,omitempty"`
	ClusterID     *string               `json:"cluster_id,omitempty"`
	SharedCluster *bool                 `json:"shared_cluster,omitempty"`
	VHostZones    []L4VHostZoneInput    `json:"vhost_zones,omitempty"`
	UpstreamZones []L4UpstreamZoneInput `json:"upstream_zones,omitempty"`
	Labels        map[string]string     `json:"labels,omitempty"`
}

// L4LoadBalancerUpdateInput represents l4 load balancer create input
type L4LoadBalancerCreateInput struct {
	Name          string                `json:"name"`
	LocationID    int64                 `json:"location_id"`
	StoreLogs     *bool                 `json:"store_logs,omitempty"`
	ClusterID     *string               `json:"cluster_id,omitempty"`
	VHostZones    []L4VHostZoneInput    `json:"vhost_zones"`
	UpstreamZones []L4UpstreamZoneInput `json:"upstream_zones"`
	Labels        map[string]string     `json:"labels,omitempty"`
}

// L7LoadBalancer represents l7 load balancer
type L7LoadBalancer struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	Domains           []string          `json:"domains"`
	Status            string            `json:"status"`
	ExternalAddresses []string          `json:"external_addresses"`
	LocationID        int64             `json:"location_id"`
	Geoip             bool              `json:"geoip"`
	StoreLogs         bool              `json:"store_logs"`
	StoreLogsRegionID int64             `json:"store_logs_region_id"`
	ClusterID         *string           `json:"cluster_id"`
	Labels            map[string]string `json:"labels"`
	Created           time.Time         `json:"created_at"`
	Updated           time.Time         `json:"updated_at"`
}

// L7LocationZoneInput represents l7 location zone input
type L7LocationZoneInput struct {
	Location     string `json:"location"`
	UpstreamID   string `json:"upstream_id"`
	UpstreamPath string `json:"upstream_path"`
}

// L7VHostZoneInput represents l7 vhost zone input
type L7VHostZoneInput struct {
	ID                  string                `json:"id"`
	Ports               []int32               `json:"ports"`
	SSL                 bool                  `json:"ssl"`
	HTTP2               bool                  `json:"http2"`
	HTTPToHttpsRedirect bool                  `json:"http_to_https_redirect"`
	HTTP2PushPreload    bool                  `json:"http2_push_preload"`
	Domains             []string              `json:"domains"`
	SSLCertID           string                `json:"ssl_certificate_id"`
	LocationZones       []L7LocationZoneInput `json:"location_zones"`
}

// L7UpstreamInput represents l7 upstream input
type L7UpstreamInput struct {
	IP          string `json:"ip"`
	Port        int32  `json:"port"`
	Weight      int32  `json:"weight,omitempty"`
	MaxConns    int32  `json:"max_conns,omitempty"`
	MaxFails    int32  `json:"max_fails,omitempty"`
	FailTimeout int32  `json:"fail_timeout,omitempty"`
}

// L7UpstreamZoneInput represents l7 upstream zone input
type L7UpstreamZoneInput struct {
	ID            string            `json:"id"`
	Method        *string           `json:"method,omitempty"`
	SSL           bool              `json:"ssl"`
	Sticky        bool              `json:"sticky"`
	HCInterval    *int              `json:"hc_interval,omitempty"`
	HCJitter      *int              `json:"hc_jitter,omitempty"`
	HCFails       *int              `json:"hc_fails,omitempty"`
	HCPasses      *int              `json:"hc_passes,omitempty"`
	HCDomain      *string           `json:"hc_domain,omitempty"`
	HCPath        *string           `json:"hc_path,omitempty"`
	HCMethod      *string           `json:"hc_method,omitempty"`
	HCMandatory   bool              `json:"hc_mandatory,omitempty"`
	HCStatus      *string           `json:"hc_status,omitempty"`
	TLSPreset     *string           `json:"tls_preset,omitempty"`
	GRPC          bool              `json:"grpc,omitempty"`
	HCGRPCService *string           `json:"hc_grpc_service,omitempty"`
	HCGRPCStatus  *int              `json:"hc_grpc_status,omitempty"`
	Upstreams     []L7UpstreamInput `json:"upstreams"`
}

// L7LoadBalancerUpdateInput represents l7 load balancer update input
type L7LoadBalancerUpdateInput struct {
	Name                string                `json:"name,omitempty"`
	StoreLogs           *bool                 `json:"store_logs,omitempty"`
	StoreLogsRegionID   *int                  `json:"store_logs_region_id,,omitempty"`
	Geoip               *bool                 `json:"geoip,omitempty"`
	NewExternalIpsCount *int                  `json:"new_external_ips_count,omitempty"`
	DeleteExternalIps   []string              `json:"delete_external_ips,omitempty"`
	ClusterID           *string               `json:"cluster_id,omitempty"`
	SharedCluster       *bool                 `json:"shared_cluster,omitempty"`
	VHostZones          []L7VHostZoneInput    `json:"vhost_zones,omitempty"`
	UpstreamZones       []L7UpstreamZoneInput `json:"upstream_zones,omitempty"`
	Labels              map[string]string     `json:"labels,omitempty"`
}

// L7LoadBalancerUpdateInput represents l7 load balancer create input
type L7LoadBalancerCreateInput struct {
	Name              string                `json:"name"`
	LocationID        int64                 `json:"location_id"`
	StoreLogs         *bool                 `json:"store_logs,omitempty"`
	StoreLogsRegionID *int                  `json:"store_logs_region_id,,omitempty"`
	Geoip             *bool                 `json:"geoip,omitempty"`
	ClusterID         *string               `json:"cluster_id,omitempty"`
	VHostZones        []L7VHostZoneInput    `json:"vhost_zones"`
	UpstreamZones     []L7UpstreamZoneInput `json:"upstream_zones"`
	Labels            map[string]string     `json:"labels,omitempty"`
}

type SBMFlavor struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	CPUName                string `json:"cpu_name"`
	CPUCount               int    `json:"cpu_count"`
	CPUCoresCount          int    `json:"cpu_cores_count"`
	CPUFrequency           string `json:"cpu_frequency"`
	RAMSize                int    `json:"ram_size"`
	DrivesConfiguration    string `json:"drives_configuration"`
	PublicUplinkModelID    int    `json:"public_uplink_model_id"`
	PublicUplinkModelName  string `json:"public_uplink_model_name"`
	PrivateUplinkModelID   int    `json:"private_uplink_model_id"`
	PrivateUplinkModelName string `json:"private_uplink_model_name"`
	BandwidthID            int    `json:"bandwidth_id"`
	BandwidthName          string `json:"bandwidth_name"`
}

// LoadBalancerCluster represents load balancer cluster
type LoadBalancerCluster struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID int64     `json:"location_id"`
	Status     string    `json:"status"`
	Created    time.Time `json:"created_at"`
	Updated    time.Time `json:"updated_at"`
}

// DedicatedServerUpdateInput represents dedicated server update input
type DedicatedServerUpdateInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// KubernetesBaremetalNodeUpdateInput represents kubernetes baremetal node update input
type KubernetesBaremetalNodeUpdateInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// SBMServerUpdateInput represents sbm server update input
type SBMServerUpdateInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// Rack represents rack
type Rack struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	LocationID   int64             `json:"location_id"`
	LocationCode string            `json:"location_code"`
	Labels       map[string]string `json:"labels,omitempty"`
}

// RackUpdateInput represents rack update input
type RackUpdateInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// CloudBlockStorageVolume represents cloud block storage volume
type CloudBlockStorageVolume struct {
	ID            string            `json:"id"`
	OpenstackUUID *string           `json:"openstack_uuid"`
	RegionID      int64             `json:"region_id"`
	Size          int               `json:"size"`
	Status        string            `json:"status"`
	Bootable      bool              `json:"bootable"`
	Labels        map[string]string `json:"labels"`
	Created       *time.Time        `json:"created_at"`
	Description   *string           `json:"description"`
	Name          string            `json:"name"`
	Attachments   []Attachment      `json:"attachments"`
}

type Attachment struct {
	ID           string `json:"id"`
	InstanceID   string `json:"instance_id"`
	InstanceName string `json:"instance_name"`
	Device       string `json:"device"`
}

// CloudBlockStorageVolumeCreateInput represents cloud block storage volume create input
type CloudBlockStorageVolumeCreateInput struct {
	Name             string            `json:"name"`
	RegionID         int               `json:"region_id"`
	Size             int               `json:"size,omitempty"`
	Description      string            `json:"description,omitempty"`
	ImageID          string            `json:"image_id,omitempty"`
	SnapshotID       string            `json:"snapshot_id,omitempty"`
	AttachInstanceID string            `json:"attach_instance_id,omitempty"`
	BackupID         string            `json:"backup_id,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
}

// CloudBlockStorageVolumeUpdateInput represents cloud block storage volume update input
type CloudBlockStorageVolumeUpdateInput struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	ImageID     string            `json:"image_id,omitempty"`
	SnapshotID  string            `json:"snapshot_id,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// CloudBlockStorageVolumeAttachInput represents cloud block storage volume attach input
type CloudBlockStorageVolumeAttachInput struct {
	InstanceID string `json:"instance_id"`
}

// CloudBlockStorageVolumeDetachInput represents cloud block storage volume detach input
type CloudBlockStorageVolumeDetachInput struct {
	InstanceID string `json:"instance_id"`
}

// CloudBlockStorageBackup represents backup for cloud block storage volume
type CloudBlockStorageBackup struct {
	ID                  string            `json:"id"`
	OpenstackUUID       *string           `json:"openstack_uuid"`
	OpenstackVolumeUUID string            `json:"openstack_volume_uuid"`
	RegionID            int               `json:"region_id"`
	Size                int               `json:"size"`
	Status              string            `json:"status"`
	Labels              map[string]string `json:"labels"`
	Created             *time.Time        `json:"created_at"`
	Name                string            `json:"name"`
}

// CloudBlockStorageBackupCreateInput represents cloud block storage volume backup create input
type CloudBlockStorageBackupCreateInput struct {
	VolumeID    string            `json:"volume_id"`
	Name        string            `json:"name"`
	Incremental bool              `json:"incremental,omitempty"`
	Force       bool              `json:"force,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// CloudBlockStorageBackupUpdateInput represents cloud block storage volume backup update input
type CloudBlockStorageBackupUpdateInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// CloudBlockStorageBackupRestoreInput represents cloud block storage backup restore input
type CloudBlockStorageBackupRestoreInput struct {
	VolumeID string `json:"volume_id"`
}

// KubernetesCluster represents Kubernetes cluster
type KubernetesCluster struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Status     string            `json:"status"`
	LocationID int64             `json:"location_id"`
	Labels     map[string]string `json:"labels"`
	Created    time.Time         `json:"created_at"`
	Updated    time.Time         `json:"updated_at"`
}

// KubernetesClusterNode represents Kubernetes cluster node
type KubernetesClusterNode struct {
	ID                 string            `json:"id"`
	Number             int64             `json:"number"`
	Hostname           string            `json:"hostname"`
	Configuration      string            `json:"configuration"`
	Type               string            `json:"type"`
	Role               string            `json:"role"`
	Status             string            `json:"status"`
	PrivateIPv4Address string            `json:"private_ipv4_address"`
	PublicIPv4Address  string            `json:"public_ipv4_address"`
	RefID              string            `json:"ref_id"`
	ClusterID          string            `json:"cluster_id"`
	Labels             map[string]string `json:"labels"`
	Created            time.Time         `json:"created_at"`
	Updated            time.Time         `json:"updated_at"`
}

// KubernetesClusterUpdateInput represents Kubernetes cluster update input
type KubernetesClusterUpdateInput struct {
	Labels map[string]string `json:"labels,omitempty"`
}

// InvoiceList represents invoices list
type InvoiceList struct {
	ID       string  `json:"id"`
	Number   int64   `json:"number"`
	ParentID *string `json:"parent_id"`
	Status   string  `json:"status"`
	Date     string  `json:"date"`
	Type     string  `json:"type"`
	TotalDue float64 `json:"total_due"`
	Currency string  `json:"currency"`
}

// Invoice represents an invoice
type Invoice struct {
	ID       string  `json:"id"`
	Number   int64   `json:"number"`
	ParentID *string `json:"parent_id"`
	Status   string  `json:"status"`
	Date     string  `json:"date"`
	Type     string  `json:"type"`
	TotalDue float64 `json:"total_due"`
	Currency string  `json:"currency"`
	CsvUrl   string  `json:"csv_url"`
	PdfUrl   string  `json:"pdf_url"`
}
