# Implementing Locales

A locale is authored **once, as slim JSON** (extracted from CLDR). From that one
file GoTime derives both a generated Go subpackage (fast path) and the runtime
`RegisterJSON` load path; UExL's host consumes the same JSON as a context map.
You do not hand-write Go locale tables, and you do not hand-write logic in JSON.

## 1. The Source: one slim JSON file

Add `locales/data/<code>.json` conforming to the schema in
[architecture.md §4.1](architecture.md) (validated against the published
`locales/data/schema.json`). It carries:

- `schemaVersion`, `code`
- `months` and `weekdays`, each with **`format`** and **`stand-alone`** contexts,
  each `short`/`long` (12 / 7 entries, Sunday = 0). Languages without a
  distinction repeat the same arrays — but the structure is mandatory so
  genitive languages (Russian, Polish, Greek, …) are correct.
- `dayPeriods` — `am`/`pm`, short + long (for `a`/`aa`)
- `ordinal` — an ordinal rule **id** (`"english"`, `"none"`, …)
- `plural.id` — a plural rule **id**, plus `plural.rules` (UExL-syntax condition
  expressions, for the UExL host and documentation)
- `relative` — phrase templates `[unit][tense][category]` with a `{n}` slot, plus
  `specials` for auto-style natural phrases
- `presets` — `@name` → NITES layout, using CLDR length names
  (`datefull/datelong/datemedium/dateshort`, `time/timeshort/time24`, …) plus
  `datetimeGlue` (the `{date}`/`{time}` combiner for `@datetime`)
- *optional, forward-looking:* `firstDay`, `weekend` (carried from CLDR now;
  Phase 1 does not yet wire them into week APIs)

Normally this file is produced by the Stage-0 extractor from CLDR (so names are
authoritative, not hand-typed). Hand-authoring is only for custom/non-CLDR
locales.

## 2. Plural & ordinal rules are referenced by id (logic lives in core)

Names are data; **plural selection is logic** and lives once in
`locales/plural.go` (see architecture §5). Your JSON only names the rule:

```jsonc
"plural": {
  "id": "west-slavic",
  "rules": { "one": "i = 1 and v = 0",
             "few": "v = 0 and i % 10 = 2..4 and i % 100 != 12..14",
             "many": "true" }
}
```

- **GoTime generated path:** the generator emits `plural: pluralWestSlavic` — a
  **symbol**, so a typo/unknown id is a *compile error*.
- **GoTime JSON path:** `RegisterJSON` resolves `plural.id` against the
  `pluralFuncs` map; an unknown id is a validation error caught by CI.
- **UExL path:** the host evaluates `plural.rules` (UExL syntax) with the CLDR
  operands as context — UExL is the expression engine; it never reads the file.

If your locale needs a plural rule that does not yet exist in core, add the
function to `locales/plural.go` and register it in `pluralFuncs` (a small,
reviewed change). Many languages share an existing rule and need no new code.

### Phrase tables must cover every category the rule can return

`west-slavic` can return `one/few/many`, so the `relative` templates for each
unit must define `one`, `few`, and `many`. A missing category is caught by the
per-locale test (JSON path) or compile/test (generated path), never silently
falling back to English.

## 3. Named format presets

Define each preset as a NITES layout in `presets`. This is where locale-driven
**ordering** comes from:

```jsonc
"presets": {
  "shortdate": "dd.mm.yyyy",     // de: "07.07.2025"
  "longdate":  "d mmmm yyyy",    // es: "7 de julio de 2025"  (use \\ to escape literal words)
  "time24":    "hhhh:ii"
}
```

Use the canonical preset names from [spec §3.3](spec.md) so callers are portable
across locales. Literal words inside a layout use NITES escaping (`\`).

## 4. TinyGo: build-gate the JSON path; keep `en` as Go literals

`encoding/json` uses reflection, which is heavy on TinyGo. The rule:

- **`en` is always built into core as Go literals** — no JSON, no reflection, no
  parse cost. It works on every target including TinyGo.
- **Generated-Go locales** are plain Go literals too — TinyGo-safe — but each is
  gated so it is only compiled when imported (it already is, via subpackage
  import) and excluded from size-critical builds if desired:

  ```go
  //go:build !tinygo
  package es
  ```

- **`RegisterJSON` (runtime JSON)** pulls in `encoding/json`; on TinyGo, prefer
  the generated-Go subpackage instead, or build-gate the JSON loader out. Do
  **not** hand-roll a reflection-free JSON decoder unless "non-English on
  TinyGo via JSON" is a hard requirement — it is bug-prone and usually
  unnecessary.

Net: a TinyGo build gets `en` (and any generated locale it imports) with zero
reflection; the `RegisterJSON` loader is for normal Go builds. UExL's host parses
the JSON on its own side — independent of GoTime's TinyGo constraints.

## 5. Required Tests

Every locale must include (build all times with explicit `time.UTC` + a fixed
base time; never depend on system locale/timezone):

**Correctness**
- `TimeAgoIn`: ≥1 past and ≥1 future sample, in **both** `StyleAuto` and
  `StyleNumeric`.
- `FormatIn`: ≥1 sample with `mmm`/`mmmm`, ≥1 with `www`/`wwww`, ≥1 with `a`/`aa`
  (day period), ≥1 `@preset` (incl. `@datetime` to exercise the glue).
- For languages with distinct grammatical context: ≥1 sample proving the
  **format** (in-date) name differs from the **stand-alone** name.
- `Ordinal`: ≥1 `dt` or `mt` sample.
- **Plural:** counts hitting every category the rule uses (Polish: 1→one, 2→few,
  5→many, 22→few, 112→many, 0→many), validated against CLDR plural test data.

**Repository-wide (not per-locale)**
- **Fallback:** subtag truncation — `es-MX → es`, `zh-Hant-HK → zh-Hant`
  (registered) and `→ en` (not) — via both `TimeAgoIn` and `FormatIn`.
- **Fuzz:** `RegisterJSON` against malformed/adversarial input must error, never
  panic.
- **Dual-path conformance:** load each locale via its generated-Go subpackage and
  via `RegisterJSON` of the same `data/<code>.json`; assert identical output. (So
  the two GoTime readers can never disagree; UExL, consuming the same JSON, stays
  consistent by construction.)
- **Schema validation:** load every `data/*.json`; assert `schemaVersion`,
  required fields, and `plural.id`/`ordinal` resolve.

**Allocation gate**
- `Benchmark*` with `b.ReportAllocs()` for `FormatIn` and `TimeAgoIn` meeting the
  budget in architecture §7; an English-path benchmark proving no regression.

## 6. Parsing-Readiness (forward-looking)

Localized **parsing** is Phase 3, but author tables so a reverse lookup is
trivial: names must be unique within their (kind, width) set, and note the
case-folding rule the locale needs (fold to lowercase; strip diacritics or not).
Phase 3 then builds `map[foldedName]time.Month` from existing data with no
redesign.

## 7. Complex Languages

Some languages inflect names by grammatical context (Polish "lipiec" vs.
"lipca"). For Phase 1 choose **one** consistent display form (typically
nominative) and document it in the JSON via a comment in the extractor mapping.
Context-aware forms are a later track via an optional extension interface
(architecture §9) — they must not change the core `Locale` contract.
