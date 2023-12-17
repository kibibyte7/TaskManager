package main

//Les importations
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"github.com/tealeg/xlsx"
)

//La structure de la tâche
type Task struct {
	Nom        string
	Heures     float64
	Facturable bool
	Type 	   string
}

//La liste des tâches
var listeTaches []Task

//Pour mettre des couleurs dans le terminal
const escapeCodeRed = "\x1b[38;5;196m"
const escapeCodeYellow = "\x1b[38;5;226m"
const escapeCodeGreen = "\x1b[38;5;82m"
const escapeCodeRose = "\x1b[38;5;205;1m"
const escapeCodeBlue = "\x1b[38;5;87m"
const escapeCodeViolet = "\x1b[38;5;99;1m"
const escapeCodeColorReset = "\033[0m"

//La fonction principale de mon programme
func main() {
	
	//L'équivalent d'un while
	for {

		//Lire l'entrée de l'utilisateur
		fmt.Printf("%vTaskManager%v %v➜%v ", escapeCodeViolet, escapeCodeColorReset , escapeCodeRose , escapeCodeColorReset)
		input := readLine()

		//On distingue des commandes rentré
		switch(strings.ToLower(input)){

			//Cette commande permet d'ajouter une tâche
			case "add task":

				//Ajoute la tâche
				addTask()

				break;

			//Cette commande permet de lister les tâches
			case "list tasks":

				//Liste les tâches
				listTask()

				break;

			//Cette commande permet de quitter l'application
			case "exit":

				//On dit au revoir à l'utilisateur
				fmt.Println("Au revoir !")

				//On ferme le programme
				os.Exit(0)

				break;

			case "export":

				fmt.Println("Le fichier est en cours d'exportation...")

				erreur := exportToXlsx()

				if erreur == nil {

					fmt.Printf("%vLe fichier a été exporté avec succès !%v\n",escapeCodeGreen,escapeCodeColorReset)

				} else {

					fmt.Printf("%vUne erreur est survenue lors de l'exportation !%v\n", escapeCodeRed, escapeCodeColorReset)

				}

			//Commande invalide
			default:

				//Dire à l'utilisateur que sa commande n'est pas valide
				fmt.Printf("%vcommande invalide: %s.%v\n", escapeCodeRed, input, escapeCodeColorReset)
		}

	}

}

//Permet de lire une chaine de caractère
func readLine() string {

	//On se fait un nouveau reader
	reader := bufio.NewReader(os.Stdin)

	//On lit la chaine de caractère
	line, _ := reader.ReadString('\n')

	// Capitaliser la première lettre du premier mot
	runes := []rune(line)
	runes[0] = unicode.ToUpper(runes[0])

	//Reprendre mon string mais avec la première lettre en majuscule
	line = string(runes)

	//On retourne le résultat
	return strings.TrimSpace(line)
}

//Permet de lire un nombre à virgule donner par l'utilisateur
func readFloat() float64 {

	//Tant que ce n'est pas un nombre valide
	for {

		//Je lis ma ligne
		line := readLine()

		//Je parse ce que j'ai lu en float64
		value, err := strconv.ParseFloat(line, 64)

		//S'il n'y a pas d'erreur
		if err == nil {

			//Je retourne ma valeur
			return value

		}

		//Sinon je dis à l'utilisateur d'entrer un nombre
		fmt.Printf("%sVeuillez entrer un nombre valide.%s\n",escapeCodeRed,escapeCodeColorReset)
	}
}

//Permet de lire le fait qu'une tâche soit facturable ou pas
func readBool() bool {

	//Tant que j'ai pas dis oui ou non
	for {

		//Je lis mon entrée
		line := strings.ToLower(readLine())

		//Si j'ai dit oui je retourne vrai
		if line == "oui" {

			//Retour
			return true

		//Si j'ai dit non, je retourne false
		} else if line == "non" {

			//Retour
			return false
		
		}
		
		//Je signale à l'utilisateur qu'il faut répondre par oui ou non
		fmt.Printf("%sVeuillez entrer 'oui' ou 'non'.%s\n", escapeCodeRed, escapeCodeColorReset)

	}
}

