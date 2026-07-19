# GoTime API & Format-Specifier Symmetry Review

Purpose: lay out **every** public function and **every** NITES format specifier
(plus the predefined/built-in layouts) in a comparable, grid form so symmetry
gaps and naming inconsistencies are easy to spot and discuss. This is a
**review/analysis** document — findings in §7 are candidates for decision, not
decided changes.

Method: tables enumerate what exists (✓ / actual name) vs. what is absent (✗).
"Absent" is not automatically a defect — some gaps are intentional. The point is
to make every gap *visible* so it can be judged deliberately. Sourced from
`relative_functions.go`, `time_arithmetic.go`, `quarters.go`,
`age_calculation.go`, `time_calc.go`, `business_calendar.go`, `calendar_math.go`,
`misc.go`, `range.go`, `weekday_*.go`, `format.go`, `parse.go`, `convert.go`,
and `internal/nites/convert.go`, `internal/utils/builtins.go`.

Legend: ✓ exists · ✗ absent · ⚠ exists but inconsistent (see note) · — N/A.

---

## 1. NITES Format-Specifier Ladder (by field)

Each row is a date/time field; columns are the *width/representation* variants.
A complete, symmetric field would offer the same rungs as its peers.

| Field | no-pad | zero-pad | space-pad | short name | long name | ordinal | other | Parseable? |
|---|---|---|---|---|---|---|---|---|
| Year | ✗ | `yy` (06), `yyyy` (2006) | ✗ | — | — | — | — | ✓ |
| Month | `m` (1) | `mm` (01) | ✗ | `mmm` (Jan) | `mmmm` (January) | `mt` (1st) | — | ✓ (names/num); `mt` ✗ |
| Day-of-month | `d` (2) | `dd` (02) | `db` (_2) | — | — | `dt` (2nd) | — | ✓ (`d/dd/db`); `dt` ✗ |
| Day-of-year | ✗ | `ddd` (002) | ✗ | — | — | — | — | ⚠ (Go parse of DOY is unreliable) |
| Weekday | ✗ (no numeric) | ✗ | — | `www` (Mon) | `wwww` (Monday) | — | ✗ (narrow) | ✓ (names) |
| Hour (12h) | `h` (3) | `hh` (03) | ✗ | — | — | — | — | ✓ |
| Hour (24h) | ✗ | `hhhh` (15) | ✗ | — | — | — | — | ✓ |
| AM/PM | — | — | — | `a` (pm) | `aa` (PM) | — | — | ✓ |
| Minute | `i` (4) | `ii` (04) | ✗ | — | — | — | — | ✓ |
| Second | `s` (5) | `ss` (05) | ✗ | — | — | — | — | ✓ |
| Fractional sec | ✗ | ✗ | — | — | — | — | — | ✗ (not implemented — see §7-F) |
| Timezone name | — | — | — | `zz` (MST) | ✗ (long/localized) | — | `z` (Z literal) | ✓ |
| TZ offset | `o` (±07) | `oo` (±0700) | — | — | — | — | `ooo` (±07:00) | ✓ |

**Visible ladder gaps** (for §7 discussion):
- **Year** has no non-padded form (`m`/`d` exist, no `y`).
- **Hour-24** has only the padded form (`hhhh`); 12h has both `h`/`hh`. No
  24h-no-pad, and no rung between `hh` and `hhhh` (`hhh` undefined).
- **Weekday** has no numeric form (Month/Day do): no `w`/`ww` (1–7 / 01–07).
- **Day-of-year** reuses the `d` prefix (`ddd`) for a *different* concept than
  day-of-month — overloaded naming; and exists only zero-padded.
- **Ordinals** (`dt`, `mt`) are the only **format-only** tokens (cannot parse).
- **Narrow** names (e.g. "J", "M") and **localized long TZ names** are absent
  (the former is also a stated localization non-goal).
- **Fractional seconds** (`0`/`9`) are *documented in the `convert.go` comments*
  (lines 82–83) but **not in the conversion map**, and the only test is
  commented out (`convert_test.go:180`) — i.e. **unimplemented**. See §7-F.

---

## 2. Format ↔ Parse ↔ Convert Symmetry

| Capability | Format side | Parse side | Convert |
|---|---|---|---|
| Core | `Format(dt, layout)` | `Parse(layout, value)` | `Convert(value, from, to)` |
| Location variant | ✗ (`FormatInLocation`) | `ParseInLocation` | ✗ |
| Unix variant | `FormatUnix`, `FormatTimestamp` | ✗ (`ParseUnix`) | — |
| Localized variant (planned) | `FormatIn` (spec) | ✗ (`ParseIn`?) | ✗ (`ConvertIn`?) |
| Ordinal tokens | supported (`dt/mt`) | **not** supported (errors) | from-side errors on `dt/mt` |

**Visible gaps:** Format has Unix/Timestamp variants but Parse has none; Parse has
an `InLocation` variant but Format has none; the planned localization adds
`FormatIn`/`TimeAgoIn` but no `ParseIn`/`ConvertIn` (decide whether parity is
wanted).

