package main

//Les importations
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
	"github.com/tealeg/xlsx/v3"
)

//La structure de la tâche
type Task struct {
	Nom        string
	Heures     float64
	Facturable bool
	Type 	   string
}

//La structure d'une commande
type Commande struct{
	Nom string
	Usage string
	Aliases []string
	Description string
	Run func()
}

//La liste des tâches
var listeTaches []Task

//La liste des commandes
var listeCommandes []Commande

//Pour mettre des couleurs dans le terminal
var (
	escapeCodeGreen  = color.New(color.FgGreen).SprintFunc()
	escapeCodeRed    = color.New(color.FgRed).SprintFunc()
	escapeCodeViolet = color.New(color.FgHiMagenta).SprintFunc()
	escapeCodeRose   = color.New(color.FgHiCyan).SprintFunc()
	escapeCodeBlue = color.New(color.FgCyan).SprintFunc()
	escapeCodeYellow = color.New(color.FgYellow).SprintFunc()
	escapeCodeColorReset = color.New(color.Reset).SprintFunc()
)

//La fonction principale de mon programme
func main() {

	//Charger les commandes
	loadCommands()

	//L'équivalent d'un while
	for {

		//Lire la commande
		readCommand()
			
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
		fmt.Printf("%sVeuillez entrer un nombre valide.%s\n",escapeCodeRed(), escapeCodeColorReset())
	}
}

//Permet de lire le fait qu'une tâche soit facturable ou pas
func readBool() bool {

	//Tant que j'ai pas dis oui ou non
	for {

		//Je lis mon entrée
		line := strings.ToLower(readLine())

		if len(line) == 0 {

			line = "oui"

		}

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
		fmt.Printf("%sVeuillez entrer 'oui' ou 'non'.%s\n", escapeCodeRed(), escapeCodeColorReset())
	}
}

//Permet de le choix fait sur le type de la tâche
func readType() string {

	//Tant que je n'ai pas donné un nombre entre 1 et 5
	for {

		//On lit la chaine de caractère
		line := readLine()

		if len(line) == 0 {

			line = "1"

		}

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
					fmt.Printf("%vEntrez un nombre entre 1 et 5.\n%v", escapeCodeRed(), escapeCodeColorReset())
			}
		
		//S'il y a une erreur
		} else {

			//Le faire savoir à l'utilisateur
			fmt.Printf("%vCette entrée n'est pas un nombre valide.\n%v", escapeCodeRed(), escapeCodeColorReset())

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
	fmt.Println("Cette tâche est-elle facturable ? Choix par défaut [Oui]")
	tache.Facturable = readBool()
	fmt.Println()

	//On demande le type de tâche que c'est
	fmt.Printf("Quel est le type de la tâche ?\n")
	
	fmt.Println("[1] Normal")
	color.New(color.FgRed).Println("[2] Problème")
	color.New(color.FgGreen).Println("[3] Solution")
	color.New(color.FgCyan).Println("[4] Réunion")
	color.New(color.FgYellow).Println("[5] Divers")
	fmt.Println()
	fmt.Println("Choix par défaut [1]")
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
		color.New(color.FgRed).Println("La liste des tâches est vide.")

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
					nom = color.WhiteString("[%v] %v",i, tache.Nom)
					break;
				
				//Si le type de la tâche est un problème
				case "problème":

					//On met le nom en rouge
					nom = color.RedString("[%v] %v",i, tache.Nom)
					break;
				
				//Si la tâche est une solution à un problème
				case "solution":

					//On met son texte en vert
					nom = color.GreenString("[%v] %v",i, tache.Nom)
					break;
				
				//Si cette tâche est une réunion avec les chefs de projets ou le client
				case "réunion":

					//On met le texte en bleu
					nom = color.CyanString("[%v] %v",i, tache.Nom)
					break;
				
				//Si cette tâche est une chose divers
				case "divers":

					//On met cette tâche en jaune
					nom = color.YellowString("[%v] %v",i, tache.Nom)
					break;

			}

			//Donner les infos de la tâche à l'utilisateur
			fmt.Printf("%v\nHeures: %v\nFacturable: %v\nType de tâche: %v\n\n", nom , tache.Heures, facturable, tache.Type)
			
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

