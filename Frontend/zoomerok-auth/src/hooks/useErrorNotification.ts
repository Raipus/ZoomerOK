import { useState } from "react";

export const useErrorNotification = () => {
  const [error, setError] = useState<string | null>(null);
  const [showError, setShowError] = useState(false);

  const showNotification = (errorMessage: string) => {
    setError(errorMessage);
    setShowError(true);
  };

  const hideNotification = () => {
    setShowError(false);
    setError(null);
  };

  return {
    error,
    showError,
    showNotification,
    hideNotification,
  };
};
