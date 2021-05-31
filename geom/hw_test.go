package geom

import "testing"

func TestGeom_Distance(t *testing.T) {
	x1, y1, x2, y2 := 1.0, 1.0, 4.0, 5.0 // тестовый пример
	got, _ := Distance(x1, y1, x2, y2)   // вызов тестируемого кода
	want := 5.0                          // заранее вычисленный результат
	if got != want {                     // сравнение результата с правильным значением
		t.Errorf("получили %f, ожидалось %f", got, want)
	}
	t.Log("OK")
}
