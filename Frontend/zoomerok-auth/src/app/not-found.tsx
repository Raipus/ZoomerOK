import Link from "next/link";

export default function NotFoundCatchAll() {
  return (
    <div className="grid place-items-center h-screen">
      <div className="grid place-items-center h-[335px] w-[1150px] bg-white rounded-[60px]">
        <div className="grid place-items-center h-[308px] w-[1118px] bg-[#7500DB] rounded-[50px]">
          <div className="justify-items-center">
            <div className="grid place-items-center">
              <p className="text-white font-[1000] italic text-[40px]">404</p>
              <p className="text-white font-[1000] italic text-[40px]">
                Этой страницы не существует
              </p>
            </div>
            <div className="grid place-items-center hover:scale-102 rounded-[60px] bg-[#FF00A9] text-[32px] text-white font-[1000] italic duration-300 hover:bg-[#ff00aaa9] mt-[20px] h-[63px] w-[600px]">
              <Link href="/signin">↩ Вернуться на страницу входа</Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
