import { getCookie, hasCookie } from "cookies-next";

type Token = null | undefined | string;

export async function getToken(): Promise<Token> {
  let token;

  if (hasCookie("access_token")) token = getCookie("access_token") as Token;

  if (!token) {
    return null;
  }

  return token;
}
