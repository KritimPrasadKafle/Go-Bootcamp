// package main

// func main(){
// 	layout := "2006-01-02T15:04:05Z07:00"
// str := "2024-07-04T14:30:18Z"

// t, err := time.Parse(layout, str)
// if err != nil {
//     fmt.Println("Error parsing time:", err)
//     return
// }
// fmt.Println(t)  // 2024-07-04 14:30:18 +0000 UTC

// // Example 2: Custom format with 12-hour clock
// str1 := "Jul 03, 2024 03:18 PM"
// layout1 := "Jan 02, 2006 03:04 PM"

// t1, err := time.Parse(layout1, str1)
// fmt.Println(t1)  // 2024-07-03 15:18:00 +0000 UTC

// }

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     t := time.Date(2024, time.July, 4, 14, 30, 18, 123456789, time.UTC)

//     fmt.Println("=== DATE COMPONENTS ===")
//     fmt.Println()

//     dateFormats := []struct {
//         layout string
//         desc   string
//     }{
//         // Year
//         {"2006", "Year (4 digit)"},
//         {"06", "Year (2 digit)"},

//         // Month
//         {"01", "Month (padded)"},
//         {"1", "Month (no pad)"},
//         {"Jan", "Month (abbrev)"},
//         {"January", "Month (full)"},

//         // Day
//         {"02", "Day (padded)"},
//         {"2", "Day (no pad)"},
//         {"002", "Day of year"},

//         // Weekday
//         {"Mon", "Weekday (abbrev)"},
//         {"Monday", "Weekday (full)"},
//     }

//     for _, f := range dateFormats {
//         fmt.Printf("%-20s %-15s → %s\n", f.desc, f.layout, t.Format(f.layout))
//     }

//     fmt.Println("\n=== TIME COMPONENTS ===")
//     fmt.Println()

//     timeFormats := []struct {
//         layout string
//         desc   string
//     }{
//         // Hour
//         {"15", "Hour 24h (padded)"},
//         {"03", "Hour 12h (padded)"},
//         {"3", "Hour 12h (no pad)"},

//         // Minute
//         {"04", "Minute (padded)"},
//         {"4", "Minute (no pad)"},

//         // Second
//         {"05", "Second (padded)"},
//         {"5", "Second (no pad)"},

//         // AM/PM
//         {"PM", "AM/PM (upper)"},
//         {"pm", "AM/PM (lower)"},
//     }

//     for _, f := range timeFormats {
//         fmt.Printf("%-20s %-15s → %s\n", f.desc, f.layout, t.Format(f.layout))
//     }

//     fmt.Println("\n=== FRACTIONAL SECONDS ===")
//     fmt.Println()

//     fracFormats := []struct {
//         layout string
//         desc   string
//     }{
//         {"15:04:05.000", "Milliseconds"},
//         {"15:04:05.000000", "Microseconds"},
//         {"15:04:05.000000000", "Nanoseconds"},
//         {"15:04:05.999999999", "Trailing zeros removed"},
//     }

//     for _, f := range fracFormats {
//         fmt.Printf("%-25s %-25s → %s\n", f.desc, f.layout, t.Format(f.layout))
//     }

//     fmt.Println("\n=== TIMEZONE COMPONENTS ===")
//     fmt.Println()

//     // Use a time with non-UTC timezone for better examples
//     loc, _ := time.LoadLocation("Asia/Kolkata")
//     tLocal := t.In(loc)

//     tzFormats := []struct {
//         layout string
//         desc   string
//     }{
//         {"MST", "Timezone name"},
//         {"-0700", "Offset ±hhmm"},
//         {"-07:00", "Offset ±hh:mm"},
//         {"-07", "Offset ±hh"},
//         {"Z0700", "Z for UTC, else offset"},
//         {"Z07:00", "Z for UTC, else ±hh:mm"},
//     }

//     fmt.Println("UTC time:")
//     for _, f := range tzFormats {
//         fmt.Printf("  %-25s %-15s → %s\n", f.desc, f.layout, t.Format(f.layout))
//     }

//     fmt.Println("\nLocal time (IST):")
//     for _, f := range tzFormats {
//         fmt.Printf("  %-25s %-15s → %s\n", f.desc, f.layout, tLocal.Format(f.layout))
//     }
// }

// func main() {
//     t := time.Now()

//     fmt.Println("=== COMMON DATE FORMATS ===")
//     fmt.Println()

//     dateLayouts := []struct {
//         name   string
//         layout string
//     }{
//         {"YYYY-MM-DD", "2006-01-02"},
//         {"DD/MM/YYYY", "02/01/2006"},
//         {"MM/DD/YYYY", "01/02/2006"},
//         {"DD-MM-YYYY", "02-01-2006"},
//         {"Month Day, Year", "January 2, 2006"},
//         {"Mon Day, Year", "Jan 2, 2006"},
//         {"Day Month Year", "2 January 2006"},
//         {"Weekday, Month Day, Year", "Monday, January 2, 2006"},
//         {"Short Date", "1/2/06"},
//     }

