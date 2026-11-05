import { useState } from "react";

type Status = "idle" | "loading" | "success" | "error";

const messages: Record<Status, string> = {
  idle: "ボタンを押して接続を確認してください。",
  loading: "接続を確認しています。",
  success: "接続できました。",
  error:
    "接続できませんでした。APIとDBの起動を確認して、もう一度お試しください。",
};

export function App() {
  const [status, setStatus] = useState<Status>("idle");

  async function checkConnection() {
    setStatus("loading");
    try {
      const response = await fetch("/api/ready", {
        signal: AbortSignal.timeout(5000),
        cache: "no-store",
      });
      if (!response.ok) throw new Error("Connection failed");
      const body: unknown = await response.json();
      if (
        typeof body !== "object" ||
        body === null ||
        !("status" in body) ||
        body.status !== "ready"
      ) {
        throw new Error("Unexpected response");
      }
      setStatus("success");
    } catch {
      setStatus("error");
    }
  }

  return (
    <main className="mx-auto my-[10vh] max-w-160 rounded-2xl border border-[#d8e0eb] bg-white p-[clamp(24px,5vw,48px)] wrap-anywhere">
      <p className="my-[1em] font-semibold text-[#475569]">AIジャンキー</p>
      <h1 className="my-[0.67em] text-[clamp(1.6rem,4vw,2rem)] font-bold">
        接続確認
      </h1>
      <p className="my-[1em]">
        開発環境の画面・API・データベースの接続を確認します。
      </p>
      <button
        className="min-h-11 max-w-full cursor-pointer rounded-lg border-0 bg-accent px-6 py-3 font-semibold text-white focus-visible:outline-3 focus-visible:outline-offset-4 focus-visible:outline-accent disabled:cursor-wait disabled:opacity-65"
        onClick={checkConnection}
        disabled={status === "loading"}
      >
        {status === "loading" ? "確認中…" : "接続を確認"}
      </button>
      <p className="my-[1em] min-h-[3.4em]" role="status" aria-live="polite">
        {messages[status]}
      </p>
    </main>
  );
}
