# stringInter


Zadatak 5:

napravi main.go fajl i u njemu funkciju koja radi sledece:
    * sa komandne linije uz pomoc flag-ova primi ulazne parametre
    * validira ih
    * ulaz je nekakav string "[ 1, 2, 3, 4, 5 ]" koji treba da pretvoris u niz int-ova
    * treba da uradis deduplikaciju
    * sortiraj izlaz
    * prikazi izlaz na stdout (konzoli)

Pitanja na koja treba dati odgovore:
    * Koja je razlika izmedju 'var' i ':=' definisanja varijabli?


Dodatno (nije obavezno):
    * uradi deduplikaciju 'in place'
    * Go unit testovi

  ------------------------------------------------------------------------------------------------------------------
 * Koja je razlika izmedju 'var' i ':=' definisanja varijabli?
:= je kratka deklaracija varijabli, dok bi var i int bila duza deklaracija varijabli. Kratka deklaracija nije dozvoljena
na top-level statementima dok izmedju blokova mozemo skinuti var deklaraciju ( duzu deklaraciju ) .
Kratka deklaracija dozvoljava
compiler-u da automatski shvati koji tip podataka je u pitanju, dok se uz duzu definisu.
Duza deklaracija:
var i int
i = 42

Kraca (deklaracija i assign)
i := 42
