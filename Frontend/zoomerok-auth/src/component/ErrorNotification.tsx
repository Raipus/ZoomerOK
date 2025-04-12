import React, { useEffect, useRef, useState } from "react";

interface ErrorNotificationProps {
  message: string;
  show: boolean;
  duration?: number;
  onClose?: () => void;
}

const ErrorNotification: React.FC<ErrorNotificationProps> = ({
  message,
  show,
  duration = 5000,
  onClose,
}) => {
  const [isVisible, setIsVisible] = useState(false);
  const [currentMessage, setCurrentMessage] = useState("");

  useEffect(() => {
    if (show) {
      setCurrentMessage(message);
      setIsVisible(true);
      const hideTimer = setTimeout(() => {
        setIsVisible(false);
        onClose?.();
      }, duration);

      return () => clearTimeout(hideTimer);
    } else {
      setIsVisible(false);
    }
  }, [show, message, duration, onClose]);

  const handleClose = () => {
    setIsVisible(false);
    onClose?.();
  };

  return (
    <div
      className={`fixed inset-x-0 top-0 z-50 flex justify-center transition-all duration-300 ease-in-out ${
        isVisible
          ? "translate-y-0 opacity-100"
          : "-translate-y-full opacity-0 pointer-events-none"
      }`}
      style={{
        display: show || isVisible ? "flex" : "none",
      }}
    >
      <div className="mt-4 px-6 py-3 bg-red-500 text-white rounded-lg shadow-lg">
        <div className="flex items-center justify-between">
          <div className="flex items-center">
            <svg
              className="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <span>{currentMessage}</span>
          </div>
          <button
            onClick={handleClose}
            className="ml-4 text-white hover:text-red-200 focus:outline-none"
          >
            <svg
              className="w-5 h-5"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
  );
};

export default ErrorNotification;