---

## 3. Predefined / Built-in Layouts

Today the only "predefined" layouts are **Go stdlib reference strings** passed
through verbatim (version-gated), from `internal/utils/builtins.go`:

| Group | Layouts |
|---|---|
| RFC | RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, RFC3339, RFC3339Nano |
| Classic | Layout, ANSIC, UnixDate, RubyDate, Kitchen |
| Stamp | Stamp, StampMilli, StampMicro, StampNano |
| Go 1.20+ | DateTime, DateOnly, TimeOnly |

There are **no GoTime-native semantic named presets today.** The localization
spec introduces `@`-presets (`@date`, `@time24`, `@datetime`, …). Cross-check:

| Concept | Today | Localization spec (planned) |
|---|---|---|
| Date only | `time.DateOnly` (fixed ISO) | `@datefull/long/medium/short` (locale) |
| Time only | `time.TimeOnly` | `@time/timeshort/time24` |
| Date+time | `time.DateTime` | `@datetime/datetime24` (locale glue) |
| Components | ✗ | `@monthyear/month/weekday` |

**Visible gap:** the predefined surface jumps from fixed Go constants straight to
locale presets — there is no locale-neutral named-preset layer (decide whether
e.g. `@iso`, `@rfc3339` NITES aliases are wanted for symmetry with the locale
presets).

---

## 4. Function Family Matrix (by time unit)

Columns are units; rows are operation families. This is the core symmetry view.

| Operation | Second | Minute | Hour | Day | Week | Month | Quarter | Year |
|---|---|---|---|---|---|---|---|---|
| Start-of | ✗ | ✗ | ✗ | `SoD` ⚠ | `WeekStart` | `MonthStart` | `QuarterStart` | `YearStart` |
| End-of | ✗ | ✗ | ✗ | `EoD` ⚠ | `WeekEnd` | `MonthEnd` | `QuarterEnd` | `YearEnd` |
| Start/End *On* | — | — | — | — | `WeekStartOn`/`WeekEndOn` | ✗ | ✗ | ✗ |
| Last | ✗ | ✗ | ✗ | `Yesterday` ⚠ | `LastWeek` | `LastMonth` | `LastQuarter` | `LastYear` |
| Next | ✗ | ✗ | ✗ | `Tomorrow` ⚠ | `NextWeek` | `NextMonth` | `NextQuarter` | `NextYear` |
| "Current" | ✗ | ✗ | ✗ | ✗ (`Today`) | ✗ | ✗ | ✗ | ✗ |
| Shift N | `Seconds` | `Minutes` | `Hours` | `Days` | `Weeks` | `Months` | `Quarters` | `Years` |
| Length (DaysIn…) | — | — | — | — | (7) | `DaysInMonth` | `DaysInQuarter` | `DaysInYear` |
| Between (frac) | ✗ | ✗ | ✗ | `DaysBetween` ⚠int | `WeeksBetween` | `MonthsBetween` | ✗ | `YearsBetween` |
| Between (cal int) | ✗ | ✗ | ✗ | (`DaysBetween`) | ✗ | `CalendarMonthsBetween` | ✗ | `CalendarYearsBetween` |
| X-of-Y | — | — | — | `DayOfYear` | `WeekOfMonth` | — | `QuarterOfYear` | ✗ (`WeekOfYear`) |

**Strong points (good symmetry):**
- **Shift-N** is complete and consistent across all 8 units (`Seconds … Years`).
- **Start/End** and **Last/Next** are complete and consistent for
  Week/Month/Quarter/Year.

**Visible asymmetries (for §7):**
- **Day** breaks the naming convention: `SoD`/`EoD` instead of `DayStart`/`DayEnd`,
  and `Yesterday`/`Tomorrow` instead of `LastDay`/`NextDay` (semantically fine,
  lexically inconsistent).
- **No `Today()`** despite `Yesterday()`/`Tomorrow()`.
- **Between return types are mixed:** `DaysBetween → int`, but
  `WeeksBetween/MonthsBetween/YearsBetween → float64`. No `Calendar*` for
  days/weeks/quarters; no `Quarter`/`Hour`/`Minute`/`Second` Between at all
  (though `Diff()` covers arbitrary units).
- **`WeekOfYear` missing** while `DayOfYear`, `WeekOfMonth`, `QuarterOfYear`
  exist (ISO week-of-year is a common ask).
- **Start/End *On*** exists only for Week; no `MonthStartOn` etc. (likely fine).

---

## 5. Predicates (`Is…`) Matrix

| Predicate | Exists | Symmetric counterpart missing? |
|---|---|---|
| `IsLeapYear` | ✓ | — |
| `IsBusinessDay` | ✓ | `IsWeekend` / `IsWeekday` ✗ |
| `IsFirstDayOfMonth` | ✓ | `IsFirstDayOfYear/Quarter/Week` ✗ |
| `IsLastDayOfMonth` | ✓ | `IsLastDayOfYear/Quarter/Week` ✗ |
| `IsBetween` | ✓ | — |
| `IsBetweenDates` | ✓ | — |
| `IsWeekdayPresentInRange` | ✓ | — |
| `IsValidAge` | ✓ | — |
| `IsToday/IsYesterday/IsTomorrow` | ✗ | (pairs with Today/Yesterday/Tomorrow) |
| `IsSameDay/IsSameMonth/IsSameYear` | ✗ | common convenience set |
| `IsPast/IsFuture` | ✗ | — |

