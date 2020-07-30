import request from '@/utils/request'

export function donateBook(data) {
  return request({
    url: '/books',
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

export function getMyBooks(username) {
  return request({
    url: `/books/${username}`,
    method: 'get',
  })
}

export function returnBook(data) {
  return request({
    url: '/books/return',
    method: 'post',
    data
  })
}

export function borrowBook(data) {
  return request({
    url: '/books/borrow',
    method: 'post',
    data
  })
}
