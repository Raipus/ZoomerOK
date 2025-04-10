"use client";

import { setCookie } from "cookies-next";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";

interface IFormStateLogin {
  login: string;
  password: string;
}

export default function SigninPage() {
  const { register, handleSubmit } = useForm<IFormStateLogin>();
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState();

  const onSubmit: SubmitHandler<IFormStateLogin> = async (data) => {
    setLoading(true);
    const response = await fetch("http://localhost:3001/auth/signin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    const data1 = await response.json();

    if (!response.ok) {
      setLoading(false);
      setError(data1.message);
    } else {
      setCookie("access_token", data1.accessToken, { maxAge: 60 * 15 });
      setCookie("refresh_token", data1.refreshToken, {
        maxAge: 60 * 60 * 24 * 7,
      });
      router.push("http://localhost:5173/");
    }
  };
  return (
    <div className="bg-cover">
      <main className="grid justify-items-center">
        {loading ? (
          <div>
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
          </div>
        ) : (
          <div className="grid place-items-center h-screen">
            {error && (
              <p className="mb-3 mt-3 text-center text-2xl text-red-400">
                {error}
              </p>
            )}
            <div>
              <div>
                <form
                  onSubmit={handleSubmit(onSubmit)}
                  className="grid space-y-4 justify-self-center"
                >
                  <h1 className="mb-3 justify-self-center text-2xl">Вход</h1>
                  <input
                    className="rounded-md border-[1px] border-black p-1 text-black bg-white"
                    placeholder="Логин"
                    type="text"
                    {...register("login", { required: true })}
                  />
                  <input
                    className="rounded-md border-[1px] border-black p-1 text-black bg-white"
                    placeholder="Пароль"
                    type="password"
                    {...register("password", { required: true })}
                  />
                  <button
                    type="submit"
                    className="hover:scale-102 rounded-md bg-[#3D8361] px-5 py-2 text-xl text-white duration-300 hover:bg-[#2F6A4E]"
                  >
                    Войти
                  </button>
                </form>
              </div>
              <div className="mt-3 grid justify-items-center hover:scale-102 duration-300">
                <Link href="http://localhost:3000/signup">
                  Еще нет аккаунта?
                </Link>
              </div>
              <div className="mt-3 grid justify-items-center hover:scale-102 duration-300">
                <Link href="http://localhost:3000/signin">Забыли пароль?</Link>
              </div>
            </div>
          </div>
        )}
      </main>
    </div>
  );
}
