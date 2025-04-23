import Cookies from 'js-cookie'
import { jwtDecode } from 'jwt-decode'

export const authGuard = (to: any, from: any, next: any) => {
  const token = Cookies.get('access_token')
  const authUrl = import.meta.env.VITE_AUTH_URL
  const isEmailConfirmed = (token: any) => {
    try {
      const decoded = jwtDecode(token) as JwtPayload
      return decoded?.ConfirmEmail
    } catch {
      return false
    }
  }

  interface JwtPayload {
    ConfirmEmail: boolean
    // Другие поля
  }

  if (!token) {
    window.location.replace(`${authUrl}/signin`)
    return
  } else if (isEmailConfirmed(token) === false) {
    window.location.replace(`${authUrl}/confirm-email`)
  } else {
    next()
  }
}
