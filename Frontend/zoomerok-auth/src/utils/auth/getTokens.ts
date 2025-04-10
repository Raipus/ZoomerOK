import { getCookie, hasCookie, setCookie } from "cookies-next";

type Token = null | undefined | string;

interface Tokens {
  access: Token;
  refresh: Token;
}

export async function getTokens(): Promise<Tokens> {
  const tokens: Tokens = {
    access: null,
    refresh: null,
  };

  if (hasCookie("refresh_token"))
    tokens.refresh = getCookie("refresh_token") as Token;
  if (hasCookie("access_token"))
    tokens.access = getCookie("access_token") as Token;

  if (!tokens.refresh) {
    setTokens(null, null);
    return tokens;
  }

  if (!tokens.access) {
    const newTokens = await getTokensOrNullFromServer(tokens.refresh);

    if (newTokens) {
      setTokens(newTokens.accessToken, newTokens.refreshToken);
      tokens.access = newTokens.accessToken;
      tokens.refresh = newTokens.refreshToken;
    } else {
      setTokens(null, null);
    }
  }

  return tokens;
}

export function setTokens(access: string | null, refresh: string | null) {
  if (access) {
    setCookie("access_token", access, { maxAge: 60 * 15 });
    setCookie("refresh_token", refresh, { maxAge: 60 * 60 * 24 * 7 });
  }
}

export async function getTokensOrNullFromServer(refreshToken: string) {
  try {
    var access = "";
    var refresh = "";
    await fetch("http://localhost:3001/auth/refresh", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${refreshToken}`,
      },
    })
      .then((response) => response.json())
      .then((data1) => {
        access = data1.accessToken;
        refresh = data1.refreshToken;
      });
    return { accessToken: access, refreshToken: refresh };
  } catch (error) {
    console.error("Error:", error);
    return { accessToken: null, refreshToken: null };
  }
}
