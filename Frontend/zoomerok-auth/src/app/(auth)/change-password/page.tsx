"use client";

import Link from "next/link";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import ErrorNotification from "@/component/ErrorNotification";
import { useErrorNotification } from "@/hooks/useErrorNotification";

interface IFormStateLogin {
  email: string;
}

export default function ChangePassPage() {
  const { register, handleSubmit } = useForm<IFormStateLogin>();
  const [loading, setLoading] = useState(false);
  const [finished, setFinished] = useState(false);
  const { error, showError, showNotification, hideNotification } =
    useErrorNotification();

  const onSubmit: SubmitHandler<IFormStateLogin> = async (data) => {
    setLoading(true);
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/account/want_change_password`,
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
        setFinished(true);
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
          ) : finished ? (
            <div className="div">
              <p>
                На указанную почту отправлено письмо с ссылкой на восстановление
                пароля!
              </p>
              <div className="mt-3 grid justify-items-center hover:scale-102 duration-300">
                <Link href="/signin">↩ Вернуться назад</Link>
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
              <div className="grid place-items-center h-[550px] w-[460px] bg-white rounded-[60px]">
                <div className="h-[524px] w-[431px] bg-[#7500DB] rounded-[50px]">
                  <div>
                    <form
                      onSubmit={handleSubmit(onSubmit)}
                      className="grid justify-self-center"
                    >
                      <div className="h-[100px] mt-[60px] leading-[50px]">
                        <p className="justify-self-center text-[40px] font-[1000] text-white italic">
                          Восстановление
                        </p>
                        <p className="justify-self-center text-[40px] font-[1000] text-white italic">
                          пароля
                        </p>
                      </div>
                      <input
                        className="rounded-[60px] text-gray-950 text-[22px] font-[800] bg-white mt-[34px] h-[61px] w-[343px] px-[30px]"
                        placeholder="Ваша почта"
                        type="text"
                        {...register("email", { required: true })}
                      />
                      <button
                        type="submit"
                        className="hover:scale-102 rounded-[60px] bg-[#FF00A9] text-[30px] text-white font-[1000] italic duration-300 hover:bg-[#ff00aaa9] mt-[50px] h-[108px] w-[343px]"
                      >
                        <p>Восстановить</p>
                        <p>пароль</p>
                      </button>
                    </form>
                  </div>
                  <div className="mt-[28px] grid justify-items-center hover:scale-102 duration-300 text-[16px] font-[800] underline decoration-[4.5%] underline-offset-[11%] text-white/80 h-[20px]">
                    <Link href="/signin">↩ Вернуться назад</Link>
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
