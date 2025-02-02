package piscine

//Fonctions qui vont nous permettrent de changer le fond de la page web

//Fonction qui gère l'arrière plan
func Back(etape int) string {
	tab := []string{
		"background0",
		"background1",
		"background2",
		"background3",
		"background4",
		"background5",
		"background6",
		"background7",
		"background8",
		"background9",
		"background10",
	}

	return tab[etape]
}

//Fonction qui permet de gérer les bordures d'un bloc
func Border(etape int) string {
	tab := []string{
		"text-box0",
		"text-box1",
		"text-box2",
		"text-box3",
		"text-box4",
		"text-box5",
		"text-box6",
		"text-box7",
		"text-box8",
		"text-box9",
		"text-box10",
	}

	return tab[etape]
}

//Fonction qui permet de gérer les bordures d'un second bloc
func Border1(etape int) string {
	tab := []string{
		"text-box-bottom0",
		"text-box-bottom1",
		"text-box-bottom2",
		"text-box-bottom3",
		"text-box-bottom4",
		"text-box-bottom5",
		"text-box-bottom6",
		"text-box-bottom7",
		"text-box-bottom8",
		"text-box-bottom9",
		"text-box-bottom10",
	}
	return tab[etape]
}

//Fonction qui permet de gérer les bordure d'un boutton
func Input(etape int) string {
	tab := []string{
		"zone-input0",
		"zone-input1",
		"zone-input2",
		"zone-input3",
		"zone-input4",
		"zone-input5",
		"zone-input6",
		"zone-input7",
		"zone-input8",
		"zone-input9",
		"zone-input10",
	}
	return tab[etape]
}
