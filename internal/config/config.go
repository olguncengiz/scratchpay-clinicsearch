package config

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-env"
	"github.com/olguncengiz/scratchpay-clinicsearch/pkg/log"
	"gopkg.in/yaml.v2"
)

const (
	defaultServerPort            = 8080
	defaultJWTExpirationHours    = 72
	defaultDentalClinicsFilePath = "dental.json"
	defaultVetClinicsFilePath    = "vet.json"
)

// Config represents an application configuration.
type Config struct {
	// the server port. Defaults to 8080
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// JWT signing key. required.
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
	// the dental clinics file path
	DentalClinicsFilePath string `yaml:"dental_clinics_file_path"`
	// the vet clinics file path
	VetClinicsFilePath string `yaml:"vet_clinics_file_path"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.JWTSigningKey, validation.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config{
		ServerPort:            defaultServerPort,
		JWTExpiration:         defaultJWTExpirationHours,
		DentalClinicsFilePath: defaultDentalClinicsFilePath,
		VetClinicsFilePath:    defaultVetClinicsFilePath,
	}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.New("APP_", logger.Infof).Load(&c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
