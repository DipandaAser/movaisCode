package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(is_there_mon_BFF_dans_ma_classe([]string{"Coddity", "Paul", "Andre", "Moise", "jules"}))
	fmt.Println(is_there_mon_BFF_dans_ma_classe([]string{"mathieu", "azerty", "Jean-BFF <3", "jetbrains", "golang"}))
	fmt.Println(is_there_mon_BFF_dans_ma_classe([]string{"mathieu", "azerty", "Jean-BFF<", "jetbrains", "golang"}))
	fmt.Println(is_there_mon_BFF_dans_ma_classe([]string{"mathieu", "azerty", "Jean-BFF<3", "jetbrains", "golang"}))
	fmt.Println(is_there_mon_BFF_dans_ma_classe([]string{"mathieu", "azerty", "Jean-BFF>3", "jetbrains", "golang"}))
}

func is_there_mon_BFF_dans_ma_classe(listeDesElevesQuiSontDansCetteClasseALaRentree []string) bool {

	for _, nomDunEleveQuiSortDesVacancesEtViensALecole := range listeDesElevesQuiSontDansCetteClasseALaRentree {

		var cestLeNomDeMonBFF bool = false
		lettresQuiConstutieLeNOmDeLeleve := strings.Split(nomDunEleveQuiSortDesVacancesEtViensALecole, "")
		for indexDeLaLettre, lettre := range lettresQuiConstutieLeNOmDeLeleve {

			// we make sure that the index ne nous trompe pas
			if indexDeLaLettre+1 == 1 {
				if isLetterEqualToLaPremiereLettreDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 2 {
				if isLetterEqualToLaDeuxiemeLettreDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 3 {
				if isLetterEqualToLaTroisiemeLettreDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 4 {
				if isLetterEqualToLa4emeLettreDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 5 {
				if isLetterEqualToLa5emeLettreOuSymboleWhateverDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 6 {
				if isLetterEqualToLa6emeLettreDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 7 {
				if isLetterEqualToLa7emeLettreDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 8 {
				if isLetterEqualToLa8emeLettreDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 9 {
				if isLetterEqualToLa9emeLettreOuSymboleouSpaceDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 10 {
				if isLetterEqualToLa10emeLettreOuSymboleInfeurieurDuNomDeMonBFF(lettre) {
					continue
				} else {
					//on evite les injections SQL
					break
				}
			}

			if indexDeLaLettre+1 == 11 {
				if isLetterEqualToLa11emeLettrequiEstUnChiffreDuNomDeMonBFF(lettre) {
					if len(nomDunEleveQuiSortDesVacancesEtViensALecole) > indexDeLaLettre + 1 {
						break
					} else {
						cestLeNomDeMonBFF = true
					}
				} else {
					//on evite les injections SQL
					break
				}
			}
		}

		if cestLeNomDeMonBFF {
			return true
		}
	}

	return false
}

func isLetterEqualToLaPremiereLettreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "J" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLaDeuxiemeLettreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "e" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLaTroisiemeLettreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "a" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa4emeLettreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "n" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa5emeLettreOuSymboleWhateverDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "-" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa6emeLettreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "B" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa7emeLettreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "F" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa8emeLettreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "F" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa9emeLettreOuSymboleouSpaceDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == " " {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa10emeLettreOuSymboleInfeurieurDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "<" {
		return true
	} else {
		return false
	}
}

func isLetterEqualToLa11emeLettrequiEstUnChiffreDuNomDeMonBFF(lettermunisculeuAVairifier string) bool {
	if lettermunisculeuAVairifier == "3" {
		return true
	} else {
		return false
	}
}

//lettermanusculeAVairifier
