package service

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
	"golangProject/cinemaBuy/repository"
	"time"
)

//电影五天内有档期的影院
func GetCinemaByFilm(filmId int) [5]gdb.Result {
	var cinema [5]gdb.Result
	now := time.Now()
	cinema[0] = repository.GetCinemaByFilm(filmId, now.Format("2006-01-02"), now.Format("15:04:05"))
	tomorrow := now.AddDate(0, 0, 1)
	cinema[1] = repository.GetCinemaByFilm(filmId, tomorrow.Format("2006-01-02"), "00:00:00")
	afterTomorrow := tomorrow.AddDate(0, 0, 1)
	cinema[2] = repository.GetCinemaByFilm(filmId, afterTomorrow.Format("2006-01-02"), "00:00:00")
	other1 := afterTomorrow.AddDate(0, 0, 1)
	cinema[3] = repository.GetCinemaByFilm(filmId, other1.Format("2006-01-02"), "00:00:00")
	other2 := other1.AddDate(0, 0, 1)
	cinema[4] = repository.GetCinemaByFilm(filmId, other2.Format("2006-01-02"), "00:00:00")
	return cinema
}

//某影院的某电影五天内的档期
func GetMovieByCinemaFilm(cinemaId int) [][5]gdb.Result {
	var result [][5]gdb.Result
	film := repository.GetFilmPic(cinemaId)
	for _, value := range film {
		var movie [5]gdb.Result
		now := time.Now()
		movie[0] = repository.GetMovieInfo(gconv.Int(value["filmId"]), cinemaId, now.Format("2006-01-02"), now.Format("15:04:05"))
		tomorrow := now.AddDate(0, 0, 1)
		movie[1] = repository.GetMovieInfo(gconv.Int(value["filmId"]), cinemaId, tomorrow.Format("2006-01-02"), "00:00:00")
		afterTomorrow := tomorrow.AddDate(0, 0, 1)
		movie[2] = repository.GetMovieInfo(gconv.Int(value["filmId"]), cinemaId, afterTomorrow.Format("2006-01-02"), "00:00:00")
		other1 := afterTomorrow.AddDate(0, 0, 1)
		movie[3] = repository.GetMovieInfo(gconv.Int(value["filmId"]), cinemaId, other1.Format("2006-01-02"), "00:00:00")
		other2 := other1.AddDate(0, 0, 1)
		movie[4] = repository.GetMovieInfo(gconv.Int(value["filmId"]), cinemaId, other2.Format("2006-01-02"), "00:00:00")
		result = append(result, movie)
	}
	return result
}
