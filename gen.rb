require_relative 'generator/collection'

CollectionGenerator.new(
  name: 'SSLCertificate',
  path: '/ssl_certificates',
  entity: 'SSLCertificate',
  var_prefix: 'sslCertificate',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/SSLCertificates',
  params: {search_pattern: 'string', type: 'string'}
).render_to_file('pkg/ssl_certificates_collection.go')

CollectionGenerator.new(
  name: 'Host',
  path: '/hosts',
  entity: 'Host',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllHosts'
).render_to_file('pkg/hosts_collection.go')

CollectionGenerator.new(
  name: 'Location',
  path: '/locations',
  entity: 'Location',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/Locations'
).render_to_file('pkg/locations_collection.go')

CollectionGenerator.new(
  name: 'ServerModelOption',
  path: '/locations/%d/order_options/server_models',
  entity: 'ServerModelOption',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllServerModelsForLocation',
  args: {LocationID: 'int64'}
).render_to_file('pkg/server_model_options_collection.go')

CollectionGenerator.new(
  name: 'RAMOption',
  path: '/locations/%d/order_options/server_models/%d/ram',
  entity: 'RAMOption',
  args: {LocationID: 'int64', ServerModelID: 'int64'}
).render_to_file('pkg/ram_options_collection.go')

CollectionGenerator.new(
  name: 'OperatingSystemOption',
  path: '/locations/%d/order_options/server_models/%d/operating_systems',
  entity: 'OperatingSystemOption',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllOperatingSystemsForServerModel',
  args: {LocationID: 'int64', ServerModelID: 'int64'}
).render_to_file('pkg/operating_system_options_collection.go')

CollectionGenerator.new(
  name: 'UplinkOption',
  path: '/locations/%d/order_options/server_models/%d/uplink_models',
  entity: 'UplinkOption',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllUplinksForServerModel',
  args: {LocationID: 'int64', ServerModelID: 'int64'}
).render_to_file('pkg/uplink_options_collection.go')

CollectionGenerator.new(
  name: 'DriveModelOption',
  path: '/locations/%d/order_options/server_models/%d/drive_models',
  entity: 'DriveModel',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllDriveModelsOptionsForServerModel',
  args: {LocationID: 'int64', ServerModelID: 'int64'}
).render_to_file('pkg/drive_model_options_collection.go')

CollectionGenerator.new(
  name: 'BandwidthOption',
  path: '/locations/%d/order_options/server_models/%d/uplink_models/%d/bandwidth',
  entity: 'BandwidthOption',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllBandwidthForUplink',
  args: {LocationID: 'int64', ServerModelID: 'int64', uplinkModelID: 'int64'}
).render_to_file('pkg/bandwidth_options_collection.go')

CollectionGenerator.new(
  name: 'SSHKey',
  path: '/ssh_keys',
  entity: 'SSHKey',
  var_prefix: 'sshKey',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllSshKeys'
).render_to_file('pkg/ssh_keys_collection.go')

CollectionGenerator.new(
  name: 'CloudComputingInstance',
  path: '/cloud_computing/instances',
  entity: 'CloudComputingInstance',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListCloudComputingInstances'
).render_to_file('pkg/cloud_computing_instances_collection.go')

CollectionGenerator.new(
  name: 'L2Segment',
  path: '/l2_segments',
  entity: 'L2Segment',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllL2Segments'
).render_to_file('pkg/l2_segments_collection.go')

CollectionGenerator.new(
  name: 'L2Member',
  path: '/l2_segments/%s/members',
  entity: 'L2Member',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllL2SegmentMembers',
  args: {segmentID: 'string'}
).render_to_file('pkg/l2_members_collection.go')

CollectionGenerator.new(
  name: 'L2Network',
  path: '/l2_segments/%s/networks',
  entity: 'Network',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllL2SegmentNetworks',
  var_prefix: 'l2Networks',
  args: {segmentID: 'string'}
).render_to_file('pkg/l2_networks_collection.go')

CollectionGenerator.new(
  name: 'L2LocationGroup',
  path: '/l2_segments/location_groups',
  entity: 'L2LocationGroup',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllLocationGroups'
).render_to_file('pkg/l2_location_groups_collection.go')

CollectionGenerator.new(
  name: 'HostConnection',
  path: '/hosts/%s/%s/connections',
  entity: 'HostConnection',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllConnectionsForAnExistingDedicatedServer',
  args: {hostType: 'string', hostID: 'string'}
).render_to_file('pkg/host_connections_collection.go')

CollectionGenerator.new(
  name: 'HostNetwork',
  path: '/hosts/%s/%s/networks',
  entity: 'Network',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllNetworksForAnExistingDedicatedServer',
  var_prefix: 'hostNetworks',
  args: {hostType: 'string', hostID: 'string'}
).render_to_file('pkg/host_networks_collection.go')

CollectionGenerator.new(
  name: 'HostPTRRecord',
  path: '/hosts/%s/%s/ptr_records',
  entity: 'PTRRecord',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllPtrRecordsForServerNetworks',
  var_prefix: 'hostPTRs',
  args: {hostType: 'string', hostID: 'string'}
).render_to_file('pkg/host_ptr_records_collection.go')

CollectionGenerator.new(
  name: 'HostDriveSlot',
  path: '/hosts/%s/%s/drive_slots',
  entity: 'HostDriveSlot',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllDriveSlotsForAnExistingDedicatedServer',
  var_prefix: 'hostDriveSlot',
  args: {hostType: 'string', hostID: 'string'}
).render_to_file('pkg/host_drive_slots_collection.go')

CollectionGenerator.new(
  name: 'CloudComputingInstancePTRRecord',
  path: '/cloud_computing/instances/%s/ptr_records',
  entity: 'PTRRecord',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ReturnsInstancePtrRecords',
  var_prefix: 'cloudInstancePTRs',
  args: {cloudInstanceID: 'string'}
).render_to_file('pkg/cloud_computing_instance_ptr_records_collection.go')

CollectionGenerator.new(
  name: 'CloudComputingRegion',
  path: '/cloud_computing/regions',
  entity: 'CloudComputingRegion',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListCloudRegions'
).render_to_file('pkg/cloud_computing_regions_collection.go')

CollectionGenerator.new(
  name: 'CloudComputingImage',
  path: '/cloud_computing/regions/%d/images',
  entity: 'CloudComputingImage',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListCloudImages',
  args: {regionID: 'int64'}
).render_to_file('pkg/cloud_computing_images_collection.go')

CollectionGenerator.new(
  name: 'CloudComputingFlavor',
  path: '/cloud_computing/regions/%d/flavors',
  entity: 'CloudComputingFlavor',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListCloudFlavors',
  args: {regionID: 'int64'}
).render_to_file('pkg/cloud_computing_flavors_collection.go')

CollectionGenerator.new(
  name: 'NetworkPool',
  path: '/network_pools',
  entity: 'NetworkPool',
  params: {search_pattern: 'string'},
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListNetworkPools',
).render_to_file('pkg/network_pools_collection.go')

CollectionGenerator.new(
  name: 'Subnetwork',
  path: '/network_pools/%s/subnetworks',
  entity: 'Subnetwork',
  args: {networkPoolID: 'string'},
  params: {search_pattern: 'string'},
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllSubnetworks',
).render_to_file('pkg/subnetworks_collection.go')