//Permet de le choix fait sur le type de la tâche
func readType() string {

	//Tant que je n'ai pas donné un nombre entre 1 et 5
	for {

		//On lit la chaine de caractère
		line := readLine()

		//Je parse ce que j'ai lu en int
		number, err := strconv.Atoi(line)

		//S'il n'y a pas d'erreur
		if err == nil {

			//On lit le choix qui est fait
			switch(number){

				//Normal
				case 1:
					return "normal"
				
				//Problème
				case 2:
					return "problème"
				
				//Solution
				case 3:
					return "solution"
				
				//Réunion
				case 4:
					return "réunion"
				
				//Divers
				case 5:
					return "divers"
				
				//Pas entre 1 et 5, donc pas valide
				default:
					fmt.Printf("%vEntrez un nombre entre 1 et 5.\n%v", escapeCodeRed, escapeCodeColorReset)
			}
		
		//S'il y a une erreur
		} else {

			//Le faire savoir à l'utilisateur
			fmt.Printf("%vCette entrée n'est pas un nombre valide.\n%v", escapeCodeRed, escapeCodeColorReset)

		}
	}

}

//Permet d'ajouter une tâche
func addTask() {

	//Bienvenue tout ce qu'il y a de plus classique
	fmt.Println("Bienvenue dans l'interface d'ajout d'une tâche")
	fmt.Println()

	//La tache
	var tache Task

	//On prend le nom de la tâche
	fmt.Println("Le nom de la tâche : ")
	tache.Nom = readLine()
	fmt.Println()

	//On prend le nombre d'heures passé dessus
	fmt.Println("Le nombre d'heures passées dessus : ")
	tache.Heures = readFloat()
	fmt.Println()

	//Est-ce que cette tâche est facturable
	fmt.Println("Cette tâche est-elle facturable ? ")
	tache.Facturable = readBool()
	fmt.Println()

	//On demande le type de tâche que c'est
	fmt.Println("Quel est le type de la tâche ?\n\n[1] Normal\n[2] Problème\n[3] Solution\n[4] Réunion\n[5] Divers")
	tache.Type = readType()
	fmt.Println()

	//On print un petit summary de la tâche
	fmt.Printf("Nom de la tâche: %s\nHeures: %.2f\nFacturable: %t\nType de la tâche: %v\n", tache.Nom, tache.Heures, tache.Facturable, tache.Type)

	//On ajoute la tâche à la liste globale de l'application
	listeTaches = append(listeTaches, tache)
}

//Permet de lister les tâches
func listTask(){

	//Si la liste est vide 
	if len(listeTaches) == 0 {

		//Dire à l'utilisateur que la liste dest vide
		fmt.Printf("%vLa liste des tâches est vide.%v\n", escapeCodeRed, escapeCodeColorReset)

	//Si la liste n'est pas vide
	} else {

		//Mettre un index à 1 directement
		i := 1

		//Ajout d'une ligne
		fmt.Println()

		//Le titre de la sortie attendue
		fmt.Println("Liste des tâches:")

		//Ajout d'une ligne
		fmt.Println()

		var heuresTotal float64

		heuresTotal = 0

		//Pour cahque tâches dans la liste
		for _, tache := range listeTaches {
			
			//On calcule les heures qui ont été faites
			heuresTotal += tache.Heures

			//On défini un nom vide pour l'instant, on doit le traiter
			nom := ""

			//On doit aussi taiter cette info, on la laise vide pour l'instant
			facturable := ""

			//On traite l'info de la facturation
			if tache.Facturable {

				//C'est facturable
				facturable = "Oui"

			} else {

				//Cette tâche n'est pas facturable
				facturable = "Non"

			}

			//On va ensuite regarder le type de la tâche
			switch tache.Type {

				//Dans une tâche qui est classique
				case "normal":

					//On met pas de couleur mais on rajoute le ":"
					nom = strings.Join([]string{tache.Nom, ":"}, "")
					break;
				
				//Si le type de la tâche est un problème
				case "problème":

					//On met le nom en rouge
					nom = strings.Join([]string{escapeCodeRed, tache.Nom, ":", escapeCodeColorReset}, "")
					break;
				
				//Si la tâche est une solution à un problème
				case "solution":

					//On met son texte en vert
					nom = strings.Join([]string{escapeCodeGreen, tache.Nom, ":", escapeCodeColorReset}, "")
					break;
				
				//Si cette tâche est une réunion avec les chefs de projets ou le client
				case "réunion":

					//On met le texte en bleu
					nom = strings.Join([]string{escapeCodeBlue, tache.Nom, ":", escapeCodeColorReset}, "")
					break;
				
				//Si cette tâche est une chose divers
				case "divers":

					//On met cette tâche en jaune
					nom = strings.Join([]string{escapeCodeYellow, tache.Nom, ":", escapeCodeColorReset}, "")
					break;

			}

			//Donner les infos de la tâche à l'utilisateur
			fmt.Printf("[%v] %v\nHeures: %v\nFacturable: %v\nType de tâche: %v\n\n", i, nom , tache.Heures, facturable, tache.Type)
			
			//On incrémente l'index
			i++
		}

		//On gère le pluiriel ici, pas d'opérateur ternaires car pas supporté en go
		heuresText := ""

		//On le gère en une ligne de if else
		if heuresTotal == 0 || heuresTotal == 1 {heuresText = "heure"} else {heuresText = "heures"}

		//On affiche enfin les heures total de l'utilisateur
		fmt.Printf("Heures totales de ce jour: %v %v\n", heuresTotal, heuresText)

	}
}

