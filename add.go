package token

func AddCipherAdapter(encryption_key ...string) *config {

	c := config{
		encryption_key: global_key,
	}

	for _, v := range encryption_key {
		c.encryption_key = v
	}

	return &c
}
