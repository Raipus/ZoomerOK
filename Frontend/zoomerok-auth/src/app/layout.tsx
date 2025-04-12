import type { Metadata } from "next";
import { Mulish } from "next/font/google";
import "./globals.css";

const mulish = Mulish({
  subsets: ["cyrillic", "latin"],
});

export const metadata: Metadata = {
  title: "ZoomerOk",
  description: "Социальная сеть для настоящих зумеров!",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ru">
      <body className={mulish.className}>{children}</body>
    </html>
  );
}
