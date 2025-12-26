[Home](../README.md) > [API Reference](README.md) > Conversion

# Date Conversion

GoTime provides seamless conversion between different date formats using the same intuitive NITES format specifiers.

## Core Function

### Convert

Converts a date string from one format to another.

```go
func Convert(value, fromLayout, toLayout string) (string, error)
```

**Parameters:**
- `value`: The input date string to convert
- `fromLayout`: NITES format specifier describing the input format
- `toLayout`: NITES format specifier describing the desired output format

**Returns:**
- `string`: The converted date string
- `error`: Error if conversion fails

**Basic Examples:**

```go
// ISO to US format
result, err := gotime.Convert("2025-07-07", "yyyy-mm-dd", "mm/dd/yyyy")
// Result: "07/07/2025"

// US to European format
result, err = gotime.Convert("07/07/2025", "mm/dd/yyyy", "dd/mm/yyyy")
// Result: "07/07/2025"

// Short to readable format
result, err = gotime.Convert("2025-07-07", "yyyy-mm-dd", "mmmm dt, yyyy")
// Result: "July 7th, 2025"

// Add weekday
result, err = gotime.Convert("2025-07-07", "yyyy-mm-dd", "wwww, mmmm dt, yyyy")
// Result: "Monday, July 7th, 2025"
```

## Common Conversion Patterns

### 1. Database to API Format

```go
// Convert database timestamp to API date
dbTimestamp := "2025-07-07 14:30:45"
apiDate, err := gotime.Convert(dbTimestamp, "yyyy-mm-dd hh:ii:ss", "yyyy-mm-dd")
// Result: "2025-07-07"

// Convert to ISO 8601 with timezone
isoDate, err := gotime.Convert(dbTimestamp, "yyyy-mm-dd hh:ii:ss", "yyyy-mm-dd'T'hh:ii:ss'Z'")
// Result: "2025-07-07T14:30:45Z"
```

### 2. User Input Normalization

```go
func normalizeUserDate(input string) (string, error) {
    // Common user input formats
    formats := []string{
        "mm/dd/yyyy",      // US format
        "dd/mm/yyyy",      // European format
        "yyyy-mm-dd",      // ISO format
        "m/d/yyyy",        // US short
        "d/m/yyyy",        // European short
        "mmmm dt, yyyy",   // Readable format
    }

    // Try each format and convert to standard ISO
    for _, format := range formats {
        if result, err := gotime.Convert(input, format, "yyyy-mm-dd"); err == nil {
            return result, nil
        }
    }

    return "", fmt.Errorf("unable to parse date: %s", input)
}

// Usage
normalized, err := normalizeUserDate("July 7th, 2025")
// Result: "2025-07-07"
```

### 3. Localization

```go
package localization

import "github.com/maniartech/gotime/v2"

type DateLocalizer struct {
    formats map[string]string
}

func NewDateLocalizer() *DateLocalizer {
    return &DateLocalizer{
        formats: map[string]string{
            "us":   "mm/dd/yyyy",
            "eu":   "dd/mm/yyyy",
            "iso":  "yyyy-mm-dd",
            "uk":   "dd/mm/yyyy",
            "ca":   "yyyy-mm-dd", // Canadian
            "au":   "dd/mm/yyyy", // Australian
        },
    }
}

func (dl *DateLocalizer) LocalizeDate(date, locale string) (string, error) {
    targetFormat, exists := dl.formats[locale]
    if !exists {
        return "", fmt.Errorf("unsupported locale: %s", locale)
    }

    // Convert from ISO standard to locale format
    return gotime.Convert(date, "yyyy-mm-dd", targetFormat)
}

// Usage
localizer := NewDateLocalizer()
usDate, _ := localizer.LocalizeDate("2025-07-07", "us")      // "07/07/2025"
euDate, _ := localizer.LocalizeDate("2025-07-07", "eu")      // "07/07/2025"
```

### 4. Report Generation

```go
package reports

type ReportFormatter struct {
    inputFormat  string
    outputFormat string
}

func NewReportFormatter(inputFmt, outputFmt string) *ReportFormatter {
    return &ReportFormatter{
        inputFormat:  inputFmt,
        outputFormat: outputFmt,
    }
}

func (rf *ReportFormatter) FormatDateColumn(dates []string) ([]string, error) {
    formatted := make([]string, len(dates))

    for i, date := range dates {
        converted, err := gotime.Convert(date, rf.inputFormat, rf.outputFormat)
        if err != nil {
            return nil, fmt.Errorf("error formatting date %s: %v", date, err)
        }
        formatted[i] = converted
    }

    return formatted, nil
}

// Usage
formatter := NewReportFormatter("yyyy-mm-dd", "mmmm yyyy")
dates := []string{"2025-01-01", "2025-02-15", "2025-03-30"}
formatted, _ := formatter.FormatDateColumn(dates)
// Result: ["January 2025", "February 2025", "March 2025"]
```

