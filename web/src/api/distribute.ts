
import request from '@/utils/http'

export function fetchGetDistributeList(params: Distribute.Params.Query) {
    return request.get<Distribute.Response.List>({
      url: '/distribute/list',
      params
    })
}
export function fetchGetDistributeDetail(params: {id:number}) {
  return request.get<Distribute.Response.Detail>({
    url: '/distribute/detail',
    params
  })
}
export function fetchPostDistributeStart(data:{id:number}) {
  return request.post({
    url: '/distribute/start',
    data
  })
}
export function fetchPostDistributeComplete(data:{id:number}) {
  return request.post({
    url: '/distribute/complete',
    data
  })
}
export function fetchPostDistributeSettlement(data:Distribute.Params.Settlement) {
  return request.post({
    url: '/distribute/settlement',
    data
  })
}