---

## 6. Business-Day vs Work-Day APIs (representation symmetry)

Two parallel subsystems model the same concept with **different signatures**:

| Aspect | Business-day API | Work-day API |
|---|---|---|
| Predicate | `IsBusinessDay(t, weekends []time.Weekday, holidays…)` | — |
| Next | `NextBusinessDay(t, weekends []time.Weekday, holidays…)` | — |
| Prev | `PrevBusinessDay(t, weekends []time.Weekday, holidays…)` | `PrevWorkDay(start, days, workingDays [7]bool, holidays…)` |
| Add N | — | `WorkDay(start, days, workingDays [7]bool, holidays…)` |
| Count | — | `NetWorkDays(start, end, workingDays [7]bool, holidays…)` |
| **Weekend model** | **`weekends []time.Weekday`** | **`workingDays [7]bool`** |
| Error return | none (returns time) | `(time.Time, error)` / `(int, error)` |

**Visible asymmetry:** the two families use **incompatible weekend
representations** and **different error conventions**, and neither is a superset
(e.g. no `NetBusinessDays`, no `WorkDay` equivalent that uses `[]time.Weekday`).

---

## 7. Findings — candidates for discussion

Each is an *observation*; the "consider" column is a starting point, not a
decision. Severity = impact on consistency/usability, not urgency.

| ID | Area | Observation | Consider | Severity |
|---|---|---|---|---|
| A | Between types | `DaysBetween` returns `int`; `Weeks/Months/YearsBetween` return `float64` | unify, or document the split intentionally (`*Between` frac vs `Calendar*` int) | High |
| B | Between coverage | No `HoursBetween/MinutesBetween/SecondsBetween/QuartersBetween`; no `CalendarDaysBetween/WeeksBetween` | decide the canonical "between" surface (or point users to `Diff`) | Medium |
| C | Day naming | `SoD`/`EoD`/`Yesterday`/`Tomorrow` vs `*Start`/`*End`/`Last*`/`Next*` | add `DayStart`/`DayEnd` aliases for convention symmetry | Medium |
| D | Missing `Today()` | `Yesterday`/`Tomorrow` exist, no `Today` | add `Today()` | Low |
| E | `WeekOfYear` | `DayOfYear`, `WeekOfMonth`, `QuarterOfYear` exist; ISO `WeekOfYear` absent | add ISO-8601 `WeekOfYear` (and `IsoYear`?) | Medium |
| F | Fractional tokens | **Confirmed unimplemented:** `convert.go:82-83` documents `0`/`9` (fractional seconds) but they are absent from the conversion map and the only test is commented out (`convert_test.go:180`). A layout with `0`/`9` silently emits literals, and there is **no fractional-second token at all** | implement fractional-second tokens (and parse), or remove the misleading doc lines | High |
| G | Hour-24 ladder | only `hhhh` (padded); no 24h-no-pad; `h`/`hh` are 12h | add a 24h-no-pad token if wanted (naming TBD) | Low |
| H | Year no-pad | no non-padded year token | likely fine; note intentional | Low |
| I | Weekday numeric | no `w`/`ww` numeric weekday (Month/Day have numerics) | decide if numeric weekday is wanted | Low |
| J | Business vs Work | incompatible weekend models (`[]time.Weekday` vs `[7]bool`) + error conventions | converge on one weekend representation; fill the matrix (`NetBusinessDays`, etc.) | High |
| K | Format/Parse parity | `FormatUnix/Timestamp` but no `Parse*`; `ParseInLocation` but no `FormatInLocation` | decide which parity gaps to close | Medium |
| L | Localization parity | `FormatIn`/`TimeAgoIn` planned; no `ParseIn`/`ConvertIn` | decide whether localized parse/convert are in scope (Phase 3 covers parse) | Medium |
| M | `Is*` set | no `IsWeekend/IsWeekday/IsToday/IsSameDay/IsPast/IsFuture` | add the common convenience predicates | Medium |
| N | First/Last-day predicates | only Month variant of `IsFirst/LastDayOf…` | add Year/Quarter/Week variants for symmetry | Low |
| O | Day-of-year naming | `ddd` overloads the `d` prefix for day-of-year vs day-of-month | document clearly; consider a distinct token | Low |

---

## 8. How to use this document

- §1 / §4 / §5 / §6 are the **comparison grids** — scan a row for missing cells.
- §7 is the **issue backlog** — triage each into keep-as-is / add / fix / defer.
- Tie-in with localization: items **F, G, I, L** and the §3 preset layer
  interact with the localization work; resolve them before freezing the
  `@`-preset vocabulary and the `Localizer` token set so the localized and
  non-localized surfaces stay symmetric.
