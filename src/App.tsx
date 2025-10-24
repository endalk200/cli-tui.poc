import { useEffect, useMemo, useState } from "react";
import { SpinnerMessage } from "./components/SpinnerMessage.tsx";
import { FileChangeList } from "./components/FileChangeList.tsx";
import { ReviewFindingsList } from "./components/ReviewFindingsList.tsx";
import { Summary } from "./components/Summary.tsx";
import { generateMockDiff, generateMockFindings } from "./mock.ts";

type Step = "fetching" | "diff" | "reviewing" | "findings" | "summary";

export function App() {
  const steps: Step[] = [
    "fetching",
    "diff",
    "reviewing",
    "findings",
    "summary",
  ];
  const [stepIndex, setStepIndex] = useState(0);
  const step = steps[stepIndex];

  const diffItems = useMemo(() => generateMockDiff(), []);
  const findingItems = useMemo(() => generateMockFindings(), []);

  useEffect(() => {
    const durations = [1200, 1600, 1200, 2500]; // fetching, show diff, reviewing, show findings
    if (stepIndex < durations.length) {
      const id = setTimeout(
        () => setStepIndex((i) => i + 1),
        durations[stepIndex]
      );
      return () => clearTimeout(id);
    }
  }, [stepIndex]);

  return (
    <>
      {step === "fetching" && (
        <SpinnerMessage message="Getting code to review" />
      )}
      {step === "diff" && <FileChangeList items={diffItems} />}
      {step === "reviewing" && <SpinnerMessage message="Reviewing code" />}
      {step === "findings" && <ReviewFindingsList items={findingItems} />}
      {step === "summary" && (
        <>
          <text>Review complete.</text>
          <Summary
            filesReviewed={diffItems.length}
            suggestions={findingItems.reduce(
              (n, f) => n + f.findings.length,
              0
            )}
          />
        </>
      )}
    </>
  );
}
