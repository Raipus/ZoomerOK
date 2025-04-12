"use client";

import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";

import { getTokens } from "@/utils/auth/getTokens";

const Layout = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const [hasTokens, setHasTokens] = useState(true);
  const networkUrl = process.env.NEXT_PUBLIC_NETWORK_URL;

  useEffect(() => {
    async function TokenCheck() {
      let refreshToken = (await getTokens()).refresh;
      if (!refreshToken) {
        setHasTokens(false);
      } else {
        setHasTokens(true);
        router.push(networkUrl);
      }
    }
    TokenCheck();
  }, [router]);

  return <div style={{ height: "100vh" }}>{hasTokens ? null : children}</div>;
};

export default Layout;
