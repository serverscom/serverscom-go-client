# Servers.com go client

[![GoDoc](https://godoc.org/github.com/serverscom/serverscom-go-client/pkg?status.svg)](https://godoc.org/github.com/serverscom/serverscom-go-client/pkg)

serverscom-go-client is a Go client library for accessing the [Servers.com API](https://developers.servers.com/api-documentation/v1/)

## Development status

Unstable, in development

## Example

```go
client := serverscom.NewClient("my-jwt-token")

hosts, err := client.Hosts.Collection().Collect()
if err != nil {
  panic(err.Error())
}

for h, i := range hosts {
  log.Println(fmt.Sprintf("Host: %s, with title: %s, private ipv4: %s, public ipv4: %s", h.ID, h.Title, h.PrivateIPv4Address, h.PublicIPv4Address))
}
```

## Copyright

The library is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
