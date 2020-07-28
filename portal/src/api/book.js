import request from '@/utils/request'

export function contributeBook(data) {
  return request({
    url: '/contribute/book',
    method: 'post',
    data
  })
}

