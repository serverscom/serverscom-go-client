package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	dnsDomainID = "aXmDk39e"
	dnsRecordID = "rEc0rd01"
)

func TestDNSCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.DNS.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDNSCreateDomain(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/dns/create_domain_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dnsDomain, err := client.DNS.CreateDomain(ctx, DNSDomainCreateInput{
		Name:   "example.com",
		Email:  "admin@example.com",
		TTL:    3600,
		Labels: map[string]string{"env": "test"},
	})

	g.Expect(err).To(BeNil())
	g.Expect(dnsDomain).ToNot(BeNil())

	g.Expect(dnsDomain.ID).To(Equal(dnsDomainID))
	g.Expect(dnsDomain.Name).To(Equal("example.com"))
	g.Expect(dnsDomain.Email).To(Equal("admin@example.com"))
	g.Expect(dnsDomain.DelegationStatus).To(Equal(DNSDomainUndelegated))
	g.Expect(dnsDomain.Labels).To(Equal(map[string]string{"env": "test"}))
}

func TestDNSGetDomain(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/dns/get_domain_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dnsDomain, err := client.DNS.GetDomain(ctx, dnsDomainID)

	g.Expect(err).To(BeNil())
	g.Expect(dnsDomain).ToNot(BeNil())

	g.Expect(dnsDomain.ID).To(Equal(dnsDomainID))
	g.Expect(dnsDomain.Name).To(Equal("example.com"))
	g.Expect(dnsDomain.DelegationStatus).To(Equal(DNSDomainDelegated))
	g.Expect(dnsDomain.TTL).To(Equal(3600))
}

func TestDNSUpdateDomain(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/dns/update_domain_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()
	newLabels := map[string]string{"env": "new-test"}

	dnsDomain, err := client.DNS.UpdateDomain(ctx, dnsDomainID, DNSDomainUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(dnsDomain).ToNot(BeNil())

	g.Expect(dnsDomain.ID).To(Equal(dnsDomainID))
	g.Expect(dnsDomain.Labels).To(Equal(newLabels))
	g.Expect(dnsDomain.DelegationStatus).To(Equal(DNSDomainVerified))
}

func TestDNSDeleteDomain(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID).
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.DNS.DeleteDomain(ctx, dnsDomainID)

	g.Expect(err).To(BeNil())
}

func TestDNSGetDelegationData(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/delegation_data").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/dns/delegation_data_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	delegationData, err := client.DNS.GetDelegationData(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(delegationData).ToNot(BeNil())

	g.Expect(delegationData.Nameservers).To(Equal([]string{"ns1.servers.com", "ns2.servers.com"}))
	g.Expect(delegationData.RequiredTxt).To(Equal("servers-com-verification=abc123"))
}

func TestDNSRecordsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID + "/records").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/dns/list_records_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.DNS.Records(dnsDomainID)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(HaveLen(1))
	g.Expect(list[0].ID).To(Equal(dnsRecordID))
	g.Expect(list[0].DomainID).To(Equal(dnsDomainID))
	g.Expect(list[0].Type).To(Equal(DNSRecordTypeA))
}

func TestDNSCreateRecord(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID + "/records").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/dns/create_record_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dnsRecord, err := client.DNS.CreateRecord(ctx, dnsDomainID, DNSRecordCreateInput{
		Name: "www",
		Type: DNSRecordTypeA,
		Data: "192.0.2.1",
		TTL:  3600,
	})

	g.Expect(err).To(BeNil())
	g.Expect(dnsRecord).ToNot(BeNil())

	g.Expect(dnsRecord.ID).To(Equal(dnsRecordID))
	g.Expect(dnsRecord.Name).To(Equal("www"))
	g.Expect(dnsRecord.Type).To(Equal(DNSRecordTypeA))
	g.Expect(dnsRecord.Data).ToNot(BeNil())
	g.Expect(*dnsRecord.Data).To(Equal("192.0.2.1"))
}

func TestDNSGetRecord(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID + "/records/" + dnsRecordID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/dns/get_record_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dnsRecord, err := client.DNS.GetRecord(ctx, dnsDomainID, dnsRecordID)

	g.Expect(err).To(BeNil())
	g.Expect(dnsRecord).ToNot(BeNil())

	g.Expect(dnsRecord.ID).To(Equal(dnsRecordID))
	g.Expect(dnsRecord.DomainID).To(Equal(dnsDomainID))
	g.Expect(dnsRecord.Type).To(Equal(DNSRecordTypeA))
}

func TestDNSUpdateRecord(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID + "/records/" + dnsRecordID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/dns/update_record_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	dnsRecord, err := client.DNS.UpdateRecord(ctx, dnsDomainID, dnsRecordID, DNSRecordUpdateInput{
		Data: "192.0.2.2",
		Name: "www",
		TTL:  7200,
	})

	g.Expect(err).To(BeNil())
	g.Expect(dnsRecord).ToNot(BeNil())

	g.Expect(dnsRecord.ID).To(Equal(dnsRecordID))
	g.Expect(dnsRecord.Data).ToNot(BeNil())
	g.Expect(*dnsRecord.Data).To(Equal("192.0.2.2"))
	g.Expect(dnsRecord.TTL).ToNot(BeNil())
	g.Expect(*dnsRecord.TTL).To(Equal(7200))
}

func TestDNSDeleteRecord(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/dns/domains/" + dnsDomainID + "/records/" + dnsRecordID).
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.DNS.DeleteRecord(ctx, dnsDomainID, dnsRecordID)

	g.Expect(err).To(BeNil())
}
