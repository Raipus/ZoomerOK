"use client";

import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import { getToken } from "@/utils/auth/getToken";
import jwt from "jsonwebtoken";

interface JwtPayload {
  ConfirmEmail: boolean;
}

const Layout = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const [hasTokens, setHasTokens] = useState(true);
  const [confirmEmail, setConfirmEmail] = useState<boolean | null>(null);

  useEffect(() => {
    async function TokenCheck() {
      const accessToken = await getToken();
      if (!accessToken) {
        setHasTokens(false);
      } else {
        try {
          const decoded = jwt.decode(accessToken) as JwtPayload;
          setConfirmEmail(decoded?.ConfirmEmail);
        } catch (error) {
          setHasTokens(false);
        }
        if (confirmEmail) {
          setHasTokens(true);
          router.push(`${process.env.NEXT_PUBLIC_NETWORK_URL}`);
        } else {
          setHasTokens(false);
          router.push("/confirm-email");
        }
      }
    }
    TokenCheck();
  }, [router]);

  return <div style={{ height: "100vh" }}>{hasTokens ? null : children}</div>;
};

export default Layout;
