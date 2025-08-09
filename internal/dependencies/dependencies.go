package dependencies

import (
	"database/sql"
	"net/http"
	"sync"

	"github.com/eduufreire/url-shortner/internal/database"
	"github.com/eduufreire/url-shortner/internal/user"
)

type injector struct {
	database *sql.DB

	UserRoutes     *http.ServeMux
	userHandler    user.UserHandler
	userService    user.UserService
	userRepository user.UserRepository
}

var (
	instance *injector
	once     sync.Once
)

func Init() *injector {
	once.Do(func() {
		instance = &injector{}
	})
	return instance
}

func (i *injector) Wire() *injector {

	if i.database == nil {
		i.database = database.CreateDatabase()
	}

	if i.userRepository == nil {
		i.userRepository = user.NewUserRepository(i.database)
	}

	if i.userService == nil {
		i.userService = user.NewUserService(i.userRepository)
	}

	if i.userHandler == nil {
		i.userHandler = user.NewUserHandler(i.userService)
	}

	if(i.UserRoutes == nil) {
		i.UserRoutes = user.Routes(i.userHandler)
	}

	return i
}
