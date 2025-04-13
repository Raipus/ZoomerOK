"use client";

import { useState, useEffect, useRef } from "react";
import ErrorNotification from "@/component/ErrorNotification";
import { useErrorNotification } from "@/hooks/useErrorNotification";
import { useParams, useRouter } from "next/navigation";
import Link from "next/link";

export default function ChangingPassPage() {
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const { error, showError, showNotification, hideNotification } =
    useErrorNotification();
  const [error1, setError1] = useState(false);
  const params = useParams<{ slug: string }>();
  const initialSlug = useRef(params.slug);
  const requestSent = useRef(false);

  useEffect(() => {
    const fetchChangingPass = async () => {
      if (requestSent.current) return;
      requestSent.current = true;

      try {
        setLoading(true);
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/auth/confirm_password/${initialSlug.current}`,
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          }
        );

        const data1 = await response.json();

        if (!response.ok) {
          setError1(true);
          showNotification(data1.message || "Произошла ошибка");
          return;
        }
      } catch (error) {
        setError1(true);
        showNotification(
          error instanceof Error ? error.message : "Неизвестная ошибка"
        );
      } finally {
        setLoading(false);
      }
    };

    fetchChangingPass();
  }, []);

  useEffect(() => {
    const timer = setTimeout(() => {
      router.push("/signin");
    }, 10 * 1000);

    return () => clearTimeout(timer);
  }, [router]);

  return (
    <div className="bg-cover">
      <main className="grid justify-items-center">
        <div className="grid place-items-center h-screen">
          {loading ? (
            <button
              type="button"
              className="m-[300px] inline-flex items-center"
            >
              <svg
                className="-ml-1 mr-3 h-5 w-5 animate-spin text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  className="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  strokeWidth="4"
                ></circle>
                <path
                  className="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              <p className="text-2xl">Загрузка...</p>
            </button>
          ) : error1 ? (
            <div className="grid place-items-center">
              <ErrorNotification
                message={error || ""}
                show={showError}
                onClose={hideNotification}
                duration={5000}
              />
              <p>Не получилось сменить Ваш пароль!</p>
              <p>Сейчас перенаправим Вас на страницу входа...</p>
              <div className="mt-3 grid justify-items-center hover:scale-102 duration-300">
                <Link href="/signin">↩ Вернуться на страницу входа</Link>
              </div>
            </div>
          ) : (
            <div className="grid place-items-center">
              <p>Ваш пароль успешно изменен!</p>
              <p>Сейчас перенаправим Вас на страницу входа...</p>
              <div className="mt-3 grid justify-items-center hover:scale-102 duration-300">
                <Link href="/signin">↩ Вернуться на страницу входа</Link>
              </div>
            </div>
          )}
        </div>
      </main>
    </div>
  );
}
