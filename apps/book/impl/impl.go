package impl

import (
	"database/sql"

	"github.com/zginkgo/ginkgo_pkg/app"
	"github.com/zginkgo/ginkgo_pkg/logger"
	"github.com/zginkgo/ginkgo_pkg/logger/zap"
	"google.golang.org/grpc"

	"github.com/zginkgo/ginkgo_auth/apps/book"
	"github.com/zginkgo/ginkgo_auth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db *sql.DB

	log  logger.Logger
	book book.ServiceServer
	book.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db

	s.log = zap.L().Named(s.Name())
	s.book = app.GetGrpcApp(book.AppName).(book.ServiceServer)
	return nil
}

func (s *service) Name() string {
	return book.AppName
}

func (s *service) Registry(server *grpc.Server) {
	book.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
