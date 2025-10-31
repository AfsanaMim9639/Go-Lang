package main

import (
    "fmt"
    "time"
)

func main() {
    // Current time
    now := time.Now()
    fmt.Println("Current time (default):", now)

    // 1. Full datetime formatting
    fmt.Println("Full datetime:", now.Format("2006-01-02 15:04:05"))

    // 2. Date only
    fmt.Println("Date only:", now.Format("2006-01-02"))

    // 3. Time only (24-hour)
    fmt.Println("Time only (24h):", now.Format("15:04:05"))

    // 4. Time only (12-hour)
    fmt.Println("Time only (12h):", now.Format("03:04:05 PM"))

    // 5. Month names
    fmt.Println("Month (short):", now.Format("Jan"))
    fmt.Println("Month (full):", now.Format("January"))

    // 6. Weekday names
    fmt.Println("Weekday (short):", now.Format("Mon"))
    fmt.Println("Weekday (full):", now.Format("Monday"))

    // 7. Time with timezone
    fmt.Println("With timezone:", now.Format("2006-01-02 15:04:05 MST"))
    fmt.Println("With numeric offset:", now.Format("2006-01-02 15:04:05 -0700"))

    // -------------------------
    // Parsing examples
    // -------------------------

    // 8. Parse full datetime
    dtStr := "2025-10-30 14:45:00"
    dt, err := time.Parse("2006-01-02 15:04:05", dtStr)
    if err != nil {
        fmt.Println("Error parsing datetime:", err)
    } else {
        fmt.Println("Parsed datetime:", dt)
    }

    // 9. Parse date only
    dateStr := "2025-10-30"
    date, _ := time.Parse("2006-01-02", dateStr)
    fmt.Println("Parsed date only:", date)

    // 10. Parse time only (assumes today’s date)
    timeStr := "14:30:00"
    t, _ := time.Parse("15:04:05", timeStr)
    fmt.Println("Parsed time only:", t.Format("15:04:05"))

    // 11. Parse datetime with timezone
    tzStr := "2025-10-30 14:45:00 +0600"
    tzTime, _ := time.Parse("2006-01-02 15:04:05 -0700", tzStr)
    fmt.Println("Parsed with timezone:", tzTime)
}
