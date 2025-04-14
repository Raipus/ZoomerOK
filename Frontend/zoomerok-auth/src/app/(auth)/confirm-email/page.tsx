"use client";

export default function ConfirmEmailPage() {
  return (
    <div className="bg-cover grid place-items-center h-screen justify-items-center">
      <main className="grid place-items-center h-[335px] w-[1350px] bg-white rounded-[60px]">
        <div className="grid place-items-center h-[308px] w-[1318px] bg-[#7500DB] rounded-[50px]">
          <div className="grid place-items-center">
            <p className="text-white font-[1000] italic text-[30px]">
              Благодарим за регистрацию!
            </p>
            <p className="text-white font-[1000] italic text-[30px]">
              На указанную Вами почту отправлено письмо с ссылкой на её
              подтверждение.
            </p>
          </div>
        </div>
      </main>
    </div>
  );
}
