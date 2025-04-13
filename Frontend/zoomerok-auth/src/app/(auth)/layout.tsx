"use client";

import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";

import { getTokens } from "@/utils/auth/getToken";

const Layout = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const [hasTokens, setHasTokens] = useState(true);

  useEffect(() => {
    async function TokenCheck() {
      const accessToken = await getTokens();
      if (!accessToken) {
        setHasTokens(false);
      } else {
        setHasTokens(true);
        router.push(
          `${process.env.NEXT_PUBLIC_NETWORK_URL}/callback#token=${accessToken}`
        );
      }
    }
    TokenCheck();
  }, [router]);

  return <div style={{ height: "100vh" }}>{hasTokens ? null : children}</div>;
};

export default Layout;
