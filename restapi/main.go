package restapi

import (
	"errors"
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/structureDAO"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/verify"
	"github.com/go-playground/validator"
	config "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
	"time"
)

var (
	factory FactoryInterface
	once    sync.Once
)

type FactoryInterface interface {
	//Role
	GetRoleListDB() (response []structureDAO.Role, Error error)
	GetRoleDBByName(req structureDAO.Role) (response structureDAO.Role, Error error)
	GetRoleDBById(req structureDAO.Role) (response structureDAO.Role, Error error)
	AddRoleDB(req structureDAO.Role) (Error error)

	// OTP
	SendOTPDB(req structureDAO.OTP) (Error error)
	GetOTPDB(req structureDAO.OTP) (response *structureDAO.OTP, Error error)
	UpdateOTPDB(req structureDAO.OTP) (Error error)

	// CreateUser
	CreateUserDB(req structureDAO.Account) (Error error)

	// Account
	GetAccountDB(req structureDAO.Account) (response *structureDAO.Account, Error error)
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

	var CheckRole []structureDAO.Role
	db.Find(&CheckRole)
	if len(CheckRole) == 0 {
		dataAdmin := structureDAO.Role{
			Name: "admin",
		}
		dataUser := structureDAO.Role{
			Name: "user",
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dataAdmin)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dataUser)

		role := structureDAO.Role{}
		address := structureDAO.Address{
			Address:     "",
			SubDistrict: "",
			District:    "",
			Province:    "",
			PostalCode:  "",
			Country:     "",
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&address)
		db.Where("name = ?", "admin").Take(&role)
		Password, _ := verify.Hash("BELLkub4424506")
		data := structureDAO.Account{
			Email:       nil,
			PhoneNumber: "0815476439",
			Password:    string(Password),
			FirstName:   "admin",
			LastName:    "admin",
			Birthday:    time.Now(),
			Gender:      "M",
			IDCard:      "1349900833347",
			Photo:       nil,
			RoleID:      role.ID,
			AddressID:   &address.ID,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data)
	}
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
	GormName string `env:"GORM_NAME,default=postgresdb"`
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
	env = "dev"
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
