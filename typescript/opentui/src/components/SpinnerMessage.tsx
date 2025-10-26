import { useEffect, useState } from "react";

type SpinnerMessageProps = {
  message: string;
};

export function SpinnerMessage({ message }: SpinnerMessageProps) {
  const frames = ["|", "/", "-", "\\"];
  const [index, setIndex] = useState(0);

  useEffect(() => {
    const id = setInterval(() => {
      setIndex((i) => (i + 1) % frames.length);
    }, 120);
    return () => clearInterval(id);
  }, []);

  return <text>{`${frames[index]} ${message}`}</text>;
}
