package main

import (
	fetch "Func/API"
	Func "Func/Routes"
	Funca "Func/funcs"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Requete HTTP test de la route "/"
	requete, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Enregistreur de reponse HTTP de test
	enregistre := httptest.NewRecorder()

	// Appelle du handler de la route "/"
	Func.Home(enregistre, requete)

	// Vérifier si les codes de statut correspondent
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v, got %v", http.StatusOK, status)
	}
}
func TestSearchHandler(t *testing.T) {
	requete, err := http.NewRequest("POST", "/search", nil)
	if err != nil {
		t.Fatal(err)
	}

	enregistre := httptest.NewRecorder()
	Func.Search(enregistre, requete)
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestLocalizationHandler(t *testing.T) {
	requete, err := http.NewRequest("POST", "/localization", nil)
	if err != nil {
		t.Fatal(err)
	}

	enregistre := httptest.NewRecorder()
	Func.Localization(enregistre, requete)
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestNorepeatart(t *testing.T) {
	test := []struct {
		Case     []fetch.Artists
		Expected []fetch.Artists
	}{
		//case 1
		{[]fetch.Artists{{1, "image1", "vincent", []string{"1", "2", "3"}, 1980, "2005"},
			{2, "image2", "felix", []string{"4", "5", "6"}, 1970, "2009"},
			{1, "image1", "vincent", []string{"1", "2", "3"}, 1980, "2005"},
			{2, "image2", "felix", []string{"4", "5", "6"}, 1970, "2009"},
			{3, "image3", "nabou", []string{"7", "8", "8"}, 1986, "2007"},
			{3, "image3", "nabou", []string{"7", "8", "8"}, 1986, "2007"}},
			// expected 1
			[]fetch.Artists{{1, "image1", "vincent", []string{"1", "2", "3"}, 1980, "2005"},
				{2, "image2", "felix", []string{"4", "5", "6"}, 1970, "2009"},
				{3, "image3", "nabou", []string{"7", "8", "8"}, 1986, "2007"}}},

		//case 2
		{[]fetch.Artists{{0, "", "", []string{"", "", ""}, 0, ""},
			{0, "", "", []string{"", "", ""}, 0, ""},
			{0, "", "", []string{"", "", ""}, 0, ""},
			{0, "", "", []string{"", "", ""}, 0, ""},
		},
			// expected 2
			[]fetch.Artists{{0, "", "", []string{"", "", ""}, 0, ""}}},

		//case 3
		{[]fetch.Artists{{1, "image1", "vincent", []string{"1", "2", "3"}, 1980, "2005"},
			{2, "image2", "felix", []string{"4", "5", "6"}, 1970, "2009"},
			{3, "image3", "nabou", []string{"7", "8", "8"}, 1986, "2007"}},
			// expected 3
			[]fetch.Artists{{1, "image1", "vincent", []string{"1", "2", "3"}, 1980, "2005"},
				{2, "image2", "felix", []string{"4", "5", "6"}, 1970, "2009"},
				{3, "image3", "nabou", []string{"7", "8", "8"}, 1986, "2007"}}},
	}

	for _, v := range test {
		if len(Funca.Norepeatart(v.Case)) != len(v.Expected) {
			t.Fatalf("expected: %v, got: %v", v.Expected, Funca.Norepeatart(v.Case))
		} else {
			t.Log("Test succeedeed")
		}
	}
}

func TestNorepeat(t *testing.T) {
	test1 := []string{"a", "b", "c", "d", "c", "t"}
	want1 := []string{"a", "b", "c", "d", "t"}
	test2 := []string{"", "", ""}
	want2 := []string{""}

	got1 := Funca.Norepeat(test1)
	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}
	got2 := Funca.Norepeat(test2)
	if !reflect.DeepEqual(want2, got2) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}
}

func TestReverse(t *testing.T) {
	str := "Masseck-Thiaw-le-grand-dev"
	want := "dev-grand-le-Thiaw-Masseck"
	str2 := "Essayons pour-voir"
	want2 := "voir-Essayons pour"
	got := Funca.Reverse(str)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	got2 := Funca.Reverse(str2)
	if !reflect.DeepEqual(want2, got2) {
		t.Fatalf("expected: %v, got: %v", want2, got2)
	}
}

func TestValidtab(t *testing.T) {
	idtest1 := []int{2, 4, 3, 6, 9, 7}
	restest1 := []int{1, 5, 4, 7, 9}
	idtest2 := []int{}
	restest2 := []int{}
	want1 := []int{4, 9, 7}
	want2 := []int{}
	idtest3 := []int{2, 3, 6, 9, 7}
	restest3 := []int{1, 5, 4, 8}
	var want3 []int
	got1 := Funca.Validtab(idtest1, restest1)
	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}
	got2 := Funca.Validtab(idtest2, restest2)
	if !reflect.DeepEqual(want2, got2) {
		t.Fatalf("expected: %v, got: %v", want2, got2)
	}
	got3 := Funca.Validtab(idtest3, restest3)
	if !reflect.DeepEqual(want3, got3) {
		t.Fatalf("expected: %v, got: %v", want3, got3)
	}
}

func TestArtistsHandler(t *testing.T) {
	// Requete HTTP test de la route "/artists"
	requete, err := http.NewRequest("GET", "/artists", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Enregistreur de reponse HTTP de test
	enregistre := httptest.NewRecorder()

	// Appelle du handler de la route "/artists"
	Func.Artists(enregistre, requete)

	// Vérifier si les codes de statut correspondent
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestFilter(t *testing.T) {

	requete, err := http.NewRequest("POST", "/filter", nil)
	if err != nil {
		t.Fatal(err)
	}
	enregistre := httptest.NewRecorder()

	Func.Filter(enregistre, requete)

	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestInfoHandler(t *testing.T) {
	// Requete HTTP pour "/info/{id}"
	requete, err := http.NewRequest("GET", "/info/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Enregistreur de reponse HTTP
	enregistre := httptest.NewRecorder()

	// Appelle du handler de la route "/info/{id}"
	Func.Info(enregistre, requete)

	// Vérifier si les codes de statut correspondent
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong code returned. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestHandlers(t *testing.T) {
	// Applle des fonctions de test handler
	t.Run("Home", TestHomeHandler)
	t.Run("Artists", TestArtistsHandler)
	t.Run("Info", TestInfoHandler)
	t.Run("Search", TestSearchHandler)
	t.Run("Filter", TestFilter)
	t.Run("Localization", TestLocalizationHandler)
}

func TestMain(t *testing.T) {
	t.Run("main", TestHandlers)
}
