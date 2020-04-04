package environment

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/joho/godotenv"

	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"

	"google.golang.org/appengine"
)

var errGetSharedEnvironments = xerrors.New("infrastructure-environment: Failed to GetSharedEnvironments")

const (
	devEnv  = "development"
	prodEnv = "production"
)

type Environments struct {
	IPAddress   string `default:"localhost" split_words:"true" envconfig:"IP_ADDRESS"`
	Port        string `default:"8080" split_words:"true"`
	GinMode     string `default:"debug" split_words:"true"`
	GormLogMode bool   `default:"false" split_words:"true"`

	MysqlUser     string `required:"true" split_words:"true"`
	MysqlPassword string `required:"true" split_words:"true"`
	MysqlProtocol string `required:"true" split_words:"true"`
	MysqlHost     string `required:"true" split_words:"true"`
	MysqlDbName   string `required:"true" split_words:"true"`
}

var environments *Environments
var once sync.Once

// GetSharedEnvironments は環境変数を保持するシングルトンを返す
func GetSharedEnvironments() *Environments {
	once.Do(func() {
		switch os.Getenv("ENV") {
		case devEnv, prodEnv:
			_, filename, _, ok := runtime.Caller(0)
			if !ok {
				panic(xerrors.Errorf(errGetSharedEnvironments.Error() + "cause No caller information"))
			}
			if IsAppEngine() {
				envPath := fmt.Sprintf("./config/%s.env", os.Getenv("ENV"))
				if err := godotenv.Load(envPath); err != nil {
					panic(xerrors.Errorf(errGetSharedEnvironments.Error()+": %s", err))
				}
			} else {
				envPath := fmt.Sprintf("%s/../config/%s.env", path.Dir(filename), os.Getenv("ENV"))
				if err := godotenv.Load(envPath); err != nil {
					panic(xerrors.Errorf(errGetSharedEnvironments.Error()+": %s", err))
				}
			}
		default:
			err := xerrors.Errorf(errGetSharedEnvironments.Error()+" cause Unexpected value '%s' by ENV", os.Getenv("ENV"))
			panic(err)
		}

		environments = &Environments{}
		if err := envconfig.Process("", environments); err != nil {
			panic(xerrors.Errorf(errGetSharedEnvironments.Error()+": %s", err))
		}

		context.Background()
	})

	return environments
}

func IsAppEngine() bool {
	return appengine.IsAppEngine()
}

func IsDev() bool {
	return os.Getenv("ENV") == devEnv
}

func IsProd() bool {
	return os.Getenv("ENV") == prodEnv
}