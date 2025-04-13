"use client";

import { setCookie } from "cookies-next";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import ErrorNotification from "@/component/ErrorNotification";
import { useErrorNotification } from "@/hooks/useErrorNotification";

interface IFormStateSignIn {
  login: string;
  email: string;
  name: string;
  password: string;
  password2: string;
}

export default function SignupPage() {
  const { register, handleSubmit } = useForm<IFormStateSignIn>();
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const { error, showError, showNotification, hideNotification } =
    useErrorNotification();

  const onSubmit: SubmitHandler<IFormStateSignIn> = async (data) => {
    setLoading(true);
    try {
      const { password2, ...dataWithoutPass2 } = data;
      if (data.password != password2) {
        showNotification("Пароли должны совпадать!");
      } else {
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/auth/signup`,
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(dataWithoutPass2),
          }
        );

        const data1 = await response.json();

        if (!response.ok) {
          showNotification(data1.message || "Произошла ошибка");
          return;
        } else {
          setCookie("access_token", data1.accessToken, { maxAge: 60 * 60 });
          router.push("/confirm-email");
        }
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
    <div>
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
          ) : (
            <div>
              <ErrorNotification
                message={error || ""}
                show={showError}
                onClose={hideNotification}
                duration={5000}
              />
              <div>
                <form
                  onSubmit={handleSubmit(onSubmit)}
                  className="grid space-y-4 justify-self-center"
                >
                  <h1 className="mb-3 text-center text-2xl">Регистрация</h1>
                  <input
                    className="rounded-md border-[1px] border-black p-1 text-black bg-white"
                    placeholder="Логин"
                    type="text"
                    {...register("login", { required: true })}
                  />
                  <input
                    className="rounded-md border-[1px] border-black p-1 text-black bg-white"
                    placeholder="Почта"
                    type="email"
                    {...register("email", { required: true })}
                  />
                  <input
                    className="rounded-md border-[1px] border-black p-1 text-black bg-white"
                    placeholder="Имя"
                    type="text"
                    {...register("name", { required: true })}
                  />
                  <input
                    className="rounded-md border-[1px] border-black p-1 text-black bg-white"
                    placeholder="Пароль"
                    type="password"
                    {...register("password", { required: true })}
                  />
                  <input
                    className="rounded-md border-[1px] border-black p-1 text-black bg-white"
                    placeholder="Повторите пароль"
                    type="password"
                    {...register("password2", { required: true })}
                  />
                  <button
                    type="submit"
                    className="hover:scale-102 rounded-md bg-[#3D8361] px-5 py-2 text-xl text-white duration-300 hover:bg-[#2F6A4E]"
                  >
                    Зарегистрироваться
                  </button>
                </form>
              </div>
              <div className="mt-3 grid justify-items-center hover:scale-102 duration-300">
                <Link href="/signin">Уже есть аккаунт?</Link>
              </div>
            </div>
          )}
        </div>
      </main>
    </div>
  );
}
