package config

type JWTConfig struct {
	SecretKey []byte
}

func DefaultJWTConfig() *JWTConfig {
	config := &JWTConfig{
		SecretKey: []byte("my-secret-key"),
	}
	return config
}
