import { AppRouteRecord } from '@/types/router'

export const dashboardRoutes: AppRouteRecord = {
  name: 'Dashboard',
  path: '/dashboard',
  component: '/index/index',
  meta: {
    title: '仪表盘',
    icon: 'solar:pie-chart-2-bold',
  },
  children: [
    {
      path: 'console',
      name: 'Console',
      component: '/dashboard/console',
      meta: {
        title: '工作台',
        keepAlive: false,
      }
    },
    {
      path: 'distribute',
      name: 'DashboardDistribute',
      component: '/distribute',
      meta: {
          title: '我的服务',
          keepAlive: true,
      },
    },
    {
      path: 'withdraw',
      name: 'DashboardWithdraw',
      component: '/withdraw',
      meta: {
          title: '佣金提现',
          keepAlive: true,
      },
  },
  ]
}
