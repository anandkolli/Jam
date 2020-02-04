package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"jam/pkg/pg"
)

type AllActivity struct {
	Count string `json: "count"`
	Time  string `json: "time"`
}

type BreakOuts struct {
	P []AllActivity `json: "activity"`
}

func GetAllActivity(db *pg.PgClient, w http.ResponseWriter, r *http.Request) error {
	log.Println("GetAllActivity")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	type Person struct {
		ID     string `json: "id"`
		FaceID string `json: "faceID"`
		Ts     string `json: "timeStamp"`
	}

	type Activity struct {
		Persons []Person `json: "persons"`
	}

	var (
		a []Person
		rows *sql.Rows
		err error
	)

	breakout := r.FormValue("breakout")

	log.Println("breakout=", breakout)
	if breakout != "" {
		ret, err := GetTimeSpent(breakout, db, w, r)
		if err == nil {
			json.NewEncoder(w).Encode(ret)
			return nil
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}

	if rows, err = db.Pdb.Query("SELECT * FROM face_activity ORDER BY aws_face_id ASC, the_time ASC"); err != nil {
		log.Println("Error querying DB", err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	defer rows.Close()

	log.Println(rows, err)
	var data Activity
	for rows.Next() {
		var aa Person
		var id, faceID, ts string
		if err = rows.Scan(&id, &faceID, &ts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		aa.ID = id
		aa.FaceID = faceID
		aa.Ts = ts
		a = append(a, aa)
		log.Println("ID | FaceID | timeStamp")
		fmt.Printf("%3v | %8v | %6v \n", id, faceID, ts)
	}
	data.Persons = a
	log.Println("Data", data)
	json.NewEncoder(w).Encode(data)
	return nil

}

func GetTimeSpentPerPerson(db *pg.PgClient, w http.ResponseWriter, r *http.Request) error {
	log.Println("GetTimeSpentPerPerson")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	type Person struct {
		Email string `json:"email"`
		Count string `json:"count"`
	}
	var (
		per  []Person
		rows *sql.Rows
		err error
	)

	duration := r.FormValue("duration")

	log.Println("duration=", duration)
	if duration == "" {
		if rows, err = db.Pdb.Query("SELECT faces.email, COUNT(*) AS c FROM face_activity INNER JOIN faces ON faces.aws_face_id = face_activity.aws_face_id GROUP BY faces.email"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	} else {
		query := fmt.Sprintf("SELECT faces.email, COUNT(*) AS c FROM face_activity INNER JOIN faces ON faces.aws_face_id = face_activity.aws_face_id GROUP BY faces.email HAVING COUNT(*) > %s",duration)
		if rows, err = db.Pdb.Query(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}

	defer rows.Close()

	for rows.Next() {
		var email string
		var count string
		var p Person
		if err = rows.Scan(&email, &count); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		p.Email = email
		p.Count = count
		per = append(per, p)
		log.Println("Email | Count")
		fmt.Printf("%3v | %8v \n", email, count)
	}
	log.Println("Data", per)
	json.NewEncoder(w).Encode(per)
	return nil
}

func GetPersonsSpentAnyTime(db *pg.PgClient, w http.ResponseWriter, r *http.Request) error {
	log.Println("GetPersonsSpentAnyTime")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	type Person struct {
		Count string `json:"count"`
	}
	var (
		data Person
		rows *sql.Rows
		err error
	)

	duration := r.FormValue("duration")

	log.Println("duration=", duration)
	if duration == "" {
		if rows, err = db.Pdb.Query("WITH foo AS(SELECT aws_face_id, COUNT(*) FROM face_activity GROUP BY aws_face_id) SELECT COUNT(*) FROM foo"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	} else {
		query := fmt.Sprintf("WITH foo AS(SELECT aws_face_id, COUNT(*) FROM face_activity GROUP BY aws_face_id HAVING COUNT(*) > %s) SELECT COUNT(*) FROM foo",duration)
		if rows, err = db.Pdb.Query(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}

	defer rows.Close()

	for rows.Next() {
		var count string
		if err = rows.Scan(&count); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		data.Count = count
		log.Println("Count")
		fmt.Printf("%3v\n", count)
	}
	log.Println("Data", data)
	json.NewEncoder(w).Encode(data)
	return nil
}

func GetTimeSpent(breakout string, db *pg.PgClient, w http.ResponseWriter, r *http.Request) (BreakOuts, error) {
	log.Println("GetTimeSpent")
	var (
		data []AllActivity
		rows *sql.Rows
		err error
		b   BreakOuts
	)

	log.Println("breakout=", breakout)
	query := fmt.Sprintf("SELECT COUNT(*), date_trunc('%s', the_time) FROM face_activity GROUP BY date_trunc",breakout)
	if rows, err = db.Pdb.Query(query); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return BreakOuts{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count string
			time  string
			d AllActivity
		)
		if err = rows.Scan(&count, &time); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return BreakOuts{}, err
		}
		d.Count = count
		d.Time  = time
		data = append(data, d)
		log.Println("Count | Time")
		fmt.Printf("%3v | %8v \n", count, time)
	}
	b.P = data
	return b , nil
}

func GetSessions(db *pg.PgClient, w http.ResponseWriter, r *http.Request) error {
	log.Println("GetSessions")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	type sessionInfo struct {
		FaceID string `json: "faceID"`
		Count string `json: "count"`
	}

	type sessions struct {
		Sinfo []sessionInfo `json: "sessions"`
	}

	var (
		sesSlice []sessionInfo
		rows *sql.Rows
		err error
	)

	if rows, err = db.Pdb.Query("with foo as(select *, lag(the_time) over (partition by aws_face_id order by the_time ASC) from face_activity), choco as (select *, EXTRACT(EPOCH FROM (foo.the_time - foo.lag)) as time_passage from foo),coffee as (select *, case COALESCE(time_passage,601) > 600 when true then 'new' else 'exisiting' end as session_status from choco)select aws_face_id, count(*) as session_count from coffee where coffee.session_status = 'new' group by aws_face_id"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	defer rows.Close()

	var data sessions
	for rows.Next() {
		var ses sessionInfo
		var faceID, count string
		if err = rows.Scan(&faceID, &count); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		ses.FaceID = faceID
		ses.Count = count
		sesSlice = append(sesSlice, ses)
		log.Println("FaceID | Count")
		fmt.Printf("%3v | %8v \n", faceID, count)
	}
	data.Sinfo = sesSlice
	log.Println("Data", data)
	json.NewEncoder(w).Encode(data)
	return nil

}

func GetFunnel(db *pg.PgClient, w http.ResponseWriter, r *http.Request) error {
	log.Println("GetFunnel")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	type funnelInfo struct {
		Value string `json:"page"`
		Name  string `json:"name"`
		Fill  string `json:"fill"`
	}

	var (
		fun []funnelInfo
		rows *sql.Rows
		err error
		query string
	)

	for i:=0;i<3;i++ {
		var f funnelInfo

		switch i {
		case 0:
			f.Name = "less than 1 minute"
			f.Fill = "#8884d8"
			query = "with foo as(select aws_face_id, count(*) from face_activity group by aws_face_id having count(*) < 1) select count(*) from foo"
		case 1:
			f.Name = "1 to 10 minutes"
			f.Fill = "#83a6ed"
			query = "with foo as(select aws_face_id, count(*) from face_activity group by aws_face_id having count(*) <= 10 and count(*) > 1) select count(*) from foo"
		case 2:
			f.Name = "greater than 10 minutes"
			f.Fill = "#8dd1e1"
			query = "with foo as(select aws_face_id, count(*) from face_activity group by aws_face_id having count(*) > 10) select count(*) from foo"
		}
		if rows, err = db.Pdb.Query(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}

		defer rows.Close()

		for rows.Next() {
			var count string
			if err = rows.Scan(&count); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return err
			}
			f.Value = count
			fun = append(fun, f)
			log.Println("Count")
			fmt.Printf("%3v\n", count)
		}
	}

	log.Println("Data", fun)
	json.NewEncoder(w).Encode(fun)
	return nil
}
