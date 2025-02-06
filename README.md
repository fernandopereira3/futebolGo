# Football Tournament Simulator

## 🎯 Overview
A football tournament simulator developed in Go that simulates matches between 4 international teams: Argentina, Germany, Japan, and Brazil. The program runs multiple rounds where each team faces all others, accumulating points based on match results.

## 🎮 How It Works

### Scoring System
- Win: 3 points
- Draw: 1.5 points
- Loss: 0 points

### Tournament Structure
- 4 participating teams
- Each team plays against all other teams
- 10 rounds of matches
- Goals are randomly generated (0-5 goals per team)

### Participating Teams
1. 🇦🇷 Argentina
2. 🇩🇪 Germany
3. 🇯🇵 Japan
4. 🇧🇷 Brazil

## 🛠️ Technologies Used
- Language: Go
- Packages:
  - `fmt`: For data output
  - `math/rand`: For random goal generation

## 📊 Program Logic
1. Program defines constants for team names
2. Global variables store each team's score
3. The `rodadas()` function simulates all matches:
   - Generates random goals for each team
   - Calculates scores based on results
   - Executes 10 complete rounds
4. Finally returns team names and their respective scores

## 🚀 How to Run
1. Make sure you have Go installed on your system
2. Clone the repository
3. Run the command:
```
go run futebolGo.go
```

## 🔍 Output
The program will display results in the format:
```
[Team 1 Name] [Score 1] [Team 2 Name] [Score 2] [Team 3 Name] [Score 3] [Team 4 Name] [Score 4]
```

## 🤝 Contributions
Contributions are welcome! Feel free to propose improvements through pull requests.

## ⚠️ Notes
- Results are randomly generated each time the program runs
- The program simulates a total of 10 complete rounds
- Each team faces all other teams in each round

| :placard: Showcase |     |
| -------------  | --- |
| :sparkles: Developer | Fernando Pereira
| :label: Technologies | Golang

## Project Details
A small program in Go created to simulate football matches with a unique scoring system: Victory (3 points), Draw (1.5 points), and Defeat (0 points). Created for entertainment and learning purposes!