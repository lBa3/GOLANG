package main

//modelo o clase o molde
type Movie struct {
  Name string  `json:"name"`
  Year int     `json:"anio"`
  Ower string  `json:"director"`
}

type Movies []Movie