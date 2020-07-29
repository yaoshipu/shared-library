import request from '@/utils/request'

export function donateBook(data) {
  return request({
    url: '/donate/book',
    method: 'post',
    data
  })
}

export function getBooks(params) {
  return request({
    url: '/books',
    method: 'get',
    params
  })
}

export function getMyBooks(params) {
  return request({
    url: '/mybooks',
    method: 'get',
    params
  })
}