//Permet d'exporter la liste des commandes en fichier classeur pour Excel ou LibreOffice
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

		//Couleurs
		rouge := *xlsx.NewFill("color", "FFFF0000", "FFFF0000")
		vert := *xlsx.NewFill("color", "FF23B84B","FF23B84B")
		bleu := *xlsx.NewFill("color", "FF00D5FF","FF00D5FF")
		jaune := *xlsx.NewFill("color", "FFD4FC21","FFD4FC21")

		//Style des problèmes
		styleProbleme := xlsx.NewStyle()
		styleProbleme.Fill = rouge
		styleProbleme.Font.Name = "Arial"
		styleProbleme.Font.Size = 8
		styleProbleme.Font.Bold = true
		styleProbleme.ApplyFill = true
		styleProbleme.ApplyFont = true

		//Style des solutions
		styleSolution := xlsx.NewStyle()
		styleSolution.Fill = vert
		styleSolution.Font.Name = "Arial"
		styleSolution.Font.Size = 8
		styleSolution.Font.Bold = true
		styleSolution.ApplyFill = true
		styleSolution.ApplyFont = true

		//Style des réunions
		styleReunion := xlsx.NewStyle()
		styleReunion.Fill = bleu
		styleReunion.Font.Name = "Arial"
		styleReunion.Font.Size = 8
		styleReunion.Font.Bold = true
		styleReunion.ApplyFill = true
		styleReunion.ApplyFont = true

		//Style des choses diverses
		styleDivers := xlsx.NewStyle()
		styleDivers.Fill = jaune
		styleDivers.Font.Name = "Arial"
		styleDivers.Font.Size = 8
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

					//Je met le style que la cellule doit avoit (Gras + rouge)
					cell.SetStyle(styleProbleme)

					//Je formate la tâche en "Poblème #numéro: nom de la tâche"
					cell.Value = strings.Join([]string{"Problème #", fmt.Sprintf("%v",problemID), ": ", tache.Nom}, "")
				

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
		color.New(color.FgRed).Println("La liste est vide.")

		//On crée une erreur et on la renvoie plus haut
		return errors.New("La liste des tâches est vide")

	}
	
}

func exportToCsv() {



}

//Permet de supprimer un élément de la liste
func deleteElement() {

	//Tant que j'ai pas rentré un numéro valide ou que la liste n'est pas vide
	for {

		//Si la liste n'est pas vide
		if len(listeTaches) != 0 {

			//Message initial
			fmt.Println("Supprimer un élément parmi cette liste")

			//Lire les tâches pour aider notre utilisateur
			listTask()

			//Message d'erreur, l'utilisateur n'a pas pris un nombre entre 1 et la taille de la liste
			fmt.Printf("Choisissez un nombre entre 1 et %v ou \"cancel\" pour annuler.\n", len(listeTaches))

			//On lit la chaîne de caractères
			line := readLine()

			if strings.ToLower(line) == "cancel" {

				//Message d'erreur, l'utilisateur n'a pas pris un nombre entre 1 et la taille de la liste
				color.New(color.FgRed).Println("Commande annulée.")

				//On sort de cette boucle
				break;

			} else {

				//Je parse ce que j'ai lu en int
				number, err := strconv.Atoi(line)

				//Si le nombre se trouve dans la range de la liste
				if err == nil && number >= 1 && number <= len(listeTaches) {

					// Supprimer l'élément de la liste
					listeTaches = append(listeTaches[:number-1], listeTaches[number:]...)

					//On dit à l'utilisateur que l'élément est supprimé
					color.New(color.FgGreen).Println("Élément supprimé avec succès !")

					//On sort de la boucle
					break;

				} else  {

					//Message d'erreur, l'utilisateur n'a pas pris un nombre entre 1 et la taille de la liste
					color.New(color.FgRed).Printf("Choisissez un nombre entre 1 et %v ou \"cancel\" pour annuler.\n", len(listeTaches))

				}	
			}
		} else {

			//Message d'erreur pour dire que la liste est vide
			color.New(color.FgRed).Println("La liste est vide.")

			//On sort de la boucle
			break;
		}
	}
}

//Permet de lire la commande
func readCommand(){

	//Lire l'entrée de l'utilisateur
	color.New(color.FgHiMagenta).Print("TaskManager")

	color.New(color.FgCyan).Print(" \u21E2 ")
	
	//On prend le texte que l'utilisateur a rentré
	input := strings.ToLower(readLine())

	//On a pas trouvé la commande pour l'instant
	commandFounded := false

	//Je défini la cible si la commande est trouvé
	var targetCommand Commande

	//On parcourt la liste des commandes
	for _, command := range listeCommandes {

		//Si le nom de la commande est la même que mon input et de manière case unsensitive
		if command.Nom == input || contains(command.Aliases, input){

			//La commande est trouvée
			commandFounded = true

			//On prend la commande cible
			targetCommand = command

			//On sort de la boucle
			break;

		}

	}

	//Si la commande est trouvée
	if commandFounded {

		//On démarre la commande
		targetCommand.Run()

	} else {

		//Dire à l'utilisateur que sa commande n'est pas valide
		color.New(color.FgRed).Printf("Commande invalide: %s.\n", strings.ToLower(input)) 

	}

}

