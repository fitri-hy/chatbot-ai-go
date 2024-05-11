package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
)

type Response struct {
	Text string `json:"text"`
}

func callAPI(question string) (string, error) {
	apiURL := fmt.Sprintf("https://api.hy-tech.my.id/api/gemini/%s", question)
	
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Text, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Chatbot AI Golang</title>
			<script src="https://cdn.tailwindcss.com"></script>
		</head>
		<body class="bg-slate-100 max-w-4xl m-auto pt-28 px-4">
			<nav class="fixed inset-x-0 top-0 z-10 w-full px-4 py-1 bg-white shadow-md border-slate-500 transition duration-700 ease-out">
				<div class="flex justify-between p-4">
					<div class="flex gap-2 items-center text-[2rem] leading-[3rem] tracking-tight font-bold text-black">
						<img class="w-10 h-10" src="https://hy-tech.my.id/images/logo.png">
						<h2 class="hidden lg:flex"><span class="text-indigo-600">Ask</span>AI</h2>
					</div>
					<div class="flex items-center space-x-4 text-lg font-semibold tracking-tight">
						<a href="https://github.com/fitri-hy" class="flex gap-2 items-center px-6 py-2 text-black transition duration-700 ease-out bg-white border border-black rounded-lg hover:bg-gray-200 hover:border">
							<svg width="20" height="20" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
							  <image xlink:href="http://hytech-icons.vercel.app/icons/pro/brands/github.svg" width="24" height="24"/>
							</svg>
							Github
						</a>
						<a href="https://hy-tech.my.id/" class="flex gap-2 items-center px-6 py-2 text-white transition duration-700 ease-out bg-indigo-600 rounded-lg hover:bg-indigo-500">
							<svg width="20" height="20" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
							  <image xlink:href="http://hytech-icons.vercel.app/icons/pro/sharp-solid/globe.svg?color=ffffff" width="24" height="24"/>
							</svg>
							Website
						</a>
					</div>
				</div>
			</nav>
			<div class="bg-white p-4 rounded-lg shadow">
				<form action="/" method="post" class="flex flex-col gap-4">
					<label class="text-3xl font-bold text-gray-700" for="question">Enter Question</label>
					<textarea class="p-2 rounded-md bg-gray-100 border" type="text" rows="4" id="question" name="question"></textarea>
					<button class="bg-indigo-600 w-[100px] text-white rounded-lg shadow px-4 py-2" type="submit">Ask AI</button>
				</form>
				<div id="result" class="py-6">
					{{if .Answer}}
						<p>{{.Answer}}</p>
					{{end}}
				</div>
			</div>
			<footer class="bg-white mt-20 mb-4 rounded-lg shadow">
				<div class="container px-6 py-8 mx-auto">
					<div class="flex flex-col items-center sm:flex-row sm:justify-center">
						<p class="text-sm text-gray-500">©2024 <a href="https://hy-tech.my.id/" target="_blank" class="text-blue-500">Hy-Tech Group</a>. All Rights Reserved.</p>
					</div>
				</div>
			</footer>
		</body>
		</html>
		`))
		tmpl.Execute(w, map[string]interface{}{"Answer": ""})
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	question := r.Form.Get("question")
	if question == "" {
		http.Error(w, "Please provide a question", http.StatusBadRequest)
		return
	}

	answer, err := callAPI(question)
	if err != nil {
		http.Error(w, "Failed to get response from API", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Chatbot AI Golang</title>
			<script src="https://cdn.tailwindcss.com"></script>
		</head>
		<body class="bg-slate-100 max-w-4xl m-auto pt-28 px-4">
			<nav class="fixed inset-x-0 top-0 z-10 w-full px-4 py-1 bg-white shadow-md border-slate-500 transition duration-700 ease-out">
				<div class="flex justify-between p-4">
					<div class="flex gap-2 items-center text-[2rem] leading-[3rem] tracking-tight font-bold text-black">
						<img class="w-10 h-10" src="https://hy-tech.my.id/images/logo.png">
						<h2 class="hidden lg:flex"><span class="text-indigo-600">Ask</span>AI</h2>
					</div>
					<div class="flex items-center space-x-4 text-lg font-semibold tracking-tight">
						<a href="https://github.com/fitri-hy" class="flex gap-2 items-center px-6 py-2 text-black transition duration-700 ease-out bg-white border border-black rounded-lg hover:bg-gray-200 hover:border">
							<svg width="20" height="20" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
							  <image xlink:href="http://hytech-icons.vercel.app/icons/pro/brands/github.svg" width="24" height="24"/>
							</svg>
							Github
						</a>
						<a href="https://hy-tech.my.id/" class="flex gap-2 items-center px-6 py-2 text-white transition duration-700 ease-out bg-indigo-600 rounded-lg hover:bg-indigo-500">
							<svg width="20" height="20" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
							  <image xlink:href="http://hytech-icons.vercel.app/icons/pro/sharp-solid/globe.svg?color=ffffff" width="24" height="24"/>
							</svg>
							Website
						</a>
					</div>
				</div>
			</nav>
			<div class="bg-white p-4 rounded-lg shadow">
				<form action="/" method="post" class="flex flex-col gap-4">
					<label class="text-3xl font-bold text-gray-700" for="question">Enter Question</label>
					<textarea class="p-2 rounded-md bg-gray-100 border" type="text" rows="4" id="question" name="question"></textarea>
					<button class="bg-indigo-600 w-[100px] text-white rounded-lg shadow px-4 py-2" type="submit">Ask AI</button>
				</form>
				<div id="result" class="py-6">
					{{if .Answer}}
						<p>{{.Answer}}</p>
					{{end}}
				</div>
			</div>
			<footer class="bg-white mt-20 mb-4 rounded-lg shadow">
				<div class="container px-6 py-8 mx-auto">
					<div class="flex flex-col items-center sm:flex-row sm:justify-center">
						<p class="text-sm text-gray-500">©2024 <a href="https://hy-tech.my.id/" target="_blank" class="text-blue-500">Hy-Tech Group</a>. All Rights Reserved.</p>
					</div>
				</div>
			</footer>
		</body>
		</html>
		`))

	tmpl.Execute(w, map[string]interface{}{"Answer": answer})
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
