import request from '@/utils/http'

export function fetchSite() {
    return request.get<Site.Response.Info>({
      url: '/site',
    })
}