## Advanced Conversion Examples

### 1. Time Format Conversion

```go
// 12-hour to 24-hour
time24, err := gotime.Convert("2:30 PM", "h:ii aa", "hhhh:ii")
// Result: "14:30"

// 24-hour to 12-hour
time12, err := gotime.Convert("14:30", "hhhh:ii", "h:ii aa")
// Result: "2:30 PM"

// Add seconds
withSeconds, err := gotime.Convert("14:30", "hhhh:ii", "hhhh:ii:ss")
// Result: "14:30:00"
```

### 2. Complex Format Transformations

```go
// Transform log format to readable format
logEntry := "2025-07-07 14:30:45.123"
readable, err := gotime.Convert(logEntry, "yyyy-mm-dd hh:ii:ss.0", "wwww, mmmm dt at h:ii aa")
// Result: "Monday, July 7th at 2:30 PM"

// Extract date from datetime
dateOnly, err := gotime.Convert(logEntry, "yyyy-mm-dd hh:ii:ss.0", "yyyy-mm-dd")
// Result: "2025-07-07"

// Extract time from datetime
timeOnly, err := gotime.Convert(logEntry, "yyyy-mm-dd hh:ii:ss.0", "hh:ii:ss")
// Result: "14:30:45"
```

### 3. Format Validation and Conversion

```go
func validateAndConvert(input, expectedFormat, outputFormat string) (string, error) {
    // First, validate that input matches expected format
    _, err := gotime.Parse(input, expectedFormat)
    if err != nil {
        return "", fmt.Errorf("input doesn't match expected format %s: %v", expectedFormat, err)
    }

    // If valid, convert to output format
    return gotime.Convert(input, expectedFormat, outputFormat)
}

// Usage
result, err := validateAndConvert("2025-07-07", "yyyy-mm-dd", "mmmm dt, yyyy")
if err != nil {
    log.Printf("Validation failed: %v", err)
} else {
    fmt.Println(result) // "July 7th, 2025"
}
```

## Real-World Use Cases

### 1. API Gateway Date Transformation

```go
package gateway

import (
    "encoding/json"
    "github.com/maniartech/gotime/v2"
    "net/http"
)

type DateTransformer struct {
    inputFormat  string
    outputFormat string
}

func (dt *DateTransformer) TransformResponse(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
    // Transform all date fields in response
    for key, value := range data {
        if str, ok := value.(string); ok && dt.isDateField(key) {
            if converted, err := gotime.Convert(str, dt.inputFormat, dt.outputFormat); err == nil {
                data[key] = converted
            }
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func (dt *DateTransformer) isDateField(key string) bool {
    dateFields := []string{"created_at", "updated_at", "start_date", "end_date", "birth_date"}
    for _, field := range dateFields {
        if key == field {
            return true
        }
    }
    return false
}
```

### 2. Data Migration

```go
package migration

import (
    "bufio"
    "github.com/maniartech/gotime/v2"
    "os"
    "strings"
)

func migrateDateFormats(inputFile, outputFile, fromFormat, toFormat string) error {
    input, err := os.Open(inputFile)
    if err != nil {
        return err
    }
    defer input.Close()

    output, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer output.Close()

    scanner := bufio.NewScanner(input)
    writer := bufio.NewWriter(output)
    defer writer.Flush()

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, ",")

        // Assume first field is date
        if len(fields) > 0 {
            if converted, err := gotime.Convert(fields[0], fromFormat, toFormat); err == nil {
                fields[0] = converted
            }
        }

        newLine := strings.Join(fields, ",")
        writer.WriteString(newLine + "\n")
    }

    return scanner.Err()
}

// Usage
err := migrateDateFormats(
    "old_data.csv",
    "new_data.csv",
    "mm/dd/yyyy",    // From US format
    "yyyy-mm-dd",    // To ISO format
)
```

### 3. Template Date Formatting

```go
package templates

import (
    "github.com/maniartech/gotime/v2"
    "html/template"
    "strings"
)

// Custom template function for date conversion
func createTemplateFuncs() template.FuncMap {
    return template.FuncMap{
        "dateFormat": func(date, fromFormat, toFormat string) string {
            result, err := gotime.Convert(date, fromFormat, toFormat)
            if err != nil {
                return date // Return original on error
            }
            return result
        },
        "humanDate": func(date, format string) string {
            result, err := gotime.Convert(date, format, "mmmm dt, yyyy")
            if err != nil {
                return date
            }
            return result
        },
    }
}

// Template usage:
// {{ .StartDate | humanDate "yyyy-mm-dd" }}
// {{ .CreatedAt | dateFormat "yyyy-mm-dd hh:ii:ss" "mm/dd/yyyy" }}
```

### 4. Configuration File Processing

