package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	sshPublicKey = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDxRD2uaK8Dn4/AiUP8EkRj9M1LLpPGjDg8VMPwDlRH4RyoMA29kRG+IVg6LdiWqhxG0FjeHhQqS+Xnz1eFH5lbaf6+UWL2EkrFMp43SRZuowheIwn8xeZDbhSxxUYDTRABNWPSx4F5+MU3WuerAI44Gy3nz0xjkJCIo3cqsHeVGyqtsHmO05THeQwQq9TaOTTwnB92RiNgruHS7DbAPfAqDxZznLDncIwSSt7QPDzeQc42bA4Leysy0Y6ymgGfwJMhqiddQvRHtQrAQ6MH4Db/f6bkFU3/FCTp9LtTZaD84c7DaiezdzQTh9BF59vw/76HCS8+UVFLIYYkD8U6SsnP Generated by Nova`
)

func TestSSHKeyCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/ssh_keys/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := SSHKeyCreateInput{
		Name:      "test-key",
		PublicKey: sshPublicKey,
	}

	ctx := context.TODO()

	SSHKey, err := client.SSHKeys.Create(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(SSHKey).ToNot(BeNil())

	g.Expect(SSHKey.Name).To(Equal("test-key"))
	g.Expect(SSHKey.Fingerprint).To(Equal("48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8"))
	g.Expect(SSHKey.Created.String()).To(Equal("2020-04-22 06:23:09 +0000 UTC"))
	g.Expect(SSHKey.Updated.String()).To(Equal("2020-04-22 06:23:09 +0000 UTC"))
}

func TestSSHKeyGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys/48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/ssh_keys/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	SSHKey, err := client.SSHKeys.Get(ctx, "48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8")

	g.Expect(err).To(BeNil())
	g.Expect(SSHKey).ToNot(BeNil())

	g.Expect(SSHKey.Name).To(Equal("test-key"))
	g.Expect(SSHKey.Fingerprint).To(Equal("48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8"))
	g.Expect(SSHKey.Created.String()).To(Equal("2020-04-22 06:23:09 +0000 UTC"))
	g.Expect(SSHKey.Updated.String()).To(Equal("2020-04-22 06:23:09 +0000 UTC"))
}

func TestSSHKeyUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys/48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8").
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/ssh_keys/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	input := SSHKeyUpdateInput{Name: "new-name"}

	ctx := context.TODO()

	SSHKey, err := client.SSHKeys.Update(ctx, "48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8", input)

	g.Expect(err).To(BeNil())
	g.Expect(SSHKey).ToNot(BeNil())

	g.Expect(SSHKey.Name).To(Equal("new-name"))
	g.Expect(SSHKey.Fingerprint).To(Equal("48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8"))
	g.Expect(SSHKey.Created.String()).To(Equal("2020-04-22 06:23:09 +0000 UTC"))
	g.Expect(SSHKey.Updated.String()).To(Equal("2020-04-22 06:23:09 +0000 UTC"))
}

func TestSSHKeyDelete(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys/48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.SSHKeys.Delete(ctx, "48:81:0c:43:99:12:71:5e:ba:fd:e7:2f:20:d7:95:e8")

	g.Expect(err).To(BeNil())
}
