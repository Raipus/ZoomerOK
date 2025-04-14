"use client";

import { setCookie } from "cookies-next";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import ErrorNotification from "@/component/ErrorNotification";
import { useErrorNotification } from "@/hooks/useErrorNotification";

interface IFormStateLogin {
  somelogin: string;
  email: string;
  login: string;
  password: string;
}

export default function SigninPage() {
  const { register, handleSubmit } = useForm<IFormStateLogin>();
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const { error, showError, showNotification, hideNotification } =
    useErrorNotification();

  const onSubmit: SubmitHandler<IFormStateLogin> = async (data) => {
    setLoading(true);
    try {
      if (data.somelogin.includes("@")) {
        const { somelogin, login, ...dataWithEmail } = data;
        dataWithEmail.email = somelogin;
      } else {
        const { somelogin, email, ...dataWithEmail } = data;
        dataWithEmail.login = somelogin;
      }

      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/auth/login`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        }
      );

      const data1 = await response.json();

      if (!response.ok) {
        showNotification(data1.message || "Произошла ошибка");
        return;
      } else {
        setCookie("access_token", data1.accessToken, { maxAge: 60 * 60 });
        router.push(`${process.env.NEXT_PUBLIC_NETWORK_URL}`);
      }
    } catch (error) {
      showNotification(
        error instanceof Error ? error.message : "Неизвестная ошибка"
      );
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className="bg-cover">
      <main className="grid justify-items-center">
        <div className="grid place-items-center h-screen">
          {loading ? (
            <div className="grid place-items-center h-[138px] w-[514px] bg-white rounded-[60px]">
              <div className="h-[122px] w-[495px] bg-[#FF00A9] rounded-[50px]">
                <div className="grid place-items-center my-[30px]">
                  <button type="button" className="inline-flex items-center">
                    <svg
                      className="mr-[40px] h-5 w-5 animate-spin text-white scale-[200%]"
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
                    <p className="text-[40px] font-[1000] italic">
                      Загрузка...
                    </p>
                  </button>
                </div>
              </div>
            </div>
          ) : (
            <div>
              <ErrorNotification
                message={error || ""}
                show={showError}
                onClose={hideNotification}
                duration={5000}
              />
              <div className="grid place-items-center h-[550px] w-[455px] bg-white rounded-[60px]">
                <div className="h-[524px] w-[425px] bg-[#7500DB] rounded-[50px]">
                  <div>
                    <form
                      onSubmit={handleSubmit(onSubmit)}
                      className="grid place-items-center"
                    >
                      <h1 className="mt-[59px] text-[48px] italic font-[1000] text-white">
                        Вход
                      </h1>
                      <input
                        className="rounded-[60px] text-gray-950 text-[22px] font-[800] bg-white mt-[34px] h-[61px] w-[368px] px-[30px]"
                        placeholder="Логин / Почта"
                        type="text"
                        {...register("somelogin", { required: true })}
                      />
                      <input
                        className="rounded-[60px] text-gray-950 text-[22px] font-[800] bg-white mt-[40px] h-[61px] w-[368px] px-[30px]"
                        placeholder="Пароль"
                        type="password"
                        {...register("password", { required: true })}
                      />
                      <button
                        type="submit"
                        className="hover:scale-102 rounded-[60px] bg-[#FF00A9] px-[45px] py-2 text-[32px] text-white font-[1000] italic duration-300 hover:bg-[#ff00aaa9] mt-[34px] h-[61px] w-[205px]"
                      >
                        Войти
                      </button>
                    </form>
                  </div>
                  <div className="mt-[24px] grid justify-items-center hover:scale-102 duration-300 text-[14px] font-[900] underline decoration-[4.5%] underline-offset-[11%] text-white/80">
                    <Link href="/signup">Нет аккаунта?</Link>
                  </div>
                  <div className="mt-[12px] grid justify-items-center hover:scale-102 duration-300 text-[14px] font-[900] underline decoration-[4.5%] underline-offset-[11%] text-white/80">
                    <Link href="/change-password">Забыли пароль?</Link>
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>
      </main>
    </div>
  );
}
