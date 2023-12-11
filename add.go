package token

func AddCipherAdapter(encryption_key ...string) *config {

	c := config{
		encryption_key: public_key_client_and_server,
	}

	for _, v := range encryption_key {
		c.encryption_key = v
	}

	return &c
}