//     for _, l := range dateLayouts {
//         fmt.Printf("%-30s → %s\n", l.name, t.Format(l.layout))
//     }

//     fmt.Println("\n=== COMMON TIME FORMATS ===")
//     fmt.Println()

//     timeLayouts := []struct {
//         name   string
//         layout string
//     }{
//         {"24-hour HH:MM", "15:04"},
//         {"24-hour HH:MM:SS", "15:04:05"},
//         {"12-hour H:MM AM/PM", "3:04 PM"},
//         {"12-hour HH:MM:SS AM/PM", "03:04:05 PM"},
//         {"With milliseconds", "15:04:05.000"},
//         {"With timezone", "15:04:05 MST"},
//         {"With offset", "15:04:05 -07:00"},
//     }

//     for _, l := range timeLayouts {
//         fmt.Printf("%-30s → %s\n", l.name, t.Format(l.layout))
//     }

//     fmt.Println("\n=== COMBINED DATE + TIME ===")
//     fmt.Println()

//     combinedLayouts := []struct {
//         name   string
//         layout string
//     }{
//         {"ISO 8601 / RFC3339", "2006-01-02T15:04:05Z07:00"},
//         {"SQL datetime", "2006-01-02 15:04:05"},
//         {"Log format", "2006/01/02 15:04:05"},
//         {"Readable", "Jan 2, 2006 at 3:04 PM"},
//         {"Full", "Monday, January 2, 2006 3:04:05 PM MST"},
//         {"Compact", "20060102150405"},
//         {"File timestamp", "2006-01-02_15-04-05"},
//     }

//     for _, l := range combinedLayouts {
//         fmt.Printf("%-30s → %s\n", l.name, t.Format(l.layout))
//     }
// }

// func main() {
//     t := time.Now()

//     fmt.Println("=== GO'S PREDEFINED LAYOUTS ===")
//     fmt.Println()

//     predefined := []struct {
//         name   string
//         layout string
//     }{
//         {"time.Layout", time.Layout},
//         {"time.ANSIC", time.ANSIC},
//         {"time.UnixDate", time.UnixDate},
//         {"time.RubyDate", time.RubyDate},
//         {"time.RFC822", time.RFC822},
//         {"time.RFC822Z", time.RFC822Z},
//         {"time.RFC850", time.RFC850},
//         {"time.RFC1123", time.RFC1123},
//         {"time.RFC1123Z", time.RFC1123Z},
//         {"time.RFC3339", time.RFC3339},
//         {"time.RFC3339Nano", time.RFC3339Nano},
//         {"time.Kitchen", time.Kitchen},
//         {"time.Stamp", time.Stamp},
//         {"time.StampMilli", time.StampMilli},
//         {"time.StampMicro", time.StampMicro},
//         {"time.StampNano", time.StampNano},
//         {"time.DateTime", time.DateTime},
//         {"time.DateOnly", time.DateOnly},
//         {"time.TimeOnly", time.TimeOnly},
//     }

//     for _, p := range predefined {
//         fmt.Printf("%-22s %-35s → %s\n", p.name, p.layout, t.Format(p.layout))
//     }
// }
// ```

// **Output:**
// ```
// === GO'S PREDEFINED LAYOUTS ===

// time.Layout            01/02 03:04:05PM '06 -0700          → 07/30 02:30:45PM '24 +0530
// time.ANSIC             Mon Jan _2 15:04:05 2006            → Tue Jul 30 14:30:45 2024
// time.UnixDate          Mon Jan _2 15:04:05 MST 2006        → Tue Jul 30 14:30:45 IST 2024
// time.RubyDate          Mon Jan 02 15:04:05 -0700 2006      → Tue Jul 30 14:30:45 +0530 2024
// time.RFC822            02 Jan 06 15:04 MST                 → 30 Jul 24 14:30 IST
// time.RFC822Z           02 Jan 06 15:04 -0700               → 30 Jul 24 14:30 +0530
// time.RFC850            Monday, 02-Jan-06 15:04:05 MST      → Tuesday, 30-Jul-24 14:30:45 IST
// time.RFC1123           Mon, 02 Jan 2006 15:04:05 MST       → Tue, 30 Jul 2024 14:30:45 IST
// time.RFC1123Z          Mon, 02 Jan 2006 15:04:05 -0700     → Tue, 30 Jul 2024 14:30:45 +0530
// time.RFC3339           2006-01-02T15:04:05Z07:00           → 2024-07-30T14:30:45+05:30
// time.RFC3339Nano       2006-01-02T15:04:05.999999999Z07:00 → 2024-07-30T14:30:45.123456789+05:30
// time.Kitchen           3:04PM                              → 2:30PM
// time.Stamp             Jan _2 15:04:05                     → Jul 30 14:30:45
// time.StampMilli        Jan _2 15:04:05.000                 → Jul 30 14:30:45.123
// time.StampMicro        Jan _2 15:04:05.000000              → Jul 30 14:30:45.123456
// time.StampNano         Jan _2 15:04:05.000000000           → Jul 30 14:30:45.123456789
// time.DateTime          2006-01-02 15:04:05                 → 2024-07-30 14:30:45
// time.DateOnly          2006-01-02                          → 2024-07-30
// time.TimeOnly          15:04:05                            → 14:30:45

