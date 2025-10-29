package main

import (
    "fmt"
    "time"
)

func main() {
    // 1️⃣ বর্তমান সময়
    now := time.Now()
    fmt.Println("বর্তমান সময়:", now)

    // 2️⃣ আলাদা অংশগুলো পাওয়া
    fmt.Println("Year:", now.Year())
    fmt.Println("Month:", now.Month())
    fmt.Println("Day:", now.Day())
    fmt.Println("Hour:", now.Hour())
    fmt.Println("Minute:", now.Minute())
    fmt.Println("Second:", now.Second())

    // 3️⃣ নির্দিষ্ট সময় তৈরি করা
    specific := time.Date(2025, time.October, 28, 10, 30, 0, 0, time.Local)
    fmt.Println("\nনির্দিষ্ট সময়:", specific)

    // 4️⃣ সময় ফরম্যাট করা (format)
    fmt.Println("\n📅 Format Example:")
    fmt.Println("Default:", now.String())
    fmt.Println("Custom:", now.Format("02-Jan-2006 15:04:05"))
    fmt.Println("Only Time:", now.Format("03:04:05 PM"))
    fmt.Println("Only Date:", now.Format("Monday, 02 Jan 2006"))

    // 5️⃣ সময় পার্থক্য বের করা (Duration)
    duration := now.Sub(specific)
    fmt.Println("\nসময় পার্থক্য:", duration)
    fmt.Println("ঘণ্টায়:", duration.Hours())
    fmt.Println("মিনিটে:", duration.Minutes())

    // 6️⃣ ভবিষ্যতের সময় পাওয়া (Add)
    future := now.Add(48 * time.Hour)
    fmt.Println("\n২ দিন পরের সময়:", future)

    // 7️⃣ সময়ের তুলনা (Before / After / Equal)
    fmt.Println("\nComparison:")
    fmt.Println("specific আগে?", specific.Before(now))
    fmt.Println("specific পরে?", specific.After(now))
    fmt.Println("specific সমান?", specific.Equal(now))

    // 8️⃣ সময় পার্স করা (string → time)
    layout := "02-01-2006 15:04"
    parsed, _ := time.Parse(layout, "28-10-2025 09:30")
    fmt.Println("\nParsed Time:", parsed)

    // 9️⃣ টাইম জোন দেখা
    fmt.Println("\nTime Zone:", now.Location())
    utc := now.UTC()
    fmt.Println("UTC Time:", utc)

    // 🔟 Sleep (Delay)
    fmt.Println("\n3 সেকেন্ড অপেক্ষা করো...")
    time.Sleep(3 * time.Second)
    fmt.Println("অপেক্ষা শেষ ✅")
}
