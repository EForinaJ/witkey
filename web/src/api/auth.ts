import request from '@/utils/http'

/**
 * 登录
 * @param params 登录参数
 * @returns 登录响应
 */
export function fetchLogin(params: Auth.Params.Login) {
  return request.post<Auth.Response.Login>({
    url: '/login',
    params
  })
}

