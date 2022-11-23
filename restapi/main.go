package restapi

import (
	"errors"
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/structureDAO"
	"github.com/go-playground/validator"
	config "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	factory FactoryInterface
	once    sync.Once
)

type FactoryInterface interface {
	//Role
	GetRoleListDB() (response []structureDAO.Role, Error error)
	GetRoleDB(req structureDAO.Role) (response structureDAO.Role, Error error)
	AddRoleDB(req structureDAO.Role) (Error error)

	// OTP
	SendOTPDB(req structureDAO.OTP) (Error error)
	GetOTPDB(req structureDAO.OTP) (response *structureDAO.OTP, Error error)

	// CreateUser
	CreateUserDB(req structureDAO.Account) (Error error)
}

func Create(env *Properties) FactoryInterface {
	once.Do(func() {
		factory = gormInstance(env)
	})
	return factory
}

type GORMFactory struct {
	env    *Properties
	client *gorm.DB
}

func gormInstance(env *Properties) GORMFactory {
	databaseSet := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		env.GormHost, env.GormPort, env.GormUser, env.GormName, env.GormPass, "disable")

	db, err := gorm.Open(postgres.Open(databaseSet), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
		//panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
	}

	//if env.Flavor != environment.Production {
	//	db = db.Debug()
	//}
	_ = db.AutoMigrate(
		&structureDAO.Role{},
		&structureDAO.Account{},
		&structureDAO.Address{},
		&structureDAO.OTP{},
		&structureDAO.LogLogin{},
	)
	return GORMFactory{env: env, client: db}
}

type Controller struct {
	Properties *Properties
	Access     *Access
}

type Access struct {
	ENV   *Properties
	RDBMS FactoryInterface
	//GRPC grpc.FactoryInterface
}

func Initial(properties *Properties) *Access {
	return &Access{
		ENV:   properties,
		RDBMS: Create(properties),
		//GRPC: grpc.Create(properties),
	}
}

func NewController() Controller {
	build := Build()
	access := Initial(build)

	return Controller{
		Properties: build,
		Access:     access,
	}
}

type Flavor string
type URL string

const (
	Develop    Flavor = "DEVELOP"
	Devspace   Flavor = "DEVSPACE"
	Production Flavor = "PRODUCTION"
)

type Properties struct {
	// -- core
	Flavor Flavor `env:"FLAVOR,default=DEVELOP"`
	// --

	// -- Gorm
	//GormHost string `env:"GORM_HOST,default=access"`
	GormHost string `env:"GORM_HOST,default=localhost"`
	//GormHost string `env:"GORM_HOST,default=access"`
	GormPort string `env:"GORM_PORT,default=5432"`
	GormName string `env:"GORM_NAME,default=postgres_db"`
	GormUser string `env:"GORM_USER,default=postgres"`
	GormPass string `env:"GORM_PASS,default=pgpassword"`
}

func Build() *Properties {
	var prop Properties
	if _, err := env.UnmarshalFromEnviron(&prop); err != nil {
		panic(err)
	}
	return &prop
}

func (ctrl Controller) LoadConfigFile() error {
	env := os.Getenv("ENV")
	//env = "dev"
	if env == "" {
		env = os.Args[1]
	}

	//ctrl.Logger.Info(transID, fmt.Sprintf("Server start running on %s environment configuration", env))
	config.SetConfigName(env)
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")
	err := config.ReadInConfig()
	if err != nil {
		errMsg := fmt.Sprintf("Read config file %s.yml occur error: %s", env, err.Error())
		panic(errMsg)
		return err
	}
	return err
}

func ValidateStruct(dataStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("%s: %s", err.StructField(), err.Tag()))
		}
	} else {
		return nil
	}
	return err
}
