package app

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/redis/go-redis/v9"
	"github.com/singhdurgesh/rednote/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Cache *redis.Client
var Broker *machinery.Server
var Config *configs.Config
var Logger *logrus.Logger
