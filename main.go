package main

//Les importations
import (
	"bufio"
	"errors"
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

			case "delete task":

				deleteElement()

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

	//Si la liste n'est pas vide
	if len(listeTaches) != 0 {

		//On prend l'ID du problème
		problemID := 1

		//On prend l'ID de la solution
		solutionID := 1

		//On crée le fichier
		fichier := xlsx.NewFile()

		//On ajoute une feuille pour mettre nos tâches dedans
		feuille, err := fichier.AddSheet("tasks")

		//S'il y a une erreur
		if err != nil {

			//Retourne l'erreur
			return err

		}

		// Je crée les entêtes de ce que je dois avoir comme info pour mes tâches
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
		
		//Je parcours ma liste de tâches
		for _, tache := range listeTaches {

			//J'ajoute une ligne
			ligne := feuille.AddRow()

			//Je déclare une variable de cellule avec un pointeur
			var cell *xlsx.Cell

			//Je traite le type de la tâche
			switch tache.Type {
				
				//Dans un cas classique
				case "normal":

					//On met juste le nom normal
					cell = ligne.AddCell()
					cell.Value = tache.Nom
					break;

				//Si le type de la tâche est un problème
				case "problème":

					//On met le nom en rouge
					cell = ligne.AddCell()

					//Je formate la tâche en "Poblème #numéro: nom de la tâche"
					cell.Value = strings.Join([]string{"Problème #", fmt.Sprintf("%v",problemID), ": ", tache.Nom}, "")

					//Je met le style que la cellule doit avoit (Gras + rouge)
					cell.SetStyle(styleProbleme)

					//J'incrémente l'ID du problème
					problemID++

					break;
				
				//Si la tâche est une solution à un problème
				case "solution":

					//On met le nom en vert
					cell = ligne.AddCell()

					//Je formate la tâche en "Solution pour le problème #numéro: nim de la tâche"
					cell.Value = strings.Join([]string{"Solution au problème #", fmt.Sprintf("%v",solutionID), ": ", tache.Nom}, "")

					//Je mets le style que la cellule doit avoir (Gras + vert)
					cell.SetStyle(styleSolution)

					//J'incrémente l'ID de la solution
					solutionID++
					
					break;
				
				//Si cette tâche est une réunion avec les chefs de projets ou le client
				case "réunion":

					//On met le texte en bleu
					cell = ligne.AddCell()

					//Je mets le nom de la tâche
					cell.Value = tache.Nom

					//Je mets le style Gras et bleu sur la cellule
					cell.SetStyle(styleReunion)
					
					break;
				
				//Si cette tâche est une chose divers
				case "divers":

					//On met cette tâche en jaune
					cell = ligne.AddCell()

					//Je mets le nom de la tâche
					cell.Value = tache.Nom

					//Je mets le style Gras et Jaune sur la cellule
					cell.SetStyle(styleDivers)
					
					break;

			}

			//Si la tâche est facturable
			if tache.Facturable {

				//J'ajoute d'abord l'heure de la tâche
				ligne.AddCell().Value = fmt.Sprintf("%.2f", tache.Heures)

				//Enfin j'ajoute l'heure non-facturable, qui est 0 dans ce cas-ci
				ligne.AddCell().Value = fmt.Sprintf("%v", 0)

			//Si la tâche n'est pas facturable
			} else {

				//J'ajoute un 0 dans l'heure facturable
				ligne.AddCell().Value = fmt.Sprintf("%v", 0)

				//Je mets l'heure non-facturable de la tâche
				ligne.AddCell().Value = fmt.Sprintf("%.2f", tache.Heures)

			}

		}

		//Pour finir je sauve le fichier et je gère les erreurs avec ce retour
		return fichier.Save("./tasks.xlsx")

	//Dans ce cas-ci la liste est vide
	} else {

		//On dit que la liste est vide
		fmt.Printf("%vLa liste est vide.%v\n", escapeCodeRed,  escapeCodeColorReset)

		//On crée une erreur et on la renvoie plus haut
		return errors.New("La liste des tâches est vide")

	}
	
}

//Permet de supprimer un élément de la liste
func deleteElement() bool{

	//Est-ce que tout s'est bien passé ?
	ok := false

	//Tant que j'ai pas rentré un numéro valide ou que la liste n'est pas vide
	for {

		//Si la liste n'est pas vide
		if len(listeTaches) != 0 {

			//Message initial
			fmt.Println("Supprimer un élément parmi cette liste")

			//Lire les tâches pour aider notre utilisateur
			listTask()

			//On lit la chaîne de caractères
			line := readLine()

			//Je parse ce que j'ai lu en int
			number, err := strconv.Atoi(line)

			//Si le nombre se trouve dans la range de la liste
			if err == nil && number >= 1 && number <= len(listeTaches) {

				// Supprimer l'élément de la liste
				listeTaches = append(listeTaches[:number-1], listeTaches[number:]...)

				//Tout s'est bien passé
				ok = true

				//On dit à l'utilisateur que l'élément est supprimé
				fmt.Printf("%vÉlément supprimé avec succès !%v\n", escapeCodeGreen, escapeCodeColorReset)

				//On retourne ok, par principe
				return ok

			} else {

				//Message d'erreur, l'utilisateur n'a pas pris un nombre entre 1 et la taille de la liste
				fmt.Printf("%vChoisissez un nombre entre 1 et %v.%v\n", escapeCodeRed, len(listeTaches), escapeCodeColorReset)

			}	
		} else {

			//Message d'erreur pour dire que la liste est vide
			fmt.Printf("%vLa liste est vide.%v\n", escapeCodeRed,  escapeCodeColorReset)

			//Retourner false
			return ok

		}
	}
}