package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetAccountBalance(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/account/balance").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/account/get_balance_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	balance, err := client.Account.GetBalance(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(balance).ToNot(BeNil())
	g.Expect(balance.Currency).To(Equal("EUR"))
	g.Expect(balance.CurrentBalance).To(Equal(float64(123.456)))
	g.Expect(balance.NextInvoiceTotalDue).To(Equal(float64(0.0)))
}
