package v5

import "testing"

func TestGenerateUUIDv5(t *testing.T) {
	uuidv5DNS, err := GenerateUUIDv5(DNSNamespace, []byte("example.com"))
	if err != nil {
		t.Fatalf("error generated uuidv5: %s", err.Error())
	}

	t.Logf("generated uuidv5 DNS: %s", uuidv5DNS.String())

	uuidv5URL, err := GenerateUUIDv5(URLNamespace, []byte("example.com"))
	if err != nil {
		t.Fatalf("error generated uuidv5: %s", err.Error())
	}

	t.Logf("generated uuidv5 URL: %s", uuidv5URL.String())



	uuidv5OID, err := GenerateUUIDv5(OIDNamespace, []byte("example.com"))

	if err != nil {
		t.Fatalf("error generated uuidv5: %s", err.Error())
	}

	t.Logf("generated uuidv5 OID: %s", uuidv5OID.String())

}
