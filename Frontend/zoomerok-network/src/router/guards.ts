import Cookies from 'js-cookie'

export const authGuard = (to: any, from: any, next: any) => {
  const token = Cookies.get('access_token')
  const authUrl = import.meta.env.VITE_AUTH_URL

  if (!token) {
    window.location.replace(`${authUrl}/signin`)
    return
  } else {
    next()
  }
}
