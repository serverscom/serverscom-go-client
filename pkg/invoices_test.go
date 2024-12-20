package serverscom

import (
	"context"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInvoicesCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/billing/invoices").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Invoices.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestGetInvoice(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/billing/invoices/9qzw9Hj4").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/invoices/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	invoice, err := client.Invoices.GetBillingInvoice(ctx, "9qzw9Hj4")

	g.Expect(err).To(BeNil())
	g.Expect(invoice).ToNot(BeNil())
	g.Expect(invoice.ID).To(Equal("9qzw9Hj4"))
	g.Expect(invoice.Number).To(Equal(int64(843071)))
	g.Expect(invoice.ParentID).To(BeNil())
	g.Expect(invoice.Status).To(Equal("paid"))
	g.Expect(invoice.Date).To(Equal("2024-06-23"))
	g.Expect(invoice.Type).To(Equal("invoice"))
	g.Expect(invoice.TotalDue).To(Equal(11.90))
	g.Expect(invoice.Currency).To(Equal("EUR"))
	g.Expect(invoice.CsvUrl).To(Equal("url"))
	g.Expect(invoice.PdfUrl).To(Equal("url"))
}
