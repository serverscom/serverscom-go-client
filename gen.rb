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
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllHosts',
).render_to_file('pkg/hosts_collection.go')

CollectionGenerator.new(
  name: 'Location',
  path: '/locations',
  entity: 'Location',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/Locations',
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
  entity: 'DriveModelOption',
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
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListAllSshKeys',
).render_to_file('pkg/ssh_keys_collection.go')

CollectionGenerator.new(
  name: 'CloudInstance',
  path: '/cloud_computing/instances',
  entity: 'CloudInstance',
  api_url: 'https://developers.servers.com/api-documentation/v1/#operation/ListCloudInstances',
).render_to_file('pkg/cloud_instances_collection.go')
