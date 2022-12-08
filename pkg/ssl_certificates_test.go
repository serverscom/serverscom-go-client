package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	sslCertificateCustomPublicKey  = `-----BEGIN CERTIFICATE-----\nMIICuDCCAiGgAwIBAgIBADANBgkqhkiG9w0BAQUFADBEMQswCQYDVQQGEwJCRTEO\nMAwGA1UECgwFVGVzdE8xDzANBgNVBAsMBlRlc3RPVTEUMBIGA1UEAwwLc2VydmVy\ncy5jb20wHhcNMjAwNDIyMDYyMzMzWhcNMzAwNDIyMDYyMzMzWjBEMQswCQYDVQQG\nEwJCRTEOMAwGA1UECgwFVGVzdE8xDzANBgNVBAsMBlRlc3RPVTEUMBIGA1UEAwwL\nc2VydmVycy5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAMmQwhWIB3Hr\nOK0jCbuVEC8jirudk53gxyskBhxwUNSCtvGciJtyhKgqmvMi8xrWOuI1PmCaq/wl\n5IfYvGoBW9p96KqffaL46lbLe9CiOfRiWY7KzuI9LsSKhJgMmr+Hc5uyToH/pJeE\n/vvfRtPRkoxSAkmNeUQqIIq9Hj0s0dYvAgMBAAGjgbkwgbYwDwYDVR0TAQH/BAUw\nAwEB/zAdBgNVHQ4EFgQUDYb5EZBj5areFGPOPnyan7AqkPkwFgYDVR0RBA8wDYIL\nc2VydmVycy5jb20wbAYDVR0jBGUwY4AUDYb5EZBj5areFGPOPnyan7AqkPmhSKRG\nMEQxCzAJBgNVBAYTAkJFMQ4wDAYDVQQKDAVUZXN0TzEPMA0GA1UECwwGVGVzdE9V\nMRQwEgYDVQQDDAtzZXJ2ZXJzLmNvbYIBADANBgkqhkiG9w0BAQUFAAOBgQAWPdlV\nWIX1gRJUUAFm0Pt6Vr+och/SlfRySQniIsc1EjofGUc42Ljg5hyAOVA8cK94bC+e\nd/9oIxHHtA4qAcZeiuhp22Lgws+GVrTM1WU+PKtvSlpYy0YSs24hNWrYg/TOf9gj\njtx4XivUYVTTbosLqZPGz9BY/lLG1SZw+EjC1Q==\n-----END CERTIFICATE-----\n`
	sslCertificateCustomPrivateKey = `-----BEGIN RSA PRIVATE KEY-----\nMIICXQIBAAKBgQDJkMIViAdx6zitIwm7lRAvI4q7nZOd4McrJAYccFDUgrbxnIib\ncoSoKprzIvMa1jriNT5gmqv8JeSH2LxqAVvafeiqn32i+OpWy3vQojn0YlmOys7i\nPS7EioSYDJq/h3Obsk6B/6SXhP7730bT0ZKMUgJJjXlEKiCKvR49LNHWLwIDAQAB\nAoGAZgLtNxhxLCZvuLBS7Ky0VCcYv3swimaIZj0FGr99KRA+pmkiegmrObDWWtcF\nUj+57WIk/59IC4Th9B6svLmTUkplLYb7a+98v42CEiyEdnqrj5o7YLmAX4KSXWic\nQMOH0L6bY1nr39gmB4lUUo2ITaX+d/N4LjMKTk1KM9FVhQECQQD1l7U01mhObDN1\nyFgIkcGSl70ItFBte08TCnGoUvVI3NyHK+l86AcPvkHXOIqIkFhlFoh2TvJ8exrS\nEhjwfWFnAkEA0hts3Q/OCCgSNrEgj5pYHOyXjGxcTSkzf7yWzwbiU8NrBUC2enfz\nWC/PhixzQiH6RIHQQ95bt9SZ2NQRQJR/+QJBAOJhNx0/TeKMBltZkxxDDsWLrSnq\n3AAvG9KXW/EmlwbU30qSBaWnU2sFmzdB3SDIvVJhFOTJHOf1qeFd4TM8v50CQEQB\n5onjXMAdIFAdozl4Lv8lwaQNSw8av/WfuHzIcKfGQDTSDn6zpsurphN6/c+xKc3U\n6vGc3rkxEp8xfWgW6RkCQQDKeBVsmMNtTDYJp7wMlq54IPV/Lduqc9m4dE7Y0JI+\nxWs4TAVA8zTiDsCCxN/WsX5yECDoWipQvnOF6QXP+SIQ\n-----END RSA PRIVATE KEY-----\n`
)

func TestSSLCertificatesCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssl_certificates").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.SSLCertificates.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestSSLCertificatesCreateCustom(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssl_certificates/custom").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/ssl_certificates/create_custom_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := SSLCertificateCreateCustomInput{
		Name:       "name157",
		PublicKey:  sslCertificateCustomPublicKey,
		PrivateKey: sslCertificateCustomPrivateKey,
	}

	ctx := context.TODO()

	SSLCertificateCustom, err := client.SSLCertificates.CreateCustom(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(SSLCertificateCustom).ToNot(BeNil())

	g.Expect(SSLCertificateCustom.Name).To(Equal("name157"))
}

func TestSSLCertificatesGetCustom(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssl_certificates/custom/LDdwpRe1").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/ssl_certificates/get_custom_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	SSLCertificateCustom, err := client.SSLCertificates.GetCustom(ctx, "LDdwpRe1")

	g.Expect(err).To(BeNil())
	g.Expect(SSLCertificateCustom).ToNot(BeNil())

	g.Expect(SSLCertificateCustom.Name).To(Equal("name156"))
}
