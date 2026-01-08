package dashboard

import (
	"server/internal/service"
)

type sDashborad struct{}

func init() {
	service.RegisterDashboard(&sDashborad{})
}
