package main


import ("fmt");
import ("net/http");
import ("github.com/gorilla/mux");
import ("encoding/json");
import ("log");
import ("gopkg.in/mgo.v2");
import ("gopkg.in/mgo.v2/bson");


//conector a mongo
 func getSession() *mgo.Session {
  session, err := mgo.Dial("mongodb://localhost")
    if err != nil{
       panic(err);
    }
    return session;
 } 

func responseMovie(w http.ResponseWriter, status int, results Movie){
  w.Header().Set("Content-Type", "application/json");
  w.WriteHeader(200);
  json.NewEncoder(w).Encode(results);
}

func responseMovies(w http.ResponseWriter, status int, results []Movie){
  w.Header().Set("Content-Type", "application/json");
  w.WriteHeader(200);
  json.NewEncoder(w).Encode(results);
}


var collection = getSession().DB("curso_Go").C("movies");


func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"::::::::::::::MUNDO PELUCULAS:::::::::::::::::::::");
}


func MovieList(w http.ResponseWriter, r *http.Request){
  //consulta BD
  var results []Movie
  err := collection.Find(nil).All(&results);

  if err != nil{
    log.Fatal(err);
  }else{
    fmt.Println("Resultados", results);
  }


  responseMovies(w, 200, results);


}

func MovieShow(w http.ResponseWriter, r *http.Request){
  parametro := mux.Vars(r);
  movie_id := parametro["id"];

  if !bson.IsObjectIdHex(movie_id){
    w.WriteHeader(404);
    return;
  }

  oid := bson.ObjectIdHex(movie_id);
  fmt.Println(movie_id);
  fmt.Println(oid);
  results := Movie{}
  err := collection.FindId(oid).One(&results);
  fmt.Println(err);

  if err != nil{
    w.WriteHeader(404);
    return;
  }else{
      responseMovie(w, 200, results);
  }

}




func MovieAdd(w http.ResponseWriter, r *http.Request){
  decoder := json.NewDecoder(r.Body)
  var movie_data Movie;
  err := decoder.Decode(&movie_data);
  if (err != nil){
     panic(err);
  }
  defer r.Body.Close();
  log.Println(movie_data);
  
  //llamada de BD e insert
  err = collection.Insert(movie_data);
  if err != nil{
    w.WriteHeader(500);
    return; 
  }


  //llamado 
  responseMovie(w, 200, movie_data);

}

func MovieUpdate(w http.ResponseWriter, r *http.Request){
  parametro := mux.Vars(r);
  movie_id := parametro["id"];
  
  if !bson.IsObjectIdHex(movie_id){
    w.WriteHeader(404);
    return;
  }

  oid := bson.ObjectIdHex(movie_id);
  decoder := json.NewDecoder(r.Body);

  
  fmt.Println(movie_id);
  fmt.Println(oid);

  var movie_data Movie;
  err := decoder.Decode(&movie_data);

  if err != nil{
    panic(err);
    w.WriteHeader(404);
    return;
  }

  defer r.Body.Close();

  document :=  bson.M{"_id":oid}
  change := bson.M{"$set":movie_data}
  err = collection.Update(document,change);

  if err != nil{
    panic(err);
    w.WriteHeader(404);
    return; 
  }


  responseMovie(w, 200, movie_data);

}


type Message struct{
  Status string  `json:"status"`
  Message string `json:"message"`
}

func MovieDelete(w http.ResponseWriter, r *http.Request){
  parametro := mux.Vars(r);
  movie_id := parametro["id"];

  if !bson.IsObjectIdHex(movie_id){
    w.WriteHeader(404);
    return;
  }

  oid := bson.ObjectIdHex(movie_id);
  fmt.Println(movie_id);
  fmt.Println(oid);
  //results := Movie{}
  err := collection.RemoveId(oid);
  fmt.Println(err);

  if err != nil{
    w.WriteHeader(404);
    return;
  }else{
      results := Message{"success", "la pelicula con el ID " + movie_id + " Se elimino correctamente"}
      w.Header().Set("Content-Type", "application/json");
      w.WriteHeader(200);
      json.NewEncoder(w).Encode(results);

  }

}