// package main

// import (
//     "fmt"
//     "time"
// )

// // FormatRelative formats time relative to now
// func FormatRelative(t time.Time) string {
//     now := time.Now()
//     diff := now.Sub(t)

//     switch {
//     case diff < time.Minute:
//         return "just now"
//     case diff < time.Hour:
//         mins := int(diff.Minutes())
//         if mins == 1 {
//             return "1 minute ago"
//         }
//         return fmt.Sprintf("%d minutes ago", mins)
//     case diff < 24*time.Hour:
//         hours := int(diff.Hours())
//         if hours == 1 {
//             return "1 hour ago"
//         }
//         return fmt.Sprintf("%d hours ago", hours)
//     case diff < 48*time.Hour:
//         return "yesterday"
//     case diff < 7*24*time.Hour:
//         return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
//     case diff < 30*24*time.Hour:
//         weeks := int(diff.Hours() / 24 / 7)
//         if weeks == 1 {
//             return "1 week ago"
//         }
//         return fmt.Sprintf("%d weeks ago", weeks)
//     case diff < 365*24*time.Hour:
//         months := int(diff.Hours() / 24 / 30)
//         if months == 1 {
//             return "1 month ago"
//         }
//         return fmt.Sprintf("%d months ago", months)
//     default:
//         years := int(diff.Hours() / 24 / 365)
//         if years == 1 {
//             return "1 year ago"
//         }
//         return fmt.Sprintf("%d years ago", years)
//     }
// }

// // FormatSmart chooses format based on how old the time is
// func FormatSmart(t time.Time) string {
//     now := time.Now()
//     diff := now.Sub(t)

//     switch {
//     case diff < 24*time.Hour && t.Day() == now.Day():
//         return t.Format("3:04 PM")  // Today: just time
//     case diff < 48*time.Hour:
//         return "Yesterday " + t.Format("3:04 PM")
//     case diff < 7*24*time.Hour:
//         return t.Format("Monday 3:04 PM")  // This week: day name
//     case t.Year() == now.Year():
//         return t.Format("Jan 2")  // This year: month and day
//     default:
//         return t.Format("Jan 2, 2006")  // Older: full date
//     }
// }

// // FormatOrdinal formats day with ordinal suffix (1st, 2nd, 3rd, etc.)
// func FormatOrdinal(t time.Time) string {
//     day := t.Day()
//     suffix := "th"

//     switch day {
//     case 1, 21, 31:
//         suffix = "st"
//     case 2, 22:
//         suffix = "nd"
//     case 3, 23:
//         suffix = "rd"
//     }

//     return fmt.Sprintf("%s %d%s, %d", t.Format("January"), day, suffix, t.Year())
// }

// func main() {
//     fmt.Println("=== RELATIVE TIME ===")
//     times := []time.Time{
//         time.Now().Add(-30 * time.Second),
//         time.Now().Add(-5 * time.Minute),
//         time.Now().Add(-2 * time.Hour),
//         time.Now().Add(-36 * time.Hour),
//         time.Now().Add(-5 * 24 * time.Hour),
//         time.Now().Add(-2 * 7 * 24 * time.Hour),
//         time.Now().Add(-3 * 30 * 24 * time.Hour),
//         time.Now().Add(-2 * 365 * 24 * time.Hour),
//     }

//     for _, t := range times {
//         fmt.Printf("%-30s → %s\n", t.Format("2006-01-02 15:04:05"), FormatRelative(t))
//     }

//     fmt.Println("\n=== SMART FORMAT ===")
//     for _, t := range times {
//         fmt.Printf("%-30s → %s\n", t.Format("2006-01-02 15:04:05"), FormatSmart(t))
//     }

//     fmt.Println("\n=== ORDINAL FORMAT ===")
//     dates := []time.Time{
//         time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC),
//         time.Date(2024, 7, 2, 0, 0, 0, 0, time.UTC),
//         time.Date(2024, 7, 3, 0, 0, 0, 0, time.UTC),
//         time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC),
//         time.Date(2024, 7, 21, 0, 0, 0, 0, time.UTC),
//         time.Date(2024, 7, 22, 0, 0, 0, 0, time.UTC),
//         time.Date(2024, 7, 23, 0, 0, 0, 0, time.UTC),
//     }

//     for _, d := range dates {
//         fmt.Printf("%s → %s\n", d.Format("2006-01-02"), FormatOrdinal(d))
//     }
// }