func exportToXlsx() error {
	problemID := 1
	solutionID := 1
	fichier := xlsx.NewFile()
	feuille, err := fichier.AddSheet("tasks")

	if err != nil {
		return err
	}

	// En-têtes XLSX
	titres := feuille.AddRow()
	titres.AddCell().Value = "Nom"
	titres.AddCell().Value = "Heures facturable"
	titres.AddCell().Value = "Heures non-facturable"

	//Style des problèmes
	styleProbleme := xlsx.NewStyle()
	styleProbleme.Fill = *xlsx.NewFill("solid", "#000000", "#FF0000")
	styleProbleme.Font.Bold = true
	styleProbleme.ApplyFill = true
	styleProbleme.ApplyFont = true

	//Style des solutions
	styleSolution := xlsx.NewStyle()
	styleSolution.Fill = *xlsx.NewFill("solid", "#23B84B","#000000")
	styleSolution.Font.Bold = true
	styleSolution.ApplyFill = true
	styleSolution.ApplyFont = true

	//Style des réunions
	styleReunion := xlsx.NewStyle()
	styleReunion.Fill = *xlsx.NewFill("solid", "#000000", "#00D5FF")
	styleReunion.Font.Bold = true
	styleReunion.ApplyFill = true
	styleReunion.ApplyFont = true

	//Style des choses diverses
	styleDivers := xlsx.NewStyle()
	styleDivers.Fill = *xlsx.NewFill("solid", "#000000", "#D4FC21")
	styleDivers.Font.Bold = true
	styleDivers.ApplyFill = true
	styleDivers.ApplyFont = true
	
	for _, tache := range listeTaches {

		ligne := feuille.AddRow()

		var cell *xlsx.Cell

		switch tache.Type {
			
			case "normal":

				cell = ligne.AddCell()
				cell.Value = tache.Nom
				break;

			//Si le type de la tâche est un problème
			case "problème":

				//On met le nom en rouge
				cell = ligne.AddCell()
				cell.Value = strings.Join([]string{"Problème #", fmt.Sprintf("%v",problemID), ": ", tache.Nom}, "")

				cell.SetStyle(styleProbleme)

				problemID++

				break;
			
			//Si la tâche est une solution à un problème
			case "solution":

				//On met son texte en vert
				cell = ligne.AddCell()
				cell.Value = strings.Join([]string{"Solution au problème #", fmt.Sprintf("%v",solutionID), ": ", tache.Nom}, "")

				cell.SetStyle(styleSolution)

				solutionID++
				
				break;
			
			//Si cette tâche est une réunion avec les chefs de projets ou le client
			case "réunion":

				//On met le texte en bleu
				cell = ligne.AddCell()
				cell.Value = tache.Nom

				cell.SetStyle(styleReunion)
				
				break;
			
			//Si cette tâche est une chose divers
			case "divers":

				//On met cette tâche en jaune
				cell = ligne.AddCell()
				cell.Value = tache.Nom

				cell.SetStyle(styleDivers)
				
				break;

		}


		if tache.Facturable {

			ligne.AddCell().Value = fmt.Sprintf("%.2f", tache.Heures)
			ligne.AddCell().Value = fmt.Sprintf("%v", 0)

		} else {

			ligne.AddCell().Value = fmt.Sprintf("%v", 0)
			ligne.AddCell().Value = fmt.Sprintf("%.2f", tache.Heures)

		}

	}

	return fichier.Save("./tasks.xlsx")
}