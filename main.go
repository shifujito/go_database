package main

import (
	"database/sql"
	"strconv"
	"strings"

	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)
func main(){
	application := app.New()
	window := application.NewWindow("app")
	application.Settings().SetTheme(theme.DarkTheme())
	edit := widget.NewMultiLineEntry()
	sc := widget.NewScrollContainer(edit)
	fnd := widget.NewEntry()
	inf := widget.NewLabel("information bar.")

	showInfo := func(s string){
		inf.SetText(s)
		dialog.ShowInformation("info", s, window)
	}

	err := func (er error) bool {
		if er != nil {
			inf.SetText(er.Error())
		}
		return false
	}

	setDB := func() *sql.DB{
		con, er := sql.Open("sqlite3", "data.sqlite3")
		if err(er){
			return nil
		}
		return con
	}

	nf := func() {
		dialog.ShowConfirm("Alert", "Clear form?", func(f bool){
			if f {
				fnd.SetText("")
				window.SetTitle("App")
				edit.SetText("")
				inf.SetText("Clear form.")
			}
		}, window)
	}

	wf := func(){
		fstr := fnd.Text
		if !strings.HasPrefix(fstr, "http"){
			fstr = "http://" + fstr
			fnd.SetText(fstr)
		}
		dc, er:= goquery.NewDocument(fstr)
		if err(er){
			return
		}
		ttl := dc.Find("title")
		window.SetTitle(ttl.Text())
		html, er := dc.Html()
		if err(er){
			return
		}
		cvtr := md.NewConverter("", true, nil)
		mkdn, er := cvtr.ConvertString(html)
		if err(er){
			return
		}
		edit.SetText(mkdn)
		inf.SetText("get web data.")
	}

	ff := func(){
		var qry string = "select * from md_data where title like ?"
		con := setDB()
		if con == nil{
			return
		}
		defer con.Close()

		rs, er := con.Query(qry, "%"+fnd.Text+"%")
		if err(er){
			return
		}
		res := ""
		for rs.Next(){
			var ID int
			var TT string
			var UR string
			var MR string
			er := rs.Scan(&ID, &TT, &UR, &MR)
			if err(er){
				return
			}
			res += strconv.Itoa(ID) + ":" + TT + "\n"
		}
		edit.SetText(res)
		inf.SetText("Find:" + fnd.Text)
	}
	idf := func (id int)  {
		var qry string = "select * from md_data where id = ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		rs := con.QueryRow(qry, id)

		var ID int
		var TT string
		var UR string
		var MR string
		rs.Scan(TT)
		fnd.SetText(UR)
		edit.SetText(MR)
		inf.SetText("Find id=" + strconv.Itoa(ID) + ".")
	}

	sf := func ()  {
		dialog.ShowConfirm("Alert", "Save data?", func(f bool) {
			if f{
				con := setDB()
				if con == nil{
					return
				}
				defer con.Close()
				qry := "insert into md_data (title,url,markdown) values (?,?,?)"
				_, er := con.Exec(qry, window.Title(), fnd.Text, edit.Text)
				if err(er){
					return
				}
				showInfo("Save data to daatabase!")
			}
		}, window)
	}
}
