package main

import (
    "os"
    "strings"
    "text/template"
)

// 🧠 Custom function (FuncMap)
func toUpper(s string) string {
    return strings.ToUpper(s)
}

// 🧱 Structs
type Subject struct {
    Name   string
    Marks  int
}

type Student struct {
    Name     string
    Passed   bool
    Subjects []Subject
    Address  struct {
        City, Country string
    }
}

func main() {
    // 🧾 Data তৈরি
    student := Student{
        Name:   "Afsana Mim",
        Passed: true,
        Subjects: []Subject{
            {"Math", 95},
            {"English", 87},
            {"Science", 92},
        },
    }
    student.Address.City = "Dhaka"
    student.Address.Country = "Bangladesh"

    // 🧰 Template টেক্সট
    const report = `
📘 STUDENT REPORT
=====================
নাম: {{.Name | upper}}
ঠিকানা: {{with .Address}}{{.City}}, {{.Country}}{{end}}

{{if .Passed}}
অভিনন্দন! তুমি পাশ করেছো! 🎉
{{else}}
দুঃখিত, তুমি ফেল করেছো 😢
{{end}}

বিষয় অনুযায়ী নম্বরসমূহ:
{{range .Subjects}}
- {{.Name}}: {{.Marks}}
{{end}}

{{if gt (len .Subjects) 2}}
তুমি ২টির বেশি বিষয় পড়ছো! 👏
{{end}}

=====================
`

    // 🧩 Custom function map যোগ করা
    funcMap := template.FuncMap{
        "upper": toUpper,
    }

    // 🛠️ Template তৈরি ও execute
    t := template.Must(template.New("report").Funcs(funcMap).Parse(report))
    t.Execute(os.Stdout, student)
}
