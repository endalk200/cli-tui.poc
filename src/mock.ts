export type DiffItem = {
  filename: string;
  added: number;
  changed: number;
  removed: number;
};

export type FindingItem = {
  filename: string;
  findings: { location: string; finding: string; suggestion: string }[];
};

export function generateMockDiff(): DiffItem[] {
  return [
    { filename: "src/cli/review.ts", added: 42, changed: 7, removed: 5 },
    { filename: "src/utils/diff.ts", added: 10, changed: 3, removed: 1 },
    { filename: "src/analyzers/ts-lint.ts", added: 18, changed: 6, removed: 0 },
    { filename: "README.md", added: 4, changed: 1, removed: 2 },
  ];
}

export function generateMockFindings(): FindingItem[] {
  return [
    {
      filename: "src/cli/review.ts",
      findings: [
        {
          location: "L120-132",
          finding: "Deeply nested condition reduces readability",
          suggestion: "Refactor into guard clauses and extract helper function",
        },
        {
          location: "L191",
          finding: "Unhandled promise rejection in async path",
          suggestion: "Wrap with try/catch and surface error via Result type",
        },
      ],
    },
    {
      filename: "src/utils/diff.ts",
      findings: [
        {
          location: "L45",
          finding: "Potential O(n^2) loop on large inputs",
          suggestion: "Use map for lookups and break early when possible",
        },
      ],
    },
    {
      filename: "src/analyzers/ts-lint.ts",
      findings: [
        {
          location: "L72",
          finding: "Magic numbers for severity levels",
          suggestion: "Introduce enum and centralize mapping",
        },
      ],
    },
  ];
}