//Affiche une page d'aide avec toutes les commandes de l'application
func printCommands(){

	//Déclaration vide pour la prise en charge des pluriels
	commandText := ""

	//Parcourir la liste des commandes
	for _, command := range listeCommandes {

		//Afficher les infos de la commande
		fmt.Printf("%v: %v\nUsage: %v\nAliases: %v\n\n", command.Nom, command.Description, command.Usage, strings.Join(command.Aliases, ", "))

	}

	//Si le nombre de commande est différent de 1 ou 0
	if len(listeCommandes) != 1 || len(listeCommandes) != 0 {

		//Mettre commande au pluriel
		commandText = "commandes"

	} else {

		//Mettre command au singulier
		commandText = "commande"

	}

	//Afficher le nombre de commandes chargées
	fmt.Println()
	fmt.Printf("Il y a actuellement %v %v !\n", len(listeCommandes), commandText)
}

//Permet de quitter l'application
func exit(){

	//On dit au revoir à l'utilisateur
	fmt.Println("Au revoir !")

	//Attendre une seconde
	time.Sleep(1 * time.Second)

	//On ferme le programme
	os.Exit(0)

}

//Permet de charger les commandes
func loadCommands(){

	//Commande pour ajouter une tâche
	addTaskCommand := Commande{
		Nom: "add task",
		Usage:"add task",
		Aliases: []string{"addt", "at"},
		Description:"Permet d'ajouter une tâche en mémoire",
		Run: addTask,
	}

	//Commande pour supprimer une tâche
	deleteTaskCommand := Commande{
		Nom: "delete task",
		Aliases: []string{"delt", "dt"},
		Usage: "delete task",
		Description: "Permet de supprimer une tâche",
		Run: deleteElement,
	}

	//Commande pour lister les tâches
	listTasksCommand := Commande{
		Nom: "list tasks",
		Aliases: []string{"list", "lt"},
		Usage: "list tasks",
		Description: "Permet de lister les tâches",
		Run: listTask,
	}

	//Commande d'exportation du ficher classeur
	exportCommand := Commande{
		Nom: "export",
		Aliases: []string{"exp", "save"},
		Usage: "export",
		Description: "Exporte la liste des tâches de la console dans un fichier classeur de type Excel",
		Run: func() {

			//Dire à l'utilisateur que le fichier est en cours d'exportation
			fmt.Println("Le fichier est en cours d'exportation...")

			//Exporter la liste des tâches avec les heures dans le fichier
			erreur := exportToXlsx()

			//S'il n'y a pas d'erreur
			if erreur == nil {

				//Message de succès
				color.New(color.FgGreen).Println("Ce fichier a été exporté avec succès !")

			//Il y a eu une erreur
			} else {

				//Message d'erreur
				color.New(color.FgRed).Println("Une erreur est survenue lors de l'exportation !")

			}
		},
	}

	//Commande pour quitter l'application
	exitCommand := Commande{
		Nom:"exit",
		Aliases: []string{"ex"},
		Usage:"exit",
		Description: "Permet de quitter l'application.",
		Run: exit,
	}

	//Commande qui permet de mettre à jour la tâche


	//Commande qui affiche la page d'aide
	helpCommand := Commande{
		Nom: "help",
		Aliases: []string{"h"},
		Usage: "help",
		Description:"Affiche la page d'aide",
		Run: printCommands,
	}

	//Chargement des commandes
	listeCommandes = append(listeCommandes, addTaskCommand)
	listeCommandes = append(listeCommandes, deleteTaskCommand)
	listeCommandes = append(listeCommandes, listTasksCommand)
	listeCommandes = append(listeCommandes, exitCommand)
	listeCommandes = append(listeCommandes, helpCommand)
	listeCommandes = append(listeCommandes, exportCommand)

	//Si part hasard la liste de commandes n'est pas vide
	if len(listeCommandes) != 0 {

		//Message de succès
		color.New(color.FgGreen).Printf("Les commandes ont été chargées avec succès. Actuellement il y a %x commandes!\n", len(listeCommandes))

	} else {

		//Message d'erreur
		color.New(color.FgRed).Println("Le chargement des commandes a échoué, la liste est vide.")

	}
}

// contains retourne true si la chaîne existe dans la liste, sinon false
func contains(list []string, str string) bool {
    for _, s := range list {
        if s == str {
            return true
        }
    }
    return false
}