```go
package config

import (
    "github.com/maniartech/gotime/v2"
    "gopkg.in/yaml.v2"
)

type ConfigProcessor struct {
    dateFields map[string]string // field -> format mapping
}

func NewConfigProcessor() *ConfigProcessor {
    return &ConfigProcessor{
        dateFields: map[string]string{
            "start_date":    "yyyy-mm-dd",
            "end_date":      "yyyy-mm-dd",
            "created_at":    "yyyy-mm-dd hh:ii:ss",
            "schedule_time": "hh:ii",
        },
    }
}

func (cp *ConfigProcessor) NormalizeDates(config map[string]interface{}) error {
    for field, expectedFormat := range cp.dateFields {
        if value, exists := config[field]; exists {
            if str, ok := value.(string); ok {
                // Try to detect format and convert to expected
                if normalized, err := cp.detectAndConvert(str, expectedFormat); err == nil {
                    config[field] = normalized
                }
            }
        }
    }
    return nil
}

func (cp *ConfigProcessor) detectAndConvert(input, targetFormat string) (string, error) {
    formats := []string{
        "yyyy-mm-dd",
        "mm/dd/yyyy",
        "dd/mm/yyyy",
        "yyyy-mm-dd hh:ii:ss",
        "hh:ii",
        "h:ii aa",
    }

    for _, format := range formats {
        if result, err := gotime.Convert(input, format, targetFormat); err == nil {
            return result, nil
        }
    }

    return "", fmt.Errorf("unable to convert date: %s", input)
}
```

## Error Handling Best Practices

### 1. Graceful Degradation

```go
func safeConvert(value, from, to string) string {
    result, err := gotime.Convert(value, from, to)
    if err != nil {
        // Log error and return original value
        log.Printf("Date conversion failed: %v", err)
        return value
    }
    return result
}
```

### 2. Validation with Conversion

```go
func validateAndConvertDate(input string, allowedFormats []string, outputFormat string) (string, error) {
    for _, format := range allowedFormats {
        if result, err := gotime.Convert(input, format, outputFormat); err == nil {
            return result, nil
        }
    }
    return "", fmt.Errorf("date '%s' doesn't match any allowed format: %v", input, allowedFormats)
}
```

### 3. Batch Conversion with Error Reporting

```go
type ConversionResult struct {
    Original  string
    Converted string
    Error     error
}

func batchConvert(dates []string, fromFormat, toFormat string) []ConversionResult {
    results := make([]ConversionResult, len(dates))

    for i, date := range dates {
        converted, err := gotime.Convert(date, fromFormat, toFormat)
        results[i] = ConversionResult{
            Original:  date,
            Converted: converted,
            Error:     err,
        }
    }

    return results
}
```

## Performance Considerations

### 1. Format Caching

```go
// Cache frequently used format combinations
var conversionCache = make(map[string]func(string) (string, error))

func cachedConvert(value, from, to string) (string, error) {
    key := from + "->" + to

    if converter, exists := conversionCache[key]; exists {
        return converter(value)
    }

    // Create and cache converter function
    converter := func(v string) (string, error) {
        return gotime.Convert(v, from, to)
    }

    conversionCache[key] = converter
    return converter(value)
}
```

### 2. Batch Operations

```go
func convertDateColumn(dates []string, from, to string) ([]string, error) {
    results := make([]string, len(dates))

    for i, date := range dates {
        converted, err := gotime.Convert(date, from, to)
        if err != nil {
            return nil, fmt.Errorf("error converting date at index %d: %v", i, err)
        }
        results[i] = converted
    }

    return results, nil
}
```

## Testing Conversions

```go
func TestDateConversions(t *testing.T) {
    tests := []struct {
        input      string
        fromFormat string
        toFormat   string
        expected   string
    }{
        {"2025-07-07", "yyyy-mm-dd", "mm/dd/yyyy", "07/07/2025"},
        {"07/07/2025", "mm/dd/yyyy", "dd/mm/yyyy", "07/07/2025"},
        {"2025-07-07", "yyyy-mm-dd", "mmmm dt, yyyy", "July 7th, 2025"},
        {"14:30", "hhhh:ii", "h:ii aa", "2:30 PM"},
    }

    for _, test := range tests {
        result, err := gotime.Convert(test.input, test.fromFormat, test.toFormat)
        if err != nil {
            t.Errorf("Conversion failed: %v", err)
            continue
        }

        if result != test.expected {
            t.Errorf("Expected %s, got %s", test.expected, result)
        }
    }
}
```

---

## Summary

GoTime's Convert function provides:

- **Seamless format transformation** between any NITES formats
- **Automatic validation** during conversion process
- **Error handling** for invalid inputs or format mismatches
- **Performance optimization** through internal caching
- **Real-world utility** for data processing, APIs, and user interfaces

The Convert function is particularly powerful for:
- Data migration and ETL processes
- API response transformation
- User input normalization
- Internationalization and localization
- Template and report generation

---

Next: [Relative Time Functions](relative-time.md) | [Back to API Reference](../api-reference/)
